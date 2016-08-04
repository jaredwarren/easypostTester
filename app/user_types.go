//************************************************************************//
// API "easypost": Application User Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/easypostTester/design
// --out=$(GOPATH)\src\github.com\jaredwarren\easypostTester
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.
type addressPayload struct {
	// The specific designation for the address (only relevant if the address is a carrier facility)
	CarrierFacility *string `form:"carrier_facility,omitempty" json:"carrier_facility,omitempty" xml:"carrier_facility,omitempty"`
	// City the address is located in
	City *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// Name of the organization. Both name and company can be included
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// ISO 3166 country code for the country the address is located in
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Email to reach the person or organization
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Federal tax identifier of the person or organization
	FederalTaxID *string `form:"federal_tax_id,omitempty" json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// Name of the person. Both name and company can be included
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Phone number to reach the person or organization
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Whether or not this address would be considered residential
	Residential *bool `form:"residential,omitempty" json:"residential,omitempty" xml:"residential,omitempty"`
	// State or province the address is located in
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// 	State tax identifier of the person or organization
	StateTaxID *string `form:"state_tax_id,omitempty" json:"state_tax_id,omitempty" xml:"state_tax_id,omitempty"`
	// First line of the address
	Street1 *string `form:"street1,omitempty" json:"street1,omitempty" xml:"street1,omitempty"`
	// Second line of the address
	Street2 *string `form:"street2,omitempty" json:"street2,omitempty" xml:"street2,omitempty"`
	// The verifications to perform when creating. verify_strict takes precedence
	Verify []string `form:"verify,omitempty" json:"verify,omitempty" xml:"verify,omitempty"`
	// The verifications to perform when creating. The failure of any of these verifications causes the whole request to fail
	VerifyStrict []string `form:"verify_strict,omitempty" json:"verify_strict,omitempty" xml:"verify_strict,omitempty"`
	// ZIP or postal code the address is located in
	Zip *string `form:"zip,omitempty" json:"zip,omitempty" xml:"zip,omitempty"`
}

// Validate validates the addressPayload type instance.
func (ut *addressPayload) Validate() (err error) {
	if ut.Street1 == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street1"))
	}
	if ut.City == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if ut.State == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}
	if ut.Zip == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "zip"))
	}
	if ut.Country == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}

	return
}

// Publicize creates AddressPayload from addressPayload
func (ut *addressPayload) Publicize() *AddressPayload {
	var pub AddressPayload
	if ut.CarrierFacility != nil {
		pub.CarrierFacility = ut.CarrierFacility
	}
	if ut.City != nil {
		pub.City = *ut.City
	}
	if ut.Company != nil {
		pub.Company = ut.Company
	}
	if ut.Country != nil {
		pub.Country = *ut.Country
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FederalTaxID != nil {
		pub.FederalTaxID = ut.FederalTaxID
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Phone != nil {
		pub.Phone = ut.Phone
	}
	if ut.Residential != nil {
		pub.Residential = ut.Residential
	}
	if ut.State != nil {
		pub.State = *ut.State
	}
	if ut.StateTaxID != nil {
		pub.StateTaxID = ut.StateTaxID
	}
	if ut.Street1 != nil {
		pub.Street1 = *ut.Street1
	}
	if ut.Street2 != nil {
		pub.Street2 = ut.Street2
	}
	if ut.Verify != nil {
		pub.Verify = ut.Verify
	}
	if ut.VerifyStrict != nil {
		pub.VerifyStrict = ut.VerifyStrict
	}
	if ut.Zip != nil {
		pub.Zip = *ut.Zip
	}
	return &pub
}

// Address objects are used to represent people, places, and organizations in a number of contexts. For example, a Shipment requires a to_address and from_address to accurately calculate rates and generate postage.
type AddressPayload struct {
	// The specific designation for the address (only relevant if the address is a carrier facility)
	CarrierFacility *string `form:"carrier_facility,omitempty" json:"carrier_facility,omitempty" xml:"carrier_facility,omitempty"`
	// City the address is located in
	City string `form:"city" json:"city" xml:"city"`
	// Name of the organization. Both name and company can be included
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// ISO 3166 country code for the country the address is located in
	Country string `form:"country" json:"country" xml:"country"`
	// Email to reach the person or organization
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Federal tax identifier of the person or organization
	FederalTaxID *string `form:"federal_tax_id,omitempty" json:"federal_tax_id,omitempty" xml:"federal_tax_id,omitempty"`
	// Name of the person. Both name and company can be included
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Phone number to reach the person or organization
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Whether or not this address would be considered residential
	Residential *bool `form:"residential,omitempty" json:"residential,omitempty" xml:"residential,omitempty"`
	// State or province the address is located in
	State string `form:"state" json:"state" xml:"state"`
	// 	State tax identifier of the person or organization
	StateTaxID *string `form:"state_tax_id,omitempty" json:"state_tax_id,omitempty" xml:"state_tax_id,omitempty"`
	// First line of the address
	Street1 string `form:"street1" json:"street1" xml:"street1"`
	// Second line of the address
	Street2 *string `form:"street2,omitempty" json:"street2,omitempty" xml:"street2,omitempty"`
	// The verifications to perform when creating. verify_strict takes precedence
	Verify []string `form:"verify,omitempty" json:"verify,omitempty" xml:"verify,omitempty"`
	// The verifications to perform when creating. The failure of any of these verifications causes the whole request to fail
	VerifyStrict []string `form:"verify_strict,omitempty" json:"verify_strict,omitempty" xml:"verify_strict,omitempty"`
	// ZIP or postal code the address is located in
	Zip string `form:"zip" json:"zip" xml:"zip"`
}

// Validate validates the AddressPayload type instance.
func (ut *AddressPayload) Validate() (err error) {
	if ut.Street1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "street1"))
	}
	if ut.City == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "city"))
	}
	if ut.State == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}
	if ut.Zip == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "zip"))
	}
	if ut.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}

	return
}

// Buy Shipment Payload
type buyShipmentPayload struct {
	// Additionally, insurance may be added during the purchase. To specify an amount to insure, pass the insurance attribute as a string. The currency of all insurance is USD.
	Insurance *string `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Selected rate
	Rate *ratePayload `form:"rate,omitempty" json:"rate,omitempty" xml:"rate,omitempty"`
}

// Validate validates the buyShipmentPayload type instance.
func (ut *buyShipmentPayload) Validate() (err error) {
	if ut.Rate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rate"))
	}

	if ut.Rate != nil {
		if err2 := ut.Rate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates BuyShipmentPayload from buyShipmentPayload
func (ut *buyShipmentPayload) Publicize() *BuyShipmentPayload {
	var pub BuyShipmentPayload
	if ut.Insurance != nil {
		pub.Insurance = ut.Insurance
	}
	if ut.Rate != nil {
		pub.Rate = ut.Rate.Publicize()
	}
	return &pub
}

// Buy Shipment Payload
type BuyShipmentPayload struct {
	// Additionally, insurance may be added during the purchase. To specify an amount to insure, pass the insurance attribute as a string. The currency of all insurance is USD.
	Insurance *string `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Selected rate
	Rate *RatePayload `form:"rate" json:"rate" xml:"rate"`
}

// Validate validates the BuyShipmentPayload type instance.
func (ut *BuyShipmentPayload) Validate() (err error) {
	if ut.Rate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rate"))
	}

	if ut.Rate != nil {
		if err2 := ut.Rate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.
type carrierAccountPayload struct {
	// If clone is true, only the reference and description are possible to update
	Clone *bool `form:"clone,omitempty" json:"clone,omitempty" xml:"clone,omitempty"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *fieldsObjectPayload `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// An optional field that may be used in place of carrier_account_id in other API endpoints
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Finalize sets the default values for carrierAccountPayload type instance.
func (ut *carrierAccountPayload) Finalize() {
	var defaultClone = false
	if ut.Clone == nil {
		ut.Clone = &defaultClone
	}
	if ut.Fields != nil {
		var defaultAutoLink = false
		if ut.Fields.AutoLink == nil {
			ut.Fields.AutoLink = &defaultAutoLink
		}
		var defaultCustomWorkflow = false
		if ut.Fields.CustomWorkflow == nil {
			ut.Fields.CustomWorkflow = &defaultCustomWorkflow
		}
	}
}

// Validate validates the carrierAccountPayload type instance.
func (ut *carrierAccountPayload) Validate() (err error) {
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// Publicize creates CarrierAccountPayload from carrierAccountPayload
func (ut *carrierAccountPayload) Publicize() *CarrierAccountPayload {
	var pub CarrierAccountPayload
	if ut.Clone != nil {
		pub.Clone = *ut.Clone
	}
	if ut.Credentials != nil {
		pub.Credentials = ut.Credentials
	}
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Fields != nil {
		pub.Fields = ut.Fields.Publicize()
	}
	if ut.Readable != nil {
		pub.Readable = ut.Readable
	}
	if ut.Reference != nil {
		pub.Reference = ut.Reference
	}
	if ut.TestCredentials != nil {
		pub.TestCredentials = ut.TestCredentials
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	return &pub
}

// A CarrierAccount encapsulates your credentials with the carrier. The CarrierAccount object provides CRUD operations for all CarrierAccounts.
type CarrierAccountPayload struct {
	// If clone is true, only the reference and description are possible to update
	Clone bool `form:"clone" json:"clone" xml:"clone"`
	// Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// An optional, user-readable field to help distinguish accounts
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Contains "credentials" and/or "test_credentials", or may be empty
	Fields *FieldsObjectPayload `form:"fields,omitempty" json:"fields,omitempty" xml:"fields,omitempty"`
	// The name used when displaying a readable value for the type of the account
	Readable *string `form:"readable,omitempty" json:"readable,omitempty" xml:"readable,omitempty"`
	// An optional field that may be used in place of carrier_account_id in other API endpoints
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
	// The name of the carrier type. Note that "EndiciaAccount" is the current USPS integration account type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the CarrierAccountPayload type instance.
func (ut *CarrierAccountPayload) Validate() (err error) {
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// customItemPayload user type.
type customItemPayload struct {
	// 3 char currency code, default USD
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Required, description of item being shipped
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Harmonized Tariff Schedule, e.g. "6109.10.0012" for Men's T-shirts
	HsTariffNumber *string `form:"hs_tariff_number,omitempty" json:"hs_tariff_number,omitempty" xml:"hs_tariff_number,omitempty"`
	// Required, 2 char country code
	OriginCountry *string `form:"origin_country,omitempty" json:"origin_country,omitempty" xml:"origin_country,omitempty"`
	// Required, greater than zero
	Quantity *float64 `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Required, greater than zero, total value (unit value * quantity)
	Value *float64 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Required, greater than zero, total weight (unit weight * quantity)
	Weight *float64 `form:"weight,omitempty" json:"weight,omitempty" xml:"weight,omitempty"`
}

// Finalize sets the default values for customItemPayload type instance.
func (ut *customItemPayload) Finalize() {
	var defaultCurrency = "USD"
	if ut.Currency == nil {
		ut.Currency = &defaultCurrency
	}
	var defaultOriginCountry = "US"
	if ut.OriginCountry == nil {
		ut.OriginCountry = &defaultOriginCountry
	}
}

// Validate validates the customItemPayload type instance.
func (ut *customItemPayload) Validate() (err error) {
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ut.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "quantity"))
	}
	if ut.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	if ut.Weight == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "weight"))
	}
	if ut.OriginCountry == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "origin_country"))
	}

	return
}

// Publicize creates CustomItemPayload from customItemPayload
func (ut *customItemPayload) Publicize() *CustomItemPayload {
	var pub CustomItemPayload
	if ut.Currency != nil {
		pub.Currency = *ut.Currency
	}
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.HsTariffNumber != nil {
		pub.HsTariffNumber = ut.HsTariffNumber
	}
	if ut.OriginCountry != nil {
		pub.OriginCountry = *ut.OriginCountry
	}
	if ut.Quantity != nil {
		pub.Quantity = *ut.Quantity
	}
	if ut.Value != nil {
		pub.Value = *ut.Value
	}
	if ut.Weight != nil {
		pub.Weight = *ut.Weight
	}
	return &pub
}

// CustomItemPayload user type.
type CustomItemPayload struct {
	// 3 char currency code, default USD
	Currency string `form:"currency" json:"currency" xml:"currency"`
	// Required, description of item being shipped
	Description string `form:"description" json:"description" xml:"description"`
	// Harmonized Tariff Schedule, e.g. "6109.10.0012" for Men's T-shirts
	HsTariffNumber *string `form:"hs_tariff_number,omitempty" json:"hs_tariff_number,omitempty" xml:"hs_tariff_number,omitempty"`
	// Required, 2 char country code
	OriginCountry string `form:"origin_country" json:"origin_country" xml:"origin_country"`
	// Required, greater than zero
	Quantity float64 `form:"quantity" json:"quantity" xml:"quantity"`
	// Required, greater than zero, total value (unit value * quantity)
	Value float64 `form:"value" json:"value" xml:"value"`
	// Required, greater than zero, total weight (unit weight * quantity)
	Weight float64 `form:"weight" json:"weight" xml:"weight"`
}

// Validate validates the CustomItemPayload type instance.
func (ut *CustomItemPayload) Validate() (err error) {
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ut.OriginCountry == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "origin_country"))
	}

	return
}

// customsInfoPayload user type.
type customsInfoPayload struct {
	ContentsType    *string              `form:"contents_type,omitempty" json:"contents_type,omitempty" xml:"contents_type,omitempty"`
	CustomsCertify  *bool                `form:"customs_certify,omitempty" json:"customs_certify,omitempty" xml:"customs_certify,omitempty"`
	CustomsItems    []*customItemPayload `form:"customs_items,omitempty" json:"customs_items,omitempty" xml:"customs_items,omitempty"`
	CutomsSigner    *string              `form:"cutoms_signer,omitempty" json:"cutoms_signer,omitempty" xml:"cutoms_signer,omitempty"`
	EelPfc          *string              `form:"eel_pfc,omitempty" json:"eel_pfc,omitempty" xml:"eel_pfc,omitempty"`
	RestrictionType *string              `form:"restriction_type,omitempty" json:"restriction_type,omitempty" xml:"restriction_type,omitempty"`
}

// Finalize sets the default values for customsInfoPayload type instance.
func (ut *customsInfoPayload) Finalize() {
	for _, e := range ut.CustomsItems {
		var defaultCurrency = "USD"
		if e.Currency == nil {
			e.Currency = &defaultCurrency
		}
		var defaultOriginCountry = "US"
		if e.OriginCountry == nil {
			e.OriginCountry = &defaultOriginCountry
		}
	}
}

// Validate validates the customsInfoPayload type instance.
func (ut *customsInfoPayload) Validate() (err error) {
	if ut.CustomsCertify == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "customs_certify"))
	}
	if ut.CutomsSigner == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "cutoms_signer"))
	}
	if ut.ContentsType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "contents_type"))
	}
	if ut.RestrictionType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "restriction_type"))
	}
	if ut.EelPfc == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eel_pfc"))
	}
	if ut.CustomsItems == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "customs_items"))
	}

	for _, e := range ut.CustomsItems {
		if e.Description == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "description"))
		}
		if e.Quantity == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "quantity"))
		}
		if e.Value == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "value"))
		}
		if e.Weight == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "weight"))
		}
		if e.OriginCountry == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "origin_country"))
		}

	}
	return
}

// Publicize creates CustomsInfoPayload from customsInfoPayload
func (ut *customsInfoPayload) Publicize() *CustomsInfoPayload {
	var pub CustomsInfoPayload
	if ut.ContentsType != nil {
		pub.ContentsType = *ut.ContentsType
	}
	if ut.CustomsCertify != nil {
		pub.CustomsCertify = *ut.CustomsCertify
	}
	if ut.CustomsItems != nil {
		pub.CustomsItems = make([]*CustomItemPayload, len(ut.CustomsItems))
		for i2, elem2 := range ut.CustomsItems {
			pub.CustomsItems[i2] = elem2.Publicize()
		}
	}
	if ut.CutomsSigner != nil {
		pub.CutomsSigner = *ut.CutomsSigner
	}
	if ut.EelPfc != nil {
		pub.EelPfc = *ut.EelPfc
	}
	if ut.RestrictionType != nil {
		pub.RestrictionType = *ut.RestrictionType
	}
	return &pub
}

// CustomsInfoPayload user type.
type CustomsInfoPayload struct {
	ContentsType    string               `form:"contents_type" json:"contents_type" xml:"contents_type"`
	CustomsCertify  bool                 `form:"customs_certify" json:"customs_certify" xml:"customs_certify"`
	CustomsItems    []*CustomItemPayload `form:"customs_items" json:"customs_items" xml:"customs_items"`
	CutomsSigner    string               `form:"cutoms_signer" json:"cutoms_signer" xml:"cutoms_signer"`
	EelPfc          string               `form:"eel_pfc" json:"eel_pfc" xml:"eel_pfc"`
	RestrictionType string               `form:"restriction_type" json:"restriction_type" xml:"restriction_type"`
}

// Validate validates the CustomsInfoPayload type instance.
func (ut *CustomsInfoPayload) Validate() (err error) {
	if ut.CutomsSigner == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "cutoms_signer"))
	}
	if ut.ContentsType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "contents_type"))
	}
	if ut.RestrictionType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "restriction_type"))
	}
	if ut.EelPfc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eel_pfc"))
	}
	if ut.CustomsItems == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "customs_items"))
	}

	for _, e := range ut.CustomsItems {
		if e.Description == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "description"))
		}
		if e.OriginCountry == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.customs_items[*]`, "origin_country"))
		}

	}
	return
}

// Contains "credentials" and/or "test_credentials", or may be empty
type fieldsObjectPayload struct {
	// For USPS this designates that no credentials are required.
	AutoLink *bool `form:"auto_link,omitempty" json:"auto_link,omitempty" xml:"auto_link,omitempty"`
	// Credentials used in the production environment.
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// When present, a seperate authentication process will be required through the UI to link this account type.
	CustomWorkflow *bool `form:"custom_workflow,omitempty" json:"custom_workflow,omitempty" xml:"custom_workflow,omitempty"`
	// Credentials used in the test environment.
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
}

// Finalize sets the default values for fieldsObjectPayload type instance.
func (ut *fieldsObjectPayload) Finalize() {
	var defaultAutoLink = false
	if ut.AutoLink == nil {
		ut.AutoLink = &defaultAutoLink
	}
	var defaultCustomWorkflow = false
	if ut.CustomWorkflow == nil {
		ut.CustomWorkflow = &defaultCustomWorkflow
	}
}

// Publicize creates FieldsObjectPayload from fieldsObjectPayload
func (ut *fieldsObjectPayload) Publicize() *FieldsObjectPayload {
	var pub FieldsObjectPayload
	if ut.AutoLink != nil {
		pub.AutoLink = *ut.AutoLink
	}
	if ut.Credentials != nil {
		pub.Credentials = ut.Credentials
	}
	if ut.CustomWorkflow != nil {
		pub.CustomWorkflow = *ut.CustomWorkflow
	}
	if ut.TestCredentials != nil {
		pub.TestCredentials = ut.TestCredentials
	}
	return &pub
}

// Contains "credentials" and/or "test_credentials", or may be empty
type FieldsObjectPayload struct {
	// For USPS this designates that no credentials are required.
	AutoLink bool `form:"auto_link" json:"auto_link" xml:"auto_link"`
	// Credentials used in the production environment.
	Credentials *interface{} `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
	// When present, a seperate authentication process will be required through the UI to link this account type.
	CustomWorkflow bool `form:"custom_workflow" json:"custom_workflow" xml:"custom_workflow"`
	// Credentials used in the test environment.
	TestCredentials *interface{} `form:"test_credentials,omitempty" json:"test_credentials,omitempty" xml:"test_credentials,omitempty"`
}

// The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.
type insuranceListPayload struct {
	// Optional pagination parameter. Only records created after the given ID will be included. May not be used with before_id
	AfterID *addressPayload `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only records created before the given ID will be included. May not be used with after_id
	BeforeID *addressPayload `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only return records created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of records to return on each page. The maximum value is 100, and default is 20.
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Only return records created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
}

// Finalize sets the default values for insuranceListPayload type instance.
func (ut *insuranceListPayload) Finalize() {
	var defaultPageSize = 20
	if ut.PageSize == nil {
		ut.PageSize = &defaultPageSize
	}
}

// Validate validates the insuranceListPayload type instance.
func (ut *insuranceListPayload) Validate() (err error) {
	if ut.AfterID != nil {
		if err2 := ut.AfterID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.BeforeID != nil {
		if err2 := ut.BeforeID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.PageSize != nil {
		if *ut.PageSize < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, *ut.PageSize, 1, true))
		}
	}
	if ut.PageSize != nil {
		if *ut.PageSize > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, *ut.PageSize, 100, false))
		}
	}
	return
}

// Publicize creates InsuranceListPayload from insuranceListPayload
func (ut *insuranceListPayload) Publicize() *InsuranceListPayload {
	var pub InsuranceListPayload
	if ut.AfterID != nil {
		pub.AfterID = ut.AfterID.Publicize()
	}
	if ut.BeforeID != nil {
		pub.BeforeID = ut.BeforeID.Publicize()
	}
	if ut.EndDatetime != nil {
		pub.EndDatetime = ut.EndDatetime
	}
	if ut.PageSize != nil {
		pub.PageSize = *ut.PageSize
	}
	if ut.StartDatetime != nil {
		pub.StartDatetime = ut.StartDatetime
	}
	return &pub
}

// The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.
type InsuranceListPayload struct {
	// Optional pagination parameter. Only records created after the given ID will be included. May not be used with before_id
	AfterID *AddressPayload `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only records created before the given ID will be included. May not be used with after_id
	BeforeID *AddressPayload `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only return records created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of records to return on each page. The maximum value is 100, and default is 20.
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
	// Only return records created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
}

// Validate validates the InsuranceListPayload type instance.
func (ut *InsuranceListPayload) Validate() (err error) {
	if ut.AfterID != nil {
		if err2 := ut.AfterID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.BeforeID != nil {
		if err2 := ut.BeforeID.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.PageSize < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, ut.PageSize, 1, true))
	}
	if ut.PageSize > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, ut.PageSize, 100, false))
	}
	return
}

// An Insurance created via this endpoint must belong to a shipment purchased outside of EasyPost. Insurance for Shipments created within EasyPost must be created via the Shipment Buy or Insure endpoints.
type insurancePayload struct {
	// The USD value of contents you would like to insure. Currently the maximum is between $5000 and $10000 depending on insurer
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The actual origin of the package to be insured
	FromAddress *addressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// Optional. A unique value that may be used in place of ID when doing Retrieve calls for this insurance
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// The actual destination of the package to be insured
	ToAddress *addressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The tracking code associated with the non-EasyPost-purchased package you'd like to insure
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}

// Validate validates the insurancePayload type instance.
func (ut *insurancePayload) Validate() (err error) {
	if ut.ToAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "to_address"))
	}
	if ut.FromAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "from_address"))
	}
	if ut.TrackingCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tracking_code"))
	}
	if ut.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}

	if ut.FromAddress != nil {
		if err2 := ut.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.ToAddress != nil {
		if err2 := ut.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates InsurancePayload from insurancePayload
func (ut *insurancePayload) Publicize() *InsurancePayload {
	var pub InsurancePayload
	if ut.Amount != nil {
		pub.Amount = *ut.Amount
	}
	if ut.Carrier != nil {
		pub.Carrier = ut.Carrier
	}
	if ut.FromAddress != nil {
		pub.FromAddress = ut.FromAddress.Publicize()
	}
	if ut.Reference != nil {
		pub.Reference = ut.Reference
	}
	if ut.ToAddress != nil {
		pub.ToAddress = ut.ToAddress.Publicize()
	}
	if ut.TrackingCode != nil {
		pub.TrackingCode = *ut.TrackingCode
	}
	return &pub
}

// An Insurance created via this endpoint must belong to a shipment purchased outside of EasyPost. Insurance for Shipments created within EasyPost must be created via the Shipment Buy or Insure endpoints.
type InsurancePayload struct {
	// The USD value of contents you would like to insure. Currently the maximum is between $5000 and $10000 depending on insurer
	Amount string `form:"amount" json:"amount" xml:"amount"`
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The actual origin of the package to be insured
	FromAddress *AddressPayload `form:"from_address" json:"from_address" xml:"from_address"`
	// Optional. A unique value that may be used in place of ID when doing Retrieve calls for this insurance
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// The actual destination of the package to be insured
	ToAddress *AddressPayload `form:"to_address" json:"to_address" xml:"to_address"`
	// The tracking code associated with the non-EasyPost-purchased package you'd like to insure
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// Validate validates the InsurancePayload type instance.
func (ut *InsurancePayload) Validate() (err error) {
	if ut.ToAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "to_address"))
	}
	if ut.FromAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "from_address"))
	}
	if ut.TrackingCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tracking_code"))
	}
	if ut.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}

	if ut.FromAddress != nil {
		if err2 := ut.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.ToAddress != nil {
		if err2 := ut.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Label Shipment Payload
type labelShipmentPayload struct {
	// Selected rate
	FileFormat *string `form:"file_format,omitempty" json:"file_format,omitempty" xml:"file_format,omitempty"`
}

// Validate validates the labelShipmentPayload type instance.
func (ut *labelShipmentPayload) Validate() (err error) {
	if ut.FileFormat == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "file_format"))
	}

	if ut.FileFormat != nil {
		if !(*ut.FileFormat == "png" || *ut.FileFormat == "zpl" || *ut.FileFormat == "epl2" || *ut.FileFormat == "pdf") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.file_format`, *ut.FileFormat, []interface{}{"png", "zpl", "epl2", "pdf"}))
		}
	}
	return
}

// Publicize creates LabelShipmentPayload from labelShipmentPayload
func (ut *labelShipmentPayload) Publicize() *LabelShipmentPayload {
	var pub LabelShipmentPayload
	if ut.FileFormat != nil {
		pub.FileFormat = *ut.FileFormat
	}
	return &pub
}

// Label Shipment Payload
type LabelShipmentPayload struct {
	// Selected rate
	FileFormat string `form:"file_format" json:"file_format" xml:"file_format"`
}

// Validate validates the LabelShipmentPayload type instance.
func (ut *LabelShipmentPayload) Validate() (err error) {
	if ut.FileFormat == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "file_format"))
	}

	if !(ut.FileFormat == "png" || ut.FileFormat == "zpl" || ut.FileFormat == "epl2" || ut.FileFormat == "pdf") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.file_format`, ut.FileFormat, []interface{}{"png", "zpl", "epl2", "pdf"}))
	}
	return
}

// Shipments can have a variety of additional options which you can specify when creating a shipment. The Options object can be populated with the keys below.
type optionsPayload struct {
	// Setting this option to true, will add an additional handling charge.
	AdditionalHandling *bool `form:"additional_handling,omitempty" json:"additional_handling,omitempty" xml:"additional_handling,omitempty"`
	// Setting this option to "0", will allow the minimum amount of address information to pass the validation check. Only for USPS postage.
	AddressValidationLevel *string `form:"address_validation_level,omitempty" json:"address_validation_level,omitempty" xml:"address_validation_level,omitempty"`
	// Set this option to true if your shipment contains alcohol. UPS - only supported for US Domestic shipments. FedEx - only supported for US Domestic shipments. Canada Post - Requires adult signature 19 years or older. If you want adult signature 18 years or older, instead use delivery_confirmation: ADULT_SIGNATURE
	Alcohol *bool `form:"alcohol,omitempty" json:"alcohol,omitempty" xml:"alcohol,omitempty"`
	// Setting an account number of the receiver who is to receive and buy the postage. UPS - bill_receiver_postal_code is also required
	BillReceiverAccount *string `form:"bill_receiver_account,omitempty" json:"bill_receiver_account,omitempty" xml:"bill_receiver_account,omitempty"`
	// Setting a postal code of the receiver account you want to buy postage. UPS - bill_receiver_account also required
	BillReceiverPostalCode *string `form:"bill_receiver_postal_code,omitempty" json:"bill_receiver_postal_code,omitempty" xml:"bill_receiver_postal_code,omitempty"`
	// Setting an account number of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_postal_code also required
	BillThirdPartyAccount *string `form:"bill_third_party_account,omitempty" json:"bill_third_party_account,omitempty" xml:"bill_third_party_account,omitempty"`
	// Setting a country of the third party account you want to buy postage. UPS - bill_third_party_account and bill_third_party_postal_code also required
	BillThirdPartyCountry *string `form:"bill_third_party_country,omitempty" json:"bill_third_party_country,omitempty" xml:"bill_third_party_country,omitempty"`
	// Setting a postal code of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_account also required
	BillThirdPartyPostalCode *bool `form:"bill_third_party_postal_code,omitempty" json:"bill_third_party_postal_code,omitempty" xml:"bill_third_party_postal_code,omitempty"`
	// Setting this option to true will indicate to the carrier to prefer delivery by drone, if the carrier supports drone delivery.
	ByDrone *bool `form:"by_drone,omitempty" json:"by_drone,omitempty" xml:"by_drone,omitempty"`
	// Setting this to true will add a charge to reduce carbon emissions.
	CarbonNeutral *bool `form:"carbon_neutral,omitempty" json:"carbon_neutral,omitempty" xml:"carbon_neutral,omitempty"`
	// Adding an amount will have the carrier collect the specified amount from the recipient.
	CodAmount *string `form:"cod_amount,omitempty" json:"cod_amount,omitempty" xml:"cod_amount,omitempty"`
	// Method for payment. "CASH", "CHECK", "MONEY_ORDER"
	CodMethod *string `form:"cod_method,omitempty" json:"cod_method,omitempty" xml:"cod_method,omitempty"`
	// Which currency this shipment will show for rates if carrier allows.
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Setting this value to false will pass the cost of duties on to the recipient of the package(s).
	DeliveredDutyPaid *bool `form:"delivered_duty_paid,omitempty" json:"delivered_duty_paid,omitempty" xml:"delivered_duty_paid,omitempty"`
	// If you want to request a signature, you can pass "ADULT_SIGNATURE" or "SIGNATURE". You may also request "NO_SIGNATURE" to leave the package at the door. All - some options may be limited for international shipments
	DeliveryConfirmation *string `form:"delivery_confirmation,omitempty" json:"delivery_confirmation,omitempty" xml:"delivery_confirmation,omitempty"`
	// Package contents contain dry ice. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set. FedEx - Need dry_ice_weight to be set
	DryIce *bool `form:"dry_ice,omitempty" json:"dry_ice,omitempty" xml:"dry_ice,omitempty"`
	// If the dry ice is for medical use, set this option to true. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set
	DryIceMedical *bool `form:"dry_ice_medical,omitempty" json:"dry_ice_medical,omitempty" xml:"dry_ice_medical,omitempty"`
	// Weight of the dry ice in ounces. UPS - Need dry_ice to be set. UPS MailInnovations - Need dry_ice to be set. FedEx - Need dry_ice to be set
	DryIceWeight *string `form:"dry_ice_weight,omitempty" json:"dry_ice_weight,omitempty" xml:"dry_ice_weight,omitempty"`
	// Additional cost to be added to the invoice of this shipment. Only applies to UPS currently.
	FreightCharge *float64 `form:"freight_charge,omitempty" json:"freight_charge,omitempty" xml:"freight_charge,omitempty"`
	// This is to designate special instructions for the carrier like "Do not drop!".
	HandlingInstructions *string `form:"handling_instructions,omitempty" json:"handling_instructions,omitempty" xml:"handling_instructions,omitempty"`
	// Package will wait at carrier facility for pickup.
	HoldForPickup *bool `form:"hold_for_pickup,omitempty" json:"hold_for_pickup,omitempty" xml:"hold_for_pickup,omitempty"`
	// This will print an invoice number on the postage label.
	InvoiceNumber *string `form:"invoice_number,omitempty" json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// Set the date that will appear on the postage label. Accepts ISO 8601 formatted string including time zone offset.
	LabelDate *string `form:"label_date,omitempty" json:"label_date,omitempty" xml:"label_date,omitempty"`
	// Supported label formats include "PNG", "PDF", "ZPL", and "EPL2". "PNG" is the only format that allows for conversion.
	LabelFormat *string `form:"label_format,omitempty" json:"label_format,omitempty" xml:"label_format,omitempty"`
	// Whether or not the parcel can be processed by the carriers equipment.
	Machinable *bool `form:"machinable,omitempty" json:"machinable,omitempty" xml:"machinable,omitempty"`
	// You can optionally print custom messages on labels. The locations of these fields show up on different spots on the carrier's labels.
	PrintCustom1 *string `form:"print_custom_1,omitempty" json:"print_custom_1,omitempty" xml:"print_custom_1,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom1Barcode *bool `form:"print_custom_1_barcode,omitempty" json:"print_custom_1_barcode,omitempty" xml:"print_custom_1_barcode,omitempty"`
	// Specify the type of print_custom_1.
	PrintCustom1Code *string `form:"print_custom_1_code,omitempty" json:"print_custom_1_code,omitempty" xml:"print_custom_1_code,omitempty"`
	// An additional message on the label. Same restrictions as print_custom_1
	PrintCustom2 *string `form:"print_custom_2,omitempty" json:"print_custom_2,omitempty" xml:"print_custom_2,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom2Barcode *bool `form:"print_custom_2_barcode,omitempty" json:"print_custom_2_barcode,omitempty" xml:"print_custom_2_barcode,omitempty"`
	// See print_custom_1_code.
	PrintCustom2Code *string `form:"print_custom_2_code,omitempty" json:"print_custom_2_code,omitempty" xml:"print_custom_2_code,omitempty"`
	// An additional message on the label. Same restrictions as print_custom_1
	PrintCustom3 *string `form:"print_custom_3,omitempty" json:"print_custom_3,omitempty" xml:"print_custom_3,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom3Barcode *bool `form:"print_custom_3_barcode,omitempty" json:"print_custom_3_barcode,omitempty" xml:"print_custom_3_barcode,omitempty"`
	// See print_custom_1_code.
	PrintCustom3Code *string `form:"print_custom_3_code,omitempty" json:"print_custom_3_code,omitempty" xml:"print_custom_3_code,omitempty"`
	// Set this value to true for delivery on Saturday. Can't be combined with Next Day Air. When setting the saturday_delivery option, you will only get rates for services that are eligible for saturday delivery. If no services are available for saturday delivery, then you will not be returned any rates. You may need to create 2 shipments, one with the saturday_delivery option set on one without to get all your eligible rates.
	SaturdayDelivery *bool `form:"saturday_delivery,omitempty" json:"saturday_delivery,omitempty" xml:"saturday_delivery,omitempty"`
	// You can use this to override the hub ID you have on your account.
	SmartpostHub *string `form:"smartpost_hub,omitempty" json:"smartpost_hub,omitempty" xml:"smartpost_hub,omitempty"`
	// The manifest ID is used to group SmartPost packages onto a manifest for each trailer.
	SmartpostManifest *bool `form:"smartpost_manifest,omitempty" json:"smartpost_manifest,omitempty" xml:"smartpost_manifest,omitempty"`
	// This option allows you to request restrictive rates from USPS. Can set to 'USPS.MEDIAMAIL' or 'USPS.LIBRARYMAIL'.
	SpecialRatesEligibility *string `form:"special_rates_eligibility,omitempty" json:"special_rates_eligibility,omitempty" xml:"special_rates_eligibility,omitempty"`
}

// Validate validates the optionsPayload type instance.
func (ut *optionsPayload) Validate() (err error) {
	if ut.CodMethod != nil {
		if !(*ut.CodMethod == "CASH" || *ut.CodMethod == "CHECK" || *ut.CodMethod == "MONEY_ORDER") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.cod_method`, *ut.CodMethod, []interface{}{"CASH", "CHECK", "MONEY_ORDER"}))
		}
	}
	if ut.HandlingInstructions != nil {
		if !(*ut.HandlingInstructions == "ORMD" || *ut.HandlingInstructions == "LIMITED_QUANTITY") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.handling_instructions`, *ut.HandlingInstructions, []interface{}{"ORMD", "LIMITED_QUANTITY"}))
		}
	}
	if ut.LabelFormat != nil {
		if !(*ut.LabelFormat == "PNG" || *ut.LabelFormat == "PDF" || *ut.LabelFormat == "ZPL" || *ut.LabelFormat == "EPL2") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.label_format`, *ut.LabelFormat, []interface{}{"PNG", "PDF", "ZPL", "EPL2"}))
		}
	}
	return
}

// Publicize creates OptionsPayload from optionsPayload
func (ut *optionsPayload) Publicize() *OptionsPayload {
	var pub OptionsPayload
	if ut.AdditionalHandling != nil {
		pub.AdditionalHandling = ut.AdditionalHandling
	}
	if ut.AddressValidationLevel != nil {
		pub.AddressValidationLevel = ut.AddressValidationLevel
	}
	if ut.Alcohol != nil {
		pub.Alcohol = ut.Alcohol
	}
	if ut.BillReceiverAccount != nil {
		pub.BillReceiverAccount = ut.BillReceiverAccount
	}
	if ut.BillReceiverPostalCode != nil {
		pub.BillReceiverPostalCode = ut.BillReceiverPostalCode
	}
	if ut.BillThirdPartyAccount != nil {
		pub.BillThirdPartyAccount = ut.BillThirdPartyAccount
	}
	if ut.BillThirdPartyCountry != nil {
		pub.BillThirdPartyCountry = ut.BillThirdPartyCountry
	}
	if ut.BillThirdPartyPostalCode != nil {
		pub.BillThirdPartyPostalCode = ut.BillThirdPartyPostalCode
	}
	if ut.ByDrone != nil {
		pub.ByDrone = ut.ByDrone
	}
	if ut.CarbonNeutral != nil {
		pub.CarbonNeutral = ut.CarbonNeutral
	}
	if ut.CodAmount != nil {
		pub.CodAmount = ut.CodAmount
	}
	if ut.CodMethod != nil {
		pub.CodMethod = ut.CodMethod
	}
	if ut.Currency != nil {
		pub.Currency = ut.Currency
	}
	if ut.DeliveredDutyPaid != nil {
		pub.DeliveredDutyPaid = ut.DeliveredDutyPaid
	}
	if ut.DeliveryConfirmation != nil {
		pub.DeliveryConfirmation = ut.DeliveryConfirmation
	}
	if ut.DryIce != nil {
		pub.DryIce = ut.DryIce
	}
	if ut.DryIceMedical != nil {
		pub.DryIceMedical = ut.DryIceMedical
	}
	if ut.DryIceWeight != nil {
		pub.DryIceWeight = ut.DryIceWeight
	}
	if ut.FreightCharge != nil {
		pub.FreightCharge = ut.FreightCharge
	}
	if ut.HandlingInstructions != nil {
		pub.HandlingInstructions = ut.HandlingInstructions
	}
	if ut.HoldForPickup != nil {
		pub.HoldForPickup = ut.HoldForPickup
	}
	if ut.InvoiceNumber != nil {
		pub.InvoiceNumber = ut.InvoiceNumber
	}
	if ut.LabelDate != nil {
		pub.LabelDate = ut.LabelDate
	}
	if ut.LabelFormat != nil {
		pub.LabelFormat = ut.LabelFormat
	}
	if ut.Machinable != nil {
		pub.Machinable = ut.Machinable
	}
	if ut.PrintCustom1 != nil {
		pub.PrintCustom1 = ut.PrintCustom1
	}
	if ut.PrintCustom1Barcode != nil {
		pub.PrintCustom1Barcode = ut.PrintCustom1Barcode
	}
	if ut.PrintCustom1Code != nil {
		pub.PrintCustom1Code = ut.PrintCustom1Code
	}
	if ut.PrintCustom2 != nil {
		pub.PrintCustom2 = ut.PrintCustom2
	}
	if ut.PrintCustom2Barcode != nil {
		pub.PrintCustom2Barcode = ut.PrintCustom2Barcode
	}
	if ut.PrintCustom2Code != nil {
		pub.PrintCustom2Code = ut.PrintCustom2Code
	}
	if ut.PrintCustom3 != nil {
		pub.PrintCustom3 = ut.PrintCustom3
	}
	if ut.PrintCustom3Barcode != nil {
		pub.PrintCustom3Barcode = ut.PrintCustom3Barcode
	}
	if ut.PrintCustom3Code != nil {
		pub.PrintCustom3Code = ut.PrintCustom3Code
	}
	if ut.SaturdayDelivery != nil {
		pub.SaturdayDelivery = ut.SaturdayDelivery
	}
	if ut.SmartpostHub != nil {
		pub.SmartpostHub = ut.SmartpostHub
	}
	if ut.SmartpostManifest != nil {
		pub.SmartpostManifest = ut.SmartpostManifest
	}
	if ut.SpecialRatesEligibility != nil {
		pub.SpecialRatesEligibility = ut.SpecialRatesEligibility
	}
	return &pub
}

// Shipments can have a variety of additional options which you can specify when creating a shipment. The Options object can be populated with the keys below.
type OptionsPayload struct {
	// Setting this option to true, will add an additional handling charge.
	AdditionalHandling *bool `form:"additional_handling,omitempty" json:"additional_handling,omitempty" xml:"additional_handling,omitempty"`
	// Setting this option to "0", will allow the minimum amount of address information to pass the validation check. Only for USPS postage.
	AddressValidationLevel *string `form:"address_validation_level,omitempty" json:"address_validation_level,omitempty" xml:"address_validation_level,omitempty"`
	// Set this option to true if your shipment contains alcohol. UPS - only supported for US Domestic shipments. FedEx - only supported for US Domestic shipments. Canada Post - Requires adult signature 19 years or older. If you want adult signature 18 years or older, instead use delivery_confirmation: ADULT_SIGNATURE
	Alcohol *bool `form:"alcohol,omitempty" json:"alcohol,omitempty" xml:"alcohol,omitempty"`
	// Setting an account number of the receiver who is to receive and buy the postage. UPS - bill_receiver_postal_code is also required
	BillReceiverAccount *string `form:"bill_receiver_account,omitempty" json:"bill_receiver_account,omitempty" xml:"bill_receiver_account,omitempty"`
	// Setting a postal code of the receiver account you want to buy postage. UPS - bill_receiver_account also required
	BillReceiverPostalCode *string `form:"bill_receiver_postal_code,omitempty" json:"bill_receiver_postal_code,omitempty" xml:"bill_receiver_postal_code,omitempty"`
	// Setting an account number of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_postal_code also required
	BillThirdPartyAccount *string `form:"bill_third_party_account,omitempty" json:"bill_third_party_account,omitempty" xml:"bill_third_party_account,omitempty"`
	// Setting a country of the third party account you want to buy postage. UPS - bill_third_party_account and bill_third_party_postal_code also required
	BillThirdPartyCountry *string `form:"bill_third_party_country,omitempty" json:"bill_third_party_country,omitempty" xml:"bill_third_party_country,omitempty"`
	// Setting a postal code of the third party account you want to buy postage. UPS - bill_third_party_country and bill_third_party_account also required
	BillThirdPartyPostalCode *bool `form:"bill_third_party_postal_code,omitempty" json:"bill_third_party_postal_code,omitempty" xml:"bill_third_party_postal_code,omitempty"`
	// Setting this option to true will indicate to the carrier to prefer delivery by drone, if the carrier supports drone delivery.
	ByDrone *bool `form:"by_drone,omitempty" json:"by_drone,omitempty" xml:"by_drone,omitempty"`
	// Setting this to true will add a charge to reduce carbon emissions.
	CarbonNeutral *bool `form:"carbon_neutral,omitempty" json:"carbon_neutral,omitempty" xml:"carbon_neutral,omitempty"`
	// Adding an amount will have the carrier collect the specified amount from the recipient.
	CodAmount *string `form:"cod_amount,omitempty" json:"cod_amount,omitempty" xml:"cod_amount,omitempty"`
	// Method for payment. "CASH", "CHECK", "MONEY_ORDER"
	CodMethod *string `form:"cod_method,omitempty" json:"cod_method,omitempty" xml:"cod_method,omitempty"`
	// Which currency this shipment will show for rates if carrier allows.
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Setting this value to false will pass the cost of duties on to the recipient of the package(s).
	DeliveredDutyPaid *bool `form:"delivered_duty_paid,omitempty" json:"delivered_duty_paid,omitempty" xml:"delivered_duty_paid,omitempty"`
	// If you want to request a signature, you can pass "ADULT_SIGNATURE" or "SIGNATURE". You may also request "NO_SIGNATURE" to leave the package at the door. All - some options may be limited for international shipments
	DeliveryConfirmation *string `form:"delivery_confirmation,omitempty" json:"delivery_confirmation,omitempty" xml:"delivery_confirmation,omitempty"`
	// Package contents contain dry ice. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set. FedEx - Need dry_ice_weight to be set
	DryIce *bool `form:"dry_ice,omitempty" json:"dry_ice,omitempty" xml:"dry_ice,omitempty"`
	// If the dry ice is for medical use, set this option to true. UPS - Need dry_ice_weight to be set. UPS MailInnovations - Need dry_ice_weight to be set
	DryIceMedical *bool `form:"dry_ice_medical,omitempty" json:"dry_ice_medical,omitempty" xml:"dry_ice_medical,omitempty"`
	// Weight of the dry ice in ounces. UPS - Need dry_ice to be set. UPS MailInnovations - Need dry_ice to be set. FedEx - Need dry_ice to be set
	DryIceWeight *string `form:"dry_ice_weight,omitempty" json:"dry_ice_weight,omitempty" xml:"dry_ice_weight,omitempty"`
	// Additional cost to be added to the invoice of this shipment. Only applies to UPS currently.
	FreightCharge *float64 `form:"freight_charge,omitempty" json:"freight_charge,omitempty" xml:"freight_charge,omitempty"`
	// This is to designate special instructions for the carrier like "Do not drop!".
	HandlingInstructions *string `form:"handling_instructions,omitempty" json:"handling_instructions,omitempty" xml:"handling_instructions,omitempty"`
	// Package will wait at carrier facility for pickup.
	HoldForPickup *bool `form:"hold_for_pickup,omitempty" json:"hold_for_pickup,omitempty" xml:"hold_for_pickup,omitempty"`
	// This will print an invoice number on the postage label.
	InvoiceNumber *string `form:"invoice_number,omitempty" json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// Set the date that will appear on the postage label. Accepts ISO 8601 formatted string including time zone offset.
	LabelDate *string `form:"label_date,omitempty" json:"label_date,omitempty" xml:"label_date,omitempty"`
	// Supported label formats include "PNG", "PDF", "ZPL", and "EPL2". "PNG" is the only format that allows for conversion.
	LabelFormat *string `form:"label_format,omitempty" json:"label_format,omitempty" xml:"label_format,omitempty"`
	// Whether or not the parcel can be processed by the carriers equipment.
	Machinable *bool `form:"machinable,omitempty" json:"machinable,omitempty" xml:"machinable,omitempty"`
	// You can optionally print custom messages on labels. The locations of these fields show up on different spots on the carrier's labels.
	PrintCustom1 *string `form:"print_custom_1,omitempty" json:"print_custom_1,omitempty" xml:"print_custom_1,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom1Barcode *bool `form:"print_custom_1_barcode,omitempty" json:"print_custom_1_barcode,omitempty" xml:"print_custom_1_barcode,omitempty"`
	// Specify the type of print_custom_1.
	PrintCustom1Code *string `form:"print_custom_1_code,omitempty" json:"print_custom_1_code,omitempty" xml:"print_custom_1_code,omitempty"`
	// An additional message on the label. Same restrictions as print_custom_1
	PrintCustom2 *string `form:"print_custom_2,omitempty" json:"print_custom_2,omitempty" xml:"print_custom_2,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom2Barcode *bool `form:"print_custom_2_barcode,omitempty" json:"print_custom_2_barcode,omitempty" xml:"print_custom_2_barcode,omitempty"`
	// See print_custom_1_code.
	PrintCustom2Code *string `form:"print_custom_2_code,omitempty" json:"print_custom_2_code,omitempty" xml:"print_custom_2_code,omitempty"`
	// An additional message on the label. Same restrictions as print_custom_1
	PrintCustom3 *string `form:"print_custom_3,omitempty" json:"print_custom_3,omitempty" xml:"print_custom_3,omitempty"`
	// Create a barcode for this custom reference if supported by carrier.
	PrintCustom3Barcode *bool `form:"print_custom_3_barcode,omitempty" json:"print_custom_3_barcode,omitempty" xml:"print_custom_3_barcode,omitempty"`
	// See print_custom_1_code.
	PrintCustom3Code *string `form:"print_custom_3_code,omitempty" json:"print_custom_3_code,omitempty" xml:"print_custom_3_code,omitempty"`
	// Set this value to true for delivery on Saturday. Can't be combined with Next Day Air. When setting the saturday_delivery option, you will only get rates for services that are eligible for saturday delivery. If no services are available for saturday delivery, then you will not be returned any rates. You may need to create 2 shipments, one with the saturday_delivery option set on one without to get all your eligible rates.
	SaturdayDelivery *bool `form:"saturday_delivery,omitempty" json:"saturday_delivery,omitempty" xml:"saturday_delivery,omitempty"`
	// You can use this to override the hub ID you have on your account.
	SmartpostHub *string `form:"smartpost_hub,omitempty" json:"smartpost_hub,omitempty" xml:"smartpost_hub,omitempty"`
	// The manifest ID is used to group SmartPost packages onto a manifest for each trailer.
	SmartpostManifest *bool `form:"smartpost_manifest,omitempty" json:"smartpost_manifest,omitempty" xml:"smartpost_manifest,omitempty"`
	// This option allows you to request restrictive rates from USPS. Can set to 'USPS.MEDIAMAIL' or 'USPS.LIBRARYMAIL'.
	SpecialRatesEligibility *string `form:"special_rates_eligibility,omitempty" json:"special_rates_eligibility,omitempty" xml:"special_rates_eligibility,omitempty"`
}

// Validate validates the OptionsPayload type instance.
func (ut *OptionsPayload) Validate() (err error) {
	if ut.CodMethod != nil {
		if !(*ut.CodMethod == "CASH" || *ut.CodMethod == "CHECK" || *ut.CodMethod == "MONEY_ORDER") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.cod_method`, *ut.CodMethod, []interface{}{"CASH", "CHECK", "MONEY_ORDER"}))
		}
	}
	if ut.HandlingInstructions != nil {
		if !(*ut.HandlingInstructions == "ORMD" || *ut.HandlingInstructions == "LIMITED_QUANTITY") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.handling_instructions`, *ut.HandlingInstructions, []interface{}{"ORMD", "LIMITED_QUANTITY"}))
		}
	}
	if ut.LabelFormat != nil {
		if !(*ut.LabelFormat == "PNG" || *ut.LabelFormat == "PDF" || *ut.LabelFormat == "ZPL" || *ut.LabelFormat == "EPL2") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.label_format`, *ut.LabelFormat, []interface{}{"PNG", "PDF", "ZPL", "EPL2"}))
		}
	}
	return
}

// Parcel objects represent the physical container being shipped. Dimensions can be supplied either as length, width, and height dimensions, or a predefined_package string. If you provide a carrier specific predefined_package the associated Shipment will only fetch rates from that carrier.
type parcelPayload struct {
	// Required if predefined_package is empty. float (inches)
	Height *float64 `form:"height,omitempty" json:"height,omitempty" xml:"height,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Length *float64 `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Optional, one of our predefined_packages
	PredefinedPackage *string `form:"predefined_package,omitempty" json:"predefined_package,omitempty" xml:"predefined_package,omitempty"`
	// Always required. float(oz)
	Weight *float64 `form:"weight,omitempty" json:"weight,omitempty" xml:"weight,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Width *float64 `form:"width,omitempty" json:"width,omitempty" xml:"width,omitempty"`
}

// Validate validates the parcelPayload type instance.
func (ut *parcelPayload) Validate() (err error) {
	if ut.Weight == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "weight"))
	}

	return
}

// Publicize creates ParcelPayload from parcelPayload
func (ut *parcelPayload) Publicize() *ParcelPayload {
	var pub ParcelPayload
	if ut.Height != nil {
		pub.Height = ut.Height
	}
	if ut.Length != nil {
		pub.Length = ut.Length
	}
	if ut.PredefinedPackage != nil {
		pub.PredefinedPackage = ut.PredefinedPackage
	}
	if ut.Weight != nil {
		pub.Weight = *ut.Weight
	}
	if ut.Width != nil {
		pub.Width = ut.Width
	}
	return &pub
}

// Parcel objects represent the physical container being shipped. Dimensions can be supplied either as length, width, and height dimensions, or a predefined_package string. If you provide a carrier specific predefined_package the associated Shipment will only fetch rates from that carrier.
type ParcelPayload struct {
	// Required if predefined_package is empty. float (inches)
	Height *float64 `form:"height,omitempty" json:"height,omitempty" xml:"height,omitempty"`
	// Required if predefined_package is empty. float (inches)
	Length *float64 `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Optional, one of our predefined_packages
	PredefinedPackage *string `form:"predefined_package,omitempty" json:"predefined_package,omitempty" xml:"predefined_package,omitempty"`
	// Always required. float(oz)
	Weight float64 `form:"weight" json:"weight" xml:"weight"`
	// Required if predefined_package is empty. float (inches)
	Width *float64 `form:"width,omitempty" json:"width,omitempty" xml:"width,omitempty"`
}

// Buy Shipment Payload
type ratePayload struct {
	// Unique, begins with "rate_"
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the ratePayload type instance.
func (ut *ratePayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	if ut.ID != nil {
		if ok := goa.ValidatePattern(`^rate_`, *ut.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *ut.ID, `^rate_`))
		}
	}
	return
}

// Publicize creates RatePayload from ratePayload
func (ut *ratePayload) Publicize() *RatePayload {
	var pub RatePayload
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	return &pub
}

// Buy Shipment Payload
type RatePayload struct {
	// Unique, begins with "rate_"
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the RatePayload type instance.
func (ut *RatePayload) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	if ok := goa.ValidatePattern(`^rate_`, ut.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, ut.ID, `^rate_`))
	}
	return
}

// Insuring your Shipment is as simple as sending us the value of the contents. We charge 1% of the value, with a $1 minimum, and handle all the claims. All our claims are paid out within 30 days.
type shipmentInsurancePayload struct {
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// Validate validates the shipmentInsurancePayload type instance.
func (ut *shipmentInsurancePayload) Validate() (err error) {
	if ut.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}

	return
}

// Publicize creates ShipmentInsurancePayload from shipmentInsurancePayload
func (ut *shipmentInsurancePayload) Publicize() *ShipmentInsurancePayload {
	var pub ShipmentInsurancePayload
	if ut.Amount != nil {
		pub.Amount = *ut.Amount
	}
	return &pub
}

// Insuring your Shipment is as simple as sending us the value of the contents. We charge 1% of the value, with a $1 minimum, and handle all the claims. All our claims are paid out within 30 days.
type ShipmentInsurancePayload struct {
	Amount string `form:"amount" json:"amount" xml:"amount"`
}

// Validate validates the ShipmentInsurancePayload type instance.
func (ut *ShipmentInsurancePayload) Validate() (err error) {
	if ut.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}

	return
}

// Parcel objects represent the physical container being shipped. Dimensions can be supplied either as length, width, and height dimensions, or a predefined_package string. If you provide a carrier specific predefined_package the associated Shipment will only fetch rates from that carrier.
type shipmentPayload struct {
	// The ID of the batch that contains this shipment, if any
	BatchID *string `form:"batch_id,omitempty" json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The buyer's address, defaults to to_address
	BuyerAddress *addressPayload `form:"buyer_address,omitempty" json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// Information for the processing of customs
	CustomsInfo *customsInfoPayload `form:"customs_info,omitempty" json:"customs_info,omitempty" xml:"customs_info,omitempty"`
	// All associated Form objects
	Forms []interface{} `form:"forms,omitempty" json:"forms,omitempty" xml:"forms,omitempty"`
	// The origin address
	FromAddress *addressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// The associated Insurance object
	Insurance *insurancePayload `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Set true to create as a return, discussed in more depth below
	IsReturn *bool `form:"is_return,omitempty" json:"is_return,omitempty" xml:"is_return,omitempty"`
	// All of the options passed to the shipment, discussed in more depth below
	Options *optionsPayload `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// The dimensions and weight of the package
	Parcel *parcelPayload `form:"parcel,omitempty" json:"parcel,omitempty" xml:"parcel,omitempty"`
	// The shipper's address, defaults to from_address
	ReturnAddress *addressPayload `form:"return_address,omitempty" json:"return_address,omitempty" xml:"return_address,omitempty"`
	// The destination address
	ToAddress *addressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The USPS zone of the shipment, if purchased with USPS
	UspsZone *string `form:"usps_zone,omitempty" json:"usps_zone,omitempty" xml:"usps_zone,omitempty"`
}

// Finalize sets the default values for shipmentPayload type instance.
func (ut *shipmentPayload) Finalize() {
	if ut.CustomsInfo != nil {
		for _, e := range ut.CustomsInfo.CustomsItems {
			var defaultCurrency = "USD"
			if e.Currency == nil {
				e.Currency = &defaultCurrency
			}
			var defaultOriginCountry = "US"
			if e.OriginCountry == nil {
				e.OriginCountry = &defaultOriginCountry
			}
		}
	}
}

// Validate validates the shipmentPayload type instance.
func (ut *shipmentPayload) Validate() (err error) {
	if ut.BuyerAddress != nil {
		if err2 := ut.BuyerAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.CustomsInfo != nil {
		if err2 := ut.CustomsInfo.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.FromAddress != nil {
		if err2 := ut.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Insurance != nil {
		if err2 := ut.Insurance.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Options != nil {
		if err2 := ut.Options.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Parcel != nil {
		if err2 := ut.Parcel.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.ReturnAddress != nil {
		if err2 := ut.ReturnAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.ToAddress != nil {
		if err2 := ut.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates ShipmentPayload from shipmentPayload
func (ut *shipmentPayload) Publicize() *ShipmentPayload {
	var pub ShipmentPayload
	if ut.BatchID != nil {
		pub.BatchID = ut.BatchID
	}
	if ut.BuyerAddress != nil {
		pub.BuyerAddress = ut.BuyerAddress.Publicize()
	}
	if ut.CustomsInfo != nil {
		pub.CustomsInfo = ut.CustomsInfo.Publicize()
	}
	if ut.Forms != nil {
		pub.Forms = ut.Forms
	}
	if ut.FromAddress != nil {
		pub.FromAddress = ut.FromAddress.Publicize()
	}
	if ut.Insurance != nil {
		pub.Insurance = ut.Insurance.Publicize()
	}
	if ut.IsReturn != nil {
		pub.IsReturn = ut.IsReturn
	}
	if ut.Options != nil {
		pub.Options = ut.Options.Publicize()
	}
	if ut.Parcel != nil {
		pub.Parcel = ut.Parcel.Publicize()
	}
	if ut.ReturnAddress != nil {
		pub.ReturnAddress = ut.ReturnAddress.Publicize()
	}
	if ut.ToAddress != nil {
		pub.ToAddress = ut.ToAddress.Publicize()
	}
	if ut.UspsZone != nil {
		pub.UspsZone = ut.UspsZone
	}
	return &pub
}

// Parcel objects represent the physical container being shipped. Dimensions can be supplied either as length, width, and height dimensions, or a predefined_package string. If you provide a carrier specific predefined_package the associated Shipment will only fetch rates from that carrier.
type ShipmentPayload struct {
	// The ID of the batch that contains this shipment, if any
	BatchID *string `form:"batch_id,omitempty" json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The buyer's address, defaults to to_address
	BuyerAddress *AddressPayload `form:"buyer_address,omitempty" json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// Information for the processing of customs
	CustomsInfo *CustomsInfoPayload `form:"customs_info,omitempty" json:"customs_info,omitempty" xml:"customs_info,omitempty"`
	// All associated Form objects
	Forms []interface{} `form:"forms,omitempty" json:"forms,omitempty" xml:"forms,omitempty"`
	// The origin address
	FromAddress *AddressPayload `form:"from_address,omitempty" json:"from_address,omitempty" xml:"from_address,omitempty"`
	// The associated Insurance object
	Insurance *InsurancePayload `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Set true to create as a return, discussed in more depth below
	IsReturn *bool `form:"is_return,omitempty" json:"is_return,omitempty" xml:"is_return,omitempty"`
	// All of the options passed to the shipment, discussed in more depth below
	Options *OptionsPayload `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// The dimensions and weight of the package
	Parcel *ParcelPayload `form:"parcel,omitempty" json:"parcel,omitempty" xml:"parcel,omitempty"`
	// The shipper's address, defaults to from_address
	ReturnAddress *AddressPayload `form:"return_address,omitempty" json:"return_address,omitempty" xml:"return_address,omitempty"`
	// The destination address
	ToAddress *AddressPayload `form:"to_address,omitempty" json:"to_address,omitempty" xml:"to_address,omitempty"`
	// The USPS zone of the shipment, if purchased with USPS
	UspsZone *string `form:"usps_zone,omitempty" json:"usps_zone,omitempty" xml:"usps_zone,omitempty"`
}

// Validate validates the ShipmentPayload type instance.
func (ut *ShipmentPayload) Validate() (err error) {
	if ut.BuyerAddress != nil {
		if err2 := ut.BuyerAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.CustomsInfo != nil {
		if err2 := ut.CustomsInfo.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.FromAddress != nil {
		if err2 := ut.FromAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Insurance != nil {
		if err2 := ut.Insurance.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Options != nil {
		if err2 := ut.Options.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.ReturnAddress != nil {
		if err2 := ut.ReturnAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.ToAddress != nil {
		if err2 := ut.ToAddress.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// trackerListPayload user type.
type trackerListPayload struct {
	// Optional pagination parameter. Only trackers created after the given ID will be included. May not be used with before_id
	AfterID *string `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only trackers created before the given ID will be included. May not be used with after_id
	BeforeID *string `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only returns Trackers with the given carrier value
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// Only return Trackers created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of Trackers to return on each page. The maximum value is 100
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Only return Trackers created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
	// Only returns Trackers with the given tracking_code. Useful for retrieving an individual Tracker by tracking_code rather than by ID
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}

// Finalize sets the default values for trackerListPayload type instance.
func (ut *trackerListPayload) Finalize() {
	var defaultPageSize = 1
	if ut.PageSize == nil {
		ut.PageSize = &defaultPageSize
	}
}

// Validate validates the trackerListPayload type instance.
func (ut *trackerListPayload) Validate() (err error) {
	if ut.TrackingCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tracking_code"))
	}

	if ut.PageSize != nil {
		if *ut.PageSize < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, *ut.PageSize, 1, true))
		}
	}
	if ut.PageSize != nil {
		if *ut.PageSize > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, *ut.PageSize, 100, false))
		}
	}
	return
}

// Publicize creates TrackerListPayload from trackerListPayload
func (ut *trackerListPayload) Publicize() *TrackerListPayload {
	var pub TrackerListPayload
	if ut.AfterID != nil {
		pub.AfterID = ut.AfterID
	}
	if ut.BeforeID != nil {
		pub.BeforeID = ut.BeforeID
	}
	if ut.Carrier != nil {
		pub.Carrier = ut.Carrier
	}
	if ut.EndDatetime != nil {
		pub.EndDatetime = ut.EndDatetime
	}
	if ut.PageSize != nil {
		pub.PageSize = *ut.PageSize
	}
	if ut.StartDatetime != nil {
		pub.StartDatetime = ut.StartDatetime
	}
	if ut.TrackingCode != nil {
		pub.TrackingCode = *ut.TrackingCode
	}
	return &pub
}

// TrackerListPayload user type.
type TrackerListPayload struct {
	// Optional pagination parameter. Only trackers created after the given ID will be included. May not be used with before_id
	AfterID *string `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only trackers created before the given ID will be included. May not be used with after_id
	BeforeID *string `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only returns Trackers with the given carrier value
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// Only return Trackers created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of Trackers to return on each page. The maximum value is 100
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
	// Only return Trackers created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
	// Only returns Trackers with the given tracking_code. Useful for retrieving an individual Tracker by tracking_code rather than by ID
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// Validate validates the TrackerListPayload type instance.
func (ut *TrackerListPayload) Validate() (err error) {
	if ut.TrackingCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tracking_code"))
	}

	if ut.PageSize < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, ut.PageSize, 1, true))
	}
	if ut.PageSize > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.page_size`, ut.PageSize, 100, false))
	}
	return
}

// trackerPayload user type.
type trackerPayload struct {
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The tracking code associated with the package you'd like to track
	TrackingCode *string `form:"tracking_code,omitempty" json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}

// Validate validates the trackerPayload type instance.
func (ut *trackerPayload) Validate() (err error) {
	if ut.TrackingCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tracking_code"))
	}

	return
}

// Publicize creates TrackerPayload from trackerPayload
func (ut *trackerPayload) Publicize() *TrackerPayload {
	var pub TrackerPayload
	if ut.Carrier != nil {
		pub.Carrier = ut.Carrier
	}
	if ut.TrackingCode != nil {
		pub.TrackingCode = *ut.TrackingCode
	}
	return &pub
}

// TrackerPayload user type.
type TrackerPayload struct {
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The tracking code associated with the package you'd like to track
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// Validate validates the TrackerPayload type instance.
func (ut *TrackerPayload) Validate() (err error) {
	if ut.TrackingCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tracking_code"))
	}

	return
}

// userCreatePayload user type.
type userCreatePayload struct {
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First and last name required
	Name                 *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password             *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	PasswordConfirmation *string `form:"password_confirmation,omitempty" json:"password_confirmation,omitempty" xml:"password_confirmation,omitempty"`
	PhoneNumber          *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate validates the userCreatePayload type instance.
func (ut *userCreatePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.PasswordConfirmation == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password_confirmation"))
	}

	return
}

// Publicize creates UserCreatePayload from userCreatePayload
func (ut *userCreatePayload) Publicize() *UserCreatePayload {
	var pub UserCreatePayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.PasswordConfirmation != nil {
		pub.PasswordConfirmation = *ut.PasswordConfirmation
	}
	if ut.PhoneNumber != nil {
		pub.PhoneNumber = ut.PhoneNumber
	}
	return &pub
}

// UserCreatePayload user type.
type UserCreatePayload struct {
	Email string `form:"email" json:"email" xml:"email"`
	// First and last name required
	Name                 string  `form:"name" json:"name" xml:"name"`
	Password             string  `form:"password" json:"password" xml:"password"`
	PasswordConfirmation string  `form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation"`
	PhoneNumber          *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate validates the UserCreatePayload type instance.
func (ut *UserCreatePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.PasswordConfirmation == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password_confirmation"))
	}

	return
}

// userUpdatePayload user type.
type userUpdatePayload struct {
	CurrentPassword *string `form:"current_password,omitempty" json:"current_password,omitempty" xml:"current_password,omitempty"`
	Email           *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First and last name required
	Name                    *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password                *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	PasswordConfirmation    *string `form:"password_confirmation,omitempty" json:"password_confirmation,omitempty" xml:"password_confirmation,omitempty"`
	PhoneNumber             *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	RechargeAmount          *string `form:"recharge_amount,omitempty" json:"recharge_amount,omitempty" xml:"recharge_amount,omitempty"`
	RechargeThreshold       *string `form:"recharge_threshold,omitempty" json:"recharge_threshold,omitempty" xml:"recharge_threshold,omitempty"`
	SecondaryRechargeAmount *string `form:"secondary_recharge_amount,omitempty" json:"secondary_recharge_amount,omitempty" xml:"secondary_recharge_amount,omitempty"`
}

// Publicize creates UserUpdatePayload from userUpdatePayload
func (ut *userUpdatePayload) Publicize() *UserUpdatePayload {
	var pub UserUpdatePayload
	if ut.CurrentPassword != nil {
		pub.CurrentPassword = ut.CurrentPassword
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.PasswordConfirmation != nil {
		pub.PasswordConfirmation = ut.PasswordConfirmation
	}
	if ut.PhoneNumber != nil {
		pub.PhoneNumber = ut.PhoneNumber
	}
	if ut.RechargeAmount != nil {
		pub.RechargeAmount = ut.RechargeAmount
	}
	if ut.RechargeThreshold != nil {
		pub.RechargeThreshold = ut.RechargeThreshold
	}
	if ut.SecondaryRechargeAmount != nil {
		pub.SecondaryRechargeAmount = ut.SecondaryRechargeAmount
	}
	return &pub
}

// UserUpdatePayload user type.
type UserUpdatePayload struct {
	CurrentPassword *string `form:"current_password,omitempty" json:"current_password,omitempty" xml:"current_password,omitempty"`
	Email           *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First and last name required
	Name                    *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password                *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	PasswordConfirmation    *string `form:"password_confirmation,omitempty" json:"password_confirmation,omitempty" xml:"password_confirmation,omitempty"`
	PhoneNumber             *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	RechargeAmount          *string `form:"recharge_amount,omitempty" json:"recharge_amount,omitempty" xml:"recharge_amount,omitempty"`
	RechargeThreshold       *string `form:"recharge_threshold,omitempty" json:"recharge_threshold,omitempty" xml:"recharge_threshold,omitempty"`
	SecondaryRechargeAmount *string `form:"secondary_recharge_amount,omitempty" json:"secondary_recharge_amount,omitempty" xml:"secondary_recharge_amount,omitempty"`
}

// applicationEasypostGieldErrorJSON user type.
type applicationEasypostGieldErrorJSON struct {
	// Field of the request that the error describes
	Field *string `form:"field,omitempty" json:"field,omitempty" xml:"field,omitempty"`
	// Human readable description of the problem
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// Publicize creates ApplicationEasypostGieldErrorJSON from applicationEasypostGieldErrorJSON
func (ut *applicationEasypostGieldErrorJSON) Publicize() *ApplicationEasypostGieldErrorJSON {
	var pub ApplicationEasypostGieldErrorJSON
	if ut.Field != nil {
		pub.Field = ut.Field
	}
	if ut.Message != nil {
		pub.Message = ut.Message
	}
	return &pub
}

// ApplicationEasypostGieldErrorJSON user type.
type ApplicationEasypostGieldErrorJSON struct {
	// Field of the request that the error describes
	Field *string `form:"field,omitempty" json:"field,omitempty" xml:"field,omitempty"`
	// Human readable description of the problem
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}
