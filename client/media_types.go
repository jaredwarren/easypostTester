//************************************************************************//
// API "easypost": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/easypostTester/design
// --out=$(GOPATH)\src\github.com\jaredwarren\easypostTester
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// EasypostAddress media type (default view)
//
// Identifier: application/easypost.address+json
type EasypostAddress struct {
	// The specific designation for the address (only relevant if the address is a carrier facility)
	CarrierFacility *string `form:"carrier_facility,omitempty" json:"carrier_facility,omitempty" xml:"carrier_facility,omitempty"`
	// City the address is located in
	City *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// Name of the organization. Both name and company can be included
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// ISO 3166 country code for the country the address is located in
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Email to reach the person or organization
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Federal tax identifier of the person or organization
	FederalTaxID *string `form:"federal_tax_id,omitempty" json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// Unique, begins with "adr_"
	ID string `form:"id" json:"id" xml:"id"`
	// Set based on which api-key you used, either "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Name of the person. Both name and company can be included
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Always: "Address"
	Object string `form:"object" json:"object" xml:"object"`
	// Phone number to reach the person or organization
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Whether or not this address would be considered residential
	Residential *bool `form:"residential,omitempty" json:"residential,omitempty" xml:"residential,omitempty"`
	// State or province the address is located in
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// 	State tax identifier of the person or organization
	StateTaxID *string `form:"state_tax_id,omitempty" json:"state_tax_id,omitempty" xml:"state_tax_id,omitempty"`
	// First line of the address
	Street1 *string `form:"street1,omitempty" json:"street1,omitempty" xml:"street1,omitempty"`
	// Second line of the address
	Street2 *string `form:"street2,omitempty" json:"street2,omitempty" xml:"street2,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The result of any verifications requested
	Verifications *EasypostAddressVerifications `form:"verifications,omitempty" json:"verifications,omitempty" xml:"verifications,omitempty"`
	// ZIP or postal code the address is located in
	Zip *string `form:"zip,omitempty" json:"zip,omitempty" xml:"zip,omitempty"`
}

// Validate validates the EasypostAddress media type instance.
func (mt *EasypostAddress) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^adr_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^adr_`))
	}
	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^Address$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Address$`))
	}
	return
}

// DecodeEasypostAddress decodes the EasypostAddress instance encoded in resp body.
func (c *Client) DecodeEasypostAddress(resp *http.Response) (*EasypostAddress, error) {
	var decoded EasypostAddress
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostAddress_verification media type (default view)
//
// Identifier: application/easypost.address_verification+json
type EasypostAddressVerification struct {
	// All errors that caused the verification to fail
	Errors []*ApplicationEasypostGieldErrorJSON `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
	// The success of the verification
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// DecodeEasypostAddressVerification decodes the EasypostAddressVerification instance encoded in resp body.
func (c *Client) DecodeEasypostAddressVerification(resp *http.Response) (*EasypostAddressVerification, error) {
	var decoded EasypostAddressVerification
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostAddress_verifications media type (default view)
//
// Identifier: application/easypost.address_verifications+json
type EasypostAddressVerifications struct {
	// Checks that the address is deliverable and makes minor corrections to spelling/format. US addresses will also have their "residential" status checked and set.
	Delivery *EasypostAddressVerification `form:"delivery,omitempty" json:"delivery,omitempty" xml:"delivery,omitempty"`
	// Only applicable to US addresses - checks and sets the ZIP+4.
	Zip4 *EasypostAddressVerification `form:"zip4,omitempty" json:"zip4,omitempty" xml:"zip4,omitempty"`
}

// DecodeEasypostAddressVerifications decodes the EasypostAddressVerifications instance encoded in resp body.
func (c *Client) DecodeEasypostAddressVerifications(resp *http.Response) (*EasypostAddressVerifications, error) {
	var decoded EasypostAddressVerifications
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostApi_key media type (default view)
//
// Identifier: application/easypost.api_key+json
type EasypostAPIKey struct {
	// The User id of the authenticated User making the API Key request
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// The actual key value to use for authentication
	Key string `form:"key" json:"key" xml:"key"`
	// "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Always: "ApiKey"
	Object string `form:"object" json:"object" xml:"object"`
}

// Validate validates the EasypostAPIKey media type instance.
func (mt *EasypostAPIKey) Validate() (err error) {
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}
	if mt.Mode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mode"))
	}
	if mt.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "key"))
	}
	if mt.CreatedAt == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "created_at"))
	}

	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^ApiKey$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^ApiKey$`))
	}
	return
}

// DecodeEasypostAPIKey decodes the EasypostAPIKey instance encoded in resp body.
func (c *Client) DecodeEasypostAPIKey(resp *http.Response) (*EasypostAPIKey, error) {
	var decoded EasypostAPIKey
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostApi_key_child media type (default view)
//
// Identifier: application/easypost.api_key_child+json
type EasypostAPIKeyChild struct {
	// A list of all Child Users presented with ONLY id, children, and key array structures.
	Children EasypostAPIKeyChildCollection `form:"children" json:"children" xml:"children"`
	// The User id of the authenticated User making the API Key request
	ID string `form:"id" json:"id" xml:"id"`
	// The list of all API keys active for an account, both for "test" and "production" modes.
	Keys []*EasypostAPIKey `form:"keys" json:"keys" xml:"keys"`
}

// Validate validates the EasypostAPIKeyChild media type instance.
func (mt *EasypostAPIKeyChild) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Children == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "children"))
	}
	if mt.Keys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "keys"))
	}

	if err2 := mt.Children.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if ok := goa.ValidatePattern(`^user_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^user_`))
	}
	for _, e := range mt.Keys {
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "object"))
		}
		if e.Mode == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "mode"))
		}
		if e.Key == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "key"))
		}
		if e.CreatedAt == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "created_at"))
		}

		if !(e.Mode == "test" || e.Mode == "production") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.keys[*].mode`, e.Mode, []interface{}{"test", "production"}))
		}
		if ok := goa.ValidatePattern(`^ApiKey$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.keys[*].object`, e.Object, `^ApiKey$`))
		}
	}
	return
}

// DecodeEasypostAPIKeyChild decodes the EasypostAPIKeyChild instance encoded in resp body.
func (c *Client) DecodeEasypostAPIKeyChild(resp *http.Response) (*EasypostAPIKeyChild, error) {
	var decoded EasypostAPIKeyChild
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostApi_key_childCollection is the media type for an array of EasypostApi_key_child (default view)
//
// Identifier: application/easypost.api_key_child+json; type=collection
type EasypostAPIKeyChildCollection []*EasypostAPIKeyChild

// Validate validates the EasypostAPIKeyChildCollection media type instance.
func (mt EasypostAPIKeyChildCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Children == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "children"))
		}
		if e.Keys == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "keys"))
		}

		if err2 := e.Children.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if ok := goa.ValidatePattern(`^user_`, e.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response[*].id`, e.ID, `^user_`))
		}
		for _, e := range e.Keys {
			if e.Object == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].keys[*]`, "object"))
			}
			if e.Mode == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].keys[*]`, "mode"))
			}
			if e.Key == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].keys[*]`, "key"))
			}
			if e.CreatedAt == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].keys[*]`, "created_at"))
			}

			if !(e.Mode == "test" || e.Mode == "production") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response[*].keys[*].mode`, e.Mode, []interface{}{"test", "production"}))
			}
			if ok := goa.ValidatePattern(`^ApiKey$`, e.Object); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response[*].keys[*].object`, e.Object, `^ApiKey$`))
			}
		}
	}
	return
}

// DecodeEasypostAPIKeyChildCollection decodes the EasypostAPIKeyChildCollection instance encoded in resp body.
func (c *Client) DecodeEasypostAPIKeyChildCollection(resp *http.Response) (EasypostAPIKeyChildCollection, error) {
	var decoded EasypostAPIKeyChildCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// EasypostApi_keys media type (default view)
//
// Identifier: application/easypost.api_keys+json
type EasypostAPIKeys struct {
	// A list of all Child Users presented with ONLY id, children, and key array structures.
	Children []*EasypostAPIKeyChild `form:"children" json:"children" xml:"children"`
	// The User id of the authenticated User making the API Key request
	ID string `form:"id" json:"id" xml:"id"`
	// The list of all API keys active for an account, both for "test" and "production" modes.
	Keys []*EasypostAPIKey `form:"keys" json:"keys" xml:"keys"`
}

// Validate validates the EasypostAPIKeys media type instance.
func (mt *EasypostAPIKeys) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Children == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "children"))
	}
	if mt.Keys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "keys"))
	}

	for _, e := range mt.Children {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*]`, "id"))
		}
		if e.Children == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*]`, "children"))
		}
		if e.Keys == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*]`, "keys"))
		}

		if err2 := e.Children.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if ok := goa.ValidatePattern(`^user_`, e.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.children[*].id`, e.ID, `^user_`))
		}
		for _, e := range e.Keys {
			if e.Object == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*].keys[*]`, "object"))
			}
			if e.Mode == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*].keys[*]`, "mode"))
			}
			if e.Key == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*].keys[*]`, "key"))
			}
			if e.CreatedAt == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response.children[*].keys[*]`, "created_at"))
			}

			if !(e.Mode == "test" || e.Mode == "production") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.children[*].keys[*].mode`, e.Mode, []interface{}{"test", "production"}))
			}
			if ok := goa.ValidatePattern(`^ApiKey$`, e.Object); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response.children[*].keys[*].object`, e.Object, `^ApiKey$`))
			}
		}
	}
	if ok := goa.ValidatePattern(`^user_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^user_`))
	}
	for _, e := range mt.Keys {
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "object"))
		}
		if e.Mode == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "mode"))
		}
		if e.Key == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "key"))
		}
		if e.CreatedAt == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.keys[*]`, "created_at"))
		}

		if !(e.Mode == "test" || e.Mode == "production") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.keys[*].mode`, e.Mode, []interface{}{"test", "production"}))
		}
		if ok := goa.ValidatePattern(`^ApiKey$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.keys[*].object`, e.Object, `^ApiKey$`))
		}
	}
	return
}

// DecodeEasypostAPIKeys decodes the EasypostAPIKeys instance encoded in resp body.
func (c *Client) DecodeEasypostAPIKeys(resp *http.Response) (*EasypostAPIKeys, error) {
	var decoded EasypostAPIKeys
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts. (default view)
//
// Identifier: application/easypost.carrier_accounts+json
type EasypostCarrierAccounts struct {
	// If clone is true, only the reference and description are possible to update
	Clone bool `form:"clone" json:"clone" xml:"clone"`
	// The name used when displaying a readable value for the type of the account
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *FieldsObjectPayload `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
	// Unique, begins with "ca_"
	ID string `form:"id" json:"id" xml:"id"`
	// Always: "CarrierAccount"
	Object string `form:"object" json:"object" xml:"object"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// An optional field that may be used in place of carrier_account_id in other API endpoints
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The name used when displaying a readable value for the type of the account
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the EasypostCarrierAccounts media type instance.
func (mt *EasypostCarrierAccounts) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^CarrierAccount$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^CarrierAccount$`))
	}
	return
}

// DecodeEasypostCarrierAccounts decodes the EasypostCarrierAccounts instance encoded in resp body.
func (c *Client) DecodeEasypostCarrierAccounts(resp *http.Response) (*EasypostCarrierAccounts, error) {
	var decoded EasypostCarrierAccounts
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The CarrierType object provides an interface for determining the valid fields of a CarrierAccount. The list of CarrierType objects only changes when a new carrier is added to EasyPost. (default view)
//
// Identifier: application/easypost.carrier_types+json
type EasypostCarrierTypes struct {
	// Contains at least one of the following keys: "auto_link", "credentials", "test_credentials", and "custom_workflow"
	Fields *EasypostFieldsObject `form:"fields" json:"fields" xml:"fields"`
	// Always: "CarrierType"
	Object string `form:"object" json:"object" xml:"object"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the EasypostCarrierTypes media type instance.
func (mt *EasypostCarrierTypes) Validate() (err error) {
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}
	if mt.Fields == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fields"))
	}

	if ok := goa.ValidatePattern(`^CarrierType$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^CarrierType$`))
	}
	return
}

// DecodeEasypostCarrierTypes decodes the EasypostCarrierTypes instance encoded in resp body.
func (c *Client) DecodeEasypostCarrierTypes(resp *http.Response) (*EasypostCarrierTypes, error) {
	var decoded EasypostCarrierTypes
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostCarrier_typesCollection is the media type for an array of EasypostCarrier_types (default view)
//
// Identifier: application/easypost.carrier_types+json; type=collection
type EasypostCarrierTypesCollection []*EasypostCarrierTypes

// Validate validates the EasypostCarrierTypesCollection media type instance.
func (mt EasypostCarrierTypesCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Type == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "type"))
		}
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "object"))
		}
		if e.Fields == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "fields"))
		}

		if ok := goa.ValidatePattern(`^CarrierType$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response[*].object`, e.Object, `^CarrierType$`))
		}
	}
	return
}

// DecodeEasypostCarrierTypesCollection decodes the EasypostCarrierTypesCollection instance encoded in resp body.
func (c *Client) DecodeEasypostCarrierTypesCollection(resp *http.Response) (EasypostCarrierTypesCollection, error) {
	var decoded EasypostCarrierTypesCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// The CarrierDetail object contains the service and container_type of the package. Additionally, it also stores the est_delivery_date_local and est_delivery_time_local, which provide information about the local delivery time. (default view)
//
// Identifier: application/easypost.carrierdetail+json
type EasypostCarrierdetail struct {
	// The type of container the associated shipment was shipped in (if available)
	ContainerType *string `form:"container_type,omitempty" json:"container_type,omitempty" xml:"container_type,omitempty"`
	// The estimated delivery date as provided by the carrier, in the local time zone (if available)
	EstDeliveryDateLocal *string `form:"est_delivery_date_local,omitempty" json:"est_delivery_date_local,omitempty" xml:"est_delivery_date_local,omitempty"`
	// The estimated delivery time as provided by the carrier, in the local time zone (if available)
	EstDeliveryTimeLocal *string `form:"est_delivery_time_local,omitempty" json:"est_delivery_time_local,omitempty" xml:"est_delivery_time_local,omitempty"`
	// Always: "CarrierDetail"
	Object string `form:"object" json:"object" xml:"object"`
	// The service level the associated shipment was shipped with (if available)
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
}

// Validate validates the EasypostCarrierdetail media type instance.
func (mt *EasypostCarrierdetail) Validate() (err error) {
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^CarrierDetail$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^CarrierDetail$`))
	}
	return
}

// DecodeEasypostCarrierdetail decodes the EasypostCarrierdetail instance encoded in resp body.
func (c *Client) DecodeEasypostCarrierdetail(resp *http.Response) (*EasypostCarrierdetail, error) {
	var decoded EasypostCarrierdetail
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A CustomsItem object describes goods for international shipment and should be created then included in a CustomsInfo object. (default view)
//
// Identifier: application/easypost.customitem+json
type EasypostCustomitem struct {
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 3 char currency code, default USD
	Currency string `form:"currency" json:"currency" xml:"currency"`
	// Required, description of item being shipped
	Description string `form:"description" json:"description" xml:"description"`
	// Harmonized Tariff Schedule, e.g. "6109.10.0012" for Men's T-shirts
	HsTariffNumber *string `form:"hs_tariff_number,omitempty" json:"hs_tariff_number,omitempty" xml:"hs_tariff_number,omitempty"`
	// Unique, begins with "cstitem_"
	ID string `form:"id" json:"id" xml:"id"`
	// Always: "CustomsItem"
	Object string `form:"object" json:"object" xml:"object"`
	// Required, 2 char country code
	OriginCountry string `form:"origin_country" json:"origin_country" xml:"origin_country"`
	// Required, greater than zero
	Quantity string `form:"quantity" json:"quantity" xml:"quantity"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Required, greater than zero, total value (unit value * quantity)
	Value string `form:"value" json:"value" xml:"value"`
	// Required, greater than zero, total weight (unit weight * quantity)
	Weight string `form:"weight" json:"weight" xml:"weight"`
}

// Validate validates the EasypostCustomitem media type instance.
func (mt *EasypostCustomitem) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.Quantity == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "quantity"))
	}
	if mt.Value == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	if mt.Weight == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "weight"))
	}
	if mt.OriginCountry == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "origin_country"))
	}

	if ok := goa.ValidatePattern(`^cstitem_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^cstitem_`))
	}
	if ok := goa.ValidatePattern(`^CustomsItem$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^CustomsItem$`))
	}
	return
}

// DecodeEasypostCustomitem decodes the EasypostCustomitem instance encoded in resp body.
func (c *Client) DecodeEasypostCustomitem(resp *http.Response) (*EasypostCustomitem, error) {
	var decoded EasypostCustomitem
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CustomsInfo objects contain CustomsItem objects and all necessary information for the generation of customs forms required for international shipping. (default view)
//
// Identifier: application/easypost.customsinfo+json
type EasypostCustomsinfo struct {
	// Human readable description of content. Required for certain carriers and always required if contents_type is "other"
	ContentsExplanation *string `form:"contents_explanation,omitempty" json:"contents_explanation,omitempty" xml:"contents_explanation,omitempty"`
	// "documents", "gift", "merchandise", "returned_goods", "sample", or "other"
	ContentsType *string `form:"contents_type,omitempty" json:"contents_type,omitempty" xml:"contents_type,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Electronically certify the information provided
	CustomsCertify *bool `form:"customs_certify,omitempty" json:"customs_certify,omitempty" xml:"customs_certify,omitempty"`
	// Describes to products being shipped
	CustomsItems []*EasypostCustomitem `form:"customs_items,omitempty" json:"customs_items,omitempty" xml:"customs_items,omitempty"`
	// Required if customs_certify is true
	CustomsSigner *string `form:"customs_signer,omitempty" json:"customs_signer,omitempty" xml:"customs_signer,omitempty"`
	// "EEL" or "PFC" value less than $2500: "NOEEI 30.37(a)"; value greater than $2500: see Customs Guide
	EelPfc *string `form:"eel_pfc,omitempty" json:"eel_pfc,omitempty" xml:"eel_pfc,omitempty"`
	// Unique, begins with "cstinfo_"
	ID string `form:"id" json:"id" xml:"id"`
	// "abandon" or "return", defaults to "return"
	NonDeliveryOption string `form:"non_delivery_option" json:"non_delivery_option" xml:"non_delivery_option"`
	// Always: "CustomsInfo"
	Object string `form:"object" json:"object" xml:"object"`
	// Required if restriction_type is not "none"
	RestrictionComments *string `form:"restriction_comments,omitempty" json:"restriction_comments,omitempty" xml:"restriction_comments,omitempty"`
	// "none", "other", "quarantine", or "sanitary_phytosanitary_inspection"
	RestrictionType *string `form:"restriction_type,omitempty" json:"restriction_type,omitempty" xml:"restriction_type,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the EasypostCustomsinfo media type instance.
func (mt *EasypostCustomsinfo) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if mt.ContentsType != nil {
		if !(*mt.ContentsType == "documents" || *mt.ContentsType == "gift" || *mt.ContentsType == "merchandise" || *mt.ContentsType == "returned_goods" || *mt.ContentsType == "sample" || *mt.ContentsType == "other") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.contents_type`, *mt.ContentsType, []interface{}{"documents", "gift", "merchandise", "returned_goods", "sample", "other"}))
		}
	}
	for _, e := range mt.CustomsItems {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "id"))
		}
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "object"))
		}
		if e.Description == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "description"))
		}
		if e.Quantity == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "quantity"))
		}
		if e.Value == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "value"))
		}
		if e.Weight == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "weight"))
		}
		if e.OriginCountry == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "origin_country"))
		}

		if ok := goa.ValidatePattern(`^cstitem_`, e.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.customs_items[*].id`, e.ID, `^cstitem_`))
		}
		if ok := goa.ValidatePattern(`^CustomsItem$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.customs_items[*].object`, e.Object, `^CustomsItem$`))
		}
	}
	if mt.EelPfc != nil {
		if !(*mt.EelPfc == "EEL" || *mt.EelPfc == "PFC") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.eel_pfc`, *mt.EelPfc, []interface{}{"EEL", "PFC"}))
		}
	}
	if ok := goa.ValidatePattern(`^cstinfo_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^cstinfo_`))
	}
	if !(mt.NonDeliveryOption == "abandon" || mt.NonDeliveryOption == "return") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.non_delivery_option`, mt.NonDeliveryOption, []interface{}{"abandon", "return"}))
	}
	if ok := goa.ValidatePattern(`^CustomsInfo$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^CustomsInfo$`))
	}
	if mt.RestrictionType != nil {
		if !(*mt.RestrictionType == "none" || *mt.RestrictionType == "other" || *mt.RestrictionType == "quarantine" || *mt.RestrictionType == "sanitary_phytosanitary_inspection") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.restriction_type`, *mt.RestrictionType, []interface{}{"none", "other", "quarantine", "sanitary_phytosanitary_inspection"}))
		}
	}
	return
}

// DecodeEasypostCustomsinfo decodes the EasypostCustomsinfo instance encoded in resp body.
func (c *Client) DecodeEasypostCustomsinfo(resp *http.Response) (*EasypostCustomsinfo, error) {
	var decoded EasypostCustomsinfo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostFee media type (default view)
//
// Identifier: application/easypost.fee+json
type EasypostFee struct {
	// USD value with sub-cent precision
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Whether EasyPost has successfully charged your account for the fee
	Charged *bool `form:"charged,omitempty" json:"charged,omitempty" xml:"charged,omitempty"`
	// Always: "Fee"
	Object string `form:"object" json:"object" xml:"object"`
	// Whether the Fee has been refunded successfully
	Refunded *bool `form:"refunded,omitempty" json:"refunded,omitempty" xml:"refunded,omitempty"`
	// The name of the category of fee. Possible types are "LabelFee", "PostageFee", "InsuranceFee", and "TrackerFee"
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the EasypostFee media type instance.
func (mt *EasypostFee) Validate() (err error) {
	if ok := goa.ValidatePattern(`^Fee$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Fee$`))
	}
	if mt.Type != nil {
		if !(*mt.Type == "LabelFee" || *mt.Type == "PostageFee" || *mt.Type == "InsuranceFee" || *mt.Type == "TrackerFee") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.type`, *mt.Type, []interface{}{"LabelFee", "PostageFee", "InsuranceFee", "TrackerFee"}))
		}
	}
	return
}

// DecodeEasypostFee decodes the EasypostFee instance encoded in resp body.
func (c *Client) DecodeEasypostFee(resp *http.Response) (*EasypostFee, error) {
	var decoded EasypostFee
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Contains "credentials" and/or "test_credentials", or may be empty (default view)
//
// Identifier: application/easypost.field_object+json
type EasypostFieldObject struct {
	Key        string `form:"key" json:"key" xml:"key"`
	Label      string `form:"label" json:"label" xml:"label"`
	Value      string `form:"value" json:"value" xml:"value"`
	Visibility string `form:"visibility" json:"visibility" xml:"visibility"`
}

// Validate validates the EasypostFieldObject media type instance.
func (mt *EasypostFieldObject) Validate() (err error) {
	if mt.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "key"))
	}
	if mt.Value == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	if mt.Visibility == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "visibility"))
	}
	if mt.Label == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "label"))
	}

	return
}

// DecodeEasypostFieldObject decodes the EasypostFieldObject instance encoded in resp body.
func (c *Client) DecodeEasypostFieldObject(resp *http.Response) (*EasypostFieldObject, error) {
	var decoded EasypostFieldObject
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Contains "credentials" and/or "test_credentials", or may be empty (default view)
//
// Identifier: application/easypost.fields_object+json
type EasypostFieldsObject struct {
	// For USPS this designates that no credentials are required.
	AutoLink *bool `form:"auto_link,omitempty" json:"auto_link,omitempty" xml:"auto_link,omitempty"`
	// Credentials used in the production environment.
	Credentials interface{} `form:"credentials" json:"credentials" xml:"credentials"`
	// When present, a seperate authentication process will be required through the UI to link this account type.
	CustomWorkflow *bool `form:"custom_workflow,omitempty" json:"custom_workflow,omitempty" xml:"custom_workflow,omitempty"`
	// Credentials used in the test environment.
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
}

// DecodeEasypostFieldsObject decodes the EasypostFieldsObject instance encoded in resp body.
func (c *Client) DecodeEasypostFieldsObject(resp *http.Response) (*EasypostFieldsObject, error) {
	var decoded EasypostFieldsObject
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// An Insurance object represents insurance for packages purchased both va the EasyPost API as well as shipments purchased through third parties and later registered with EasyPost. An Insurance is created automatically whenever you buy a Shipment through EasyPost and pass insurance options during the Buy call or in a later call to Insure a Shipment. (default view)
//
// Identifier: application/easypost.insurance+json
type EasypostInsurance struct {
	// USD value of insured goods with sub-cent precision
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The associated InsuranceFee object if any
	Fee *EasypostFee `form:"fee,omitempty" json:"fee,omitempty" xml:"fee,omitempty"`
	// The associated Address object for origin
	FromAddress *EasypostAddress `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// Unique, begins with "ins_"
	ID string `form:"id" json:"id" xml:"id"`
	// The list of errors encountered during attempted purchase of the insurance
	Messages []string `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Set based on which api-key you used, either "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Always: "Insurance"
	Object string `form:"object" json:"object" xml:"object"`
	// The insurance provider used by EasyPost
	Provider *string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// An identifying number for some insurance providers used by EasyPost
	ProviderID *string `form:"provider_id,omitempty" json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// The unique reference for this Insurance, if any
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// The ID of the Shipment in EasyPost, if postage was purchased via EasyPost
	ShipmentID *string `form:"shipment_id,omitempty" json:"shipment_id,omitempty" xml:"shipment_id,omitempty"`
	// The current status of the insurance, possible values are "new", "pending", "purchased", "failed", or "cancelled"
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The associated Address object for destination
	ToAddress *EasypostAddress `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The associated Tracker object
	Tracker *EasypostTracker `form:"tracker,omitempty" json:"tracker,omitempty" xml:"tracker,omitempty"`
	// The tracking code of either the shipment within EasyPost, or provided by you during creation
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the EasypostInsurance media type instance.
func (mt *EasypostInsurance) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if mt.Fee != nil {
		if err2 := mt.Fee.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.FromAddress != nil {
		if err2 := mt.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`^ins_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^ins_`))
	}
	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^Insurance$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Insurance$`))
	}
	if mt.Status != nil {
		if !(*mt.Status == "cancelled" || *mt.Status == "failed" || *mt.Status == "purchased" || *mt.Status == "pending" || *mt.Status == "new") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"cancelled", "failed", "purchased", "pending", "new"}))
		}
	}
	if mt.ToAddress != nil {
		if err2 := mt.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Tracker != nil {
		if err2 := mt.Tracker.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeEasypostInsurance decodes the EasypostInsurance instance encoded in resp body.
func (c *Client) DecodeEasypostInsurance(resp *http.Response) (*EasypostInsurance, error) {
	var decoded EasypostInsurance
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins. (default view)
//
// Identifier: application/easypost.insurances+json
type EasypostInsurances struct {
	Insurances []*EasypostInsurance `form:"insurances" json:"insurances" xml:"insurances"`
}

// Validate validates the EasypostInsurances media type instance.
func (mt *EasypostInsurances) Validate() (err error) {
	if mt.Insurances == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "insurances"))
	}

	for _, e := range mt.Insurances {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.insurances[*]`, "id"))
		}
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.insurances[*]`, "object"))
		}

		if e.Fee != nil {
			if err2 := e.Fee.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if e.FromAddress != nil {
			if err2 := e.FromAddress.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if ok := goa.ValidatePattern(`^ins_`, e.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.insurances[*].id`, e.ID, `^ins_`))
		}
		if !(e.Mode == "test" || e.Mode == "production") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.insurances[*].mode`, e.Mode, []interface{}{"test", "production"}))
		}
		if ok := goa.ValidatePattern(`^Insurance$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.insurances[*].object`, e.Object, `^Insurance$`))
		}
		if e.Status != nil {
			if !(*e.Status == "cancelled" || *e.Status == "failed" || *e.Status == "purchased" || *e.Status == "pending" || *e.Status == "new") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.insurances[*].status`, *e.Status, []interface{}{"cancelled", "failed", "purchased", "pending", "new"}))
			}
		}
		if e.ToAddress != nil {
			if err2 := e.ToAddress.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if e.Tracker != nil {
			if err2 := e.Tracker.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeEasypostInsurances decodes the EasypostInsurances instance encoded in resp body.
func (c *Client) DecodeEasypostInsurances(resp *http.Response) (*EasypostInsurances, error) {
	var decoded EasypostInsurances
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostOptions media type (default view)
//
// Identifier: application/easypost.options+json
type EasypostOptions struct {
	// Setting this option to true, will add an additional handling charge.
	AdditionalHandling *bool `form:"additional_handling,omitempty" json:"additional_handling,omitempty" xml:"additional_handling,omitempty"`
	// Setting this option to "0", will allow the minimum amount of address information to pass the validation check. Only for USPS postage.
	AddressValidationLevel *string `form:"address_validation_level,omitempty" json:"address_validation_level,omitempty" xml:"address_validation_level,omitempty"`
	// Set this option to true if your shipment contains alcohol. UPS - only supported for US Domestic shipments. FedEx - only supported for US Domestic shipments. Canada Post - Requires adult signature 19 years or older. If you want adult signature 18 years or older, instead use delivery_confirmation: ADULT_SIGNATURE
	Alcohol *bool `form:"alcohol,omitempty" json:"alcohol,omitempty" xml:"alcohol,omitempty"`
	// Setting an account number of the receiver who is to receive and buy the postage. UPS - bill_receiver_postal_code is also required
	BillReceiverAccount *string `form:"bill_receiver_account,omitempty" json:"bill_receiver_account,omitempty" xml:"bill_receiver_account,omitempty"`
	// Setting a postal code of the receiver account you want to buy postage. UPS - bill_receiver_account also required
	BillReceiverPostalCode *string `form:"bill_receiver_postal_code,omitempty" json:"bill_receiver_postal_code,omitempty" xml:"bill_receiver_postal_code,omitempty"`
	// Setting an account number of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_postal_code also required
	BillThirdPartyAccount *string `form:"bill_third_party_account,omitempty" json:"bill_third_party_account,omitempty" xml:"bill_third_party_account,omitempty"`
	// Setting a country of the third party account you want to buy postage. UPS - bill_third_party_account and bill_third_party_postal_code also required
	BillThirdPartyCountry *string `form:"bill_third_party_country,omitempty" json:"bill_third_party_country,omitempty" xml:"bill_third_party_country,omitempty"`
	// Setting a postal code of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_account also required
	BillThirdPartyPostalCode *bool `form:"bill_third_party_postal_code,omitempty" json:"bill_third_party_postal_code,omitempty" xml:"bill_third_party_postal_code,omitempty"`
	// Setting this option to true will indicate to the carrier to prefer delivery by drone, if the carrier supports drone delivery.
	ByDrone *bool `form:"by_drone,omitempty" json:"by_drone,omitempty" xml:"by_drone,omitempty"`
	// Setting this to true will add a charge to reduce carbon emissions.
	CarbonNeutral *bool `form:"carbon_neutral,omitempty" json:"carbon_neutral,omitempty" xml:"carbon_neutral,omitempty"`
	// Adding an amount will have the carrier collect the specified amount from the recipient.
	CodAmount *string `form:"cod_amount,omitempty" json:"cod_amount,omitempty" xml:"cod_amount,omitempty"`
	// Method for payment. "CASH", "CHECK", "MONEY_ORDER"
	CodMethod *string `form:"cod_method,omitempty" json:"cod_method,omitempty" xml:"cod_method,omitempty"`
	// Which currency this shipment will show for rates if carrier allows.
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Setting this value to false will pass the cost of duties on to the recipient of the package(s).
	DeliveredDutyPaid *bool `form:"delivered_duty_paid,omitempty" json:"delivered_duty_paid,omitempty" xml:"delivered_duty_paid,omitempty"`
	// If you want to request a signature, you can pass "ADULT_SIGNATURE" or "SIGNATURE". You may also request "NO_SIGNATURE" to leave the package at the door. All - some options may be limited for international shipments
	DeliveryConfirmation *string `form:"delivery_confirmation,omitempty" json:"delivery_confirmation,omitempty" xml:"delivery_confirmation,omitempty"`
	// Package contents contain dry ice. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set. FedEx - Need dry_ice_weight to be set
	DryIce *bool `form:"dry_ice,omitempty" json:"dry_ice,omitempty" xml:"dry_ice,omitempty"`
	// If the dry ice is for medical use, set this option to true. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set
	DryIceMedical *bool `form:"dry_ice_medical,omitempty" json:"dry_ice_medical,omitempty" xml:"dry_ice_medical,omitempty"`
	// Weight of the dry ice in ounces. UPS - Need dry_ice to be set. UPS MailInnovations - Need dry_ice to be set. FedEx - Need dry_ice to be set
	DryIceWeight *string `form:"dry_ice_weight,omitempty" json:"dry_ice_weight,omitempty" xml:"dry_ice_weight,omitempty"`
	// Additional cost to be added to the invoice of this shipment. Only applies to UPS currently.
	FreightCharge *float64 `form:"freight_charge,omitempty" json:"freight_charge,omitempty" xml:"freight_charge,omitempty"`
	// This is to designate special instructions for the carrier like "Do not drop!".
	HandlingInstructions *string `form:"handling_instructions,omitempty" json:"handling_instructions,omitempty" xml:"handling_instructions,omitempty"`
	// Package will wait at carrier facility for pickup.
	HoldForPickup *bool `form:"hold_for_pickup,omitempty" json:"hold_for_pickup,omitempty" xml:"hold_for_pickup,omitempty"`
	// This will print an invoice number on the postage label.
	InvoiceNumber *string `form:"invoice_number,omitempty" json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// Set the date that will appear on the postage label. Accepts ISO 8601 formatted string including time zone offset.
	LabelDate *string `form:"label_date,omitempty" json:"label_date,omitempty" xml:"label_date,omitempty"`
	// Supported label formats include "PNG", "PDF", "ZPL", and "EPL2". "PNG" is the only format that allows for conversion.
	LabelFormat *string `form:"label_format,omitempty" json:"label_format,omitempty" xml:"label_format,omitempty"`
	// Whether or not the parcel can be processed by the carriers equipment.
	Machinable *bool `form:"machinable,omitempty" json:"machinable,omitempty" xml:"machinable,omitempty"`
	// You can optionally print custom messages on labels. The locations of these fields show up on different spots on the carrier's labels.
	PrintCustom1 *string `form:"print_custom_1,omitempty" json:"print_custom_1,omitempty" xml:"print_custom_1,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom1Barcode *bool `form:"print_custom_1_barcode,omitempty" json:"print_custom_1_barcode,omitempty" xml:"print_custom_1_barcode,omitempty"`
	// Specify the type of print_custom_1.
	PrintCustom1Code *string `form:"print_custom_1_code,omitempty" json:"print_custom_1_code,omitempty" xml:"print_custom_1_code,omitempty"`
	// An additional message on the label. Same restrictions as print_custom_1
	PrintCustom2 *string `form:"print_custom_2,omitempty" json:"print_custom_2,omitempty" xml:"print_custom_2,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom2Barcode *bool `form:"print_custom_2_barcode,omitempty" json:"print_custom_2_barcode,omitempty" xml:"print_custom_2_barcode,omitempty"`
	// See print_custom_1_code.
	PrintCustom2Code *string `form:"print_custom_2_code,omitempty" json:"print_custom_2_code,omitempty" xml:"print_custom_2_code,omitempty"`
	// An additional message on the label. Same restrictions as print_custom_1
	PrintCustom3 *string `form:"print_custom_3,omitempty" json:"print_custom_3,omitempty" xml:"print_custom_3,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom3Barcode *bool `form:"print_custom_3_barcode,omitempty" json:"print_custom_3_barcode,omitempty" xml:"print_custom_3_barcode,omitempty"`
	// See print_custom_1_code.
	PrintCustom3Code *string `form:"print_custom_3_code,omitempty" json:"print_custom_3_code,omitempty" xml:"print_custom_3_code,omitempty"`
	// Set this value to true for delivery on Saturday. Can't be combined with Next Day Air. When setting the saturday_delivery option, you will only get rates for services that are eligible for saturday delivery. If no services are available for saturday delivery, then you will not be returned any rates. You may need to create 2 shipments, one with the saturday_delivery option set on one without to get all your eligible rates.
	SaturdayDelivery *bool `form:"saturday_delivery,omitempty" json:"saturday_delivery,omitempty" xml:"saturday_delivery,omitempty"`
	// You can use this to override the hub ID you have on your account.
	SmartpostHub *string `form:"smartpost_hub,omitempty" json:"smartpost_hub,omitempty" xml:"smartpost_hub,omitempty"`
	// The manifest ID is used to group SmartPost packages onto a manifest for each trailer.
	SmartpostManifest *bool `form:"smartpost_manifest,omitempty" json:"smartpost_manifest,omitempty" xml:"smartpost_manifest,omitempty"`
	// This option allows you to request restrictive rates from USPS. Can set to 'USPS.MEDIAMAIL' or 'USPS.LIBRARYMAIL'.
	SpecialRatesEligibility *string `form:"special_rates_eligibility,omitempty" json:"special_rates_eligibility,omitempty" xml:"special_rates_eligibility,omitempty"`
}

// Validate validates the EasypostOptions media type instance.
func (mt *EasypostOptions) Validate() (err error) {
	if mt.CodMethod != nil {
		if !(*mt.CodMethod == "CASH" || *mt.CodMethod == "CHECK" || *mt.CodMethod == "MONEY_ORDER") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.cod_method`, *mt.CodMethod, []interface{}{"CASH", "CHECK", "MONEY_ORDER"}))
		}
	}
	if mt.HandlingInstructions != nil {
		if !(*mt.HandlingInstructions == "ORMD" || *mt.HandlingInstructions == "LIMITED_QUANTITY") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.handling_instructions`, *mt.HandlingInstructions, []interface{}{"ORMD", "LIMITED_QUANTITY"}))
		}
	}
	if mt.LabelFormat != nil {
		if !(*mt.LabelFormat == "PNG" || *mt.LabelFormat == "PDF" || *mt.LabelFormat == "ZPL" || *mt.LabelFormat == "EPL2") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.label_format`, *mt.LabelFormat, []interface{}{"PNG", "PDF", "ZPL", "EPL2"}))
		}
	}
	return
}

// DecodeEasypostOptions decodes the EasypostOptions instance encoded in resp body.
func (c *Client) DecodeEasypostOptions(resp *http.Response) (*EasypostOptions, error) {
	var decoded EasypostOptions
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostParcel media type (default view)
//
// Identifier: application/easypost.parcel+json
type EasypostParcel struct {
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Height *float64 `form:"height,omitempty" json:"height,omitempty" xml:"height,omitempty"`
	// Unique, begins with "prcl_"
	ID string `form:"id" json:"id" xml:"id"`
	// Required if predefined_package is empty. float (inches)
	Length *float64 `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Set based on which api-key you used, either "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Always: "Parcel"
	Object string `form:"object" json:"object" xml:"object"`
	// Optional, one of our predefined_packages
	PredefinedPackage *string `form:"predefined_package,omitempty" json:"predefined_package,omitempty" xml:"predefined_package,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Always required. float(oz)
	Weight float64 `form:"weight" json:"weight" xml:"weight"`
	// Required if predefined_package is empty. float (inches)
	Width *float64 `form:"width,omitempty" json:"width,omitempty" xml:"width,omitempty"`
}

// Validate validates the EasypostParcel media type instance.
func (mt *EasypostParcel) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^prcl_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^prcl_`))
	}
	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^Parcel$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Parcel$`))
	}
	return
}

// DecodeEasypostParcel decodes the EasypostParcel instance encoded in resp body.
func (c *Client) DecodeEasypostParcel(resp *http.Response) (*EasypostParcel, error) {
	var decoded EasypostParcel
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostPostagelabel media type (default view)
//
// Identifier: application/easypost.postagelabel+json
type EasypostPostagelabel struct {
	// Time Created
	CreatedAt   *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	DateAdvance *int    `form:"date_advance,omitempty" json:"date_advance,omitempty" xml:"date_advance,omitempty"`
	// Unique, begins with "pl_"
	ID             *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	IntegratedForm string  `form:"integrated_form" json:"integrated_form" xml:"integrated_form"`
	LabelDate      *string `form:"label_date,omitempty" json:"label_date,omitempty" xml:"label_date,omitempty"`
	LabelEpl2URL   *string `form:"label_epl2_url,omitempty" json:"label_epl2_url,omitempty" xml:"label_epl2_url,omitempty"`
	LabelFileType  string  `form:"label_file_type" json:"label_file_type" xml:"label_file_type"`
	LabelPdfURL    *string `form:"label_pdf_url,omitempty" json:"label_pdf_url,omitempty" xml:"label_pdf_url,omitempty"`
	// Assuming DPI
	LabelResolution int     `form:"label_resolution" json:"label_resolution" xml:"label_resolution"`
	LabelSize       string  `form:"label_size" json:"label_size" xml:"label_size"`
	LabelType       string  `form:"label_type" json:"label_type" xml:"label_type"`
	LabelURL        *string `form:"label_url,omitempty" json:"label_url,omitempty" xml:"label_url,omitempty"`
	LabelZplURL     *string `form:"label_zpl_url,omitempty" json:"label_zpl_url,omitempty" xml:"label_zpl_url,omitempty"`
	// Always: "PostageLabel"
	Object string `form:"object" json:"object" xml:"object"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the EasypostPostagelabel media type instance.
func (mt *EasypostPostagelabel) Validate() (err error) {
	if mt.ID != nil {
		if ok := goa.ValidatePattern(`^pl_`, *mt.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *mt.ID, `^pl_`))
		}
	}
	if ok := goa.ValidatePattern(`^PostageLabel$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^PostageLabel$`))
	}
	return
}

// DecodeEasypostPostagelabel decodes the EasypostPostagelabel instance encoded in resp body.
func (c *Client) DecodeEasypostPostagelabel(resp *http.Response) (*EasypostPostagelabel, error) {
	var decoded EasypostPostagelabel
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostRate media type (default view)
//
// Identifier: application/easypost.rate+json
type EasypostRate struct {
	// name of carrier
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// ID of the CarrierAccount record used to generate this rate
	CarrierAccountID *string `form:"carrier_account_id,omitempty" json:"carrier_account_id,omitempty" xml:"carrier_account_id,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// currency for the rate
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// date for delivery
	DeliveryDate *string `form:"delivery_date,omitempty" json:"delivery_date,omitempty" xml:"delivery_date,omitempty"`
	// indicates if delivery window is guaranteed (true) or not (false)
	DeliveryDateGuaranteed *bool `form:"delivery_date_guaranteed,omitempty" json:"delivery_date_guaranteed,omitempty" xml:"delivery_date_guaranteed,omitempty"`
	// delivery days for this service
	DeliveryDays *int `form:"delivery_days,omitempty" json:"delivery_days,omitempty" xml:"delivery_days,omitempty"`
	// Unique, begins with "rate_"
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// currency for the list rate
	ListCurrency *string `form:"list_currency,omitempty" json:"list_currency,omitempty" xml:"list_currency,omitempty"`
	// the list rate is the non-negotiated rate given for having an account with the carrier
	ListRate *string `form:"list_rate,omitempty" json:"list_rate,omitempty" xml:"list_rate,omitempty"`
	// Set based on which api-key you used, either "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Always: "Rate"
	Object string `form:"object" json:"object" xml:"object"`
	// the actual rate quote for this service
	Rate *string `form:"rate,omitempty" json:"rate,omitempty" xml:"rate,omitempty"`
	// currency for the retail rate
	RetailCurrency *string `form:"retail_currency,omitempty" json:"retail_currency,omitempty" xml:"retail_currency,omitempty"`
	// the retail rate is the in-store rate given with no account
	RetailRate *string `form:"retail_rate,omitempty" json:"retail_rate,omitempty" xml:"retail_rate,omitempty"`
	// service level/name
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// ID of the Shipment this rate belongs to
	ShipmentID *string `form:"shipment_id,omitempty" json:"shipment_id,omitempty" xml:"shipment_id,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the EasypostRate media type instance.
func (mt *EasypostRate) Validate() (err error) {
	if mt.ID != nil {
		if ok := goa.ValidatePattern(`^rate_`, *mt.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *mt.ID, `^rate_`))
		}
	}
	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^Rate$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Rate$`))
	}
	return
}

// DecodeEasypostRate decodes the EasypostRate instance encoded in resp body.
func (c *Client) DecodeEasypostRate(resp *http.Response) (*EasypostRate, error) {
	var decoded EasypostRate
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostScanform media type (default view)
//
// Identifier: application/easypost.scanform+json
type EasypostScanform struct {
	// Address the will be Shipments shipped from
	Address *EasypostAddress `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// The id of the associated Batch. Unique, starts with "batch_"
	BatchID *string `form:"batch_id,omitempty" json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// File format of the document
	FormFileType *string `form:"form_file_type,omitempty" json:"form_file_type,omitempty" xml:"form_file_type,omitempty"`
	// 	Url of the document
	FormURL *string `form:"form_url,omitempty" json:"form_url,omitempty" xml:"form_url,omitempty"`
	// Unique, begins with "sf_"
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Human readable message explaining any failures
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Always: "ScanForm"
	Object string `form:"object" json:"object" xml:"object"`
	// Current status. Possible values are "creating", "created" and "failed"
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Tracking codes included on the ScanForm
	TrackingCodes []string `form:"tracking_codes,omitempty" json:"tracking_codes,omitempty" xml:"tracking_codes,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the EasypostScanform media type instance.
func (mt *EasypostScanform) Validate() (err error) {
	if mt.Address != nil {
		if err2 := mt.Address.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.ID != nil {
		if ok := goa.ValidatePattern(`^sf_`, *mt.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *mt.ID, `^sf_`))
		}
	}
	if ok := goa.ValidatePattern(`^ScanForm$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^ScanForm$`))
	}
	if mt.Status != nil {
		if !(*mt.Status == "creating" || *mt.Status == "created" || *mt.Status == "failed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"creating", "created", "failed"}))
		}
	}
	return
}

// DecodeEasypostScanform decodes the EasypostScanform instance encoded in resp body.
func (c *Client) DecodeEasypostScanform(resp *http.Response) (*EasypostScanform, error) {
	var decoded EasypostScanform
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostShipment media type (default view)
//
// Identifier: application/easypost.shipment+json
type EasypostShipment struct {
	// The ID of the batch that contains this shipment, if any
	BatchID *string `form:"batch_id,omitempty" json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The current message of the associated BatchShipment
	BatchMessage *string `form:"batch_message,omitempty" json:"batch_message,omitempty" xml:"batch_message,omitempty"`
	// The current state of the associated BatchShipment
	BatchStatus *string `form:"batch_status,omitempty" json:"batch_status,omitempty" xml:"batch_status,omitempty"`
	// The buyer's address, defaults to to_address
	BuyerAddress *AddressPayload `form:"buyer_address,omitempty" json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Information for the processing of customs
	CustomsInfo *CustomsInfoPayload `form:"customs_info,omitempty" json:"customs_info,omitempty" xml:"customs_info,omitempty"`
	// The associated Fee objects charged to the billing user account
	Fees []*EasypostFee `form:"fees,omitempty" json:"fees,omitempty" xml:"fees,omitempty"`
	// All associated Form objects
	Forms []interface{} `form:"forms,omitempty" json:"forms,omitempty" xml:"forms,omitempty"`
	// The origin address
	FromAddress *AddressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// Unique, begins with "shp_"
	ID string `form:"id" json:"id" xml:"id"`
	// The associated Insurance object
	Insurance *InsurancePayload `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Set true to create as a return, discussed in more depth below
	IsReturn *bool `form:"is_return,omitempty" json:"is_return,omitempty" xml:"is_return,omitempty"`
	// Any carrier errors encountered during rating, discussed in more depth below
	Messages []*Message `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Set based on which api-key you used, either "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Always: "Shipment"
	Object string `form:"object" json:"object" xml:"object"`
	// All of the options passed to the shipment, discussed in more depth below
	Options *OptionsPayload `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// The dimensions and weight of the package
	Parcel *ParcelPayload `form:"parcel,omitempty" json:"parcel,omitempty" xml:"parcel,omitempty"`
	// The associated PostageLabel object
	PostageLabel *EasypostPostagelabel `form:"postage_label,omitempty" json:"postage_label,omitempty" xml:"postage_label,omitempty"`
	// All associated Rate objects
	Rates []*EasypostRate `form:"rates,omitempty" json:"rates,omitempty" xml:"rates,omitempty"`
	// The current status of the shipment refund process. Possible values are "submitted", "refunded", "rejected".
	RefundStatus *string `form:"refund_status,omitempty" json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// The shipper's address, defaults to from_address
	ReturnAddress *AddressPayload `form:"return_address,omitempty" json:"return_address,omitempty" xml:"return_address,omitempty"`
	// Document created to manifest and scan multiple shipments
	ScanForm *EasypostScanform `form:"scan_form,omitempty" json:"scan_form,omitempty" xml:"scan_form,omitempty"`
	// The specific rate purchased for the shipment, or null if unpurchased or purchased through another mechanism
	SelectedRate *EasypostRate `form:"selected_rate,omitempty" json:"selected_rate,omitempty" xml:"selected_rate,omitempty"`
	// The current tracking status of the shipment
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The destination address
	ToAddress *AddressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The associated Tracker object
	Tracker *EasypostTracker `form:"tracker,omitempty" json:"tracker,omitempty" xml:"tracker,omitempty"`
	// If purchased, the tracking code will appear here as well as within the Tracker object
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The USPS zone of the shipment, if purchased with USPS
	UspsZone *string `form:"usps_zone,omitempty" json:"usps_zone,omitempty" xml:"usps_zone,omitempty"`
}

// Validate validates the EasypostShipment media type instance.
func (mt *EasypostShipment) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if mt.BuyerAddress != nil {
		if err2 := mt.BuyerAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.CustomsInfo != nil {
		if err2 := mt.CustomsInfo.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range mt.Fees {
		if ok := goa.ValidatePattern(`^Fee$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.fees[*].object`, e.Object, `^Fee$`))
		}
		if e.Type != nil {
			if !(*e.Type == "LabelFee" || *e.Type == "PostageFee" || *e.Type == "InsuranceFee" || *e.Type == "TrackerFee") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.fees[*].type`, *e.Type, []interface{}{"LabelFee", "PostageFee", "InsuranceFee", "TrackerFee"}))
			}
		}
	}
	if mt.FromAddress != nil {
		if err2 := mt.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`^shp_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^shp_`))
	}
	if mt.Insurance != nil {
		if err2 := mt.Insurance.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^Shipment$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Shipment$`))
	}
	if mt.Options != nil {
		if err2 := mt.Options.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.PostageLabel != nil {
		if err2 := mt.PostageLabel.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range mt.Rates {
		if e.ID != nil {
			if ok := goa.ValidatePattern(`^rate_`, *e.ID); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response.rates[*].id`, *e.ID, `^rate_`))
			}
		}
		if !(e.Mode == "test" || e.Mode == "production") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.rates[*].mode`, e.Mode, []interface{}{"test", "production"}))
		}
		if ok := goa.ValidatePattern(`^Rate$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.rates[*].object`, e.Object, `^Rate$`))
		}
	}
	if mt.RefundStatus != nil {
		if !(*mt.RefundStatus == "submitted" || *mt.RefundStatus == "rejected" || *mt.RefundStatus == "refunded") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.refund_status`, *mt.RefundStatus, []interface{}{"submitted", "rejected", "refunded"}))
		}
	}
	if mt.ReturnAddress != nil {
		if err2 := mt.ReturnAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.ScanForm != nil {
		if err2 := mt.ScanForm.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.SelectedRate != nil {
		if err2 := mt.SelectedRate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.ToAddress != nil {
		if err2 := mt.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Tracker != nil {
		if err2 := mt.Tracker.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeEasypostShipment decodes the EasypostShipment instance encoded in resp body.
func (c *Client) DecodeEasypostShipment(resp *http.Response) (*EasypostShipment, error) {
	var decoded EasypostShipment
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The Tracker object contains both the current information about the package as well as previous updates. All of the previous updates are stored in the tracking_details array. Each TrackingDetail object contains the status, the message from the carrier, and a TrackingLocation. (default view)
//
// Identifier: application/easypost.tracker+json
type EasypostTracker struct {
	// The name of the carrier handling the shipment
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The associated CarrierDetail object (if available)
	CarrierDetail []*EasypostCarrierdetail `form:"carrier_detail,omitempty" json:"carrier_detail,omitempty" xml:"carrier_detail,omitempty"`
	// Time Created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The estimated delivery date provided by the carrier (if available)
	EstDeliveryDate *string `form:"est_delivery_date,omitempty" json:"est_delivery_date,omitempty" xml:"est_delivery_date,omitempty"`
	// Array of the associated Fee objects
	Fees []*EasypostFee `form:"fees,omitempty" json:"fees,omitempty" xml:"fees,omitempty"`
	// Unique, begins with "trk_"
	ID string `form:"id" json:"id" xml:"id"`
	// Set based on which api-key you used, either "test" or "production"
	Mode string `form:"mode" json:"mode" xml:"mode"`
	// Always: "Tracker"
	Object string `form:"object" json:"object" xml:"object"`
	// URL to a publicly-accessible html page that shows tracking details for this tracker
	PublicURL *string `form:"public_url,omitempty" json:"public_url,omitempty" xml:"public_url,omitempty"`
	// The id of the EasyPost Shipment object associated with the Tracker (if any)
	ShipmentID *string `form:"shipment_id,omitempty" json:"shipment_id,omitempty" xml:"shipment_id,omitempty"`
	// The name of the person who signed for the package (if available)
	SignedBy *string `form:"signed_by,omitempty" json:"signed_by,omitempty" xml:"signed_by,omitempty"`
	// The current status of the package, possible values are "unknown", "pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled" or "error"
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The tracking code provided by the carrier
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Array of the associated TrackingDetail objects
	TrackingDetails []*EasypostTrackingdetail `form:"tracking_details,omitempty" json:"tracking_details,omitempty" xml:"tracking_details,omitempty"`
	// Time Last Updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The weight of the package as measured by the carrier in ounces (if available)
	Weight *float64 `form:"weight,omitempty" json:"weight,omitempty" xml:"weight,omitempty"`
}

// Validate validates the EasypostTracker media type instance.
func (mt *EasypostTracker) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	for _, e := range mt.CarrierDetail {
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.carrier_detail[*]`, "object"))
		}

		if ok := goa.ValidatePattern(`^CarrierDetail$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.carrier_detail[*].object`, e.Object, `^CarrierDetail$`))
		}
	}
	for _, e := range mt.Fees {
		if ok := goa.ValidatePattern(`^Fee$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.fees[*].object`, e.Object, `^Fee$`))
		}
		if e.Type != nil {
			if !(*e.Type == "LabelFee" || *e.Type == "PostageFee" || *e.Type == "InsuranceFee" || *e.Type == "TrackerFee") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.fees[*].type`, *e.Type, []interface{}{"LabelFee", "PostageFee", "InsuranceFee", "TrackerFee"}))
			}
		}
	}
	if ok := goa.ValidatePattern(`^trk_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^trk_`))
	}
	if !(mt.Mode == "test" || mt.Mode == "production") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.mode`, mt.Mode, []interface{}{"test", "production"}))
	}
	if ok := goa.ValidatePattern(`^Tracker$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^Tracker$`))
	}
	if mt.Status != nil {
		if !(*mt.Status == "pre_transit" || *mt.Status == "in_transit" || *mt.Status == "out_for_delivery" || *mt.Status == "delivered" || *mt.Status == "available_for_pickup" || *mt.Status == "return_to_sender" || *mt.Status == "failure" || *mt.Status == "cancelled" || *mt.Status == "error" || *mt.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown"}))
		}
	}
	for _, e := range mt.TrackingDetails {
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.tracking_details[*]`, "object"))
		}

		if ok := goa.ValidatePattern(`^TrackingDetail$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.tracking_details[*].object`, e.Object, `^TrackingDetail$`))
		}
		if e.Status != nil {
			if !(*e.Status == "pre_transit" || *e.Status == "in_transit" || *e.Status == "out_for_delivery" || *e.Status == "delivered" || *e.Status == "available_for_pickup" || *e.Status == "return_to_sender" || *e.Status == "failure" || *e.Status == "cancelled" || *e.Status == "error" || *e.Status == "unknown") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.tracking_details[*].status`, *e.Status, []interface{}{"pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown"}))
			}
		}
		if e.TrackingLocation != nil {
			if err2 := e.TrackingLocation.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeEasypostTracker decodes the EasypostTracker instance encoded in resp body.
func (c *Client) DecodeEasypostTracker(resp *http.Response) (*EasypostTracker, error) {
	var decoded EasypostTracker
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostTrackers media type (default view)
//
// Identifier: application/easypost.trackers+json
type EasypostTrackers struct {
	Trackers []*EasypostTracker `form:"trackers" json:"trackers" xml:"trackers"`
}

// Validate validates the EasypostTrackers media type instance.
func (mt *EasypostTrackers) Validate() (err error) {
	if mt.Trackers == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "trackers"))
	}

	for _, e := range mt.Trackers {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.trackers[*]`, "id"))
		}
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.trackers[*]`, "object"))
		}

		for _, e := range e.CarrierDetail {
			if e.Object == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response.trackers[*].carrier_detail[*]`, "object"))
			}

			if ok := goa.ValidatePattern(`^CarrierDetail$`, e.Object); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response.trackers[*].carrier_detail[*].object`, e.Object, `^CarrierDetail$`))
			}
		}
		for _, e := range e.Fees {
			if ok := goa.ValidatePattern(`^Fee$`, e.Object); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response.trackers[*].fees[*].object`, e.Object, `^Fee$`))
			}
			if e.Type != nil {
				if !(*e.Type == "LabelFee" || *e.Type == "PostageFee" || *e.Type == "InsuranceFee" || *e.Type == "TrackerFee") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.trackers[*].fees[*].type`, *e.Type, []interface{}{"LabelFee", "PostageFee", "InsuranceFee", "TrackerFee"}))
				}
			}
		}
		if ok := goa.ValidatePattern(`^trk_`, e.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.trackers[*].id`, e.ID, `^trk_`))
		}
		if !(e.Mode == "test" || e.Mode == "production") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.trackers[*].mode`, e.Mode, []interface{}{"test", "production"}))
		}
		if ok := goa.ValidatePattern(`^Tracker$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.trackers[*].object`, e.Object, `^Tracker$`))
		}
		if e.Status != nil {
			if !(*e.Status == "pre_transit" || *e.Status == "in_transit" || *e.Status == "out_for_delivery" || *e.Status == "delivered" || *e.Status == "available_for_pickup" || *e.Status == "return_to_sender" || *e.Status == "failure" || *e.Status == "cancelled" || *e.Status == "error" || *e.Status == "unknown") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.trackers[*].status`, *e.Status, []interface{}{"pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown"}))
			}
		}
		for _, e := range e.TrackingDetails {
			if e.Object == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response.trackers[*].tracking_details[*]`, "object"))
			}

			if ok := goa.ValidatePattern(`^TrackingDetail$`, e.Object); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response.trackers[*].tracking_details[*].object`, e.Object, `^TrackingDetail$`))
			}
			if e.Status != nil {
				if !(*e.Status == "pre_transit" || *e.Status == "in_transit" || *e.Status == "out_for_delivery" || *e.Status == "delivered" || *e.Status == "available_for_pickup" || *e.Status == "return_to_sender" || *e.Status == "failure" || *e.Status == "cancelled" || *e.Status == "error" || *e.Status == "unknown") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.trackers[*].tracking_details[*].status`, *e.Status, []interface{}{"pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown"}))
				}
			}
			if e.TrackingLocation != nil {
				if err2 := e.TrackingLocation.Validate(); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
	}
	return
}

// DecodeEasypostTrackers decodes the EasypostTrackers instance encoded in resp body.
func (c *Client) DecodeEasypostTrackers(resp *http.Response) (*EasypostTrackers, error) {
	var decoded EasypostTrackers
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Each TrackingDetail object contains the status, the message from the carrier, and a TrackingLocation. (default view)
//
// Identifier: application/easypost.trackingdetail+json
type EasypostTrackingdetail struct {
	// The timestamp when the tracking scan occurred
	Datetime *string `form:"datetime,omitempty" json:"datetime,omitempty" xml:"datetime,omitempty"`
	// Description of the scan event
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Always: "TrackingDetail"
	Object string `form:"object" json:"object" xml:"object"`
	// The original source of the information for this scan event, usually the carrier
	Source *string `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// The current status of the package, possible values are "unknown", "pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled" or "error"
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// The location associated with the scan event
	TrackingLocation *EasypostTrackinglocation `form:"tracking_location,omitempty" json:"tracking_location,omitempty" xml:"tracking_location,omitempty"`
}

// Validate validates the EasypostTrackingdetail media type instance.
func (mt *EasypostTrackingdetail) Validate() (err error) {
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^TrackingDetail$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^TrackingDetail$`))
	}
	if mt.Status != nil {
		if !(*mt.Status == "pre_transit" || *mt.Status == "in_transit" || *mt.Status == "out_for_delivery" || *mt.Status == "delivered" || *mt.Status == "available_for_pickup" || *mt.Status == "return_to_sender" || *mt.Status == "failure" || *mt.Status == "cancelled" || *mt.Status == "error" || *mt.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown"}))
		}
	}
	if mt.TrackingLocation != nil {
		if err2 := mt.TrackingLocation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeEasypostTrackingdetail decodes the EasypostTrackingdetail instance encoded in resp body.
func (c *Client) DecodeEasypostTrackingdetail(resp *http.Response) (*EasypostTrackingdetail, error) {
	var decoded EasypostTrackingdetail
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The TrackingLocation contains city, state, country, and zip information about the location where the package was scanned. The information each carrier provides is different, so some carriers may not make use of all of these fields. (default view)
//
// Identifier: application/easypost.trackinglocation+json
type EasypostTrackinglocation struct {
	// The city where the scan event occurred (if available)
	City *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// 	The country where the scan event occurred (if available)
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Always: "TrackingLocation"
	Object string `form:"object" json:"object" xml:"object"`
	// The state where the scan event occurred (if available)
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// The postal code where the scan event occurred (if available)
	Zip *string `form:"zip,omitempty" json:"zip,omitempty" xml:"zip,omitempty"`
}

// Validate validates the EasypostTrackinglocation media type instance.
func (mt *EasypostTrackinglocation) Validate() (err error) {
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^TrackingLocation$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^TrackingLocation$`))
	}
	return
}

// DecodeEasypostTrackinglocation decodes the EasypostTrackinglocation instance encoded in resp body.
func (c *Client) DecodeEasypostTrackinglocation(resp *http.Response) (*EasypostTrackinglocation, error) {
	var decoded EasypostTrackinglocation
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// An Insurance object represents insurance for packages purchased both va the EasyPost API as well as shipments purchased through third parties and later registered with EasyPost. An Insurance is created automatically whenever you buy a Shipment through EasyPost and pass insurance options during the Buy call or in a later call to Insure a Shipment. (default view)
//
// Identifier: application/easypost.user+json
type EasypostUser struct {
	// Formatted as string "XX.XXXXX"
	Balance *string `form:"balance,omitempty" json:"balance,omitempty" xml:"balance,omitempty"`
	// All associated children
	Children EasypostUserCollection `form:"children,omitempty" json:"children,omitempty" xml:"children,omitempty"`
	// Required
	Email string `form:"email" json:"email" xml:"email"`
	// Unique, begins with "user_"
	ID string `form:"id" json:"id" xml:"id"`
	// First and last name required
	Name string `form:"name" json:"name" xml:"name"`
	// Always: "User"
	Object string `form:"object" json:"object" xml:"object"`
	// The ID of the parent user object. Top-level users are defined as users with no parent
	ParentID *string `form:"parent_id,omitempty" json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// Optional
	PhoneNumber *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// USD formatted dollars and cents, delimited by a decimal point
	RechargeAmount *string `form:"recharge_amount,omitempty" json:"recharge_amount,omitempty" xml:"recharge_amount,omitempty"`
	// Number of cents USD that when your balance drops below, we automatically recharge your account with your primary payment method.
	RechargeThreshold *string `form:"recharge_threshold,omitempty" json:"recharge_threshold,omitempty" xml:"recharge_threshold,omitempty"`
	// USD formatted dollars and cents, delimited by a decimal point
	SecondaryRechargeAmount *string `form:"secondary_recharge_amount,omitempty" json:"secondary_recharge_amount,omitempty" xml:"secondary_recharge_amount,omitempty"`
}

// Validate validates the EasypostUser media type instance.
func (mt *EasypostUser) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}

	if err2 := mt.Children.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if ok := goa.ValidatePattern(`^user_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^user_`))
	}
	if ok := goa.ValidatePattern(`^User$`, mt.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, mt.Object, `^User$`))
	}
	return
}

// DecodeEasypostUser decodes the EasypostUser instance encoded in resp body.
func (c *Client) DecodeEasypostUser(resp *http.Response) (*EasypostUser, error) {
	var decoded EasypostUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EasypostUserCollection is the media type for an array of EasypostUser (default view)
//
// Identifier: application/easypost.user+json; type=collection
type EasypostUserCollection []*EasypostUser

// Validate validates the EasypostUserCollection media type instance.
func (mt EasypostUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Object == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "object"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.Email == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "email"))
		}

		if err2 := e.Children.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if ok := goa.ValidatePattern(`^user_`, e.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response[*].id`, e.ID, `^user_`))
		}
		if ok := goa.ValidatePattern(`^User$`, e.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response[*].object`, e.Object, `^User$`))
		}
	}
	return
}

// DecodeEasypostUserCollection decodes the EasypostUserCollection instance encoded in resp body.
func (c *Client) DecodeEasypostUserCollection(resp *http.Response) (EasypostUserCollection, error) {
	var decoded EasypostUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// When rating a Shipment or Pickup, some carriers may fail to generate rates. These failures are returned as part of the Shipment or Pickup as part of their messages attribute, and follow a common object structure. (default view)
//
// Identifier: message
type Message struct {
	// the name of the carrier generating the error, e.g. "UPS"
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// the string from the carrier explaining the problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// the category of error that occurred. Most frequently "rate_error"
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// DecodeMessage decodes the Message instance encoded in resp body.
func (c *Client) DecodeMessage(resp *http.Response) (*Message, error) {
	var decoded Message
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
