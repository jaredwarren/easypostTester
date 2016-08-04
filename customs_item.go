package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// CustomsItemController implements the customs_item resource.
type CustomsItemController struct {
	*goa.Controller
}

// NewCustomsItemController creates a customs_item controller.
func NewCustomsItemController(service *goa.Service) *CustomsItemController {
	return &CustomsItemController{Controller: service.NewController("CustomsItemController")}
}

// Create runs the create action.
func (c *CustomsItemController) Create(ctx *app.CreateCustomsItemContext) error {
	// CustomsItemController_Create: start_implement

	// Put your logic here

	// CustomsItemController_Create: end_implement
	res := &app.EasypostCustomitem{}
	return ctx.OK(res)
}
