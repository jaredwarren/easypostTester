//************************************************************************//
// API "cellar": Application Media Types
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

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts. (default view)
//
// Identifier: application/easypost.carrier_accounts+json
type EasypostCarrierAccounts struct {
	// If clone is true, only the reference and description are possible to update
	Clone *bool `form:"clone,omitempty" json:"clone,omitempty" xml:"clone,omitempty"`
	// The name used when displaying a readable value for the type of the account
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *EasypostFieldsObject `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
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

	if ok := goa.ValidatePattern(`^ca_`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^ca_`))
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

// Contains "credentials" and/or "test_credentials", or may be empty (default view)
//
// Identifier: application/easypost.field_object+json
type EasypostFieldObject struct {
	// Each key in the sub-objects of a CarrierAccount's fields is the name of a settable field
	Key string `form:"key" json:"key" xml:"key"`
	// The label value is used in form rendering to display a more precise field name
	Label string `form:"label" json:"label" xml:"label"`
	// Checkbox fields use "0" and "1" as False and True, all other field types present plaintext, partly-masked, or masked credential data for reference
	Value string `form:"value" json:"value" xml:"value"`
	// The visibility value is used to control form field types, and is discussed in the CarrierType section
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
