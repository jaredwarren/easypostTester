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
	"time"
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

	address := &app.EasypostAddress{}

	payload := ctx.Payload

	// ID
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		fmt.Println("ID")
		return err
	}
	id := fmt.Sprintf("adr_%x", uuid)
	// TODO: make sure id is unique, and touch file
	filepath := "./data/address/" + id + ".json"
	fileHandler, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Create File")
		return err
	}
	defer fileHandler.Close()
	address.ID = id

	address.Object = "Address"
	address.Mode = "test"

	// Time
	now := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	address.CreatedAt = &now
	address.UpdatedAt = &now

	// Copy from payload to address
	address.Street1 = &payload.Street1
	address.Street2 = payload.Street2
	address.Zip = &payload.Zip
	address.StateTaxID = payload.StateTaxID
	address.State = &payload.State
	address.Residential = payload.Residential
	address.Phone = payload.Phone
	address.Name = payload.Name
	address.FederalTaxID = payload.FederalTaxID
	address.Email = payload.Email
	address.Country = &payload.Country
	address.Company = payload.Company
	address.City = &payload.City
	address.CarrierFacility = payload.CarrierFacility

	// TODO: make city, state, street1,2 upper case, fakes standardation...

	// TODO: if verify, add verificateions to. must array containing either or both "delivery, zip4", verify_strict takes precedence

	// marshall to string
	b, err := json.Marshal(address)
	if err != nil {
		fmt.Println("Marshal")
		return err
	}
	// write file
	_, err = fileHandler.Write(b)
	if err != nil {
		fmt.Println("Write")
		return err
	}

	// AddressController_Create: end_implement
	return ctx.OK(address)
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
