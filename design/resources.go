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

/**
* Carrier Types
**/
var _ = Resource("carrier_types", func() {
	BasePath("/carrier_types")
	DefaultMedia(CollectionOf(CarrierTypes))

	Action("show", func() {
		Description("List carrier types.")
		Routing(GET("/"))
		Response(OK)
	})
})

/**
* Address
**/
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

/**
* API Keys
**/
var _ = Resource("api_key", func() {
	BasePath("/api_keys")
	DefaultMedia(APIKeys)

	Action("list", func() {
		Description("Both production and test keys will be returned a User and all of its children. If the request is authenticated as a Child only the API Keys for that Child will be returned.")
		Routing(GET("/"))
		Response(OK, ArrayOf(APIKeys))
	})
})

/**
* Parcels
**/
var _ = Resource("parcel", func() {
	BasePath("/parcels")
	DefaultMedia(Parcel)

	Action("show", func() {
		Description("Get a Parcel by its id. In general you should not need to use this in your automated solution. A Parcel's id can be inlined into the creation call to other objects. This allows you to only create one Parcel for each package you will be using.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Parcel ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(POST("/"))
		Payload(ParcelPayload)
		Description("Include the weight, and either a predefined_package or length, width and height.")
		Response(OK)
	})
})

/**
* User
**/

var _ = Resource("user", func() {
	BasePath("/users")
	DefaultMedia(User)

	Action("create", func() {
		Routing(POST("/"))
		Payload(UserCreatePayload) // TODO: handles create and return if is_return = true
		Response(OK)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "User ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(UserUpdatePayload)
		Response(OK)
	})

})

/**
* Shipment
**/
var _ = Resource("shipment", func() {
	BasePath("/shipments")
	DefaultMedia(Shipment)

	Action("show", func() {
		Description("Get a Parcel by its id. In general you should not need to use this in your automated solution. A Parcel's id can be inlined into the creation call to other objects. This allows you to only create one Parcel for each package you will be using.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Parcel ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Description("A Shipment is almost exclusively a container for other objects, and thus a Shipment may reuse many of these objects. Additionally, all the objects contained within a Shipment may be created at the same time.")
		Routing(POST("/"))
		Payload(ShipmentPayload) // TODO: handles create and return if is_return = true
		Response(OK)
	})

	Action("buy", func() {
		Description("To purchase a Shipment you only need to specify the Rate to purchase. This operation populates the tracking_code and postage_label attributes. The default image format of the associated PostageLabel is PNG. To change this default see the label_format option.")
		Routing(POST("/:id/buy"))
		Payload(BuyShipmentPayload)
		Response(OK)
	})

	Action("label", func() {
		Description("A Shipment's PostageLabel can be converted from PNG to other formats. If the PostageLabel was originally generated in a format other than PNG it cannot be converted.")
		Routing(POST("/:id/label"))
		Payload(LabelShipmentPayload)
		Response(OK)
	})

	Action("rates", func() {
		Description("You can update the Rates of a Shipment at any time. This operation respects the carrier_accounts attribute.")
		Routing(POST("/:id/rates"))
		Response(OK)
	})

	Action("insure", func() {
		Description("Insuring your Shipment is as simple as sending us the value of the contents. We charge 1% of the value, with a $1 minimum, and handle all the claims. All our claims are paid out within 30 days.")
		Routing(POST("/:id/insure"))
		Payload(ShipmentInsurancePayload)
		Response(OK)
	})

	Action("refund", func() {
		Description("Refunding a Shipment is avaliable for many carriers supported by EasyPost. Once the refund has been submitted, refund_status attribute of the Shipment will be populated with one of the possible values: \"submitted\", \"refunded\", \"rejected\". The most common initial status is \"submitted\". Many carriers require that the refund be processed before the refund_status will move to \"refunded\". The length of this process depends on the carrier, but no greater than 30 days.")
		Routing(POST("/:id/refund"))
		Response(OK)
	})
})

/**
* Tracking
**/
var _ = Resource("tracker", func() {
	BasePath("/trackers")
	DefaultMedia(Tracker)

	Action("list", func() {
		Description("The Tracker List is a paginated list of all Tracker objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.")
		Routing(GET("/"))
		Payload(TrackerListPayload)
		Response(OK, Trackers)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Retrieve a Tracker by id.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Tracking ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Description("A Tracker can be created with only a tracking_code. Optionally, you can provide the carrier parameter, which indicates the carrier the package was shipped with. If no carrier is provided, EasyPost will attempt to determine the carrier based on the tracking_code provided. Providing a carrier parameter is recommended, since some tracking_codes are ambiguous and may match with more than one carrier. In addition, not having to auto-match the carrier will significantly speed up the response time.")
		Routing(POST("/"))
		Payload(TrackerPayload) // TODO: handles create and return if is_return = true
		Response(OK)
	})

	/// TESTING:   https://www.easypost.com/docs/api/curl#test-tracking-codes

})

/**
* Insurance
**/
var _ = Resource("insurance", func() {
	BasePath("/insurances")
	DefaultMedia(Tracker)

	Action("create", func() {
		Description("An Insurance created via this endpoint must belong to a shipment purchased outside of EasyPost. Insurance for Shipments created within EasyPost must be created via the Shipment Buy or Insure endpoints. When creating Insurance for a non-EasyPost shipment, you must provide to_address, from_address, tracking_code, and amount information. Optionally, you can provide the carrier parameter, which will help EasyPost identify the carrier the package was shipped with. If no carrier is provided, EasyPost will attempt to determine the carrier based on the tracking_code provided. Providing a carrier parameter is recommended, since some tracking_codes are ambiguous and may match with more than one carrier. In addition, not having to auto-match the carrier will significantly speed up the response time.")
		Routing(POST("/"))
		Payload(InsurancePayload) // TODO: handles create and return if is_return = true
		Response(OK)
	})

	Action("list", func() {
		Description("The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.")
		Routing(GET("/"))
		Payload(InsuranceListPayload)
		Response(OK, Trackers)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Retrieve an Insurance by id.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Unique, starts with \"ins_\"")
		})
		Response(OK)
		Response(NotFound)
	})
})

/**
* Customs
**/
var _ = Resource("customs_info", func() {
	BasePath("/customs_infos")
	DefaultMedia(CustomsInfo)

	Action("create", func() {
		Routing(POST("/"))
		Payload(CustomsInfoPayload)
		Response(OK)
	})
})

var _ = Resource("customs_item", func() {
	BasePath("/customs_items")
	DefaultMedia(CustomItem)

	Action("create", func() {
		Routing(POST("/"))
		Payload(CustomItemPayload)
		Response(OK)
	})
})
