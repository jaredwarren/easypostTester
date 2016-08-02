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

package app

import (
	"github.com/goadesign/goa"
	"time"
)

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
type yesNoPayload struct {
	// If clone is true, only the reference and description are possible to update
	Clone *bool `form:"clone,omitempty" json:"clone,omitempty" xml:"clone,omitempty"`
	// The name used when displaying a readable value for the type of the account
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
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
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Finalize sets the default values for yesNoPayload type instance.
func (ut *yesNoPayload) Finalize() {
	var defaultObject = "CarrierAccount"
	if ut.Object == nil {
		ut.Object = &defaultObject
	}
}

// Validate validates the yesNoPayload type instance.
func (ut *yesNoPayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ut.Object == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ut.ID != nil {
		if ok := goa.ValidatePattern(`^ca_`, *ut.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *ut.ID, `^ca_`))
		}
	}
	if ut.Object != nil {
		if ok := goa.ValidatePattern(`^CarrierAccount$`, *ut.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, *ut.Object, `^CarrierAccount$`))
		}
	}
	return
}

// Publicize creates YesNoPayload from yesNoPayload
func (ut *yesNoPayload) Publicize() *YesNoPayload {
	var pub YesNoPayload
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
	if ut.ID != nil {
		pub.ID = *ut.ID
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
		pub.Type = ut.Type
	}
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	return &pub
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.
type YesNoPayload struct {
	// If clone is true, only the reference and description are possible to update
	Clone *bool `form:"clone,omitempty" json:"clone,omitempty" xml:"clone,omitempty"`
	// The name used when displaying a readable value for the type of the account
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
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
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the YesNoPayload type instance.
func (ut *YesNoPayload) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ut.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "object"))
	}

	if ok := goa.ValidatePattern(`^ca_`, ut.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, ut.ID, `^ca_`))
	}
	if ok := goa.ValidatePattern(`^CarrierAccount$`, ut.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.object`, ut.Object, `^CarrierAccount$`))
	}
	return
}
