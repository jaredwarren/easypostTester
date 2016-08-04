package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateInsurancePayload is the insurance create action payload.
type CreateInsurancePayload struct {
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

// CreateInsurancePath computes a request path to the create action of insurance.
func CreateInsurancePath() string {
	return fmt.Sprintf("/v2/insurances")
}

// An Insurance created via this endpoint must belong to a shipment purchased outside of EasyPost. Insurance for Shipments created within EasyPost must be created via the Shipment Buy or Insure endpoints. When creating Insurance for a non-EasyPost shipment, you must provide to_address, from_address, tracking_code, and amount information. Optionally, you can provide the carrier parameter, which will help EasyPost identify the carrier the package was shipped with. If no carrier is provided, EasyPost will attempt to determine the carrier based on the tracking_code provided. Providing a carrier parameter is recommended, since some tracking_codes are ambiguous and may match with more than one carrier. In addition, not having to auto-match the carrier will significantly speed up the response time.
func (c *Client) CreateInsurance(ctx context.Context, path string, payload *CreateInsurancePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateInsuranceRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateInsuranceRequest create the request corresponding to the create action endpoint of the insurance resource.
func (c *Client) NewCreateInsuranceRequest(ctx context.Context, path string, payload *CreateInsurancePayload, contentType string) (*http.Request, error) {
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

// ListInsurancePayload is the insurance list action payload.
type ListInsurancePayload struct {
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

// ListInsurancePath computes a request path to the list action of insurance.
func ListInsurancePath() string {
	return fmt.Sprintf("/v2/insurances")
}

// The Insurance List is a paginated list of all Insurance objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.
func (c *Client) ListInsurance(ctx context.Context, path string, payload *ListInsurancePayload, contentType string) (*http.Response, error) {
	req, err := c.NewListInsuranceRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListInsuranceRequest create the request corresponding to the list action endpoint of the insurance resource.
func (c *Client) NewListInsuranceRequest(ctx context.Context, path string, payload *ListInsurancePayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("GET", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowInsurancePath computes a request path to the show action of insurance.
func ShowInsurancePath(id string) string {
	return fmt.Sprintf("/v2/insurances/%v", id)
}

// Retrieve an Insurance by id.
func (c *Client) ShowInsurance(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowInsuranceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowInsuranceRequest create the request corresponding to the show action endpoint of the insurance resource.
func (c *Client) NewShowInsuranceRequest(ctx context.Context, path string) (*http.Request, error) {
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
