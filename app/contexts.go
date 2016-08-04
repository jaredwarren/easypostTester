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
