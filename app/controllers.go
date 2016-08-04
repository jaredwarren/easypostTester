//************************************************************************//
// API "easypost": Application Controllers
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

// AddressController is the controller interface for the Address actions.
type AddressController interface {
	goa.Muxer
	Create(*CreateAddressContext) error
	Show(*ShowAddressContext) error
}

// MountAddressController "mounts" a Address resource controller on the given service.
func MountAddressController(service *goa.Service, ctrl AddressController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateAddressContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateAddressPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/addresses", ctrl.MuxHandler("Create", h, unmarshalCreateAddressPayload))
	service.LogInfo("mount", "ctrl", "Address", "action", "Create", "route", "POST /v2/addresses")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowAddressContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/addresses/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Address", "action", "Show", "route", "GET /v2/addresses/:id")
}

// unmarshalCreateAddressPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAddressPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createAddressPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// APIKeyController is the controller interface for the APIKey actions.
type APIKeyController interface {
	goa.Muxer
	List(*ListAPIKeyContext) error
}

// MountAPIKeyController "mounts" a APIKey resource controller on the given service.
func MountAPIKeyController(service *goa.Service, ctrl APIKeyController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListAPIKeyContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v2/api_keys", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "APIKey", "action", "List", "route", "GET /v2/api_keys")
}

// CarrierAccountController is the controller interface for the CarrierAccount actions.
type CarrierAccountController interface {
	goa.Muxer
	Create(*CreateCarrierAccountContext) error
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
		rctx, err := NewCreateCarrierAccountContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCarrierAccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/carrier_accounts", ctrl.MuxHandler("Create", h, unmarshalCreateCarrierAccountPayload))
	service.LogInfo("mount", "ctrl", "CarrierAccount", "action", "Create", "route", "POST /v2/carrier_accounts")

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

// unmarshalCreateCarrierAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCarrierAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCarrierAccountPayload{}
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

// CustomsInfoController is the controller interface for the CustomsInfo actions.
type CustomsInfoController interface {
	goa.Muxer
	Create(*CreateCustomsInfoContext) error
}

// MountCustomsInfoController "mounts" a CustomsInfo resource controller on the given service.
func MountCustomsInfoController(service *goa.Service, ctrl CustomsInfoController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCustomsInfoContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCustomsInfoPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/customs_infos", ctrl.MuxHandler("Create", h, unmarshalCreateCustomsInfoPayload))
	service.LogInfo("mount", "ctrl", "CustomsInfo", "action", "Create", "route", "POST /v2/customs_infos")
}

// unmarshalCreateCustomsInfoPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCustomsInfoPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCustomsInfoPayload{}
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

// CustomsItemController is the controller interface for the CustomsItem actions.
type CustomsItemController interface {
	goa.Muxer
	Create(*CreateCustomsItemContext) error
}

// MountCustomsItemController "mounts" a CustomsItem resource controller on the given service.
func MountCustomsItemController(service *goa.Service, ctrl CustomsItemController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCustomsItemContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCustomsItemPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/customs_items", ctrl.MuxHandler("Create", h, unmarshalCreateCustomsItemPayload))
	service.LogInfo("mount", "ctrl", "CustomsItem", "action", "Create", "route", "POST /v2/customs_items")
}

// unmarshalCreateCustomsItemPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCustomsItemPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCustomsItemPayload{}
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

// InsuranceController is the controller interface for the Insurance actions.
type InsuranceController interface {
	goa.Muxer
	Create(*CreateInsuranceContext) error
	List(*ListInsuranceContext) error
	Show(*ShowInsuranceContext) error
}

// MountInsuranceController "mounts" a Insurance resource controller on the given service.
func MountInsuranceController(service *goa.Service, ctrl InsuranceController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateInsuranceContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateInsurancePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/insurances", ctrl.MuxHandler("Create", h, unmarshalCreateInsurancePayload))
	service.LogInfo("mount", "ctrl", "Insurance", "action", "Create", "route", "POST /v2/insurances")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListInsuranceContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListInsurancePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v2/insurances", ctrl.MuxHandler("List", h, unmarshalListInsurancePayload))
	service.LogInfo("mount", "ctrl", "Insurance", "action", "List", "route", "GET /v2/insurances")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowInsuranceContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/insurances/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Insurance", "action", "Show", "route", "GET /v2/insurances/:id")
}

// unmarshalCreateInsurancePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateInsurancePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createInsurancePayload{}
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

// unmarshalListInsurancePayload unmarshals the request body into the context request data Payload field.
func unmarshalListInsurancePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listInsurancePayload{}
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

// ParcelController is the controller interface for the Parcel actions.
type ParcelController interface {
	goa.Muxer
	Create(*CreateParcelContext) error
	Show(*ShowParcelContext) error
}

// MountParcelController "mounts" a Parcel resource controller on the given service.
func MountParcelController(service *goa.Service, ctrl ParcelController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateParcelContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateParcelPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/parcels", ctrl.MuxHandler("Create", h, unmarshalCreateParcelPayload))
	service.LogInfo("mount", "ctrl", "Parcel", "action", "Create", "route", "POST /v2/parcels")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowParcelContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/parcels/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Parcel", "action", "Show", "route", "GET /v2/parcels/:id")
}

// unmarshalCreateParcelPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateParcelPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createParcelPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ShipmentController is the controller interface for the Shipment actions.
type ShipmentController interface {
	goa.Muxer
	Buy(*BuyShipmentContext) error
	Create(*CreateShipmentContext) error
	Insure(*InsureShipmentContext) error
	Label(*LabelShipmentContext) error
	Rates(*RatesShipmentContext) error
	Refund(*RefundShipmentContext) error
	Show(*ShowShipmentContext) error
}

// MountShipmentController "mounts" a Shipment resource controller on the given service.
func MountShipmentController(service *goa.Service, ctrl ShipmentController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewBuyShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*BuyShipmentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Buy(rctx)
	}
	service.Mux.Handle("POST", "/v2/shipments/:id/buy", ctrl.MuxHandler("Buy", h, unmarshalBuyShipmentPayload))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Buy", "route", "POST /v2/shipments/:id/buy")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateShipmentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/shipments", ctrl.MuxHandler("Create", h, unmarshalCreateShipmentPayload))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Create", "route", "POST /v2/shipments")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewInsureShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*InsureShipmentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Insure(rctx)
	}
	service.Mux.Handle("POST", "/v2/shipments/:id/insure", ctrl.MuxHandler("Insure", h, unmarshalInsureShipmentPayload))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Insure", "route", "POST /v2/shipments/:id/insure")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLabelShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*LabelShipmentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Label(rctx)
	}
	service.Mux.Handle("POST", "/v2/shipments/:id/label", ctrl.MuxHandler("Label", h, unmarshalLabelShipmentPayload))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Label", "route", "POST /v2/shipments/:id/label")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRatesShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Rates(rctx)
	}
	service.Mux.Handle("POST", "/v2/shipments/:id/rates", ctrl.MuxHandler("Rates", h, nil))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Rates", "route", "POST /v2/shipments/:id/rates")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRefundShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Refund(rctx)
	}
	service.Mux.Handle("POST", "/v2/shipments/:id/refund", ctrl.MuxHandler("Refund", h, nil))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Refund", "route", "POST /v2/shipments/:id/refund")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowShipmentContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/shipments/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Shipment", "action", "Show", "route", "GET /v2/shipments/:id")
}

// unmarshalBuyShipmentPayload unmarshals the request body into the context request data Payload field.
func unmarshalBuyShipmentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &buyShipmentPayload{}
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

// unmarshalCreateShipmentPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateShipmentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createShipmentPayload{}
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

// unmarshalInsureShipmentPayload unmarshals the request body into the context request data Payload field.
func unmarshalInsureShipmentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &insureShipmentPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalLabelShipmentPayload unmarshals the request body into the context request data Payload field.
func unmarshalLabelShipmentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &labelShipmentPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// TrackerController is the controller interface for the Tracker actions.
type TrackerController interface {
	goa.Muxer
	Create(*CreateTrackerContext) error
	List(*ListTrackerContext) error
	Show(*ShowTrackerContext) error
}

// MountTrackerController "mounts" a Tracker resource controller on the given service.
func MountTrackerController(service *goa.Service, ctrl TrackerController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTrackerContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateTrackerPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/trackers", ctrl.MuxHandler("Create", h, unmarshalCreateTrackerPayload))
	service.LogInfo("mount", "ctrl", "Tracker", "action", "Create", "route", "POST /v2/trackers")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTrackerContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListTrackerPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v2/trackers", ctrl.MuxHandler("List", h, unmarshalListTrackerPayload))
	service.LogInfo("mount", "ctrl", "Tracker", "action", "List", "route", "GET /v2/trackers")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTrackerContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/trackers/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Tracker", "action", "Show", "route", "GET /v2/trackers/:id")
}

// unmarshalCreateTrackerPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTrackerPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createTrackerPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalListTrackerPayload unmarshals the request body into the context request data Payload field.
func unmarshalListTrackerPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listTrackerPayload{}
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

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Show(*ShowUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/v2/users", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /v2/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v2/users/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Show", "route", "GET /v2/users/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/v2/users/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PUT /v2/users/:id")
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
