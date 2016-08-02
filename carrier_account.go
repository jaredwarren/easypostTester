package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

// CarrierAccountController implements the carrier_account resource.
type CarrierAccountController struct {
	*goa.Controller
}

// NewCarrierAccountController creates a carrier_account controller.
func NewCarrierAccountController(service *goa.Service) *CarrierAccountController {
	return &CarrierAccountController{Controller: service.NewController("CarrierAccountController")}
}

// Create runs the create action.
func (c *CarrierAccountController) Create(ctx *app.CreateCarrierAccountContext) error {
	// CarrierAccountController_Create: start_implement

	// Put your logic here
	//TODO: write file here...
	//randomly generate id
	//"./data/" + ctx.ID + ".json"

	// Validation
	//// make sure type matches a type

	payload := ctx.Payload
	if *payload.Clone {
		// TODO: copy everything, then override reference and description
		*payload.Clone = false // for now
	}

	if payload.Type == "EndiciaAccount" {
		// TODO: check that there is only one.
	}

	// generate id
	uuid := make([]byte, 10)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return err
	}
	id := fmt.Sprintf("ca_%x", uuid)
	// TODO: make sure id is unique
	payload.ID = &id

	if payload.Fields == nil {

	}
	now := time.Now().UTC().Format("2006-01-02T15:04:05-0700")
	payload.CreatedAt = &now
	payload.UpdatedAt = &now

	// apply defaults
	//// created_at: now
	//// updated_at: now
	//// readable: {from carrier_type if empty}
	//// credentials {if empty use fields}
	//// fields {if empty use credentials}
	////// visibility: visible, unless key contains "password"
	////// value: value
	////// key: key
	////// Label: ??

	//27839172aee03918a701
	//0d388c36db23ae6a
	//00bbaaae809fbce74bc0
	//9bc7ffdd-e765-4f00-948c-948ba2ea3293
	fmt.Printf("%+v\n", id)

	// CarrierAccountController_Create: end_implement
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

	//payload.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05-0700")

	// CarrierAccountController_Update: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}
