package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateCustomsInfoPayload is the customs_info create action payload.
type CreateCustomsInfoPayload struct {
	ContentsType    string               `form:"contents_type" json:"contents_type" xml:"contents_type"`
	CustomsCertify  bool                 `form:"customs_certify" json:"customs_certify" xml:"customs_certify"`
	CustomsItems    []*CustomItemPayload `form:"customs_items" json:"customs_items" xml:"customs_items"`
	CutomsSigner    string               `form:"cutoms_signer" json:"cutoms_signer" xml:"cutoms_signer"`
	EelPfc          string               `form:"eel_pfc" json:"eel_pfc" xml:"eel_pfc"`
	RestrictionType string               `form:"restriction_type" json:"restriction_type" xml:"restriction_type"`
}

// CreateCustomsInfoPath computes a request path to the create action of customs_info.
func CreateCustomsInfoPath() string {
	return fmt.Sprintf("/v2/customs_infos")
}

// CreateCustomsInfo makes a request to the create action endpoint of the customs_info resource
func (c *Client) CreateCustomsInfo(ctx context.Context, path string, payload *CreateCustomsInfoPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCustomsInfoRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCustomsInfoRequest create the request corresponding to the create action endpoint of the customs_info resource.
func (c *Client) NewCreateCustomsInfoRequest(ctx context.Context, path string, payload *CreateCustomsInfoPayload, contentType string) (*http.Request, error) {
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
