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

// payloads
var CarrierAccountPayload = Type("CarrierAccountPayload", func() {
	Description("A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.")

	Attribute("type", String, "The name of the carrier type. Note that \"EndiciaAccount\" is the current USPS integration account type")
	Attribute("fields", FieldsObjectPayload, "Contains \"credentials\" and/or \"test_credentials\", or may be empty")
	Attribute("clone", Boolean, "If clone is true, only the reference and description are possible to update", func() {
		Default(false)
	})
	Attribute("description", String, "An optional, user-readable field to help distinguish accounts")
	Attribute("reference", String, "An optional field that may be used in place of carrier_account_id in other API endpoints")
	Attribute("readable", String, "The name used when displaying a readable value for the type of the account")
	Attribute("credentials", Any, "Unlike the \"credentials\" object contained in \"fields\", this nullable object contains just raw credential pairs for client library consumption")
	Attribute("test_credentials", Any, "Unlike the \"test_credentials\" object contained in \"fields\", this nullable object contains just raw test_credential pairs for client library consumption")

	Required("type")
})

var FieldsObjectPayload = Type("FieldsObjectPayload", func() {
	Description("Contains \"credentials\" and/or \"test_credentials\", or may be empty")
	Attribute("credentials", Any, "Credentials used in the production environment.")
	Attribute("test_credentials", Any, "Credentials used in the test environment.")
	Attribute("auto_link", Boolean, "For USPS this designates that no credentials are required.", func() {
		Default(false)
	})
	Attribute("custom_workflow", Boolean, "When present, a seperate authentication process will be required through the UI to link this account type.", func() {
		Default(false)
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

// FieldError ...
var FieldError = Type("application/easypost.gield_error+json", func() {
	Attribute("field", String, "Field of the request that the error describes")
	Attribute("message", String, "Human readable description of the problem")
})

var AddressPayload = Type("AddressPayload", func() {
	Description("Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.")

	Attribute("street1", String, "First line of the address")
	Attribute("street2", String, "Second line of the address")
	Attribute("city", String, "City the address is located in")
	Attribute("state", String, "State or province the address is located in")
	Attribute("zip", String, "ZIP or postal code the address is located in")
	Attribute("country", String, "ISO 3166 country code for the country the address is located in")
	Attribute("residential", Boolean, "Whether or not this address would be considered residential")
	Attribute("carrier_facility", String, "The specific designation for the address (only relevant if the address is a carrier facility)")
	Attribute("name", String, "Name of the person. Both name and company can be included")
	Attribute("company", String, "Name of the organization. Both name and company can be included")
	Attribute("phone", String, "Phone number to reach the person or organization")
	Attribute("email", String, "Email to reach the person or organization")
	Attribute("federal_tax_id", String, "Federal tax identifier of the person or organization")
	Attribute("state_tax_id", String, "	State tax identifier of the person or organization")

	Attribute("verify", ArrayOf(String), "The verifications to perform when creating. verify_strict takes precedence")
	Attribute("verify_strict", ArrayOf(String), "The verifications to perform when creating. The failure of any of these verifications causes the whole request to fail")

	Required("street1", "city", "state", "zip", "country")
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

var ParcelPayload = Type("ParcelPayload", func() {
	Description("Parcel objects represent the physical container being shipped. Dimensions can be supplied either as length, width, and height dimensions, or a predefined_package string. If you provide a carrier specific predefined_package the associated Shipment will only fetch rates from that carrier.")
	Attribute("length", Number, "Required if predefined_package is empty. float (inches)")
	Attribute("width", Number, "Required if predefined_package is empty. float (inches)")
	Attribute("height", Number, "Required if predefined_package is empty. float (inches)")
	Attribute("weight", Number, "Always required. float(oz)")
	Attribute("predefined_package", String, "Optional, one of our predefined_packages")

	Required("weight")
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
		Attribute("created_at", String, "Time Created")
		Attribute("updated_at", String, "Time Last Updated")

		Required("id", "object")
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

var ShipmentPayload = Type("ShipmentPayload", func() {
	Description("Parcel objects represent the physical container being shipped. Dimensions can be supplied either as length, width, and height dimensions, or a predefined_package string. If you provide a carrier specific predefined_package the associated Shipment will only fetch rates from that carrier.")

	Attribute("to_address", Address, "The destination address")
	Attribute("from_address", Address, "The origin address")
	Attribute("return_address", Address, "The shipper's address, defaults to from_address")
	Attribute("buyer_address", Address, "The buyer's address, defaults to to_address")
	Attribute("parcel", Parcel, "The dimensions and weight of the package")

	Attribute("customs_info", CustomsInfo, "Information for the processing of customs")
	Attribute("scan_form", ScanForm, "Document created to manifest and scan multiple shipments")
	Attribute("forms", ArrayOf(Any), "All associated Form objects")
	Attribute("insurance", Insurance, "The associated Insurance object")
	Attribute("options", Options, "All of the options passed to the shipment, discussed in more depth below")
	Attribute("is_return", Boolean, "Set true to create as a return, discussed in more depth below")
	Attribute("usps_zone", String, "The USPS zone of the shipment, if purchased with USPS")
	Attribute("batch_id", String, "The ID of the batch that contains this shipment, if any")

	//Required("to_address", "from_address", "parcel")
})

var BuyShipmentPayload = Type("BuyShipmentPayload", func() {
	Attribute("rate", Rate, "Selected rate")
	Attribute("insurance", String, "Additionally, insurance may be added during the purchase. To specify an amount to insure, pass the insurance attribute as a string. The currency of all insurance is USD.")

	Required("rate")
})

var LabelShipmentPayload = Type("LabelShipmentPayload", func() {
	Attribute("file_format", String, "Selected rate", func() {
		Enum("png", "zpl", "epl2", "pdf")
	})
	Required("file_format")
})

/**
* Options
**/
var Options = MediaType("application/easypost.options+json", func() {
	Reference(OptionsPayload)
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
var OptionsPayload = Type("OptionsPayload", func() {
	Description("Shipments can have a variety of additional options which you can specify when creating a shipment. The Options object can be populated with the keys below.")

	Attribute("additional_handling", Boolean, "Setting this option to true, will add an additional handling charge.")
	Attribute("address_validation_level", String, "Setting this option to \"0\", will allow the minimum amount of address information to pass the validation check. Only for USPS postage.")
	Attribute("alcohol", Boolean, "Set this option to true if your shipment contains alcohol. UPS - only supported for US Domestic shipments. FedEx - only supported for US Domestic shipments. Canada Post - Requires adult signature 19 years or older. If you want adult signature 18 years or older, instead use delivery_confirmation: ADULT_SIGNATURE")
	Attribute("bill_receiver_account", String, "Setting an account number of the receiver who is to receive and buy the postage. UPS - bill_receiver_postal_code is also required")
	Attribute("bill_receiver_postal_code", String, "Setting a postal code of the receiver account you want to buy postage. UPS - bill_receiver_account also required")
	Attribute("bill_third_party_account", String, "Setting an account number of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_postal_code also required")
	Attribute("bill_third_party_country", String, "Setting a country of the third party account you want to buy postage. UPS - bill_third_party_account and bill_third_party_postal_code also required")
	Attribute("bill_third_party_postal_code", Boolean, "Setting a postal code of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_account also required")
	Attribute("by_drone", Boolean, "Setting this option to true will indicate to the carrier to prefer delivery by drone, if the carrier supports drone delivery.")
	Attribute("carbon_neutral", Boolean, "Setting this to true will add a charge to reduce carbon emissions.")
	Attribute("cod_amount", String, "Adding an amount will have the carrier collect the specified amount from the recipient.")
	Attribute("cod_method", String, "Method for payment. \"CASH\", \"CHECK\", \"MONEY_ORDER\"", func() {
		Enum("CASH", "CHECK", "MONEY_ORDER")
	})
	Attribute("currency", String, "Which currency this shipment will show for rates if carrier allows.")
	Attribute("delivered_duty_paid", Boolean, "Setting this value to false will pass the cost of duties on to the recipient of the package(s).")
	Attribute("delivery_confirmation", String, "If you want to request a signature, you can pass \"ADULT_SIGNATURE\" or \"SIGNATURE\". You may also request \"NO_SIGNATURE\" to leave the package at the door. All - some options may be limited for international shipments")
	Attribute("dry_ice", Boolean, "Package contents contain dry ice. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set. FedEx - Need dry_ice_weight to be set")
	Attribute("dry_ice_medical", Boolean, "If the dry ice is for medical use, set this option to true. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set")
	Attribute("dry_ice_weight", String, "Weight of the dry ice in ounces. UPS - Need dry_ice to be set. UPS MailInnovations - Need dry_ice to be set. FedEx - Need dry_ice to be set")
	Attribute("freight_charge", Number, "Additional cost to be added to the invoice of this shipment. Only applies to UPS currently.")
	Attribute("handling_instructions", String, "This is to designate special instructions for the carrier like \"Do not drop!\".", func() {
		Enum("ORMD", "LIMITED_QUANTITY")
	})
	Attribute("hold_for_pickup", Boolean, "Package will wait at carrier facility for pickup.")
	Attribute("invoice_number", String, "This will print an invoice number on the postage label.")
	Attribute("label_date", String, "Set the date that will appear on the postage label. Accepts ISO 8601 formatted string including time zone offset.")
	Attribute("label_format", String, "Supported label formats include \"PNG\", \"PDF\", \"ZPL\", and \"EPL2\". \"PNG\" is the only format that allows for conversion.", func() {
		Enum("PNG", "PDF", "ZPL", "EPL2")
	})
	Attribute("machinable", Boolean, "Whether or not the parcel can be processed by the carriers equipment.")
	Attribute("print_custom_1", String, "You can optionally print custom messages on labels. The locations of these fields show up on different spots on the carrier's labels.")
	Attribute("print_custom_2", String, "An additional message on the label. Same restrictions as print_custom_1")
	Attribute("print_custom_3", String, "An additional message on the label. Same restrictions as print_custom_1")
	Attribute("print_custom_1_barcode", Boolean, "Create a barcode for this custom reference if supported by carrier.")
	Attribute("print_custom_2_barcode", Boolean, "Create a barcode for this custom reference if supported by carrier.")
	Attribute("print_custom_3_barcode", Boolean, "Create a barcode for this custom reference if supported by carrier.")
	Attribute("print_custom_1_code", String, "Specify the type of print_custom_1.")
	Attribute("print_custom_2_code", String, "See print_custom_1_code.")
	Attribute("print_custom_3_code", String, "See print_custom_1_code.")
	Attribute("saturday_delivery", Boolean, "Set this value to true for delivery on Saturday. Can't be combined with Next Day Air. When setting the saturday_delivery option, you will only get rates for services that are eligible for saturday delivery. If no services are available for saturday delivery, then you will not be returned any rates. You may need to create 2 shipments, one with the saturday_delivery option set on one without to get all your eligible rates.")
	Attribute("special_rates_eligibility", String, "This option allows you to request restrictive rates from USPS. Can set to 'USPS.MEDIAMAIL' or 'USPS.LIBRARYMAIL'.")
	Attribute("smartpost_hub", String, "You can use this to override the hub ID you have on your account.")
	Attribute("smartpost_manifest", Boolean, "The manifest ID is used to group SmartPost packages onto a manifest for each trailer.")
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

var ShipmentInsurancePayload = Type("InsurancePayload", func() {
	Description("Insuring your Shipment is as simple as sending us the value of the contents. We charge 1% of the value, with a $1 minimum, and handle all the claims. All our claims are paid out within 30 days.")
	Attribute("amount", String)
	Required("amount")

})

var InsurancePayload = Type("InsurancePayload", func() {
	Description("An Insurance created via this endpoint must belong to a shipment purchased outside of EasyPost. Insurance for Shipments created within EasyPost must be created via the Shipment Buy or Insure endpoints.")
	Attribute("to_address", Address, "The actual destination of the package to be insured")
	Attribute("from_address", Address, "The actual origin of the package to be insured")
	Attribute("tracking_code", String, "The tracking code associated with the non-EasyPost-purchased package you'd like to insure")
	Attribute("reference", String, "Optional. A unique value that may be used in place of ID when doing Retrieve calls for this insurance")
	Attribute("amount", String, "The USD value of contents you would like to insure. Currently the maximum is between $5000 and $10000 depending on insurer")
	Attribute("carrier", String, "The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided")

	Required("to_address", "from_address", "tracking_code", "amount")
})

var InsuranceListPayload = Type("InsuranceListPayload", func() {
	Description("The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.")

	Attribute("before_id", Address, "Optional pagination parameter. Only records created before the given ID will be included. May not be used with after_id")
	Attribute("after_id", Address, "Optional pagination parameter. Only records created after the given ID will be included. May not be used with before_id")
	Attribute("start_datetime", String, "Only return records created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime")
	Attribute("end_datetime", String, "Only return records created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime")
	Attribute("page_size", String, "The number of records to return on each page. The maximum value is 100, and default is 20.", func() {
		Minimum(1)
		Maximum(100)
		Default(20)
	})
})

var Insurances = MediaType("application/easypost.insurances+json", func() {
	Description("The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.")
	Attribute("insurances", ArrayOf(Insurance))
	Required("insurances")
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
	Attribute("trackers", ArrayOf(Tracker))
	Required("trackers")
})
var TrackerPayload = Type("TrackerPayload", func() {
	Attribute("tracking_code", String, "The tracking code associated with the package you'd like to track")
	Attribute("carrier", String, "The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided")
	Required("tracking_code")
})
var TrackerListPayload = Type("TrackerListPayload", func() {
	Attribute("before_id", String, "Optional pagination parameter. Only trackers created before the given ID will be included. May not be used with after_id")
	Attribute("after_id", String, "Optional pagination parameter. Only trackers created after the given ID will be included. May not be used with before_id")
	Attribute("start_datetime", String, "Only return Trackers created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime")
	Attribute("end_datetime", String, "Only return Trackers created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime")
	Attribute("page_size", Integer, "The number of Trackers to return on each page. The maximum value is 100", func() {
		Minimum(1)
		Maximum(100)
		Default(1)
	})
	Attribute("tracking_code", String, "Only returns Trackers with the given tracking_code. Useful for retrieving an individual Tracker by tracking_code rather than by ID")
	Attribute("carrier", String, "Only returns Trackers with the given carrier value")

	Required("tracking_code")
})
var TrackingDetail = Type("TrackingDetail", func() {
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
var TrackingLocation = Type("TrackingLocation", func() {
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
var CarrierDetail = Type("CarrierDetail", func() {
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
var CustomsInfoPayload = Type("CustomsInfoPayload", func() {
	Attribute("customs_certify", Boolean)
	Attribute("cutoms_signer", String)
	Attribute("contents_type", String)
	Attribute("restriction_type", String)
	Attribute("eel_pfc", String)
	Attribute("customs_items", ArrayOf(CustomItemPayload))
	Required("customs_certify", "cutoms_signer", "contents_type", "restriction_type", "eel_pfc", "customs_items")
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

var CustomItemPayload = Type("CustomItemPayload", func() {
	Attribute("description", String, "Required, description of item being shipped")
	Attribute("quantity", Number, "Required, greater than zero")
	Attribute("value", Number, "Required, greater than zero, total value (unit value * quantity)")
	Attribute("weight", Number, "Required, greater than zero, total weight (unit weight * quantity)")
	Attribute("hs_tariff_number", String, "Harmonized Tariff Schedule, e.g. \"6109.10.0012\" for Men's T-shirts")
	Attribute("origin_country", String, "Required, 2 char country code", func() {
		Default("US")
	})
	Attribute("currency", String, "3 char currency code, default USD", func() {
		Default("USD")
	})
	Required("description", "quantity", "value", "weight", "origin_country")
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
