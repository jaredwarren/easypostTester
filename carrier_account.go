package main

import (
	"encoding/json"
	//"fmt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
	"io/ioutil"
	"os"
	"regexp"
)

// CarrierAccountController implements the carrier_account resource.
type CarrierAccountController struct {
	*goa.Controller
}

// NewCarrierAccountController creates a carrier_account controller.
func NewCarrierAccountController(service *goa.Service) *CarrierAccountController {
	return &CarrierAccountController{Controller: service.NewController("CarrierAccountController")}
}

// Careate runs the careate action.
func (c *CarrierAccountController) Careate(ctx *app.CareateCarrierAccountContext) error {
	// CarrierAccountController_Careate: start_implement

	// Put your logic here
	//TODO: write file here...
	//randomly generate id
	//"./data/" + ctx.ID + ".json"

	// CarrierAccountController_Careate: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *CarrierAccountController) Delete(ctx *app.DeleteCarrierAccountContext) error {
	// CarrierAccountController_Delete: start_implement

	// Put your logic here
	err := os.Remove("./data/" + ctx.ID + ".json")

	if err != nil {
		return err
	}

	// CarrierAccountController_Delete: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *CarrierAccountController) List(ctx *app.ListCarrierAccountContext) error {
	// CarrierAccountController_List: start_implement

	// Put your logic here
	//TODO: search for each file with ca_*.json
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		return err
	}

	for _, file := range files {
		matched, err := regexp.MatchString("^ca_(.+?)\\.json$", file.Name())
		//fmt.Println(file.Name())
	}

	// CarrierAccountController_List: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CarrierAccountController) Show(ctx *app.ShowCarrierAccountContext) error {
	// CarrierAccountController_Show: start_implement

	file, e := ioutil.ReadFile("./data/" + ctx.ID + ".json")
	if e != nil {
		//fmt.Printf("File error: %v\n", e)
		return e
	}

	resJson := &app.EasypostCarrierAccounts{}
	e = json.Unmarshal(file, resJson)
	if e != nil {
		//fmt.Printf("Json Error: %v\n", e)
		return e
	}

	// CarrierAccountController_Show: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *CarrierAccountController) Update(ctx *app.UpdateCarrierAccountContext) error {
	// CarrierAccountController_Update: start_implement

	// Put your logic here
	//TODO: update ONLY: description, reference, and any fields in credentials or test_credentials.
	//TODO: update file "./data/" + ctx.ID + ".json"

	// CarrierAccountController_Update: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}
