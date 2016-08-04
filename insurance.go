package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// InsuranceController implements the insurance resource.
type InsuranceController struct {
	*goa.Controller
}

// NewInsuranceController creates a insurance controller.
func NewInsuranceController(service *goa.Service) *InsuranceController {
	return &InsuranceController{Controller: service.NewController("InsuranceController")}
}

// Create runs the create action.
func (c *InsuranceController) Create(ctx *app.CreateInsuranceContext) error {
	// InsuranceController_Create: start_implement

	// Put your logic here

	// InsuranceController_Create: end_implement
	res := &app.EasypostTracker{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *InsuranceController) List(ctx *app.ListInsuranceContext) error {
	// InsuranceController_List: start_implement

	// Put your logic here

	// InsuranceController_List: end_implement
	res := &app.EasypostTrackers{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *InsuranceController) Show(ctx *app.ShowInsuranceContext) error {
	// InsuranceController_Show: start_implement

	// Put your logic here

	// InsuranceController_Show: end_implement
	res := &app.EasypostTracker{}
	return ctx.OK(res)
}
