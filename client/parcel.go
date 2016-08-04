package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateParcelPayload is the parcel create action payload.
type CreateParcelPayload struct {
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

// CreateParcelPath computes a request path to the create action of parcel.
func CreateParcelPath() string {
	return fmt.Sprintf("/v2/parcels")
}

// Include the weight, and either a predefined_package or length, width and height.
func (c *Client) CreateParcel(ctx context.Context, path string, payload *CreateParcelPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateParcelRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateParcelRequest create the request corresponding to the create action endpoint of the parcel resource.
func (c *Client) NewCreateParcelRequest(ctx context.Context, path string, payload *CreateParcelPayload, contentType string) (*http.Request, error) {
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

// ShowParcelPath computes a request path to the show action of parcel.
func ShowParcelPath(id string) string {
	return fmt.Sprintf("/v2/parcels/%v", id)
}

// Get a Parcel by its id. In general you should not need to use this in your automated solution. A Parcel's id can be inlined into the creation call to other objects. This allows you to only create one Parcel for each package you will be using.
func (c *Client) ShowParcel(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowParcelRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowParcelRequest create the request corresponding to the show action endpoint of the parcel resource.
func (c *Client) NewShowParcelRequest(ctx context.Context, path string) (*http.Request, error) {
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
