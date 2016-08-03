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
	// Unique, begins with "adr_"
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Set based on which api-key you used, either "test" or "production"
	Mode *string `form:"mode,omitempty" json:"mode,omitempty" xml:"mode,omitempty"`
	// Name of the person. Both name and company can be included
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Always: "Address"
	Object *string `form:"object,omitempty" json:"object,omitempty" xml:"object,omitempty"`
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
	// The result of any verifications requested
	Verifications *verificationsPayload `form:"verifications,omitempty" json:"verifications,omitempty" xml:"verifications,omitempty"`
	// ZIP or postal code the address is located in
	Zip *string `form:"zip,omitempty" json:"zip,omitempty" xml:"zip,omitempty"`
}

// Finalize sets the default values for addressPayload type instance.
func (ut *addressPayload) Finalize() {
	var defaultObject = "Address"
	if ut.Object == nil {
		ut.Object = &defaultObject
	}
}

// Validate validates the addressPayload type instance.
func (ut *addressPayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ut.Object == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ut.ID != nil {
		if ok := goa.ValidatePattern(`^adr_`, *ut.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *ut.ID, `^adr_`))
		}
	}
	if ut.Object != nil {
		if ok := goa.ValidatePattern(`^Address$`, *ut.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, *ut.Object, `^Address$`))
		}
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
		pub.City = ut.City
	}
	if ut.Company != nil {
		pub.Company = ut.Company
	}
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FederalTaxID != nil {
		pub.FederalTaxID = ut.FederalTaxID
	}
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Mode != nil {
		pub.Mode = ut.Mode
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Object != nil {
		pub.Object = *ut.Object
	}
	if ut.Phone != nil {
		pub.Phone = ut.Phone
	}
	if ut.Residential != nil {
		pub.Residential = ut.Residential
	}
	if ut.State != nil {
		pub.State = ut.State
	}
	if ut.StateTaxID != nil {
		pub.StateTaxID = ut.StateTaxID
	}
	if ut.Street1 != nil {
		pub.Street1 = ut.Street1
	}
	if ut.Street2 != nil {
		pub.Street2 = ut.Street2
	}
	if ut.Verifications != nil {
		pub.Verifications = ut.Verifications.Publicize()
	}
	if ut.Zip != nil {
		pub.Zip = ut.Zip
	}
	return &pub
}

// Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.
type AddressPayload struct {
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
	// Unique, begins with "adr_"
	ID string `form:"id" json:"id" xml:"id"`
	// Set based on which api-key you used, either "test" or "production"
	Mode *string `form:"mode,omitempty" json:"mode,omitempty" xml:"mode,omitempty"`
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
	// The result of any verifications requested
	Verifications *VerificationsPayload `form:"verifications,omitempty" json:"verifications,omitempty" xml:"verifications,omitempty"`
	// ZIP or postal code the address is located in
	Zip *string `form:"zip,omitempty" json:"zip,omitempty" xml:"zip,omitempty"`
}

// Validate validates the AddressPayload type instance.
func (ut *AddressPayload) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ut.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^adr_`, ut.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, ut.ID, `^adr_`))
	}
	if ok := goa.ValidatePattern(`^Address$`, ut.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, ut.Object, `^Address$`))
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
	// The name used when displaying a readable value for the type of the account
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *fieldsObjectPayload `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
	// Unique, begins with "ca_"
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Always: "CarrierAccount"
	Object *string `form:"object,omitempty" json:"object,omitempty" xml:"object,omitempty"`
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
	var defaultObject = "CarrierAccount"
	if ut.Object == nil {
		ut.Object = &defaultObject
	}
}

// Validate validates the carrierAccountPayload type instance.
func (ut *carrierAccountPayload) Validate() (err error) {
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	if ut.Object != nil {
		if ok := goa.ValidatePattern(`^CarrierAccount$`, *ut.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, *ut.Object, `^CarrierAccount$`))
		}
	}
	return
}

// Publicize creates CarrierAccountPayload from carrierAccountPayload
func (ut *carrierAccountPayload) Publicize() *CarrierAccountPayload {
	var pub CarrierAccountPayload
	if ut.Clone != nil {
		pub.Clone = *ut.Clone
	}
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
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
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Object != nil {
		pub.Object = *ut.Object
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
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	return &pub
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.
type CarrierAccountPayload struct {
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
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Always: "CarrierAccount"
	Object string `form:"object" json:"object" xml:"object"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// An optional field that may be used in place of carrier_account_id in other API endpoints
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type string `form:"type" json:"type" xml:"type"`
	// The name used when displaying a readable value for the type of the account
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the CarrierAccountPayload type instance.
func (ut *CarrierAccountPayload) Validate() (err error) {
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	if ok := goa.ValidatePattern(`^CarrierAccount$`, ut.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, ut.Object, `^CarrierAccount$`))
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

// verificationPayload user type.
type verificationPayload struct {
	// All errors that caused the verification to fail
	Errors []*applicationEasypostGieldErrorJSON `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
	// The success of the verification
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// Publicize creates VerificationPayload from verificationPayload
func (ut *verificationPayload) Publicize() *VerificationPayload {
	var pub VerificationPayload
	if ut.Errors != nil {
		pub.Errors = make([]*ApplicationEasypostGieldErrorJSON, len(ut.Errors))
		for i2, elem2 := range ut.Errors {
			pub.Errors[i2] = elem2.Publicize()
		}
	}
	if ut.Success != nil {
		pub.Success = ut.Success
	}
	return &pub
}

// VerificationPayload user type.
type VerificationPayload struct {
	// All errors that caused the verification to fail
	Errors []*ApplicationEasypostGieldErrorJSON `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
	// The success of the verification
	Success *bool `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// verificationsPayload user type.
type verificationsPayload struct {
	// Checks that the address is deliverable and makes minor corrections to spelling/format. US addresses will also have their "residential" status checked and set.
	Delivery *verificationPayload `form:"delivery,omitempty" json:"delivery,omitempty" xml:"delivery,omitempty"`
	// Only applicable to US addresses - checks and sets the ZIP+4.
	Zip4 *verificationPayload `form:"zip4,omitempty" json:"zip4,omitempty" xml:"zip4,omitempty"`
}

// Publicize creates VerificationsPayload from verificationsPayload
func (ut *verificationsPayload) Publicize() *VerificationsPayload {
	var pub VerificationsPayload
	if ut.Delivery != nil {
		pub.Delivery = ut.Delivery.Publicize()
	}
	if ut.Zip4 != nil {
		pub.Zip4 = ut.Zip4.Publicize()
	}
	return &pub
}

// VerificationsPayload user type.
type VerificationsPayload struct {
	// Checks that the address is deliverable and makes minor corrections to spelling/format. US addresses will also have their "residential" status checked and set.
	Delivery *VerificationPayload `form:"delivery,omitempty" json:"delivery,omitempty" xml:"delivery,omitempty"`
	// Only applicable to US addresses - checks and sets the ZIP+4.
	Zip4 *VerificationPayload `form:"zip4,omitempty" json:"zip4,omitempty" xml:"zip4,omitempty"`
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
