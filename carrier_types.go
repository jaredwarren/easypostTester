package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
	"io/ioutil"
)

// CarrierTypesController implements the carrier_types resource.
type CarrierTypesController struct {
	*goa.Controller
}

// NewCarrierTypesController creates a carrier_types controller.
func NewCarrierTypesController(service *goa.Service) *CarrierTypesController {
	return &CarrierTypesController{Controller: service.NewController("CarrierTypesController")}
}

// Show runs the show action.
func (c *CarrierTypesController) Show(ctx *app.ShowCarrierTypesContext) error {
	// CarrierTypesController_Show: start_implement

	file, e := ioutil.ReadFile("./data/carrier_types.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return e
	}

	res := &app.EasypostCarrierTypesCollection{}
	e = json.Unmarshal(file, res)
	if e != nil {
		fmt.Printf("Json Error: %v\n", e)
		return e
	}

	// CarrierTypesController_Show: end_implement
	return ctx.OK(*res)
}
