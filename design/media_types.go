package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CarrierAccount = MediaType("application/easypost.carrier_accounts+json", func() {
	Reference(CarrierAccountPayload)
	Description("A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.")
	Attributes(func() {
		Attribute("id")
		Attribute("object")
		Attribute("type")
		Attribute("fields")
		Attribute("clone")
		Attribute("description")
		Attribute("reference")
		Attribute("readable")
		//Attribute("logo") // Not supported yet
		Attribute("credentials")
		Attribute("test_credentials")
		Attribute("created_at")
		Attribute("updated_at")

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
	Attribute("id", String, "Unique, begins with \"ca_\"")
	Attribute("object", String, "Always: \"CarrierAccount\"", func() {
		Pattern("^CarrierAccount$")
		Default("CarrierAccount")
	})
	Attribute("type", String, "The name of the carrier type. Note that \"EndiciaAccount\" is the current USPS integration account type")
	Attribute("fields", FieldsObjectPayload, "Contains \"credentials\" and/or \"test_credentials\", or may be empty")
	Attribute("clone", Boolean, "If clone is true, only the reference and description are possible to update", func() {
		Default(false)
	})
	Attribute("description", String, "An optional, user-readable field to help distinguish accounts")
	Attribute("reference", String, "An optional field that may be used in place of carrier_account_id in other API endpoints")
	Attribute("readable", String, "The name used when displaying a readable value for the type of the account")
	//Attribute("logo", String, "The name used when displaying a readable value for the type of the account") // Not supported yet
	Attribute("credentials", Any, "Unlike the \"credentials\" object contained in \"fields\", this nullable object contains just raw credential pairs for client library consumption")
	Attribute("test_credentials", Any, "Unlike the \"test_credentials\" object contained in \"fields\", this nullable object contains just raw test_credential pairs for client library consumption")
	Attribute("created_at", String, "The name used when displaying a readable value for the type of the account")
	Attribute("updated_at", String, "The name used when displaying a readable value for the type of the account")

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
 */

var Address = MediaType("application/easypost.address+json", func() {
	Reference(AddressPayload)
	Attributes(func() {
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
		Attribute("verifications")

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
		Attribute("verifications")
	})
})

/*// Verifications ...
var Verifications = MediaType("application/easypost.address_verifications+json", func() {
	Reference(VerificationsPayload)
	Attributes(func() {
		Attribute("zip4")
		Attribute("delivery")
	})
	View("default", func() {
		Attribute("zip4")
		Attribute("delivery")
	})
})

// Verification ...
var Verification = MediaType("application/easypost.address_verification+json", func() {
	Reference(VerificationPayload)
	Attributes(func() {
		Attribute("success")
		Attribute("errors")
	})
	View("default", func() {
		Attribute("success")
		Attribute("errors")
	})
})


*/

var AddressPayload = Type("AddressPayload", func() {
	Description("Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.")
	Attribute("id", String, "Unique, begins with \"adr_\"", func() {
		Pattern("^adr_")
	})
	Attribute("object", String, "Always: \"Address\"", func() {
		Pattern("^Address$")
		Default("Address")
	})

	Attribute("mode", String, "Set based on which api-key you used, either \"test\" or \"production\"")
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
	Attribute("verifications", VerificationsPayload, "The result of any verifications requested")

	Required("id", "object")
})

var VerificationsPayload = Type("VerificationsPayload", func() {
	Attribute("zip4", VerificationPayload, "Only applicable to US addresses - checks and sets the ZIP+4.")
	Attribute("delivery", VerificationPayload, "Checks that the address is deliverable and makes minor corrections to spelling/format. US addresses will also have their \"residential\" status checked and set.")
})

var VerificationPayload = Type("VerificationPayload", func() {
	Attribute("success", Boolean, "The success of the verification")
	Attribute("errors", ArrayOf(FieldError), "All errors that caused the verification to fail")
})

// FieldError ...
var FieldError = Type("application/easypost.gield_error+json", func() {
	Attribute("field", String, "Field of the request that the error describes")
	Attribute("message", String, "Human readable description of the problem")
})
