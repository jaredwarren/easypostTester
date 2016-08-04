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

package app

import "github.com/goadesign/goa"

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

// EasypostAddress_verification media type (default view)
//
// Identifier: application/easypost.address_verification+json
type EasypostAddressVerification struct {
	// All errors that caused the verification to fail
	Errors []*ApplicationEasypostGieldErrorJSON `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
	// The success of the verification
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
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

// EasypostShipment media type (default view)
//
// Identifier: application/easypost.shipment+json
type EasypostShipment struct {
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

// Validate validates the EasypostShipment media type instance.
func (mt *EasypostShipment) Validate() (err error) {
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
