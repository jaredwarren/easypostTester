//************************************************************************//
// API "cellar": Application Contexts
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
		if ok := goa.ValidatePattern(`^ca_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^ca_`))
		}
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
		if ok := goa.ValidatePattern(`^ca_`, rctx.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`id`, rctx.ID, `^ca_`))
		}
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
