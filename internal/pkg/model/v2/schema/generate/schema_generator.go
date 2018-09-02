/*
Sniperkit-Bot
- Status: analyzed
*/

// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mashling-support/jsonschema"

	"github.com/sniperkit/snk.fork.tibcosoftware-mashling/internal/pkg/model/v2/types"
)

func main() {
	schema := jsonschema.Reflect(&types.Schema{})
	schemaJSON, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("schema.json", schemaJSON, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
