package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CarrierAccount = MediaType("application/easypost.carrier_accounts+json", func() {
	Description("A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.")
	Attributes(func() {
		Attribute("id", String, "Unique, begins with \"ca_\"", func() {
			Pattern("^ca_")
		})
		Attribute("object", String, "Always: \"CarrierAccount\"", func() {
			Pattern("^CarrierAccount$")
			Default("CarrierAccount")
		})
		Attribute("type", String, "The name of the carrier type. Note that \"EndiciaAccount\" is the current USPS integration account type")
		Attribute("fields", FieldsObject, "Contains \"credentials\" and/or \"test_credentials\", or may be empty")
		Attribute("clone", Boolean, "If clone is true, only the reference and description are possible to update")
		Attribute("description", String, "An optional, user-readable field to help distinguish accounts")
		Attribute("reference", String, "An optional field that may be used in place of carrier_account_id in other API endpoints")
		Attribute("readable", String, "The name used when displaying a readable value for the type of the account")
		//Attribute("logo", String, "The name used when displaying a readable value for the type of the account") // Not supported yet
		Attribute("credentials", Any, "Unlike the \"credentials\" object contained in \"fields\", this nullable object contains just raw credential pairs for client library consumption")
		Attribute("test_credentials", Any, "Unlike the \"test_credentials\" object contained in \"fields\", this nullable object contains just raw test_credential pairs for client library consumption")
		Attribute("created_at", DateTime, "The name used when displaying a readable value for the type of the account")
		Attribute("updated_at", DateTime, "The name used when displaying a readable value for the type of the account")

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
	Attributes(func() {
		Attribute("key", String, "Each key in the sub-objects of a CarrierAccount's fields is the name of a settable field")
		Attribute("visibility", String, "The visibility value is used to control form field types, and is discussed in the CarrierType section")
		Attribute("label", String, "The label value is used in form rendering to display a more precise field name")
		Attribute("value", String, "Checkbox fields use \"0\" and \"1\" as False and True, all other field types present plaintext, partly-masked, or masked credential data for reference")

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
var CarrierAccountPayload = Type("YesNoPayload", func() {
	Description("A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.")
	Attribute("id", String, "Unique, begins with \"ca_\"", func() {
		Pattern("^ca_")
	})
	Attribute("object", String, "Always: \"CarrierAccount\"", func() {
		Pattern("^CarrierAccount$")
		Default("CarrierAccount")
	})
	Attribute("type", String, "The name of the carrier type. Note that \"EndiciaAccount\" is the current USPS integration account type")
	//Attribute("fields", FieldsObject, "Contains \"credentials\" and/or \"test_credentials\", or may be empty")
	Attribute("clone", Boolean, "If clone is true, only the reference and description are possible to update")
	Attribute("description", String, "An optional, user-readable field to help distinguish accounts")
	Attribute("reference", String, "An optional field that may be used in place of carrier_account_id in other API endpoints")
	Attribute("readable", String, "The name used when displaying a readable value for the type of the account")
	//Attribute("logo", String, "The name used when displaying a readable value for the type of the account") // Not supported yet
	Attribute("credentials", Any, "Unlike the \"credentials\" object contained in \"fields\", this nullable object contains just raw credential pairs for client library consumption")
	Attribute("test_credentials", Any, "Unlike the \"test_credentials\" object contained in \"fields\", this nullable object contains just raw test_credential pairs for client library consumption")
	Attribute("created_at", DateTime, "The name used when displaying a readable value for the type of the account")
	Attribute("updated_at", DateTime, "The name used when displaying a readable value for the type of the account")

	Required("id", "object")
})
