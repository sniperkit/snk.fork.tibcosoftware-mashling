/*
Sniperkit-Bot
- Status: analyzed
*/

/*
* Copyright © 2017. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package condition

// OperatorInfo is the information for an operation
type OperatorInfo struct {
	// Name(s) of the operator
	Names []string
	// Description of the operator
	Description string
}

// HasOperatorInfo is an interface for an object that
// has Operator Information
type HasOperatorInfo interface {
	OperatorInfo() *OperatorInfo
}
