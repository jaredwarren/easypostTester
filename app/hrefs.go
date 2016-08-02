//************************************************************************//
// API "cellar": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/easypostTester/design
// --out=$(GOPATH)\src\github.com\jaredwarren\easypostTester
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// CarrierAccountHref returns the resource href.
func CarrierAccountHref(id interface{}) string {
	return fmt.Sprintf("/v2/carrier_accounts/%v", id)
}

// CarrierTypesHref returns the resource href.
func CarrierTypesHref() string {
	return fmt.Sprintf("/v2/carrier_types")
}
