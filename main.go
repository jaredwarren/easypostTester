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
	// Mount "api_key" controller
	c2 := NewAPIKeyController(service)
	app.MountAPIKeyController(service, c2)
	// Mount "carrier_account" controller
	c3 := NewCarrierAccountController(service)
	app.MountCarrierAccountController(service, c3)
	// Mount "carrier_types" controller
	c4 := NewCarrierTypesController(service)
	app.MountCarrierTypesController(service, c4)
	// Mount "customs_info" controller
	c5 := NewCustomsInfoController(service)
	app.MountCustomsInfoController(service, c5)
	// Mount "customs_item" controller
	c6 := NewCustomsItemController(service)
	app.MountCustomsItemController(service, c6)
	// Mount "insurance" controller
	c7 := NewInsuranceController(service)
	app.MountInsuranceController(service, c7)
	// Mount "parcel" controller
	c8 := NewParcelController(service)
	app.MountParcelController(service, c8)
	// Mount "shipment" controller
	c9 := NewShipmentController(service)
	app.MountShipmentController(service, c9)
	// Mount "tracker" controller
	c10 := NewTrackerController(service)
	app.MountTrackerController(service, c10)
	// Mount "user" controller
	c11 := NewUserController(service)
	app.MountUserController(service, c11)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
