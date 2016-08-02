package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BottlePayload = Type("BottlePayload", func() {
	Attribute("status", Integer, "status of bottle")
	Required("status")
})

// Carrier Accounts ...
var _ = Resource("carrier_account", func() {
	BasePath("/carrier_accounts")
	DefaultMedia(CarrierAccount)

	Action("list", func() {
		Description("Retrieve an unpaginated list of all CarrierAccounts available to the authenticated account. Only Production API keys may be used to retrieve this list, as there is no test mode equivalent.")
		Routing(GET("/"))
		Response(OK)
	})

	Action("show", func() {
		Description("Retrieve an unpaginated list of all CarrierAccounts available to the authenticated account. Only Production API keys may be used to retrieve this list, as there is no test mode equivalent.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Carrier ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(CarrierAccountPayload)
		Description("Updates can be made to description, reference, and any fields in credentials or test_credentials.")
		Response(OK)
	})

	Action("careate", func() {
		Routing(POST("/"))
		Payload(CarrierAccountPayload)
		Description("CarrierAccount objects may be managed through the EasyPost API using the Production mode API key only. Multiple accounts can be added for a single carrier (with the exception of the USPS).")
		Response(OK)
	})

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Description("CarrierAccount objects may be removed from your account when they become out of date or no longer useful.")
		Response(OK)
	})
})

// Carrier Types ...
var _ = Resource("carrier_types", func() {
	BasePath("/carrier_types")
	DefaultMedia(CollectionOf(CarrierTypes))

	Action("show", func() {
		Description("List carrier types.")
		//Routing(GET(""))
		Routing(GET("/"))
		Response(OK)
		//Response(NotFound)
	})
})
