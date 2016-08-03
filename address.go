package main

import (
	"encoding/json"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
	"io/ioutil"
)

// AddressController implements the address resource.
type AddressController struct {
	*goa.Controller
}

// NewAddressController creates a address controller.
func NewAddressController(service *goa.Service) *AddressController {
	return &AddressController{Controller: service.NewController("AddressController")}
}

// Create runs the create action.
func (c *AddressController) Create(ctx *app.CreateAddressContext) error {
	// AddressController_Create: start_implement

	// Put your logic here

	// AddressController_Create: end_implement
	res := &app.EasypostAddress{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AddressController) Show(ctx *app.ShowAddressContext) error {
	// AddressController_Show: start_implement

	// Get File
	file, e := ioutil.ReadFile("./data/address/" + ctx.ID + ".json")
	if e != nil {
		return e
	}
	// Unmarshal
	resJson := &app.EasypostAddress{}
	e = json.Unmarshal(file, resJson)
	if e != nil {
		return e
	}

	// AddressController_Show: end_implement
	return ctx.OK(resJson)
}
