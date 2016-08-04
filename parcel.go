package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// ParcelController implements the parcel resource.
type ParcelController struct {
	*goa.Controller
}

// NewParcelController creates a parcel controller.
func NewParcelController(service *goa.Service) *ParcelController {
	return &ParcelController{Controller: service.NewController("ParcelController")}
}

// Create runs the create action.
func (c *ParcelController) Create(ctx *app.CreateParcelContext) error {
	// ParcelController_Create: start_implement

	// Put your logic here

	// ParcelController_Create: end_implement
	res := &app.EasypostParcel{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ParcelController) Show(ctx *app.ShowParcelContext) error {
	// ParcelController_Show: start_implement

	// Put your logic here

	// ParcelController_Show: end_implement
	res := &app.EasypostParcel{}
	return ctx.OK(res)
}
