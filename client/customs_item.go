package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateCustomsItemPayload is the customs_item create action payload.
type CreateCustomsItemPayload struct {
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

// CreateCustomsItemPath computes a request path to the create action of customs_item.
func CreateCustomsItemPath() string {
	return fmt.Sprintf("/v2/customs_items")
}

// CreateCustomsItem makes a request to the create action endpoint of the customs_item resource
func (c *Client) CreateCustomsItem(ctx context.Context, path string, payload *CreateCustomsItemPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCustomsItemRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCustomsItemRequest create the request corresponding to the create action endpoint of the customs_item resource.
func (c *Client) NewCreateCustomsItemRequest(ctx context.Context, path string, payload *CreateCustomsItemPayload, contentType string) (*http.Request, error) {
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
