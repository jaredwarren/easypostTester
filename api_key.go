package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// APIKeyController implements the api_key resource.
type APIKeyController struct {
	*goa.Controller
}

// NewAPIKeyController creates a api_key controller.
func NewAPIKeyController(service *goa.Service) *APIKeyController {
	return &APIKeyController{Controller: service.NewController("APIKeyController")}
}

// List runs the list action.
func (c *APIKeyController) List(ctx *app.ListAPIKeyContext) error {
	// APIKeyController_List: start_implement

	// Put your logic here

	// APIKeyController_List: end_implement
	res := []*app.EasypostAPIKeys{}
	return ctx.OK(res)
}
