package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Email string `form:"email" json:"email" xml:"email"`
	// First and last name required
	Name                 string  `form:"name" json:"name" xml:"name"`
	Password             string  `form:"password" json:"password" xml:"password"`
	PasswordConfirmation string  `form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation"`
	PhoneNumber          *string `form:"phone_number,omitempty" json:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {
	return fmt.Sprintf("/v2/users")
}

// CreateUser makes a request to the create action endpoint of the user resource
func (c *Client) CreateUser(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Request, error) {
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

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(id string) string {
	return fmt.Sprintf("/v2/users/%v", id)
}

// ShowUser makes a request to the show action endpoint of the user resource
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
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

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(id string) string {
	return fmt.Sprintf("/v2/users/%v", id)
}

// UpdateUser makes a request to the update action endpoint of the user resource
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UpdateUserPayload, contentType string) (*http.Request, error) {
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
