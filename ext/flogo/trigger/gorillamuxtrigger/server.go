/*
Sniperkit-Bot
- Status: analyzed
*/

/*
* Copyright © 2017. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package gorillamuxtrigger

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// Graceful shutdown HttpServer from: https://github.com/corneldamian/httpway/blob/master/server.go

// NewServer create a new server instance
//param server - is a instance of http.Server, can be nil and a default one will be created
func NewServer(addr string, handler http.Handler, enableTLS bool, serverCert string, serverKey string, enableClientAuth bool, trustStore string) *Server {
	srv := &Server{}
	srv.Server = &http.Server{Addr: addr, Handler: handler}
	srv.enableTLS = enableTLS
	srv.serverCert = serverCert
	srv.serverKey = serverKey
	srv.enableClientAuth = enableClientAuth
	srv.trustStore = trustStore

	return srv
}

//Server the server  structure
type Server struct {
	*http.Server

	serverInstanceID string
	listener         net.Listener
	lastError        error
	serverGroup      *sync.WaitGroup
	clientsGroup     chan bool
	enableTLS        bool
	serverCert       string
	serverKey        string
	enableClientAuth bool
	trustStore       string
}

// InstanceID the server instance id
func (s *Server) InstanceID() string {
	return s.serverInstanceID
}

// Start this will start server
// command isn't blocking, will exit after run
func (s *Server) Start() error {
	if s.Handler == nil {
		return errors.New("No server handler set")
	}

	if s.listener != nil {
		return errors.New("Server already started")
	}

	addr := s.Addr
	if addr == "" {
		addr = ":http"
	}

	hostname, _ := os.Hostname()
	s.serverInstanceID = fmt.Sprintf("%x", md5.Sum([]byte(hostname+addr)))

	if s.enableTLS {
		// log.Debugf("TLS is enabled for the trigger instance - %v%v", s.serverInstanceID, addr)
		//TLS is enabled, load server certificate & key files
		cer, err := tls.LoadX509KeyPair(s.serverCert, s.serverKey)
		if err != nil {
			fmt.Printf("Error while loading certificates - %v", err)
			return err
		}

		var config *tls.Config
		if s.enableClientAuth {
			log.Debugf("TLS with client AUTH is enabled for the trigger instance - %v%v", s.serverInstanceID, addr)
			caCertPool, err := getCerts(s.trustStore)
			if err != nil {
				fmt.Printf("Error while loading client trust store - %v", err)
				return err
			}

			config = &tls.Config{Certificates: []tls.Certificate{cer},
				ClientAuth: tls.RequireAndVerifyClientCert,
				ClientCAs:  caCertPool}
			config.BuildNameToCertificate()
		} else {
			log.Debugf("TLS is enabled for the trigger instance - %v%v", s.serverInstanceID, addr)
			config = &tls.Config{Certificates: []tls.Certificate{cer}}
		}

		// bind secure listener
		listener, err := tls.Listen("tcp", addr, config)
		if err != nil {
			return err
		}
		s.listener = listener
	} else {
		log.Debugf("TLS is not enabled for the trigger instance - %v at port %v", s.serverInstanceID, addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		s.listener = listener
	}

	s.serverGroup = &sync.WaitGroup{}
	s.clientsGroup = make(chan bool, 50000)

	//if s.ErrorLog == nil {
	//    if r, ok := s.Handler.(ishttpwayrouter); ok {
	//        s.ErrorLog = log.New(&internalServerLoggerWriter{r.(*Router).Logger}, "", 0)
	//    }
	//}
	//

	s.Handler = &serverHandler{s.Handler, s.clientsGroup, s.serverInstanceID}

	s.serverGroup.Add(1)
	go func() {
		defer s.serverGroup.Done()
		err := s.Serve(s.listener)
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}

			s.lastError = err
		}
	}()

	return nil
}

// Stop sends stop command to the server
func (s *Server) Stop() error {
	if s.listener == nil {
		return errors.New("Server not started")
	}

	if err := s.listener.Close(); err != nil {
		return err
	}

	return s.lastError
}

// IsStarted checks if the server is started
// will return true even if the server is stopped but there are still some requests to finish
func (s *Server) IsStarted() bool {
	if s.listener != nil {
		return true
	}

	if len(s.clientsGroup) > 0 {
		return true
	}

	return false
}

// WaitStop waits until server is stopped and all requests are finish
// timeout - is the time to wait for the requests to finish after the server is stopped
// will return error if there are still some requests not finished
func (s *Server) WaitStop(timeout time.Duration) error {
	if s.listener == nil {
		return errors.New("Server not started")
	}

	s.serverGroup.Wait()

	checkClients := time.Tick(100 * time.Millisecond)
	timeoutTime := time.NewTimer(timeout)

	for {
		select {
		case <-checkClients:
			if len(s.clientsGroup) == 0 {
				return s.lastError
			}
		case <-timeoutTime.C:
			return fmt.Errorf("WaitStop error, timeout after %s waiting for %d client(s) to finish", timeout, len(s.clientsGroup))
		}
	}
}

type serverHandler struct {
	handler          http.Handler
	clientsGroup     chan bool
	serverInstanceID string
}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.clientsGroup <- true
	defer func() {
		<-sh.clientsGroup
	}()

	w.Header().Add("X-Server-Instance-Id", sh.serverInstanceID)

	sh.handler.ServeHTTP(w, r)
}

func getCerts(trustStore string) (*x509.CertPool, error) {
	certPool := x509.NewCertPool()
	fileInfo, err := os.Stat(trustStore)
	if err != nil {
		return certPool, fmt.Errorf("Truststore [%s] does not exist", trustStore)
	}
	switch mode := fileInfo.Mode(); {
	case mode.IsDir():
		break
	case mode.IsRegular():
		return certPool, fmt.Errorf("Truststore [%s] is not a directory.  Must be a directory containing trusted certificates in PEM format",
			trustStore)
	}
	trustedCertFiles, err := ioutil.ReadDir(trustStore)
	if err != nil || len(trustedCertFiles) == 0 {
		return certPool, fmt.Errorf("Failed to read trusted certificates from [%s]  Must be a directory containing trusted certificates in PEM format", trustStore)
	}
	for _, trustCertFile := range trustedCertFiles {
		fqfName := fmt.Sprintf("%s%c%s", trustStore, os.PathSeparator, trustCertFile.Name())
		trustCertBytes, err := ioutil.ReadFile(fqfName)
		if err != nil {
			log.Warnf("Failed to read trusted certificate [%s] ... continueing", trustCertFile.Name())
		}
		log.Debugf("Loading cert file - %v", fqfName)
		certPool.AppendCertsFromPEM(trustCertBytes)
	}
	if len(certPool.Subjects()) < 1 {
		return certPool, fmt.Errorf("Failed to read trusted certificates from [%s]  After processing all files in the directory no valid trusted certs were found", trustStore)
	}
	return certPool, nil
}
