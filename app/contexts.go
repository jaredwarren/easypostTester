//************************************************************************//
// API "easypost": Application Contexts
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
	"golang.org/x/net/context"
)

// CreateAddressContext provides the address create action context.
type CreateAddressContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateAddressPayload
}

// NewCreateAddressContext parses the incoming request URL and body, performs validations and creates the
// context used by the address controller create action.
func NewCreateAddressContext(ctx context.Context, service *goa.Service) (*CreateAddressContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateAddressContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createAddressPayload is the address create action payload.
type createAddressPayload struct {
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

// Finalize sets the default values defined in the design.
func (payload *createAddressPayload) Finalize() {
	var defaultObject = "Address"
	if payload.Object == nil {
		payload.Object = &defaultObject
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createAddressPayload) Validate() (err error) {
	if payload.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "id"))
	}
	if payload.Object == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "object"))
	}

	if payload.ID != nil {
		if ok := goa.ValidatePattern(`^adr_`, *payload.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.id`, *payload.ID, `^adr_`))
		}
	}
	if payload.Object != nil {
		if ok := goa.ValidatePattern(`^Address$`, *payload.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.object`, *payload.Object, `^Address$`))
		}
	}
	return
}

// Publicize creates CreateAddressPayload from createAddressPayload
func (payload *createAddressPayload) Publicize() *CreateAddressPayload {
	var pub CreateAddressPayload
	if payload.CarrierFacility != nil {
		pub.CarrierFacility = payload.CarrierFacility
	}
	if payload.City != nil {
		pub.City = payload.City
	}
	if payload.Company != nil {
		pub.Company = payload.Company
	}
	if payload.Country != nil {
		pub.Country = payload.Country
	}
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.FederalTaxID != nil {
		pub.FederalTaxID = payload.FederalTaxID
	}
	if payload.ID != nil {
		pub.ID = *payload.ID
	}
	if payload.Mode != nil {
		pub.Mode = payload.Mode
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	if payload.Object != nil {
		pub.Object = *payload.Object
	}
	if payload.Phone != nil {
		pub.Phone = payload.Phone
	}
	if payload.Residential != nil {
		pub.Residential = payload.Residential
	}
	if payload.State != nil {
		pub.State = payload.State
	}
	if payload.StateTaxID != nil {
		pub.StateTaxID = payload.StateTaxID
	}
	if payload.Street1 != nil {
		pub.Street1 = payload.Street1
	}
	if payload.Street2 != nil {
		pub.Street2 = payload.Street2
	}
	if payload.Verifications != nil {
		pub.Verifications = payload.Verifications.Publicize()
	}
	if payload.Zip != nil {
		pub.Zip = payload.Zip
	}
	return &pub
}

// CreateAddressPayload is the address create action payload.
type CreateAddressPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *CreateAddressPayload) Validate() (err error) {
	if payload.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "id"))
	}
	if payload.Object == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "object"))
	}

	if ok := goa.ValidatePattern(`^adr_`, payload.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.id`, payload.ID, `^adr_`))
	}
	if ok := goa.ValidatePattern(`^Address$`, payload.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.object`, payload.Object, `^Address$`))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateAddressContext) OK(r *EasypostAddress) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.address+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowAddressContext provides the address show action context.
type ShowAddressContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowAddressContext parses the incoming request URL and body, performs validations and creates the
// context used by the address controller show action.
func NewShowAddressContext(ctx context.Context, service *goa.Service) (*ShowAddressContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowAddressContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
		if ok := goa.ValidatePattern(`^adr_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^adr_`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAddressContext) OK(r *EasypostAddress) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.address+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAddressContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateCarrierAccountContext provides the carrier_account create action context.
type CreateCarrierAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateCarrierAccountPayload
}

// NewCreateCarrierAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the carrier_account controller create action.
func NewCreateCarrierAccountContext(ctx context.Context, service *goa.Service) (*CreateCarrierAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateCarrierAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createCarrierAccountPayload is the carrier_account create action payload.
type createCarrierAccountPayload struct {
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

// Finalize sets the default values defined in the design.
func (payload *createCarrierAccountPayload) Finalize() {
	var defaultClone = false
	if payload.Clone == nil {
		payload.Clone = &defaultClone
	}
	if payload.Fields != nil {
		var defaultAutoLink = false
		if payload.Fields.AutoLink == nil {
			payload.Fields.AutoLink = &defaultAutoLink
		}
		var defaultCustomWorkflow = false
		if payload.Fields.CustomWorkflow == nil {
			payload.Fields.CustomWorkflow = &defaultCustomWorkflow
		}
	}
	var defaultObject = "CarrierAccount"
	if payload.Object == nil {
		payload.Object = &defaultObject
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createCarrierAccountPayload) Validate() (err error) {
	if payload.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	if payload.Object != nil {
		if ok := goa.ValidatePattern(`^CarrierAccount$`, *payload.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.object`, *payload.Object, `^CarrierAccount$`))
		}
	}
	return
}

// Publicize creates CreateCarrierAccountPayload from createCarrierAccountPayload
func (payload *createCarrierAccountPayload) Publicize() *CreateCarrierAccountPayload {
	var pub CreateCarrierAccountPayload
	if payload.Clone != nil {
		pub.Clone = *payload.Clone
	}
	if payload.CreatedAt != nil {
		pub.CreatedAt = payload.CreatedAt
	}
	if payload.Credentials != nil {
		pub.Credentials = payload.Credentials
	}
	if payload.Description != nil {
		pub.Description = payload.Description
	}
	if payload.Fields != nil {
		pub.Fields = payload.Fields.Publicize()
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.Object != nil {
		pub.Object = *payload.Object
	}
	if payload.Readable != nil {
		pub.Readable = payload.Readable
	}
	if payload.Reference != nil {
		pub.Reference = payload.Reference
	}
	if payload.TestCredentials != nil {
		pub.TestCredentials = payload.TestCredentials
	}
	if payload.Type != nil {
		pub.Type = *payload.Type
	}
	if payload.UpdatedAt != nil {
		pub.UpdatedAt = payload.UpdatedAt
	}
	return &pub
}

// CreateCarrierAccountPayload is the carrier_account create action payload.
type CreateCarrierAccountPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *CreateCarrierAccountPayload) Validate() (err error) {
	if payload.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	if ok := goa.ValidatePattern(`^CarrierAccount$`, payload.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.object`, payload.Object, `^CarrierAccount$`))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCarrierAccountContext) OK(r *EasypostCarrierAccounts) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.carrier_accounts+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// DeleteCarrierAccountContext provides the carrier_account delete action context.
type DeleteCarrierAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteCarrierAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the carrier_account controller delete action.
func NewDeleteCarrierAccountContext(ctx context.Context, service *goa.Service) (*DeleteCarrierAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteCarrierAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteCarrierAccountContext) OK(r *EasypostCarrierAccounts) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.carrier_accounts+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ListCarrierAccountContext provides the carrier_account list action context.
type ListCarrierAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListCarrierAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the carrier_account controller list action.
func NewListCarrierAccountContext(ctx context.Context, service *goa.Service) (*ListCarrierAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListCarrierAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCarrierAccountContext) OK(r []*EasypostCarrierAccounts) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.carrier_accounts+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowCarrierAccountContext provides the carrier_account show action context.
type ShowCarrierAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowCarrierAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the carrier_account controller show action.
func NewShowCarrierAccountContext(ctx context.Context, service *goa.Service) (*ShowCarrierAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowCarrierAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCarrierAccountContext) OK(r *EasypostCarrierAccounts) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.carrier_accounts+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCarrierAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateCarrierAccountContext provides the carrier_account update action context.
type UpdateCarrierAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *UpdateCarrierAccountPayload
}

// NewUpdateCarrierAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the carrier_account controller update action.
func NewUpdateCarrierAccountContext(ctx context.Context, service *goa.Service) (*UpdateCarrierAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateCarrierAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// updateCarrierAccountPayload is the carrier_account update action payload.
type updateCarrierAccountPayload struct {
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

// Finalize sets the default values defined in the design.
func (payload *updateCarrierAccountPayload) Finalize() {
	var defaultClone = false
	if payload.Clone == nil {
		payload.Clone = &defaultClone
	}
	if payload.Fields != nil {
		var defaultAutoLink = false
		if payload.Fields.AutoLink == nil {
			payload.Fields.AutoLink = &defaultAutoLink
		}
		var defaultCustomWorkflow = false
		if payload.Fields.CustomWorkflow == nil {
			payload.Fields.CustomWorkflow = &defaultCustomWorkflow
		}
	}
	var defaultObject = "CarrierAccount"
	if payload.Object == nil {
		payload.Object = &defaultObject
	}
}

// Validate runs the validation rules defined in the design.
func (payload *updateCarrierAccountPayload) Validate() (err error) {
	if payload.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	if payload.Object != nil {
		if ok := goa.ValidatePattern(`^CarrierAccount$`, *payload.Object); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.object`, *payload.Object, `^CarrierAccount$`))
		}
	}
	return
}

// Publicize creates UpdateCarrierAccountPayload from updateCarrierAccountPayload
func (payload *updateCarrierAccountPayload) Publicize() *UpdateCarrierAccountPayload {
	var pub UpdateCarrierAccountPayload
	if payload.Clone != nil {
		pub.Clone = *payload.Clone
	}
	if payload.CreatedAt != nil {
		pub.CreatedAt = payload.CreatedAt
	}
	if payload.Credentials != nil {
		pub.Credentials = payload.Credentials
	}
	if payload.Description != nil {
		pub.Description = payload.Description
	}
	if payload.Fields != nil {
		pub.Fields = payload.Fields.Publicize()
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.Object != nil {
		pub.Object = *payload.Object
	}
	if payload.Readable != nil {
		pub.Readable = payload.Readable
	}
	if payload.Reference != nil {
		pub.Reference = payload.Reference
	}
	if payload.TestCredentials != nil {
		pub.TestCredentials = payload.TestCredentials
	}
	if payload.Type != nil {
		pub.Type = *payload.Type
	}
	if payload.UpdatedAt != nil {
		pub.UpdatedAt = payload.UpdatedAt
	}
	return &pub
}

// UpdateCarrierAccountPayload is the carrier_account update action payload.
type UpdateCarrierAccountPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *UpdateCarrierAccountPayload) Validate() (err error) {
	if payload.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	if ok := goa.ValidatePattern(`^CarrierAccount$`, payload.Object); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.object`, payload.Object, `^CarrierAccount$`))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateCarrierAccountContext) OK(r *EasypostCarrierAccounts) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.carrier_accounts+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowCarrierTypesContext provides the carrier_types show action context.
type ShowCarrierTypesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowCarrierTypesContext parses the incoming request URL and body, performs validations and creates the
// context used by the carrier_types controller show action.
func NewShowCarrierTypesContext(ctx context.Context, service *goa.Service) (*ShowCarrierTypesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowCarrierTypesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCarrierTypesContext) OK(r EasypostCarrierTypesCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.carrier_types+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
