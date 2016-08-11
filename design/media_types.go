package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CarrierAccount = MediaType("application/easypost.carrier_accounts+json", func() {
	Reference(CarrierAccountPayload)
	Description("A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.")
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"ca_\"")
		Attribute("object", String, "Always: \"CarrierAccount\"", func() {
			Pattern("^CarrierAccount$")
			Default("CarrierAccount")
		})
		Attribute("type")
		Attribute("fields")
		Attribute("clone")
		Attribute("description")
		Attribute("reference")
		Attribute("readable")
		Attribute("logo", String, "url to logo") // Not supported yet
		Attribute("credentials")
		Attribute("test_credentials")
		Attribute("created_at", String, "The name used when displaying a readable value for the type of the account")
		Attribute("updated_at", String, "The name used when displaying a readable value for the type of the account")

		Required("id", "object")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("type")
		Attribute("fields")
		Attribute("clone")
		Attribute("description")
		Attribute("reference")
		Attribute("readable")
		Attribute("credentials")
		Attribute("test_credentials")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

// FieldsObject ...
var FieldsObject = MediaType("application/easypost.fields_object+json", func() {
	Description("Contains \"credentials\" and/or \"test_credentials\", or may be empty")
	Attributes(func() {
		Attribute("credentials", Any, "Credentials used in the production environment.")
		Attribute("test_credentials", Any, "Credentials used in the test environment.")
		Attribute("auto_link", Boolean, "For USPS this designates that no credentials are required.")
		Attribute("custom_workflow", Boolean, "When present, a seperate authentication process will be required through the UI to link this account type.")

		Required("credentials")
	})
	View("default", func() {
		Attribute("credentials")
		Attribute("test_credentials")
		Attribute("auto_link")
		Attribute("custom_workflow")
	})
})

// FieldsObject ...
var FieldObject = MediaType("application/easypost.field_object+json", func() {
	Description("Contains \"credentials\" and/or \"test_credentials\", or may be empty")
	Reference(FieldsObjectPayload)
	Attributes(func() {
		Attribute("key")
		Attribute("visibility")
		Attribute("label", String)
		Attribute("value", String)

		Required("key", "value", "visibility", "label")
	})
	View("default", func() {
		Attribute("key")
		Attribute("visibility")
		Attribute("label")
		Attribute("value")
	})
})

// CarrierTypes
var CarrierTypes = MediaType("application/easypost.carrier_types+json", func() {
	Description("The CarrierType object provides an interface for determining the valid fields of a CarrierAccount. The list of CarrierType objects only changes when a new carrier is added to EasyPost.")
	Attributes(func() {
		Attribute("object", String, "Always: \"CarrierType\"", func() {
			Pattern("^CarrierType$")
			Default("CarrierType")
		})
		Attribute("type", String, "The name of the carrier type. Note that \"EndiciaAccount\" is the current USPS integration account type")
		Attribute("fields", FieldsObject, "Contains at least one of the following keys: \"auto_link\", \"credentials\", \"test_credentials\", and \"custom_workflow\"")
		Attribute("readable", String, "The name used when displaying a readable value for the type of the account")

		Required("type", "object", "fields")
	})
	View("default", func() {
		Attribute("object")
		Attribute("type")
		Attribute("readable")
		Attribute("fields")
	})
})

/**
* Address
**/

var Address = MediaType("application/easypost.address+json", func() {
	Reference(AddressPayload)
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"adr_\"", func() {
			Pattern("^adr_")
		})
		Attribute("object", String, "Always: \"Address\"", func() {
			Pattern("^Address$")
			Default("Address")
		})
		Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})
		Attribute("street1")
		Attribute("street2")
		Attribute("city")
		Attribute("state")
		Attribute("zip")
		Attribute("country")
		Attribute("residential")
		Attribute("carrier_facility")
		Attribute("name")
		Attribute("company")
		Attribute("phone")
		Attribute("email")
		Attribute("federal_tax_id")
		Attribute("state_tax_id")
		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")
		Attribute("verifications", Verifications, "The result of any verifications requested")

		Required("id", "object")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("mode")
		Attribute("street1")
		Attribute("street2")
		Attribute("city")
		Attribute("state")
		Attribute("zip")
		Attribute("country")
		Attribute("residential")
		Attribute("carrier_facility")
		Attribute("name")
		Attribute("company")
		Attribute("phone")
		Attribute("email")
		Attribute("federal_tax_id")
		Attribute("state_tax_id")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("verifications")
	})
})

// Verifications ...
var Verifications = MediaType("application/easypost.address_verifications+json", func() {
	Attributes(func() {
		Attribute("zip4", Verification, "Only applicable to US addresses - checks and sets the ZIP+4.")
		Attribute("delivery", Verification, "Checks that the address is deliverable and makes minor corrections to spelling/format. US addresses will also have their \"residential\" status checked and set.")
	})
	View("default", func() {
		Attribute("zip4")
		Attribute("delivery")
	})
})

// Verification ...
var Verification = MediaType("application/easypost.address_verification+json", func() {
	Attributes(func() {
		Attribute("success", Boolean, "The success of the verification")
		Attribute("errors", ArrayOf(FieldError), "All errors that caused the verification to fail")
	})
	View("default", func() {
		Attribute("success")
		Attribute("errors")
	})
})

/**
* Parcels
**/

var Parcel = MediaType("application/easypost.parcel+json", func() {
	Attributes(func() {
		Reference(ParcelPayload)
		Attribute("id", String, "Unique, begins with \"prcl_\"", func() {
			Pattern("^prcl_")
		})
		Attribute("object", String, "Always: \"Parcel\"", func() {
			Pattern("^Parcel$")
			Default("Parcel")
		})
		Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})
		Attribute("length", Number, "Required if predefined_package is empty. float (inches)")
		Attribute("width", Number, "Required if predefined_package is empty. float (inches)")
		Attribute("height", Number, "Required if predefined_package is empty. float (inches)")
		Attribute("weight", Number, "Always required. float(oz)")
		Attribute("predefined_package", String, "Optional, one of our predefined_packages")
		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object", "weight")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("mode")
		Attribute("length")
		Attribute("width")
		Attribute("height")
		Attribute("weight")
		Attribute("predefined_package")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

/**
* API Keys
**/
var APIKeys = MediaType("application/easypost.api_keys+json", func() {
	Attributes(func() {
		Attribute("id", String, "The User id of the authenticated User making the API Key request", func() {
			Pattern("^user_")
		})
		Attribute("children", ArrayOf(APIKeyChild), "A list of all Child Users presented with ONLY id, children, and key array structures.")
		Attribute("keys", ArrayOf(APIKey), "The list of all API keys active for an account, both for \"test\" and \"production\" modes.")

		Required("id", "children", "keys")
	})
	View("default", func() {
		Attribute("id")
		Attribute("children")
		Attribute("keys")
	})
})

var APIKeyChild = MediaType("application/easypost.api_key_child+json", func() {
	Attributes(func() {
		Attribute("id", String, "The User id of the authenticated User making the API Key request", func() {
			Pattern("^user_")
		})
		Attribute("children", CollectionOf("application/easypost.api_key_child+json"), "A list of all Child Users presented with ONLY id, children, and key array structures.")
		Attribute("keys", ArrayOf(APIKey), "The list of all API keys active for an account, both for \"test\" and \"production\" modes.")

		Required("id", "children", "keys")
	})
	View("default", func() {
		Attribute("id")
		Attribute("children")
		Attribute("keys")
	})
})

var APIKey = MediaType("application/easypost.api_key+json", func() {
	Attributes(func() {
		Attribute("object", String, "Always: \"ApiKey\"", func() {
			Pattern("^ApiKey$")
			Default("ApiKey")
		})
		Attribute("mode", String, "\"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})
		Attribute("key", String, "The actual key value to use for authentication")
		Attribute("created_at", String, "The User id of the authenticated User making the API Key request")

		Required("object", "mode", "key", "created_at")
	})
	View("default", func() {
		Attribute("object")
		Attribute("mode")
		Attribute("key")
		Attribute("created_at")
	})
})

/**
* User
**/

var User = MediaType("application/easypost.user+json", func() {
	Description("An Insurance object represents insurance for packages purchased both va the EasyPost API as well as shipments purchased through third parties and later registered with EasyPost. An Insurance is created automatically whenever you buy a Shipment through EasyPost and pass insurance options during the Buy call or in a later call to Insure a Shipment.")
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"user_\"", func() {
			Pattern("^user_")
		})
		Attribute("object", String, "Always: \"User\"", func() {
			Pattern("^User$")
			Default("User")
		})

		Attribute("parent_id", String, "The ID of the parent user object. Top-level users are defined as users with no parent")
		Attribute("name", String, "First and last name required")
		Attribute("email", String, "Required")
		Attribute("phone_number", String, "Optional")
		Attribute("balance", String, "Formatted as string \"XX.XXXXX\"")
		Attribute("recharge_amount", String, "USD formatted dollars and cents, delimited by a decimal point")
		Attribute("secondary_recharge_amount", String, "USD formatted dollars and cents, delimited by a decimal point")
		Attribute("recharge_threshold", String, "Number of cents USD that when your balance drops below, we automatically recharge your account with your primary payment method.")
		Attribute("children", CollectionOf("application/easypost.user+json"), "All associated children")

		Required("id", "object", "name", "email")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")

		Attribute("parent_id")
		Attribute("name")
		Attribute("email")
		Attribute("phone_number")
		Attribute("balance")
		Attribute("recharge_amount")
		Attribute("secondary_recharge_amount")
		Attribute("recharge_threshold")
		Attribute("children")
	})
})

/**
* Shipment
**/
var Shipment = MediaType("application/easypost.shipment+json", func() {
	Attributes(func() {
		Reference(ShipmentPayload)
		Attribute("id", String, "Unique, begins with \"shp_\"", func() {
			Pattern("^shp_")
		})
		Attribute("object", String, "Always: \"Shipment\"", func() {
			Pattern("^Shipment$")
			Default("Shipment")
		})
		Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})

		Attribute("to_address")
		Attribute("from_address")
		Attribute("return_address")
		Attribute("buyer_address")
		Attribute("parcel")
		Attribute("customs_info")
		Attribute("scan_form")
		Attribute("forms")
		Attribute("insurance")
		Attribute("options")
		Attribute("is_return")
		Attribute("usps_zone")
		Attribute("batch_id")

		Attribute("rates", ArrayOf(Rate), "All associated Rate objects")
		Attribute("selected_rate", Rate, "The specific rate purchased for the shipment, or null if unpurchased or purchased through another mechanism")
		Attribute("postage_label", PostageLabel, "The associated PostageLabel object")
		Attribute("messages", ArrayOf(Message), "Any carrier errors encountered during rating, discussed in more depth below")
		Attribute("tracking_code", String, "If purchased, the tracking code will appear here as well as within the Tracker object")
		Attribute("status", String, "The current tracking status of the shipment")
		Attribute("tracker", Tracker, "The associated Tracker object")
		Attribute("fees", ArrayOf(Fee), "The associated Fee objects charged to the billing user account")
		Attribute("refund_status", String, "The current status of the shipment refund process. Possible values are \"submitted\", \"refunded\", \"rejected\".", func() {
			Enum("submitted", "rejected", "refunded")
		})
		Attribute("batch_status", String, "The current state of the associated BatchShipment")
		Attribute("batch_message", String, "The current message of the associated BatchShipment")
		Attribute("scan_form", ScanForm, "Document created to manifest and scan multiple shipments")

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("mode")
		Attribute("rates")
		Attribute("selected_rate")
		Attribute("postage_label")
		Attribute("messages")
		Attribute("tracking_code")
		Attribute("status")
		Attribute("tracker")
		Attribute("fees")
		Attribute("refund_status")
		Attribute("batch_status")
		Attribute("batch_message")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("to_address")
		Attribute("from_address")
		Attribute("return_address")
		Attribute("buyer_address")
		Attribute("parcel")
		Attribute("customs_info")
		Attribute("scan_form")
		Attribute("forms")
		Attribute("insurance")
		Attribute("options")
		Attribute("is_return")
		Attribute("usps_zone")
		Attribute("batch_id")
	})
})

/**
* Options
**/
var Options = MediaType("application/easypost.options+json", func() {
	Reference(OptionsPayload)
	Attributes(func() {
		Attribute("additional_handling")
		Attribute("address_validation_level")
		Attribute("alcohol")
		Attribute("bill_receiver_account")
		Attribute("bill_receiver_postal_code")
		Attribute("bill_third_party_account")
		Attribute("bill_third_party_country")
		Attribute("bill_third_party_postal_code")
		Attribute("by_drone")
		Attribute("carbon_neutral")
		Attribute("cod_amount")
		Attribute("cod_method")
		Attribute("currency")
		Attribute("delivered_duty_paid")
		Attribute("delivery_confirmation")
		Attribute("dry_ice")
		Attribute("dry_ice_medical")
		Attribute("dry_ice_weight")
		Attribute("freight_charge")
		Attribute("handling_instructions")
		Attribute("hold_for_pickup")
		Attribute("invoice_number")
		Attribute("label_date")
		Attribute("label_format")
		Attribute("machinable")
		Attribute("print_custom_1")
		Attribute("print_custom_2")
		Attribute("print_custom_3")
		Attribute("print_custom_1_barcode")
		Attribute("print_custom_2_barcode")
		Attribute("print_custom_3_barcode")
		Attribute("print_custom_1_code")
		Attribute("print_custom_2_code")
		Attribute("print_custom_3_code")
		Attribute("saturday_delivery")
		Attribute("special_rates_eligibility")
		Attribute("smartpost_hub")
		Attribute("smartpost_manifest")
	})
	View("default", func() {
		Attribute("additional_handling")
		Attribute("address_validation_level")
		Attribute("alcohol")
		Attribute("bill_receiver_account")
		Attribute("bill_receiver_postal_code")
		Attribute("bill_third_party_account")
		Attribute("bill_third_party_country")
		Attribute("bill_third_party_postal_code")
		Attribute("by_drone")
		Attribute("carbon_neutral")
		Attribute("cod_amount")
		Attribute("cod_method")
		Attribute("currency")
		Attribute("delivered_duty_paid")
		Attribute("delivery_confirmation")
		Attribute("dry_ice")
		Attribute("dry_ice_medical")
		Attribute("dry_ice_weight")
		Attribute("freight_charge")
		Attribute("handling_instructions")
		Attribute("hold_for_pickup")
		Attribute("invoice_number")
		Attribute("label_date")
		Attribute("label_format")
		Attribute("machinable")
		Attribute("print_custom_1")
		Attribute("print_custom_2")
		Attribute("print_custom_3")
		Attribute("print_custom_1_barcode")
		Attribute("print_custom_2_barcode")
		Attribute("print_custom_3_barcode")
		Attribute("print_custom_1_code")
		Attribute("print_custom_2_code")
		Attribute("print_custom_3_code")
		Attribute("saturday_delivery")
		Attribute("special_rates_eligibility")
		Attribute("smartpost_hub")
		Attribute("smartpost_manifest")
	})
})

/**
* Rate
**/
var Rate = MediaType("application/easypost.rate+json", func() {
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"rate_\"", func() {
			Pattern("^rate_")
		})
		Attribute("object", String, "Always: \"Rate\"", func() {
			Pattern("^Rate$")
			Default("Rate")
		})
		Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})

		Attribute("service", String, "service level/name")
		Attribute("carrier", String, "name of carrier")
		Attribute("carrier_account_id", String, "ID of the CarrierAccount record used to generate this rate")
		Attribute("shipment_id", String, "ID of the Shipment this rate belongs to")
		Attribute("rate", String, "the actual rate quote for this service")
		Attribute("currency", String, "currency for the rate")
		Attribute("retail_rate", String, "the retail rate is the in-store rate given with no account")
		Attribute("retail_currency", String, "currency for the retail rate")
		Attribute("list_rate", String, "the list rate is the non-negotiated rate given for having an account with the carrier")
		Attribute("list_currency", String, "currency for the list rate")
		Attribute("delivery_days", Integer, "delivery days for this service")
		Attribute("delivery_date", String, "date for delivery")
		Attribute("delivery_date_guaranteed", Boolean, "indicates if delivery window is guaranteed (true) or not (false)")

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("mode")
		Attribute("service")
		Attribute("carrier")
		Attribute("carrier_account_id")
		Attribute("shipment_id")
		Attribute("rate")
		Attribute("currency")
		Attribute("retail_rate")
		Attribute("retail_currency")
		Attribute("list_rate")
		Attribute("list_currency")
		Attribute("delivery_days")
		Attribute("delivery_date")
		Attribute("delivery_date_guaranteed")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

/**
* Message
**/
var Message = MediaType("Message", func() {
	Description("When rating a Shipment or Pickup, some carriers may fail to generate rates. These failures are returned as part of the Shipment or Pickup as part of their messages attribute, and follow a common object structure.")
	Attributes(func() {
		Attribute("carrier", String, "the name of the carrier generating the error, e.g. \"UPS\"")
		Attribute("type", String, "the category of error that occurred. Most frequently \"rate_error\"")
		Attribute("message", String, "the string from the carrier explaining the problem.")
	})
	View("default", func() {
		Attribute("carrier")
		Attribute("type")
		Attribute("message")
	})
})

/**
* Insurance
**/
var Insurance = MediaType("application/easypost.insurance+json", func() {
	Description("An Insurance object represents insurance for packages purchased both va the EasyPost API as well as shipments purchased through third parties and later registered with EasyPost. An Insurance is created automatically whenever you buy a Shipment through EasyPost and pass insurance options during the Buy call or in a later call to Insure a Shipment.")
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"ins_\"", func() {
			Pattern("^ins_")
		})
		Attribute("object", String, "Always: \"Insurance\"", func() {
			Pattern("^Insurance$")
			Default("Insurance")
		})
		Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})

		Attribute("reference", String, "The unique reference for this Insurance, if any")
		Attribute("amount", String, "USD value of insured goods with sub-cent precision")
		Attribute("provider", String, "The insurance provider used by EasyPost")
		Attribute("provider_id", String, "An identifying number for some insurance providers used by EasyPost")
		Attribute("shipment_id", String, "The ID of the Shipment in EasyPost, if postage was purchased via EasyPost")
		Attribute("tracking_code", String, "The tracking code of either the shipment within EasyPost, or provided by you during creation")
		Attribute("status", String, "The current status of the insurance, possible values are \"new\", \"pending\", \"purchased\", \"failed\", or \"cancelled\"", func() {
			Enum("cancelled", "failed", "purchased", "pending", "new")
		})
		Attribute("tracker", Tracker, "The associated Tracker object")
		Attribute("to_address", Address, "The associated Address object for destination")
		Attribute("from_address", Address, "The associated Address object for origin")
		Attribute("fee", Fee, "The associated InsuranceFee object if any")
		Attribute("messages", ArrayOf(String), "The list of errors encountered during attempted purchase of the insurance")

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object")
	})

	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("mode")
		Attribute("reference")
		Attribute("amount")
		Attribute("provider")
		Attribute("provider_id")
		Attribute("shipment_id")
		Attribute("tracking_code")
		Attribute("status")
		Attribute("tracker")
		Attribute("to_address")
		Attribute("from_address")
		Attribute("fee")
		Attribute("messages")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var Insurances = MediaType("application/easypost.insurances+json", func() {
	Description("The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.")
	Attributes(func() {
		Attribute("insurances", ArrayOf(Insurance))
	})
	Required("insurances")
	View("default", func() {
		Attribute("insurances")
	})
})

/**
* Tracking
**/
var Tracker = MediaType("application/easypost.tracker+json", func() {
	Description("The Tracker object contains both the current information about the package as well as previous updates. All of the previous updates are stored in the tracking_details array. Each TrackingDetail object contains the status, the message from the carrier, and a TrackingLocation.")
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"trk_\"", func() {
			Pattern("^trk_")
		})
		Attribute("object", String, "Always: \"Tracker\"", func() {
			Pattern("^Tracker$")
			Default("Tracker")
		})
		Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"", func() {
			Enum("test", "production")
			Default("test")
		})

		Attribute("tracking_code", String, "The tracking code provided by the carrier")
		Attribute("status", String, "The current status of the package, possible values are \"unknown\", \"pre_transit\", \"in_transit\", \"out_for_delivery\", \"delivered\", \"available_for_pickup\", \"return_to_sender\", \"failure\", \"cancelled\" or \"error\"", func() {
			Enum("pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown")
		})
		Attribute("signed_by", String, "The name of the person who signed for the package (if available)")
		Attribute("weight", Number, "The weight of the package as measured by the carrier in ounces (if available)")
		Attribute("est_delivery_date", String, "The estimated delivery date provided by the carrier (if available)")
		Attribute("shipment_id", String, "The id of the EasyPost Shipment object associated with the Tracker (if any)")
		Attribute("carrier", String, "The name of the carrier handling the shipment")
		Attribute("tracking_details", ArrayOf(TrackingDetail), "Array of the associated TrackingDetail objects")
		Attribute("carrier_detail", ArrayOf(CarrierDetail), "The associated CarrierDetail object (if available)")
		Attribute("public_url", String, "URL to a publicly-accessible html page that shows tracking details for this tracker")
		Attribute("fees", ArrayOf(Fee), "Array of the associated Fee objects")

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("mode")

		Attribute("tracking_code")
		Attribute("status")
		Attribute("signed_by")
		Attribute("weight")
		Attribute("est_delivery_date")
		Attribute("shipment_id")
		Attribute("carrier")
		Attribute("tracking_details")
		Attribute("carrier_detail")
		Attribute("public_url")
		Attribute("fees")

		Attribute("created_at")
		Attribute("updated_at")
	})
})
var Trackers = MediaType("application/easypost.trackers+json", func() {
	Attributes(func() {
		Attribute("trackers", ArrayOf(Tracker))
	})
	Required("trackers")
	View("default", func() {
		Attribute("trackers")
	})
})

var TrackingDetail = MediaType("application/easypost.trackingdetail+json", func() {
	Description("Each TrackingDetail object contains the status, the message from the carrier, and a TrackingLocation.")
	Attributes(func() {
		Attribute("object", String, "Always: \"TrackingDetail\"", func() {
			Pattern("^TrackingDetail$")
			Default("TrackingDetail")
		})

		Attribute("message", String, "Description of the scan event")
		Attribute("status", String, "The current status of the package, possible values are \"unknown\", \"pre_transit\", \"in_transit\", \"out_for_delivery\", \"delivered\", \"available_for_pickup\", \"return_to_sender\", \"failure\", \"cancelled\" or \"error\"", func() {
			Enum("pre_transit", "in_transit", "out_for_delivery", "delivered", "available_for_pickup", "return_to_sender", "failure", "cancelled", "error", "unknown")
		})
		Attribute("datetime", String, "The timestamp when the tracking scan occurred")
		Attribute("source", String, "The original source of the information for this scan event, usually the carrier")
		Attribute("tracking_location", TrackingLocation, "The location associated with the scan event")

		Required("object")
	})
	View("default", func() {
		Attribute("object")
		Attribute("message")
		Attribute("status")
		Attribute("datetime")
		Attribute("source")
		Attribute("tracking_location")
	})
})
var TrackingLocation = MediaType("application/easypost.trackinglocation+json", func() {
	Description("The TrackingLocation contains city, state, country, and zip information about the location where the package was scanned. The information each carrier provides is different, so some carriers may not make use of all of these fields.")
	Attributes(func() {
		Attribute("object", String, "Always: \"TrackingLocation\"", func() {
			Pattern("^TrackingLocation$")
			Default("TrackingLocation")
		})

		Attribute("city", String, "The city where the scan event occurred (if available)")
		Attribute("state", String, "The state where the scan event occurred (if available)")
		Attribute("country", String, "	The country where the scan event occurred (if available)")
		Attribute("zip", String, "The postal code where the scan event occurred (if available)")

		Required("object")
	})
	View("default", func() {
		Attribute("object")
		Attribute("city")
		Attribute("state")
		Attribute("country")
		Attribute("zip")
	})
})
var CarrierDetail = MediaType("application/easypost.carrierdetail+json", func() {
	Description("The CarrierDetail object contains the service and container_type of the package. Additionally, it also stores the est_delivery_date_local and est_delivery_time_local, which provide information about the local delivery time.")
	Attributes(func() {
		Attribute("object", String, "Always: \"CarrierDetail\"", func() {
			Pattern("^CarrierDetail$")
			Default("CarrierDetail")
		})

		Attribute("service", String, "The service level the associated shipment was shipped with (if available)")
		Attribute("container_type", String, "The type of container the associated shipment was shipped in (if available)")
		Attribute("est_delivery_date_local", String, "The estimated delivery date as provided by the carrier, in the local time zone (if available)")
		Attribute("est_delivery_time_local", String, "The estimated delivery time as provided by the carrier, in the local time zone (if available)")

		Required("object")
	})
	View("default", func() {
		Attribute("object")
		Attribute("service")
		Attribute("container_type")
		Attribute("est_delivery_date_local")
		Attribute("est_delivery_time_local")
	})
})

/**
* Customs
**/
var CustomsInfo = MediaType("application/easypost.customsinfo+json", func() {
	Description("CustomsInfo objects contain CustomsItem objects and all necessary information for the generation of customs forms required for international shipping.")
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"cstinfo_\"", func() {
			Pattern("^cstinfo_")
		})
		Attribute("object", String, "Always: \"CustomsInfo\"", func() {
			Pattern("^CustomsInfo$")
			Default("CustomsInfo")
		})
		Attribute("eel_pfc", String, "\"EEL\" or \"PFC\" value less than $2500: \"NOEEI 30.37(a)\"; value greater than $2500: see Customs Guide", func() {
			Enum("EEL", "PFC")
		})

		Attribute("contents_type", String, "\"documents\", \"gift\", \"merchandise\", \"returned_goods\", \"sample\", or \"other\"", func() {
			Enum("documents", "gift", "merchandise", "returned_goods", "sample", "other")
		})
		Attribute("contents_explanation", String, "Human readable description of content. Required for certain carriers and always required if contents_type is \"other\"")
		Attribute("customs_certify", Boolean, "Electronically certify the information provided")
		Attribute("customs_signer", String, "Required if customs_certify is true")
		Attribute("non_delivery_option", String, "\"abandon\" or \"return\", defaults to \"return\"", func() {
			Enum("abandon", "return")
			Default("return")
		})
		Attribute("restriction_type", String, "\"none\", \"other\", \"quarantine\", or \"sanitary_phytosanitary_inspection\"", func() {
			Enum("none", "other", "quarantine", "sanitary_phytosanitary_inspection")
		})
		Attribute("restriction_comments", String, "Required if restriction_type is not \"none\"") // Probably should be "other"
		Attribute("customs_items", ArrayOf(CustomItem), "Describes to products being shipped")

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("eel_pfc")

		Attribute("contents_type")
		Attribute("contents_explanation")
		Attribute("customs_certify")
		Attribute("customs_signer")
		Attribute("non_delivery_option")
		Attribute("restriction_type")
		Attribute("restriction_comments")
		Attribute("customs_items")

		Attribute("created_at")
		Attribute("updated_at")
	})
})

var CustomItem = MediaType("application/easypost.customitem+json", func() {
	Description("A CustomsItem object describes goods for international shipment and should be created then included in a CustomsInfo object.")
	Reference(CustomItemPayload)
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"cstitem_\"", func() {
			Pattern("^cstitem_")
		})
		Attribute("object", String, "Always: \"CustomsItem\"", func() {
			Pattern("^CustomsItem$")
			Default("CustomsItem")
		})

		Attribute("description", String, "Required, description of item being shipped")
		Attribute("quantity", String, "Required, greater than zero")
		Attribute("value", String, "Required, greater than zero, total value (unit value * quantity)")
		Attribute("weight", String, "Required, greater than zero, total weight (unit weight * quantity)")
		Attribute("hs_tariff_number", String, "Harmonized Tariff Schedule, e.g. \"6109.10.0012\" for Men's T-shirts")
		Attribute("origin_country", String, "Required, 2 char country code", func() {
			Default("US")
		})
		Attribute("currency", String, "3 char currency code, default USD", func() {
			Default("USD")
		})

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object", "description", "quantity", "value", "weight", "origin_country")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")

		Attribute("description")
		Attribute("quantity")
		Attribute("value")
		Attribute("weight")
		Attribute("hs_tariff_number")
		Attribute("origin_country")
		Attribute("currency")

		Attribute("created_at")
		Attribute("updated_at")
	})
})

/**
* Fee
**/
var Fee = MediaType("application/easypost.fee+json", func() {
	Attributes(func() {
		Attribute("object", String, "Always: \"Fee\"", func() {
			Pattern("^Fee$")
			Default("Fee")
		})

		Attribute("type", String, "The name of the category of fee. Possible types are \"LabelFee\", \"PostageFee\", \"InsuranceFee\", and \"TrackerFee\"", func() {
			Enum("LabelFee", "PostageFee", "InsuranceFee", "TrackerFee")
		})
		Attribute("amount", String, "USD value with sub-cent precision")
		Attribute("charged", Boolean, "Whether EasyPost has successfully charged your account for the fee")
		Attribute("refunded", Boolean, "Whether the Fee has been refunded successfully")
	})
	View("default", func() {
		Attribute("object")
		Attribute("type")
		Attribute("amount")
		Attribute("charged")
		Attribute("refunded")
	})
})

/**
* ScanForm
**/
var ScanForm = MediaType("application/easypost.scanform+json", func() {
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"sf_\"", func() {
			Pattern("^sf_")
		})
		Attribute("object", String, "Always: \"ScanForm\"", func() {
			Pattern("^ScanForm$")
			Default("ScanForm")
		})

		Attribute("status", String, "Current status. Possible values are \"creating\", \"created\" and \"failed\"", func() {
			Enum("creating", "created", "failed")
		})
		Attribute("message", String, "Human readable message explaining any failures")
		Attribute("address", Address, "Address the will be Shipments shipped from")
		Attribute("tracking_codes", ArrayOf(String), "Tracking codes included on the ScanForm")
		Attribute("form_url", String, "	Url of the document")
		Attribute("form_file_type", String, "File format of the document")
		Attribute("batch_id", String, "The id of the associated Batch. Unique, starts with \"batch_\"")
		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")
		Attribute("status")
		Attribute("message")
		Attribute("address")
		Attribute("tracking_codes")
		Attribute("form_url")
		Attribute("form_file_type")
		Attribute("batch_id")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

/**
* PostageLabel
**/

var PostageLabel = MediaType("application/easypost.postagelabel+json", func() {
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"pl_\"", func() {
			Pattern("^pl_")
		})
		Attribute("object", String, "Always: \"PostageLabel\"", func() {
			Pattern("^PostageLabel$")
			Default("PostageLabel")
		})

		Attribute("date_advance", Integer)
		Attribute("integrated_form", String, "", func() {
			Default("none")
		})
		Attribute("label_date", String)
		Attribute("label_resolution", Integer, "Assuming DPI", func() {
			Default(200)
		})
		Attribute("label_size", String, "", func() {
			Default("PAPER_4X6")
		})
		Attribute("label_type", String, "", func() {
			Default("default")
		})
		Attribute("label_url", String)
		Attribute("label_file_type", String, "", func() {
			Default("image/png")
		})
		Attribute("label_pdf_url", String)
		Attribute("label_epl2_url", String)
		Attribute("label_zpl_url", String)

		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")
	})
	View("default", func() {
		Attribute("id")
		Attribute("object")

		Attribute("date_advance")
		Attribute("integrated_form")
		Attribute("label_date")
		Attribute("label_resolution")
		Attribute("label_size")
		Attribute("label_type")
		Attribute("label_url")
		Attribute("label_file_type")
		Attribute("label_pdf_url")
		Attribute("label_epl2_url")
		Attribute("label_zpl_url")

		Attribute("created_at")
		Attribute("updated_at")
	})
})

// TODO: Orders
// TODO: Pickups
// TODO: Events?
