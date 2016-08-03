package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateAddressPayload is the address create action payload.
type CreateAddressPayload struct {
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

// CreateAddressPath computes a request path to the create action of address.
func CreateAddressPath() string {
	return fmt.Sprintf("/v2/addresses")
}

// Depending on your use case an Address can be used in many different ways. Certain carriers allow rating between two zip codes, but full addresses are required to purchase postage. It is recommended to provide as much information as possible during creation and to reuse these objects whenever possible.
func (c *Client) CreateAddress(ctx context.Context, path string, payload *CreateAddressPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateAddressRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAddressRequest create the request corresponding to the create action endpoint of the address resource.
func (c *Client) NewCreateAddressRequest(ctx context.Context, path string, payload *CreateAddressPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowAddressPath computes a request path to the show action of address.
func ShowAddressPath(id string) string {
	return fmt.Sprintf("/v2/addresses/%v", id)
}

// An Address can be retrieved by either its id or reference. However it is recommended to use EasyPost's provided identifiers because uniqueness on reference is not enforced.
func (c *Client) ShowAddress(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAddressRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAddressRequest create the request corresponding to the show action endpoint of the address resource.
func (c *Client) NewShowAddressRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
