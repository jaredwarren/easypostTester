package main

import (
	"encoding/json"
	_ "fmt"
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

	// Delete file
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

	// Get All Files
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		return err
	}
	accounts := make([]*app.EasypostCarrierAccounts, 0)

	// Find account file and add to list
	for _, fileInfo := range files {
		matched, _ := regexp.MatchString("^ca_(.+?)\\.json$", fileInfo.Name())
		if matched {
			resJson := &app.EasypostCarrierAccounts{}
			file, e := ioutil.ReadFile("./data/" + fileInfo.Name())
			if e != nil {
				return e
			}
			e = json.Unmarshal(file, resJson)
			if e != nil {
				return e
			}
			accounts = append(accounts, resJson)
		}
	}

	// CarrierAccountController_List: end_implement
	return ctx.OK(accounts)
}

// Show runs the show action.
func (c *CarrierAccountController) Show(ctx *app.ShowCarrierAccountContext) error {
	// CarrierAccountController_Show: start_implement

	file, e := ioutil.ReadFile("./data/" + ctx.ID + ".json")
	if e != nil {
		return e
	}

	resJson := &app.EasypostCarrierAccounts{}
	e = json.Unmarshal(file, resJson)
	if e != nil {
		return e
	}

	// CarrierAccountController_Show: end_implement
	return ctx.OK(resJson)
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
