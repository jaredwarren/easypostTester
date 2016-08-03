package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateCarrierAccountPayload is the carrier_account create action payload.
type CreateCarrierAccountPayload struct {
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

// CreateCarrierAccountPath computes a request path to the create action of carrier_account.
func CreateCarrierAccountPath() string {
	return fmt.Sprintf("/v2/carrier_accounts")
}

// CarrierAccount objects may be managed through the EasyPost API using the Production mode API key only. Multiple accounts can be added for a single carrier (with the exception of the USPS).
func (c *Client) CreateCarrierAccount(ctx context.Context, path string, payload *CreateCarrierAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCarrierAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCarrierAccountRequest create the request corresponding to the create action endpoint of the carrier_account resource.
func (c *Client) NewCreateCarrierAccountRequest(ctx context.Context, path string, payload *CreateCarrierAccountPayload, contentType string) (*http.Request, error) {
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

// DeleteCarrierAccountPath computes a request path to the delete action of carrier_account.
func DeleteCarrierAccountPath(id string) string {
	return fmt.Sprintf("/v2/carrier_accounts/%v", id)
}

// CarrierAccount objects may be removed from your account when they become out of date or no longer useful.
func (c *Client) DeleteCarrierAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCarrierAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCarrierAccountRequest create the request corresponding to the delete action endpoint of the carrier_account resource.
func (c *Client) NewDeleteCarrierAccountRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListCarrierAccountPath computes a request path to the list action of carrier_account.
func ListCarrierAccountPath() string {
	return fmt.Sprintf("/v2/carrier_accounts")
}

// Retrieve an unpaginated list of all CarrierAccounts available to the authenticated account. Only Production API keys may be used to retrieve this list, as there is no test mode equivalent.
func (c *Client) ListCarrierAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCarrierAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCarrierAccountRequest create the request corresponding to the list action endpoint of the carrier_account resource.
func (c *Client) NewListCarrierAccountRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowCarrierAccountPath computes a request path to the show action of carrier_account.
func ShowCarrierAccountPath(id string) string {
	return fmt.Sprintf("/v2/carrier_accounts/%v", id)
}

// Retrieve an unpaginated list of all CarrierAccounts available to the authenticated account. Only Production API keys may be used to retrieve this list, as there is no test mode equivalent.
func (c *Client) ShowCarrierAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowCarrierAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCarrierAccountRequest create the request corresponding to the show action endpoint of the carrier_account resource.
func (c *Client) NewShowCarrierAccountRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateCarrierAccountPayload is the carrier_account update action payload.
type UpdateCarrierAccountPayload struct {
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

// UpdateCarrierAccountPath computes a request path to the update action of carrier_account.
func UpdateCarrierAccountPath(id string) string {
	return fmt.Sprintf("/v2/carrier_accounts/%v", id)
}

// Updates can be made to description, reference, and any fields in credentials or test_credentials.
func (c *Client) UpdateCarrierAccount(ctx context.Context, path string, payload *UpdateCarrierAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateCarrierAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCarrierAccountRequest create the request corresponding to the update action endpoint of the carrier_account resource.
func (c *Client) NewUpdateCarrierAccountRequest(ctx context.Context, path string, payload *UpdateCarrierAccountPayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
