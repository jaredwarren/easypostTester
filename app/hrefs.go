//************************************************************************//
// API "easypost": Application Resource Href Factories
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

// AddressHref returns the resource href.
func AddressHref(id interface{}) string {
	return fmt.Sprintf("/v2/addresses/%v", id)
}

// CarrierAccountHref returns the resource href.
func CarrierAccountHref(id interface{}) string {
	return fmt.Sprintf("/v2/carrier_accounts/%v", id)
}

// CarrierTypesHref returns the resource href.
func CarrierTypesHref() string {
	return fmt.Sprintf("/v2/carrier_types")
}

// InsuranceHref returns the resource href.
func InsuranceHref(id interface{}) string {
	return fmt.Sprintf("/v2/insurances/%v", id)
}

// ParcelHref returns the resource href.
func ParcelHref(id interface{}) string {
	return fmt.Sprintf("/v2/parcels/%v", id)
}

// ShipmentHref returns the resource href.
func ShipmentHref(id interface{}) string {
	return fmt.Sprintf("/v2/shipments/%v", id)
}

// TrackerHref returns the resource href.
func TrackerHref(id interface{}) string {
	return fmt.Sprintf("/v2/trackers/%v", id)
}

// UserHref returns the resource href.
func UserHref(id interface{}) string {
	return fmt.Sprintf("/v2/users/%v", id)
}
