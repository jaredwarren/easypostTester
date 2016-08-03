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
		Response(OK, ArrayOf(CarrierAccount))
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

	Action("create", func() {
		Routing(POST("/"))
		Payload(CarrierAccountPayload)
		Description("CarrierAccount objects may be managed through the EasyPost API using the Production mode API key only. Multiple accounts can be added for a single carrier (with the exception of the USPS).")
		Response(OK)
	})

	Action("delete", func() {
		Description("CarrierAccount objects may be removed from your account when they become out of date or no longer useful.")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Carrier ID")
		})
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

// Address ...
var _ = Resource("address", func() {
	BasePath("/addresses")
	DefaultMedia(Address)

	Action("show", func() {
		Description("An Address can be retrieved by either its id or reference. However it is recommended to use EasyPost's provided identifiers because uniqueness on reference is not enforced.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Address ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(POST("/"))
		Payload(AddressPayload)
		Description("Depending on your use case an Address can be used in many different ways. Certain carriers allow rating between two zip codes, but full addresses are required to purchase postage. It is recommended to provide as much information as possible during creation and to reuse these objects whenever possible.")
		Response(OK)
	})
})
