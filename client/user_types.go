//************************************************************************//
// API "easypost": Application User Types
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

import "github.com/goadesign/goa"

// Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.
type addressPayload struct {
	// The specific designation for the address (only relevant if the address is a carrier facility)
	CarrierFacility *string `form:"carrier_facility,omitempty" json:"carrier_facility,omitempty" xml:"carrier_facility,omitempty"`
	// City the address is located in
	City *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// Name of the organization. Both name and company can be included
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// ISO 3166 country code for the country the address is located in
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Email to reach the person or organization
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Federal tax identifier of the person or organization
	FederalTaxID *string `form:"federal_tax_id,omitempty" json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// Name of the person. Both name and company can be included
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
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
	// The verifications to perform when creating. verify_strict takes precedence
	Verify []string `form:"verify,omitempty" json:"verify,omitempty" xml:"verify,omitempty"`
	// The verifications to perform when creating. The failure of any of these verifications causes the whole request to fail
	VerifyStrict []string `form:"verify_strict,omitempty" json:"verify_strict,omitempty" xml:"verify_strict,omitempty"`
	// ZIP or postal code the address is located in
	Zip *string `form:"zip,omitempty" json:"zip,omitempty" xml:"zip,omitempty"`
}

// Validate validates the addressPayload type instance.
func (ut *addressPayload) Validate() (err error) {
	if ut.Street1 == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street1"))
	}
	if ut.City == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if ut.State == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}
	if ut.Zip == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "zip"))
	}
	if ut.Country == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}

	return
}

// Publicize creates AddressPayload from addressPayload
func (ut *addressPayload) Publicize() *AddressPayload {
	var pub AddressPayload
	if ut.CarrierFacility != nil {
		pub.CarrierFacility = ut.CarrierFacility
	}
	if ut.City != nil {
		pub.City = *ut.City
	}
	if ut.Company != nil {
		pub.Company = ut.Company
	}
	if ut.Country != nil {
		pub.Country = *ut.Country
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FederalTaxID != nil {
		pub.FederalTaxID = ut.FederalTaxID
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Phone != nil {
		pub.Phone = ut.Phone
	}
	if ut.Residential != nil {
		pub.Residential = ut.Residential
	}
	if ut.State != nil {
		pub.State = *ut.State
	}
	if ut.StateTaxID != nil {
		pub.StateTaxID = ut.StateTaxID
	}
	if ut.Street1 != nil {
		pub.Street1 = *ut.Street1
	}
	if ut.Street2 != nil {
		pub.Street2 = ut.Street2
	}
	if ut.Verify != nil {
		pub.Verify = ut.Verify
	}
	if ut.VerifyStrict != nil {
		pub.VerifyStrict = ut.VerifyStrict
	}
	if ut.Zip != nil {
		pub.Zip = *ut.Zip
	}
	return &pub
}

// Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.
type AddressPayload struct {
	// The specific designation for the address (only relevant if the address is a carrier facility)
	CarrierFacility *string `form:"carrier_facility,omitempty" json:"carrier_facility,omitempty" xml:"carrier_facility,omitempty"`
	// City the address is located in
	City string `form:"city" json:"city" xml:"city"`
	// Name of the organization. Both name and company can be included
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// ISO 3166 country code for the country the address is located in
	Country string `form:"country" json:"country" xml:"country"`
	// Email to reach the person or organization
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Federal tax identifier of the person or organization
	FederalTaxID *string `form:"federal_tax_id,omitempty" json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// Name of the person. Both name and company can be included
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Phone number to reach the person or organization
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Whether or not this address would be considered residential
	Residential *bool `form:"residential,omitempty" json:"residential,omitempty" xml:"residential,omitempty"`
	// State or province the address is located in
	State string `form:"state" json:"state" xml:"state"`
	// 	State tax identifier of the person or organization
	StateTaxID *string `form:"state_tax_id,omitempty" json:"state_tax_id,omitempty" xml:"state_tax_id,omitempty"`
	// First line of the address
	Street1 string `form:"street1" json:"street1" xml:"street1"`
	// Second line of the address
	Street2 *string `form:"street2,omitempty" json:"street2,omitempty" xml:"street2,omitempty"`
	// The verifications to perform when creating. verify_strict takes precedence
	Verify []string `form:"verify,omitempty" json:"verify,omitempty" xml:"verify,omitempty"`
	// The verifications to perform when creating. The failure of any of these verifications causes the whole request to fail
	VerifyStrict []string `form:"verify_strict,omitempty" json:"verify_strict,omitempty" xml:"verify_strict,omitempty"`
	// ZIP or postal code the address is located in
	Zip string `form:"zip" json:"zip" xml:"zip"`
}

// Validate validates the AddressPayload type instance.
func (ut *AddressPayload) Validate() (err error) {
	if ut.Street1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street1"))
	}
	if ut.City == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if ut.State == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}
	if ut.Zip == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "zip"))
	}
	if ut.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}

	return
}

// bottlePayload user type.
type bottlePayload struct {
	// status of bottle
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// Validate validates the bottlePayload type instance.
func (ut *bottlePayload) Validate() (err error) {
	if ut.Status == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "status"))
	}

	return
}

// Publicize creates BottlePayload from bottlePayload
func (ut *bottlePayload) Publicize() *BottlePayload {
	var pub BottlePayload
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	return &pub
}

// BottlePayload user type.
type BottlePayload struct {
	// status of bottle
	Status int `form:"status" json:"status" xml:"status"`
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.
type carrierAccountPayload struct {
	// If clone is true, only the reference and description are possible to update
	Clone *bool `form:"clone,omitempty" json:"clone,omitempty" xml:"clone,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *fieldsObjectPayload `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// An optional field that may be used in place of carrier_account_id in other API endpoints
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Finalize sets the default values for carrierAccountPayload type instance.
func (ut *carrierAccountPayload) Finalize() {
	var defaultClone = false
	if ut.Clone == nil {
		ut.Clone = &defaultClone
	}
	if ut.Fields != nil {
		var defaultAutoLink = false
		if ut.Fields.AutoLink == nil {
			ut.Fields.AutoLink = &defaultAutoLink
		}
		var defaultCustomWorkflow = false
		if ut.Fields.CustomWorkflow == nil {
			ut.Fields.CustomWorkflow = &defaultCustomWorkflow
		}
	}
}

// Validate validates the carrierAccountPayload type instance.
func (ut *carrierAccountPayload) Validate() (err error) {
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// Publicize creates CarrierAccountPayload from carrierAccountPayload
func (ut *carrierAccountPayload) Publicize() *CarrierAccountPayload {
	var pub CarrierAccountPayload
	if ut.Clone != nil {
		pub.Clone = *ut.Clone
	}
	if ut.Credentials != nil {
		pub.Credentials = ut.Credentials
	}
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Fields != nil {
		pub.Fields = ut.Fields.Publicize()
	}
	if ut.Readable != nil {
		pub.Readable = ut.Readable
	}
	if ut.Reference != nil {
		pub.Reference = ut.Reference
	}
	if ut.TestCredentials != nil {
		pub.TestCredentials = ut.TestCredentials
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	return &pub
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.
type CarrierAccountPayload struct {
	// If clone is true, only the reference and description are possible to update
	Clone bool `form:"clone" json:"clone" xml:"clone"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *FieldsObjectPayload `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// An optional field that may be used in place of carrier_account_id in other API endpoints
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the CarrierAccountPayload type instance.
func (ut *CarrierAccountPayload) Validate() (err error) {
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// Contains "credentials" and/or "test_credentials", or may be empty
type fieldsObjectPayload struct {
	// For USPS this designates that no credentials are required.
	AutoLink *bool `form:"auto_link,omitempty" json:"auto_link,omitempty" xml:"auto_link,omitempty"`
	// Credentials used in the production environment.
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// When present, a seperate authentication process will be required through the UI to link this account type.
	CustomWorkflow *bool `form:"custom_workflow,omitempty" json:"custom_workflow,omitempty" xml:"custom_workflow,omitempty"`
	// Credentials used in the test environment.
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
}

// Finalize sets the default values for fieldsObjectPayload type instance.
func (ut *fieldsObjectPayload) Finalize() {
	var defaultAutoLink = false
	if ut.AutoLink == nil {
		ut.AutoLink = &defaultAutoLink
	}
	var defaultCustomWorkflow = false
	if ut.CustomWorkflow == nil {
		ut.CustomWorkflow = &defaultCustomWorkflow
	}
}

// Publicize creates FieldsObjectPayload from fieldsObjectPayload
func (ut *fieldsObjectPayload) Publicize() *FieldsObjectPayload {
	var pub FieldsObjectPayload
	if ut.AutoLink != nil {
		pub.AutoLink = *ut.AutoLink
	}
	if ut.Credentials != nil {
		pub.Credentials = ut.Credentials
	}
	if ut.CustomWorkflow != nil {
		pub.CustomWorkflow = *ut.CustomWorkflow
	}
	if ut.TestCredentials != nil {
		pub.TestCredentials = ut.TestCredentials
	}
	return &pub
}

// Contains "credentials" and/or "test_credentials", or may be empty
type FieldsObjectPayload struct {
	// For USPS this designates that no credentials are required.
	AutoLink bool `form:"auto_link" json:"auto_link" xml:"auto_link"`
	// Credentials used in the production environment.
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// When present, a seperate authentication process will be required through the UI to link this account type.
	CustomWorkflow bool `form:"custom_workflow" json:"custom_workflow" xml:"custom_workflow"`
	// Credentials used in the test environment.
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
}

// applicationEasypostGieldErrorJSON user type.
type applicationEasypostGieldErrorJSON struct {
	// Field of the request that the error describes
	Field *string `form:"field,omitempty" json:"field,omitempty" xml:"field,omitempty"`
	// Human readable description of the problem
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// Publicize creates ApplicationEasypostGieldErrorJSON from applicationEasypostGieldErrorJSON
func (ut *applicationEasypostGieldErrorJSON) Publicize() *ApplicationEasypostGieldErrorJSON {
	var pub ApplicationEasypostGieldErrorJSON
	if ut.Field != nil {
		pub.Field = ut.Field
	}
	if ut.Message != nil {
		pub.Message = ut.Message
	}
	return &pub
}

// ApplicationEasypostGieldErrorJSON user type.
type ApplicationEasypostGieldErrorJSON struct {
	// Field of the request that the error describes
	Field *string `form:"field,omitempty" json:"field,omitempty" xml:"field,omitempty"`
	// Human readable description of the problem
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}
