//************************************************************************//
// API "cellar": Application User Types
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
		pub.Clone = ut.Clone
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
	Clone *bool `form:"clone,omitempty" json:"clone,omitempty" xml:"clone,omitempty"`
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

// Publicize creates FieldsObjectPayload from fieldsObjectPayload
func (ut *fieldsObjectPayload) Publicize() *FieldsObjectPayload {
	var pub FieldsObjectPayload
	if ut.AutoLink != nil {
		pub.AutoLink = ut.AutoLink
	}
	if ut.Credentials != nil {
		pub.Credentials = ut.Credentials
	}
	if ut.CustomWorkflow != nil {
		pub.CustomWorkflow = ut.CustomWorkflow
	}
	if ut.TestCredentials != nil {
		pub.TestCredentials = ut.TestCredentials
	}
	return &pub
}

// Contains "credentials" and/or "test_credentials", or may be empty
type FieldsObjectPayload struct {
	// For USPS this designates that no credentials are required.
	AutoLink *bool `form:"auto_link,omitempty" json:"auto_link,omitempty" xml:"auto_link,omitempty"`
	// Credentials used in the production environment.
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// When present, a seperate authentication process will be required through the UI to link this account type.
	CustomWorkflow *bool `form:"custom_workflow,omitempty" json:"custom_workflow,omitempty" xml:"custom_workflow,omitempty"`
	// Credentials used in the test environment.
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
}
