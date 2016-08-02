package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
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

	// Clone
	if payload.Clone {
		// TODO: copy everything, then override reference and description
		payload.Clone = false // for now
	}

	// Type
	if payload.Type == "EndiciaAccount" {
		// TODO: check that there is only one account with theis type(excluding clones). Might be able to do this async
	}
	file, e := ioutil.ReadFile("./data/carrier_types.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return e
	}
	carrierTypes := &app.EasypostCarrierTypesCollection{}
	e = json.Unmarshal(file, carrierTypes)
	if e != nil {
		fmt.Printf("Json Error: %v\n", e)
		return e
	}

	var sourceType *app.EasypostCarrierTypes
	for _, carrierType := range *carrierTypes {
		if carrierType.Type == payload.Type {
			sourceType = carrierType
			break
		}
	}
	if sourceType == nil {
		return errors.New("Invalid Carrier Type:" + payload.Type)
	}

	// ID
	uuid := make([]byte, 10)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return err
	}
	id := fmt.Sprintf("ca_%x", uuid)
	// TODO: make sure id is unique, and touch file
	payload.ID = &id

	// Readable
	if payload.Readable == nil {
		payload.Readable = sourceType.Readable
	}

	// Time
	now := time.Now().UTC().Format("2006-01-02T15:04:05-0700")
	payload.CreatedAt = &now
	payload.UpdatedAt = &now

	// Credentials & Fields
	if payload.Fields == nil {
		if payload.Credentials == nil {
			return errors.New("Missing Credentials")
		} else {
			// update fields with credentials
			a := reflect.ValueOf(*payload.Credentials).Interface().(map[string]interface{})
			fieldJsons := make([]string, len(a))
			i := 0
			for key, val := range a {
				visibility := "visible"
				matched, _ := regexp.MatchString(`(?i)password`, key)
				if matched {
					visibility = "password"
				}
				fieldJsons[i] = fmt.Sprintf("\"%s\": {\"visibility\": \"%s\", \"label\": \"\", \"value\": \"%s\"}", key, visibility, val)
				i += 1
			}

			var f interface{}
			err := json.Unmarshal([]byte("{"+strings.Join(fieldJsons, ", ")+"}"), &f)
			if err != nil {
				return err
			}

			fields := &app.FieldsObjectPayload{
				Credentials: &f,
			}
			payload.Fields = fields
		}
	} else {
		if payload.Fields.Credentials == nil {
			return errors.New("Missing Fields Credentials")
		} else {
			// update credentials with fields. Fields override credentials
			a := reflect.ValueOf(*payload.Fields.Credentials).Interface().(map[string]interface{})
			fieldJsons := make([]string, len(a))
			i := 0
			for key, val := range a {
				// Assume val is a (map[string]interface{}) for now
				a := reflect.ValueOf(val).Interface().(map[string]interface{})
				fieldJsons[i] = fmt.Sprintf("\"%s\": \"%s\"", key, a["value"])

				i += 1
			}
			var f interface{}
			err := json.Unmarshal([]byte("{"+strings.Join(fieldJsons, ", ")+"}"), &f)
			if err != nil {
				return err
			}
			payload.Credentials = &f
		}
	}
	// TODO: do same for test_credentials and fields.test_credentials

	// TODO: make sure account fields and account_type fields match

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
