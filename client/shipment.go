package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ConvertLabelShipmentPayload is the shipment convertLabel action payload.
type ConvertLabelShipmentPayload struct {
	// Selected rate
	FileFormat string `form:"file_format" json:"file_format" xml:"file_format"`
}

// ConvertLabelShipmentPath computes a request path to the convertLabel action of shipment.
func ConvertLabelShipmentPath(id string) string {
	return fmt.Sprintf("/v2/shipments/%v/label", id)
}

// A Shipment's PostageLabel can be converted from PNG to other formats. If the PostageLabel was originally generated in a format other than PNG it cannot be converted.
func (c *Client) ConvertLabelShipment(ctx context.Context, path string, payload *ConvertLabelShipmentPayload, contentType string) (*http.Response, error) {
	req, err := c.NewConvertLabelShipmentRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewConvertLabelShipmentRequest create the request corresponding to the convertLabel action endpoint of the shipment resource.
func (c *Client) NewConvertLabelShipmentRequest(ctx context.Context, path string, payload *ConvertLabelShipmentPayload, contentType string) (*http.Request, error) {
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

// CreateShipmentPayload is the shipment create action payload.
type CreateShipmentPayload struct {
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

// CreateShipmentPath computes a request path to the create action of shipment.
func CreateShipmentPath() string {
	return fmt.Sprintf("/v2/shipments")
}

// A Shipment is almost exclusively a container for other objects, and thus a Shipment may reuse many of these objects. Additionally, all the objects contained within a Shipment may be created at the same time.
func (c *Client) CreateShipment(ctx context.Context, path string, payload *CreateShipmentPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateShipmentRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateShipmentRequest create the request corresponding to the create action endpoint of the shipment resource.
func (c *Client) NewCreateShipmentRequest(ctx context.Context, path string, payload *CreateShipmentPayload, contentType string) (*http.Request, error) {
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

// InsureShipmentPayload is the shipment insure action payload.
type InsureShipmentPayload struct {
	Amount string `form:"amount" json:"amount" xml:"amount"`
}

// InsureShipmentPath computes a request path to the insure action of shipment.
func InsureShipmentPath(id string) string {
	return fmt.Sprintf("/v2/shipments/%v/insure", id)
}

// Insuring your Shipment is as simple as sending us the value of the contents. We charge 1% of the value, with a $1 minimum, and handle all the claims. All our claims are paid out within 30 days.
func (c *Client) InsureShipment(ctx context.Context, path string, payload *InsureShipmentPayload, contentType string) (*http.Response, error) {
	req, err := c.NewInsureShipmentRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewInsureShipmentRequest create the request corresponding to the insure action endpoint of the shipment resource.
func (c *Client) NewInsureShipmentRequest(ctx context.Context, path string, payload *InsureShipmentPayload, contentType string) (*http.Request, error) {
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

// PruchaseShipmentPayload is the shipment pruchase action payload.
type PruchaseShipmentPayload struct {
	// Additionally, insurance may be added during the purchase. To specify an amount to insure, pass the insurance attribute as a string. The currency of all insurance is USD.
	Insurance *string `form:"insurance,omitempty" json:"insurance,omitempty" xml:"insurance,omitempty"`
	// Selected rate
	Rate *RatePayload `form:"rate" json:"rate" xml:"rate"`
}

// PruchaseShipmentPath computes a request path to the pruchase action of shipment.
func PruchaseShipmentPath(id string) string {
	return fmt.Sprintf("/v2/shipments/%v/buy", id)
}

// To purchase a Shipment you only need to specify the Rate to purchase. This operation populates the tracking_code and postage_label attributes. The default image format of the associated PostageLabel is PNG. To change this default see the label_format option.
func (c *Client) PruchaseShipment(ctx context.Context, path string, payload *PruchaseShipmentPayload, contentType string) (*http.Response, error) {
	req, err := c.NewPruchaseShipmentRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPruchaseShipmentRequest create the request corresponding to the pruchase action endpoint of the shipment resource.
func (c *Client) NewPruchaseShipmentRequest(ctx context.Context, path string, payload *PruchaseShipmentPayload, contentType string) (*http.Request, error) {
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

// RatesShipmentPath computes a request path to the rates action of shipment.
func RatesShipmentPath(id string) string {
	return fmt.Sprintf("/v2/shipments/%v/rates", id)
}

// You can update the Rates of a Shipment at any time. This operation respects the carrier_accounts attribute.
func (c *Client) RatesShipment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRatesShipmentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRatesShipmentRequest create the request corresponding to the rates action endpoint of the shipment resource.
func (c *Client) NewRatesShipmentRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RefundShipmentPath computes a request path to the refund action of shipment.
func RefundShipmentPath(id string) string {
	return fmt.Sprintf("/v2/shipments/%v/refund", id)
}

// Refunding a Shipment is avaliable for many carriers supported by EasyPost. Once the refund has been submitted, refund_status attribute of the Shipment will be populated with one of the possible values: "submitted", "refunded", "rejected". The most common initial status is "submitted". Many carriers require that the refund be processed before the refund_status will move to "refunded". The length of this process depends on the carrier, but no greater than 30 days.
func (c *Client) RefundShipment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRefundShipmentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRefundShipmentRequest create the request corresponding to the refund action endpoint of the shipment resource.
func (c *Client) NewRefundShipmentRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowShipmentPath computes a request path to the show action of shipment.
func ShowShipmentPath(id string) string {
	return fmt.Sprintf("/v2/shipments/%v", id)
}

// Get a Parcel by its id. In general you should not need to use this in your automated solution. A Parcel's id can be inlined into the creation call to other objects. This allows you to only create one Parcel for each package you will be using.
func (c *Client) ShowShipment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowShipmentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowShipmentRequest create the request corresponding to the show action endpoint of the shipment resource.
func (c *Client) NewShowShipmentRequest(ctx context.Context, path string) (*http.Request, error) {
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
