//************************************************************************//
// API "cellar": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CarrierAccountController is the controller interface for the CarrierAccount actions.
type CarrierAccountController interface {
	goa.Muxer
	Careate(*CareateCarrierAccountContext) error
	Delete(*DeleteCarrierAccountContext) error
	List(*ListCarrierAccountContext) error
	Show(*ShowCarrierAccountContext) error
	Update(*UpdateCarrierAccountContext) error
}

// MountCarrierAccountController "mounts" a CarrierAccount resource controller on the given service.
func MountCarrierAccountController(service *goa.Service, ctrl CarrierAccountController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCareateCarrierAccountContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CareateCarrierAccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Careate(rctx)
	}
	service.Mux.Handle("POST", "/v2/carrier_accounts", ctrl.MuxHandler("Careate", h, unmarshalCareateCarrierAccountPayload))
	service.LogInfo("mount", "ctrl", "CarrierAccount", "action", "Careate", "route", "POST /v2/carrier_accounts")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCarrierAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/v2/carrier_accounts/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "CarrierAccount", "action", "Delete", "route", "DELETE /v2/carrier_accounts/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCarrierAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v2/carrier_accounts", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "CarrierAccount", "action", "List", "route", "GET /v2/carrier_accounts")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCarrierAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/carrier_accounts/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "CarrierAccount", "action", "Show", "route", "GET /v2/carrier_accounts/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCarrierAccountContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateCarrierAccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/v2/carrier_accounts/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateCarrierAccountPayload))
	service.LogInfo("mount", "ctrl", "CarrierAccount", "action", "Update", "route", "PUT /v2/carrier_accounts/:id")
}

// unmarshalCareateCarrierAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalCareateCarrierAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &careateCarrierAccountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateCarrierAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateCarrierAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateCarrierAccountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// CarrierTypesController is the controller interface for the CarrierTypes actions.
type CarrierTypesController interface {
	goa.Muxer
	Show(*ShowCarrierTypesContext) error
}

// MountCarrierTypesController "mounts" a CarrierTypes resource controller on the given service.
func MountCarrierTypesController(service *goa.Service, ctrl CarrierTypesController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCarrierTypesContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/carrier_types", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "CarrierTypes", "action", "Show", "route", "GET /v2/carrier_types")
}
