package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// ShipmentController implements the shipment resource.
type ShipmentController struct {
	*goa.Controller
}

// NewShipmentController creates a shipment controller.
func NewShipmentController(service *goa.Service) *ShipmentController {
	return &ShipmentController{Controller: service.NewController("ShipmentController")}
}

// ConvertLabel runs the convertLabel action.
func (c *ShipmentController) ConvertLabel(ctx *app.ConvertLabelShipmentContext) error {
	// ShipmentController_ConvertLabel: start_implement

	// Put your logic here

	// ShipmentController_ConvertLabel: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}

// Create runs the create action.
func (c *ShipmentController) Create(ctx *app.CreateShipmentContext) error {
	// ShipmentController_Create: start_implement

	// Put your logic here

	// ShipmentController_Create: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}

// Insure runs the insure action.
func (c *ShipmentController) Insure(ctx *app.InsureShipmentContext) error {
	// ShipmentController_Insure: start_implement

	// Put your logic here

	// ShipmentController_Insure: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}

// Pruchase runs the pruchase action.
func (c *ShipmentController) Pruchase(ctx *app.PruchaseShipmentContext) error {
	// ShipmentController_Pruchase: start_implement

	// Put your logic here

	// ShipmentController_Pruchase: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}

// Rates runs the rates action.
func (c *ShipmentController) Rates(ctx *app.RatesShipmentContext) error {
	// ShipmentController_Rates: start_implement

	// Put your logic here

	// ShipmentController_Rates: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}

// Refund runs the refund action.
func (c *ShipmentController) Refund(ctx *app.RefundShipmentContext) error {
	// ShipmentController_Refund: start_implement

	// Put your logic here

	// ShipmentController_Refund: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ShipmentController) Show(ctx *app.ShowShipmentContext) error {
	// ShipmentController_Show: start_implement

	// Put your logic here

	// ShipmentController_Show: end_implement
	res := &app.EasypostShipment{}
	return ctx.OK(res)
}
