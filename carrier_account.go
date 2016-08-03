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

	account := &app.EasypostCarrierAccounts{}
	payload := ctx.Payload

	// Clone
	if payload.Clone {
		// TODO: copy everything, then override reference and description.
	}
	account.Clone = false // for now

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
	// TODO: get list of credential fields

	// ID
	uuid := make([]byte, 10)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return err
	}
	id := fmt.Sprintf("ca_%x", uuid)
	// TODO: make sure id is unique, and touch file
	filepath := "./data/" + id + ".json"
	fileHandler, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fileHandler.Close()
	account.ID = id

	// Readable
	if payload.Readable == nil {
		account.Readable = sourceType.Readable
	}

	// Credentials & Fields
	if payload.Fields == nil {
		if payload.Credentials != nil {
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
			account.Fields = fields
		}
	} else {
		if payload.Fields.Credentials != nil {
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
			account.Credentials = &f
		}
	}

	// TestCredentials
	if payload.Fields == nil {
		if payload.TestCredentials != nil {
			// update fields with credentials
			a := reflect.ValueOf(*payload.TestCredentials).Interface().(map[string]interface{})
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
				TestCredentials: &f,
			}
			account.Fields = fields
		}
	} else {
		if payload.Fields.TestCredentials != nil {
			// update credentials with fields. Fields override credentials
			a := reflect.ValueOf(*payload.Fields.TestCredentials).Interface().(map[string]interface{})
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
			account.TestCredentials = &f
		}
	}

	// TODO: make sure account fields and account_type fields match

	// marshall to string
	b, err := json.Marshal(account)
	if err != nil {
		return err
	}
	// write file
	_, err = fileHandler.Write(b)
	if err != nil {
		return err
	}

	// CarrierAccountController_Create: end_implement
	return ctx.OK(account)
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
		fmt.Println("Error Reading Directory")
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
				fmt.Println("Error Reading File:", fileInfo.Name())
				return e
			}
			e = json.Unmarshal(file, resJson)
			if e != nil {
				fmt.Println("Error Parsing File:", fileInfo.Name())
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

	// Get File
	file, e := ioutil.ReadFile("./data/" + ctx.ID + ".json")
	if e != nil {
		return e
	}
	// Unmarshal
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

	payload := ctx.Payload

	filepath := "./data/" + ctx.ID + ".json"
	// Get File
	fileData, e := ioutil.ReadFile(filepath)
	if e != nil {
		return e
	}
	// Unmarshal
	resJson := &app.EasypostCarrierAccounts{}
	e = json.Unmarshal(fileData, resJson)
	if e != nil {
		return e
	}

	// UpdatedAt
	now := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	resJson.UpdatedAt = &now

	// Description
	if payload.Description != nil {
		resJson.Description = payload.Description
	}

	// Reference
	if payload.Reference != nil {
		resJson.Reference = payload.Reference
	}

	// Credentials & Fields
	if payload.Fields == nil {
		if payload.Credentials != nil {
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
			resJson.Fields = fields
		}
		// TestCredentials
		if payload.TestCredentials != nil {
			// update fields with TestCredentials
			a := reflect.ValueOf(*payload.TestCredentials).Interface().(map[string]interface{})
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
				TestCredentials: &f,
			}
			resJson.Fields = fields
		}
	} else {
		if payload.Fields.Credentials != nil {
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
			resJson.Credentials = &f
		}
		// TestCredentials
		if payload.Fields.TestCredentials != nil {
			// update TestCredentials with fields. Fields override TestCredentials
			a := reflect.ValueOf(*payload.Fields.TestCredentials).Interface().(map[string]interface{})
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
			resJson.TestCredentials = &f
		}
	}

	// File to string
	b, err := json.Marshal(resJson)
	if err != nil {
		return err
	}
	// Open file
	fileHandler, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fileHandler.Close()
	// Write file
	_, err = fileHandler.Write(b)
	if err != nil {
		return err
	}

	// CarrierAccountController_Update: end_implement
	res := &app.EasypostCarrierAccounts{}
	return ctx.OK(res)
}
