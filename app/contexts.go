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

// Validate runs the validation rules defined in the design.
func (payload *createAddressPayload) Validate() (err error) {
	if payload.Street1 == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "street1"))
	}
	if payload.City == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "city"))
	}
	if payload.State == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "state"))
	}
	if payload.Zip == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "zip"))
	}
	if payload.Country == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "country"))
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
		pub.City = *payload.City
	}
	if payload.Company != nil {
		pub.Company = payload.Company
	}
	if payload.Country != nil {
		pub.Country = *payload.Country
	}
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.FederalTaxID != nil {
		pub.FederalTaxID = payload.FederalTaxID
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	if payload.Phone != nil {
		pub.Phone = payload.Phone
	}
	if payload.Residential != nil {
		pub.Residential = payload.Residential
	}
	if payload.State != nil {
		pub.State = *payload.State
	}
	if payload.StateTaxID != nil {
		pub.StateTaxID = payload.StateTaxID
	}
	if payload.Street1 != nil {
		pub.Street1 = *payload.Street1
	}
	if payload.Street2 != nil {
		pub.Street2 = payload.Street2
	}
	if payload.Verify != nil {
		pub.Verify = payload.Verify
	}
	if payload.VerifyStrict != nil {
		pub.VerifyStrict = payload.VerifyStrict
	}
	if payload.Zip != nil {
		pub.Zip = *payload.Zip
	}
	return &pub
}

// CreateAddressPayload is the address create action payload.
type CreateAddressPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *CreateAddressPayload) Validate() (err error) {
	if payload.Street1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "street1"))
	}
	if payload.City == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "city"))
	}
	if payload.State == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "state"))
	}
	if payload.Zip == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "zip"))
	}
	if payload.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "country"))
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

// ListAPIKeyContext provides the api_key list action context.
type ListAPIKeyContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListAPIKeyContext parses the incoming request URL and body, performs validations and creates the
// context used by the api_key controller list action.
func NewListAPIKeyContext(ctx context.Context, service *goa.Service) (*ListAPIKeyContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListAPIKeyContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAPIKeyContext) OK(r []*EasypostAPIKeys) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.api_keys+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
}

// Validate runs the validation rules defined in the design.
func (payload *createCarrierAccountPayload) Validate() (err error) {
	if payload.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	return
}

// Publicize creates CreateCarrierAccountPayload from createCarrierAccountPayload
func (payload *createCarrierAccountPayload) Publicize() *CreateCarrierAccountPayload {
	var pub CreateCarrierAccountPayload
	if payload.Clone != nil {
		pub.Clone = *payload.Clone
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
	return &pub
}

// CreateCarrierAccountPayload is the carrier_account create action payload.
type CreateCarrierAccountPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *CreateCarrierAccountPayload) Validate() (err error) {
	if payload.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
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
}

// Validate runs the validation rules defined in the design.
func (payload *updateCarrierAccountPayload) Validate() (err error) {
	if payload.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	return
}

// Publicize creates UpdateCarrierAccountPayload from updateCarrierAccountPayload
func (payload *updateCarrierAccountPayload) Publicize() *UpdateCarrierAccountPayload {
	var pub UpdateCarrierAccountPayload
	if payload.Clone != nil {
		pub.Clone = *payload.Clone
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
	return &pub
}

// UpdateCarrierAccountPayload is the carrier_account update action payload.
type UpdateCarrierAccountPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *UpdateCarrierAccountPayload) Validate() (err error) {
	if payload.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
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

// CreateCustomsInfoContext provides the customs_info create action context.
type CreateCustomsInfoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateCustomsInfoPayload
}

// NewCreateCustomsInfoContext parses the incoming request URL and body, performs validations and creates the
// context used by the customs_info controller create action.
func NewCreateCustomsInfoContext(ctx context.Context, service *goa.Service) (*CreateCustomsInfoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateCustomsInfoContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createCustomsInfoPayload is the customs_info create action payload.
type createCustomsInfoPayload struct {
	ContentsType    *string              `form:"contents_type,omitempty" json:"contents_type,omitempty" xml:"contents_type,omitempty"`
	CustomsCertify  *bool                `form:"customs_certify,omitempty" json:"customs_certify,omitempty" xml:"customs_certify,omitempty"`
	CustomsItems    []*customItemPayload `form:"customs_items,omitempty" json:"customs_items,omitempty" xml:"customs_items,omitempty"`
	CutomsSigner    *string              `form:"cutoms_signer,omitempty" json:"cutoms_signer,omitempty" xml:"cutoms_signer,omitempty"`
	EelPfc          *string              `form:"eel_pfc,omitempty" json:"eel_pfc,omitempty" xml:"eel_pfc,omitempty"`
	RestrictionType *string              `form:"restriction_type,omitempty" json:"restriction_type,omitempty" xml:"restriction_type,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *createCustomsInfoPayload) Finalize() {
	for _, e := range payload.CustomsItems {
		var defaultCurrency = "USD"
		if e.Currency == nil {
			e.Currency = &defaultCurrency
		}
		var defaultOriginCountry = "US"
		if e.OriginCountry == nil {
			e.OriginCountry = &defaultOriginCountry
		}
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createCustomsInfoPayload) Validate() (err error) {
	if payload.CustomsCertify == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "customs_certify"))
	}
	if payload.CutomsSigner == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "cutoms_signer"))
	}
	if payload.ContentsType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "contents_type"))
	}
	if payload.RestrictionType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "restriction_type"))
	}
	if payload.EelPfc == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "eel_pfc"))
	}
	if payload.CustomsItems == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "customs_items"))
	}

	for _, e := range payload.CustomsItems {
		if e.Description == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "description"))
		}
		if e.Quantity == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "quantity"))
		}
		if e.Value == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "value"))
		}
		if e.Weight == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "weight"))
		}
		if e.OriginCountry == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "origin_country"))
		}

	}
	return
}

// Publicize creates CreateCustomsInfoPayload from createCustomsInfoPayload
func (payload *createCustomsInfoPayload) Publicize() *CreateCustomsInfoPayload {
	var pub CreateCustomsInfoPayload
	if payload.ContentsType != nil {
		pub.ContentsType = *payload.ContentsType
	}
	if payload.CustomsCertify != nil {
		pub.CustomsCertify = *payload.CustomsCertify
	}
	if payload.CustomsItems != nil {
		pub.CustomsItems = make([]*CustomItemPayload, len(payload.CustomsItems))
		for i2, elem2 := range payload.CustomsItems {
			pub.CustomsItems[i2] = elem2.Publicize()
		}
	}
	if payload.CutomsSigner != nil {
		pub.CutomsSigner = *payload.CutomsSigner
	}
	if payload.EelPfc != nil {
		pub.EelPfc = *payload.EelPfc
	}
	if payload.RestrictionType != nil {
		pub.RestrictionType = *payload.RestrictionType
	}
	return &pub
}

// CreateCustomsInfoPayload is the customs_info create action payload.
type CreateCustomsInfoPayload struct {
	ContentsType    string               `form:"contents_type" json:"contents_type" xml:"contents_type"`
	CustomsCertify  bool                 `form:"customs_certify" json:"customs_certify" xml:"customs_certify"`
	CustomsItems    []*CustomItemPayload `form:"customs_items" json:"customs_items" xml:"customs_items"`
	CutomsSigner    string               `form:"cutoms_signer" json:"cutoms_signer" xml:"cutoms_signer"`
	EelPfc          string               `form:"eel_pfc" json:"eel_pfc" xml:"eel_pfc"`
	RestrictionType string               `form:"restriction_type" json:"restriction_type" xml:"restriction_type"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateCustomsInfoPayload) Validate() (err error) {
	if payload.CutomsSigner == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "cutoms_signer"))
	}
	if payload.ContentsType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "contents_type"))
	}
	if payload.RestrictionType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "restriction_type"))
	}
	if payload.EelPfc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "eel_pfc"))
	}
	if payload.CustomsItems == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "customs_items"))
	}

	for _, e := range payload.CustomsItems {
		if e.Description == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "description"))
		}
		if e.OriginCountry == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`raw.customs_items[*]`, "origin_country"))
		}

	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCustomsInfoContext) OK(r *EasypostCustomsinfo) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.customsinfo+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateCustomsItemContext provides the customs_item create action context.
type CreateCustomsItemContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateCustomsItemPayload
}

// NewCreateCustomsItemContext parses the incoming request URL and body, performs validations and creates the
// context used by the customs_item controller create action.
func NewCreateCustomsItemContext(ctx context.Context, service *goa.Service) (*CreateCustomsItemContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateCustomsItemContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createCustomsItemPayload is the customs_item create action payload.
type createCustomsItemPayload struct {
	// 3 char currency code, default USD
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Required, description of item being shipped
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Harmonized Tariff Schedule, e.g. "6109.10.0012" for Men's T-shirts
	HsTariffNumber *string `form:"hs_tariff_number,omitempty" json:"hs_tariff_number,omitempty" xml:"hs_tariff_number,omitempty"`
	// Required, 2 char country code
	OriginCountry *string `form:"origin_country,omitempty" json:"origin_country,omitempty" xml:"origin_country,omitempty"`
	// Required, greater than zero
	Quantity *float64 `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Required, greater than zero, total value (unit value * quantity)
	Value *float64 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Required, greater than zero, total weight (unit weight * quantity)
	Weight *float64 `form:"weight,omitempty" json:"weight,omitempty" xml:"weight,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *createCustomsItemPayload) Finalize() {
	var defaultCurrency = "USD"
	if payload.Currency == nil {
		payload.Currency = &defaultCurrency
	}
	var defaultOriginCountry = "US"
	if payload.OriginCountry == nil {
		payload.OriginCountry = &defaultOriginCountry
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createCustomsItemPayload) Validate() (err error) {
	if payload.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "description"))
	}
	if payload.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "quantity"))
	}
	if payload.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "value"))
	}
	if payload.Weight == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "weight"))
	}
	if payload.OriginCountry == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "origin_country"))
	}

	return
}

// Publicize creates CreateCustomsItemPayload from createCustomsItemPayload
func (payload *createCustomsItemPayload) Publicize() *CreateCustomsItemPayload {
	var pub CreateCustomsItemPayload
	if payload.Currency != nil {
		pub.Currency = *payload.Currency
	}
	if payload.Description != nil {
		pub.Description = *payload.Description
	}
	if payload.HsTariffNumber != nil {
		pub.HsTariffNumber = payload.HsTariffNumber
	}
	if payload.OriginCountry != nil {
		pub.OriginCountry = *payload.OriginCountry
	}
	if payload.Quantity != nil {
		pub.Quantity = *payload.Quantity
	}
	if payload.Value != nil {
		pub.Value = *payload.Value
	}
	if payload.Weight != nil {
		pub.Weight = *payload.Weight
	}
	return &pub
}

// CreateCustomsItemPayload is the customs_item create action payload.
type CreateCustomsItemPayload struct {
	// 3 char currency code, default USD
	Currency string `form:"currency" json:"currency" xml:"currency"`
	// Required, description of item being shipped
	Description string `form:"description" json:"description" xml:"description"`
	// Harmonized Tariff Schedule, e.g. "6109.10.0012" for Men's T-shirts
	HsTariffNumber *string `form:"hs_tariff_number,omitempty" json:"hs_tariff_number,omitempty" xml:"hs_tariff_number,omitempty"`
	// Required, 2 char country code
	OriginCountry string `form:"origin_country" json:"origin_country" xml:"origin_country"`
	// Required, greater than zero
	Quantity float64 `form:"quantity" json:"quantity" xml:"quantity"`
	// Required, greater than zero, total value (unit value * quantity)
	Value float64 `form:"value" json:"value" xml:"value"`
	// Required, greater than zero, total weight (unit weight * quantity)
	Weight float64 `form:"weight" json:"weight" xml:"weight"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateCustomsItemPayload) Validate() (err error) {
	if payload.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "description"))
	}
	if payload.OriginCountry == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "origin_country"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCustomsItemContext) OK(r *EasypostCustomitem) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.customitem+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateInsuranceContext provides the insurance create action context.
type CreateInsuranceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateInsurancePayload
}

// NewCreateInsuranceContext parses the incoming request URL and body, performs validations and creates the
// context used by the insurance controller create action.
func NewCreateInsuranceContext(ctx context.Context, service *goa.Service) (*CreateInsuranceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateInsuranceContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createInsurancePayload is the insurance create action payload.
type createInsurancePayload struct {
	// The USD value of contents you would like to insure. Currently the maximum is between $5000 and $10000 depending on insurer
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The actual origin of the package to be insured
	FromAddress *addressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// Optional. A unique value that may be used in place of ID when doing Retrieve calls for this insurance
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// The actual destination of the package to be insured
	ToAddress *addressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The tracking code associated with the non-EasyPost-purchased package you'd like to insure
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createInsurancePayload) Validate() (err error) {
	if payload.ToAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "to_address"))
	}
	if payload.FromAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "from_address"))
	}
	if payload.TrackingCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "tracking_code"))
	}
	if payload.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}

	if payload.FromAddress != nil {
		if err2 := payload.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.ToAddress != nil {
		if err2 := payload.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates CreateInsurancePayload from createInsurancePayload
func (payload *createInsurancePayload) Publicize() *CreateInsurancePayload {
	var pub CreateInsurancePayload
	if payload.Amount != nil {
		pub.Amount = *payload.Amount
	}
	if payload.Carrier != nil {
		pub.Carrier = payload.Carrier
	}
	if payload.FromAddress != nil {
		pub.FromAddress = payload.FromAddress.Publicize()
	}
	if payload.Reference != nil {
		pub.Reference = payload.Reference
	}
	if payload.ToAddress != nil {
		pub.ToAddress = payload.ToAddress.Publicize()
	}
	if payload.TrackingCode != nil {
		pub.TrackingCode = *payload.TrackingCode
	}
	return &pub
}

// CreateInsurancePayload is the insurance create action payload.
type CreateInsurancePayload struct {
	// The USD value of contents you would like to insure. Currently the maximum is between $5000 and $10000 depending on insurer
	Amount string `form:"amount" json:"amount" xml:"amount"`
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The actual origin of the package to be insured
	FromAddress *AddressPayload `form:"from_address" json:"from_address" xml:"from_address"`
	// Optional. A unique value that may be used in place of ID when doing Retrieve calls for this insurance
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// The actual destination of the package to be insured
	ToAddress *AddressPayload `form:"to_address" json:"to_address" xml:"to_address"`
	// The tracking code associated with the non-EasyPost-purchased package you'd like to insure
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateInsurancePayload) Validate() (err error) {
	if payload.ToAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "to_address"))
	}
	if payload.FromAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "from_address"))
	}
	if payload.TrackingCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "tracking_code"))
	}
	if payload.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}

	if payload.FromAddress != nil {
		if err2 := payload.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.ToAddress != nil {
		if err2 := payload.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateInsuranceContext) OK(r *EasypostTracker) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.tracker+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ListInsuranceContext provides the insurance list action context.
type ListInsuranceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ListInsurancePayload
}

// NewListInsuranceContext parses the incoming request URL and body, performs validations and creates the
// context used by the insurance controller list action.
func NewListInsuranceContext(ctx context.Context, service *goa.Service) (*ListInsuranceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListInsuranceContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// listInsurancePayload is the insurance list action payload.
type listInsurancePayload struct {
	// Optional pagination parameter. Only records created after the given ID will be included. May not be used with before_id
	AfterID *addressPayload `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only records created before the given ID will be included. May not be used with after_id
	BeforeID *addressPayload `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only return records created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of records to return on each page. The maximum value is 100, and default is 20.
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Only return records created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *listInsurancePayload) Finalize() {
	var defaultPageSize = 20
	if payload.PageSize == nil {
		payload.PageSize = &defaultPageSize
	}
}

// Validate runs the validation rules defined in the design.
func (payload *listInsurancePayload) Validate() (err error) {
	if payload.AfterID != nil {
		if err2 := payload.AfterID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.BeforeID != nil {
		if err2 := payload.BeforeID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.PageSize != nil {
		if *payload.PageSize < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, *payload.PageSize, 1, true))
		}
	}
	if payload.PageSize != nil {
		if *payload.PageSize > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, *payload.PageSize, 100, false))
		}
	}
	return
}

// Publicize creates ListInsurancePayload from listInsurancePayload
func (payload *listInsurancePayload) Publicize() *ListInsurancePayload {
	var pub ListInsurancePayload
	if payload.AfterID != nil {
		pub.AfterID = payload.AfterID.Publicize()
	}
	if payload.BeforeID != nil {
		pub.BeforeID = payload.BeforeID.Publicize()
	}
	if payload.EndDatetime != nil {
		pub.EndDatetime = payload.EndDatetime
	}
	if payload.PageSize != nil {
		pub.PageSize = *payload.PageSize
	}
	if payload.StartDatetime != nil {
		pub.StartDatetime = payload.StartDatetime
	}
	return &pub
}

// ListInsurancePayload is the insurance list action payload.
type ListInsurancePayload struct {
	// Optional pagination parameter. Only records created after the given ID will be included. May not be used with before_id
	AfterID *AddressPayload `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only records created before the given ID will be included. May not be used with after_id
	BeforeID *AddressPayload `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only return records created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of records to return on each page. The maximum value is 100, and default is 20.
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
	// Only return records created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *ListInsurancePayload) Validate() (err error) {
	if payload.AfterID != nil {
		if err2 := payload.AfterID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.BeforeID != nil {
		if err2 := payload.BeforeID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.PageSize < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, payload.PageSize, 1, true))
	}
	if payload.PageSize > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, payload.PageSize, 100, false))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *ListInsuranceContext) OK(r *EasypostTrackers) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.trackers+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListInsuranceContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowInsuranceContext provides the insurance show action context.
type ShowInsuranceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowInsuranceContext parses the incoming request URL and body, performs validations and creates the
// context used by the insurance controller show action.
func NewShowInsuranceContext(ctx context.Context, service *goa.Service) (*ShowInsuranceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowInsuranceContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
		if ok := goa.ValidatePattern(`^trk_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^trk_`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowInsuranceContext) OK(r *EasypostTracker) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.tracker+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowInsuranceContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateParcelContext provides the parcel create action context.
type CreateParcelContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateParcelPayload
}

// NewCreateParcelContext parses the incoming request URL and body, performs validations and creates the
// context used by the parcel controller create action.
func NewCreateParcelContext(ctx context.Context, service *goa.Service) (*CreateParcelContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateParcelContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createParcelPayload is the parcel create action payload.
type createParcelPayload struct {
	// Required if predefined_package is empty. float (inches)
	Height *float64 `form:"height,omitempty" json:"height,omitempty" xml:"height,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Length *float64 `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Optional, one of our predefined_packages
	PredefinedPackage *string `form:"predefined_package,omitempty" json:"predefined_package,omitempty" xml:"predefined_package,omitempty"`
	// Always required. float(oz)
	Weight *float64 `form:"weight,omitempty" json:"weight,omitempty" xml:"weight,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Width *float64 `form:"width,omitempty" json:"width,omitempty" xml:"width,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createParcelPayload) Validate() (err error) {
	if payload.Weight == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "weight"))
	}

	return
}

// Publicize creates CreateParcelPayload from createParcelPayload
func (payload *createParcelPayload) Publicize() *CreateParcelPayload {
	var pub CreateParcelPayload
	if payload.Height != nil {
		pub.Height = payload.Height
	}
	if payload.Length != nil {
		pub.Length = payload.Length
	}
	if payload.PredefinedPackage != nil {
		pub.PredefinedPackage = payload.PredefinedPackage
	}
	if payload.Weight != nil {
		pub.Weight = *payload.Weight
	}
	if payload.Width != nil {
		pub.Width = payload.Width
	}
	return &pub
}

// CreateParcelPayload is the parcel create action payload.
type CreateParcelPayload struct {
	// Required if predefined_package is empty. float (inches)
	Height *float64 `form:"height,omitempty" json:"height,omitempty" xml:"height,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Length *float64 `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Optional, one of our predefined_packages
	PredefinedPackage *string `form:"predefined_package,omitempty" json:"predefined_package,omitempty" xml:"predefined_package,omitempty"`
	// Always required. float(oz)
	Weight float64 `form:"weight" json:"weight" xml:"weight"`
	// Required if predefined_package is empty. float (inches)
	Width *float64 `form:"width,omitempty" json:"width,omitempty" xml:"width,omitempty"`
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateParcelContext) OK(r *EasypostParcel) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.parcel+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowParcelContext provides the parcel show action context.
type ShowParcelContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowParcelContext parses the incoming request URL and body, performs validations and creates the
// context used by the parcel controller show action.
func NewShowParcelContext(ctx context.Context, service *goa.Service) (*ShowParcelContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowParcelContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
		if ok := goa.ValidatePattern(`^prcl_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^prcl_`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowParcelContext) OK(r *EasypostParcel) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.parcel+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowParcelContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ConvertLabelShipmentContext provides the shipment convertLabel action context.
type ConvertLabelShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *ConvertLabelShipmentPayload
}

// NewConvertLabelShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller convertLabel action.
func NewConvertLabelShipmentContext(ctx context.Context, service *goa.Service) (*ConvertLabelShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ConvertLabelShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// convertLabelShipmentPayload is the shipment convertLabel action payload.
type convertLabelShipmentPayload struct {
	// Selected rate
	FileFormat *string `form:"file_format,omitempty" json:"file_format,omitempty" xml:"file_format,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *convertLabelShipmentPayload) Validate() (err error) {
	if payload.FileFormat == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "file_format"))
	}

	if payload.FileFormat != nil {
		if !(*payload.FileFormat == "png" || *payload.FileFormat == "zpl" || *payload.FileFormat == "epl2" || *payload.FileFormat == "pdf") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.file_format`, *payload.FileFormat, []interface{}{"png", "zpl", "epl2", "pdf"}))
		}
	}
	return
}

// Publicize creates ConvertLabelShipmentPayload from convertLabelShipmentPayload
func (payload *convertLabelShipmentPayload) Publicize() *ConvertLabelShipmentPayload {
	var pub ConvertLabelShipmentPayload
	if payload.FileFormat != nil {
		pub.FileFormat = *payload.FileFormat
	}
	return &pub
}

// ConvertLabelShipmentPayload is the shipment convertLabel action payload.
type ConvertLabelShipmentPayload struct {
	// Selected rate
	FileFormat string `form:"file_format" json:"file_format" xml:"file_format"`
}

// Validate runs the validation rules defined in the design.
func (payload *ConvertLabelShipmentPayload) Validate() (err error) {
	if payload.FileFormat == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "file_format"))
	}

	if !(payload.FileFormat == "png" || payload.FileFormat == "zpl" || payload.FileFormat == "epl2" || payload.FileFormat == "pdf") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.file_format`, payload.FileFormat, []interface{}{"png", "zpl", "epl2", "pdf"}))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *ConvertLabelShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateShipmentContext provides the shipment create action context.
type CreateShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateShipmentPayload
}

// NewCreateShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller create action.
func NewCreateShipmentContext(ctx context.Context, service *goa.Service) (*CreateShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createShipmentPayload is the shipment create action payload.
type createShipmentPayload struct {
	// The ID of the batch that contains this shipment, if any
	BatchID *string `form:"batch_id,omitempty" json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The buyer's address, defaults to to_address
	BuyerAddress *addressPayload `form:"buyer_address,omitempty" json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// Information for the processing of customs
	CustomsInfo *customsInfoPayload `form:"customs_info,omitempty" json:"customs_info,omitempty" xml:"customs_info,omitempty"`
	// All associated Form objects
	Forms []interface{} `form:"forms,omitempty" json:"forms,omitempty" xml:"forms,omitempty"`
	// The origin address
	FromAddress *addressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// The associated Insurance object
	Insurance *insurancePayload `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Set true to create as a return, discussed in more depth below
	IsReturn *bool `form:"is_return,omitempty" json:"is_return,omitempty" xml:"is_return,omitempty"`
	// All of the options passed to the shipment, discussed in more depth below
	Options *optionsPayload `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// The dimensions and weight of the package
	Parcel *parcelPayload `form:"parcel,omitempty" json:"parcel,omitempty" xml:"parcel,omitempty"`
	// The shipper's address, defaults to from_address
	ReturnAddress *addressPayload `form:"return_address,omitempty" json:"return_address,omitempty" xml:"return_address,omitempty"`
	// The destination address
	ToAddress *addressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The USPS zone of the shipment, if purchased with USPS
	UspsZone *string `form:"usps_zone,omitempty" json:"usps_zone,omitempty" xml:"usps_zone,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *createShipmentPayload) Finalize() {
	if payload.CustomsInfo != nil {
		for _, e := range payload.CustomsInfo.CustomsItems {
			var defaultCurrency = "USD"
			if e.Currency == nil {
				e.Currency = &defaultCurrency
			}
			var defaultOriginCountry = "US"
			if e.OriginCountry == nil {
				e.OriginCountry = &defaultOriginCountry
			}
		}
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createShipmentPayload) Validate() (err error) {
	if payload.BuyerAddress != nil {
		if err2 := payload.BuyerAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.CustomsInfo != nil {
		if err2 := payload.CustomsInfo.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.FromAddress != nil {
		if err2 := payload.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Insurance != nil {
		if err2 := payload.Insurance.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Options != nil {
		if err2 := payload.Options.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Parcel != nil {
		if err2 := payload.Parcel.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.ReturnAddress != nil {
		if err2 := payload.ReturnAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.ToAddress != nil {
		if err2 := payload.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates CreateShipmentPayload from createShipmentPayload
func (payload *createShipmentPayload) Publicize() *CreateShipmentPayload {
	var pub CreateShipmentPayload
	if payload.BatchID != nil {
		pub.BatchID = payload.BatchID
	}
	if payload.BuyerAddress != nil {
		pub.BuyerAddress = payload.BuyerAddress.Publicize()
	}
	if payload.CustomsInfo != nil {
		pub.CustomsInfo = payload.CustomsInfo.Publicize()
	}
	if payload.Forms != nil {
		pub.Forms = payload.Forms
	}
	if payload.FromAddress != nil {
		pub.FromAddress = payload.FromAddress.Publicize()
	}
	if payload.Insurance != nil {
		pub.Insurance = payload.Insurance.Publicize()
	}
	if payload.IsReturn != nil {
		pub.IsReturn = payload.IsReturn
	}
	if payload.Options != nil {
		pub.Options = payload.Options.Publicize()
	}
	if payload.Parcel != nil {
		pub.Parcel = payload.Parcel.Publicize()
	}
	if payload.ReturnAddress != nil {
		pub.ReturnAddress = payload.ReturnAddress.Publicize()
	}
	if payload.ToAddress != nil {
		pub.ToAddress = payload.ToAddress.Publicize()
	}
	if payload.UspsZone != nil {
		pub.UspsZone = payload.UspsZone
	}
	return &pub
}

// CreateShipmentPayload is the shipment create action payload.
type CreateShipmentPayload struct {
	// The ID of the batch that contains this shipment, if any
	BatchID *string `form:"batch_id,omitempty" json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The buyer's address, defaults to to_address
	BuyerAddress *AddressPayload `form:"buyer_address,omitempty" json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// Information for the processing of customs
	CustomsInfo *CustomsInfoPayload `form:"customs_info,omitempty" json:"customs_info,omitempty" xml:"customs_info,omitempty"`
	// All associated Form objects
	Forms []interface{} `form:"forms,omitempty" json:"forms,omitempty" xml:"forms,omitempty"`
	// The origin address
	FromAddress *AddressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// The associated Insurance object
	Insurance *InsurancePayload `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Set true to create as a return, discussed in more depth below
	IsReturn *bool `form:"is_return,omitempty" json:"is_return,omitempty" xml:"is_return,omitempty"`
	// All of the options passed to the shipment, discussed in more depth below
	Options *OptionsPayload `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// The dimensions and weight of the package
	Parcel *ParcelPayload `form:"parcel,omitempty" json:"parcel,omitempty" xml:"parcel,omitempty"`
	// The shipper's address, defaults to from_address
	ReturnAddress *AddressPayload `form:"return_address,omitempty" json:"return_address,omitempty" xml:"return_address,omitempty"`
	// The destination address
	ToAddress *AddressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The USPS zone of the shipment, if purchased with USPS
	UspsZone *string `form:"usps_zone,omitempty" json:"usps_zone,omitempty" xml:"usps_zone,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateShipmentPayload) Validate() (err error) {
	if payload.BuyerAddress != nil {
		if err2 := payload.BuyerAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.CustomsInfo != nil {
		if err2 := payload.CustomsInfo.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.FromAddress != nil {
		if err2 := payload.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Insurance != nil {
		if err2 := payload.Insurance.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Options != nil {
		if err2 := payload.Options.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.ReturnAddress != nil {
		if err2 := payload.ReturnAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.ToAddress != nil {
		if err2 := payload.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// InsureShipmentContext provides the shipment insure action context.
type InsureShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *InsureShipmentPayload
}

// NewInsureShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller insure action.
func NewInsureShipmentContext(ctx context.Context, service *goa.Service) (*InsureShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := InsureShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// insureShipmentPayload is the shipment insure action payload.
type insureShipmentPayload struct {
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *insureShipmentPayload) Validate() (err error) {
	if payload.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}

	return
}

// Publicize creates InsureShipmentPayload from insureShipmentPayload
func (payload *insureShipmentPayload) Publicize() *InsureShipmentPayload {
	var pub InsureShipmentPayload
	if payload.Amount != nil {
		pub.Amount = *payload.Amount
	}
	return &pub
}

// InsureShipmentPayload is the shipment insure action payload.
type InsureShipmentPayload struct {
	Amount string `form:"amount" json:"amount" xml:"amount"`
}

// Validate runs the validation rules defined in the design.
func (payload *InsureShipmentPayload) Validate() (err error) {
	if payload.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *InsureShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// PruchaseShipmentContext provides the shipment pruchase action context.
type PruchaseShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *PruchaseShipmentPayload
}

// NewPruchaseShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller pruchase action.
func NewPruchaseShipmentContext(ctx context.Context, service *goa.Service) (*PruchaseShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := PruchaseShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// pruchaseShipmentPayload is the shipment pruchase action payload.
type pruchaseShipmentPayload struct {
	// Additionally, insurance may be added during the purchase. To specify an amount to insure, pass the insurance attribute as a string. The currency of all insurance is USD.
	Insurance *string `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Selected rate
	Rate *ratePayload `form:"rate,omitempty" json:"rate,omitempty" xml:"rate,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *pruchaseShipmentPayload) Validate() (err error) {
	if payload.Rate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "rate"))
	}

	if payload.Rate != nil {
		if err2 := payload.Rate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates PruchaseShipmentPayload from pruchaseShipmentPayload
func (payload *pruchaseShipmentPayload) Publicize() *PruchaseShipmentPayload {
	var pub PruchaseShipmentPayload
	if payload.Insurance != nil {
		pub.Insurance = payload.Insurance
	}
	if payload.Rate != nil {
		pub.Rate = payload.Rate.Publicize()
	}
	return &pub
}

// PruchaseShipmentPayload is the shipment pruchase action payload.
type PruchaseShipmentPayload struct {
	// Additionally, insurance may be added during the purchase. To specify an amount to insure, pass the insurance attribute as a string. The currency of all insurance is USD.
	Insurance *string `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Selected rate
	Rate *RatePayload `form:"rate" json:"rate" xml:"rate"`
}

// Validate runs the validation rules defined in the design.
func (payload *PruchaseShipmentPayload) Validate() (err error) {
	if payload.Rate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "rate"))
	}

	if payload.Rate != nil {
		if err2 := payload.Rate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *PruchaseShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// RatesShipmentContext provides the shipment rates action context.
type RatesShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewRatesShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller rates action.
func NewRatesShipmentContext(ctx context.Context, service *goa.Service) (*RatesShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RatesShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RatesShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// RefundShipmentContext provides the shipment refund action context.
type RefundShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewRefundShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller refund action.
func NewRefundShipmentContext(ctx context.Context, service *goa.Service) (*RefundShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RefundShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RefundShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowShipmentContext provides the shipment show action context.
type ShowShipmentContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowShipmentContext parses the incoming request URL and body, performs validations and creates the
// context used by the shipment controller show action.
func NewShowShipmentContext(ctx context.Context, service *goa.Service) (*ShowShipmentContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowShipmentContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
		if ok := goa.ValidatePattern(`^shp_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^shp_`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowShipmentContext) OK(r *EasypostShipment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.shipment+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowShipmentContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateTrackerContext provides the tracker create action context.
type CreateTrackerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateTrackerPayload
}

// NewCreateTrackerContext parses the incoming request URL and body, performs validations and creates the
// context used by the tracker controller create action.
func NewCreateTrackerContext(ctx context.Context, service *goa.Service) (*CreateTrackerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateTrackerContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createTrackerPayload is the tracker create action payload.
type createTrackerPayload struct {
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The tracking code associated with the package you'd like to track
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createTrackerPayload) Validate() (err error) {
	if payload.TrackingCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "tracking_code"))
	}

	return
}

// Publicize creates CreateTrackerPayload from createTrackerPayload
func (payload *createTrackerPayload) Publicize() *CreateTrackerPayload {
	var pub CreateTrackerPayload
	if payload.Carrier != nil {
		pub.Carrier = payload.Carrier
	}
	if payload.TrackingCode != nil {
		pub.TrackingCode = *payload.TrackingCode
	}
	return &pub
}

// CreateTrackerPayload is the tracker create action payload.
type CreateTrackerPayload struct {
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The tracking code associated with the package you'd like to track
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateTrackerPayload) Validate() (err error) {
	if payload.TrackingCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "tracking_code"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateTrackerContext) OK(r *EasypostTracker) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.tracker+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ListTrackerContext provides the tracker list action context.
type ListTrackerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ListTrackerPayload
}

// NewListTrackerContext parses the incoming request URL and body, performs validations and creates the
// context used by the tracker controller list action.
func NewListTrackerContext(ctx context.Context, service *goa.Service) (*ListTrackerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListTrackerContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// listTrackerPayload is the tracker list action payload.
type listTrackerPayload struct {
	// Optional pagination parameter. Only trackers created after the given ID will be included. May not be used with before_id
	AfterID *string `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only trackers created before the given ID will be included. May not be used with after_id
	BeforeID *string `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only returns Trackers with the given carrier value
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// Only return Trackers created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of Trackers to return on each page. The maximum value is 100
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Only return Trackers created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
	// Only returns Trackers with the given tracking_code. Useful for retrieving an individual Tracker by tracking_code rather than by ID
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *listTrackerPayload) Finalize() {
	var defaultPageSize = 1
	if payload.PageSize == nil {
		payload.PageSize = &defaultPageSize
	}
}

// Validate runs the validation rules defined in the design.
func (payload *listTrackerPayload) Validate() (err error) {
	if payload.TrackingCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "tracking_code"))
	}

	if payload.PageSize != nil {
		if *payload.PageSize < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, *payload.PageSize, 1, true))
		}
	}
	if payload.PageSize != nil {
		if *payload.PageSize > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, *payload.PageSize, 100, false))
		}
	}
	return
}

// Publicize creates ListTrackerPayload from listTrackerPayload
func (payload *listTrackerPayload) Publicize() *ListTrackerPayload {
	var pub ListTrackerPayload
	if payload.AfterID != nil {
		pub.AfterID = payload.AfterID
	}
	if payload.BeforeID != nil {
		pub.BeforeID = payload.BeforeID
	}
	if payload.Carrier != nil {
		pub.Carrier = payload.Carrier
	}
	if payload.EndDatetime != nil {
		pub.EndDatetime = payload.EndDatetime
	}
	if payload.PageSize != nil {
		pub.PageSize = *payload.PageSize
	}
	if payload.StartDatetime != nil {
		pub.StartDatetime = payload.StartDatetime
	}
	if payload.TrackingCode != nil {
		pub.TrackingCode = *payload.TrackingCode
	}
	return &pub
}

// ListTrackerPayload is the tracker list action payload.
type ListTrackerPayload struct {
	// Optional pagination parameter. Only trackers created after the given ID will be included. May not be used with before_id
	AfterID *string `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only trackers created before the given ID will be included. May not be used with after_id
	BeforeID *string `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only returns Trackers with the given carrier value
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// Only return Trackers created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of Trackers to return on each page. The maximum value is 100
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
	// Only return Trackers created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
	// Only returns Trackers with the given tracking_code. Useful for retrieving an individual Tracker by tracking_code rather than by ID
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// Validate runs the validation rules defined in the design.
func (payload *ListTrackerPayload) Validate() (err error) {
	if payload.TrackingCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "tracking_code"))
	}

	if payload.PageSize < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, payload.PageSize, 1, true))
	}
	if payload.PageSize > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.page_size`, payload.PageSize, 100, false))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *ListTrackerContext) OK(r *EasypostTrackers) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.trackers+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListTrackerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowTrackerContext provides the tracker show action context.
type ShowTrackerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowTrackerContext parses the incoming request URL and body, performs validations and creates the
// context used by the tracker controller show action.
func NewShowTrackerContext(ctx context.Context, service *goa.Service) (*ShowTrackerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowTrackerContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
		if ok := goa.ValidatePattern(`^trk_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^trk_`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTrackerContext) OK(r *EasypostTracker) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.tracker+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowTrackerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createUserPayload is the user create action payload.
type createUserPayload struct {
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First and last name required
	Name                 *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password             *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	PasswordConfirmation *string `form:"password_confirmation,omitempty" json:"password_confirmation,omitempty" xml:"password_confirmation,omitempty"`
	PhoneNumber          *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.PasswordConfirmation == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password_confirmation"))
	}

	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.PasswordConfirmation != nil {
		pub.PasswordConfirmation = *payload.PasswordConfirmation
	}
	if payload.PhoneNumber != nil {
		pub.PhoneNumber = payload.PhoneNumber
	}
	return &pub
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Email string `form:"email" json:"email" xml:"email"`
	// First and last name required
	Name                 string  `form:"name" json:"name" xml:"name"`
	Password             string  `form:"password" json:"password" xml:"password"`
	PasswordConfirmation string  `form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation"`
	PhoneNumber          *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.PasswordConfirmation == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password_confirmation"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateUserContext) OK(r *EasypostUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
		if ok := goa.ValidatePattern(`^user_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^user_`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *EasypostUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// updateUserPayload is the user update action payload.
type updateUserPayload struct {
	CurrentPassword *string `form:"current_password,omitempty" json:"current_password,omitempty" xml:"current_password,omitempty"`
	Email           *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First and last name required
	Name                    *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password                *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	PasswordConfirmation    *string `form:"password_confirmation,omitempty" json:"password_confirmation,omitempty" xml:"password_confirmation,omitempty"`
	PhoneNumber             *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	RechargeAmount          *string `form:"recharge_amount,omitempty" json:"recharge_amount,omitempty" xml:"recharge_amount,omitempty"`
	RechargeThreshold       *string `form:"recharge_threshold,omitempty" json:"recharge_threshold,omitempty" xml:"recharge_threshold,omitempty"`
	SecondaryRechargeAmount *string `form:"secondary_recharge_amount,omitempty" json:"secondary_recharge_amount,omitempty" xml:"secondary_recharge_amount,omitempty"`
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (payload *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if payload.CurrentPassword != nil {
		pub.CurrentPassword = payload.CurrentPassword
	}
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	if payload.Password != nil {
		pub.Password = payload.Password
	}
	if payload.PasswordConfirmation != nil {
		pub.PasswordConfirmation = payload.PasswordConfirmation
	}
	if payload.PhoneNumber != nil {
		pub.PhoneNumber = payload.PhoneNumber
	}
	if payload.RechargeAmount != nil {
		pub.RechargeAmount = payload.RechargeAmount
	}
	if payload.RechargeThreshold != nil {
		pub.RechargeThreshold = payload.RechargeThreshold
	}
	if payload.SecondaryRechargeAmount != nil {
		pub.SecondaryRechargeAmount = payload.SecondaryRechargeAmount
	}
	return &pub
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	CurrentPassword *string `form:"current_password,omitempty" json:"current_password,omitempty" xml:"current_password,omitempty"`
	Email           *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First and last name required
	Name                    *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password                *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	PasswordConfirmation    *string `form:"password_confirmation,omitempty" json:"password_confirmation,omitempty" xml:"password_confirmation,omitempty"`
	PhoneNumber             *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	RechargeAmount          *string `form:"recharge_amount,omitempty" json:"recharge_amount,omitempty" xml:"recharge_amount,omitempty"`
	RechargeThreshold       *string `form:"recharge_threshold,omitempty" json:"recharge_threshold,omitempty" xml:"recharge_threshold,omitempty"`
	SecondaryRechargeAmount *string `form:"secondary_recharge_amount,omitempty" json:"secondary_recharge_amount,omitempty" xml:"secondary_recharge_amount,omitempty"`
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *EasypostUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/easypost.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
