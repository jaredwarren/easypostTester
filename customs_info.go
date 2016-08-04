package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// CustomsInfoController implements the customs_info resource.
type CustomsInfoController struct {
	*goa.Controller
}

// NewCustomsInfoController creates a customs_info controller.
func NewCustomsInfoController(service *goa.Service) *CustomsInfoController {
	return &CustomsInfoController{Controller: service.NewController("CustomsInfoController")}
}

// Create runs the create action.
func (c *CustomsInfoController) Create(ctx *app.CreateCustomsInfoContext) error {
	// CustomsInfoController_Create: start_implement

	// Put your logic here

	// CustomsInfoController_Create: end_implement
	res := &app.EasypostCustomsinfo{}
	return ctx.OK(res)
}
