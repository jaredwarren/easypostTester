//go:generate goagen bootstrap -d github.com/jaredwarren/easypostTester/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/easypostTester/app"
)

func main() {
	// Create service
	service := goa.New("easypost")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "address" controller
	c := NewAddressController(service)
	app.MountAddressController(service, c)
	// Mount "carrier_account" controller
	c2 := NewCarrierAccountController(service)
	app.MountCarrierAccountController(service, c2)
	// Mount "carrier_types" controller
	c3 := NewCarrierTypesController(service)
	app.MountCarrierTypesController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
