package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ShowCarrierTypesPath computes a request path to the show action of carrier_types.
func ShowCarrierTypesPath() string {
	return fmt.Sprintf("/v2/carrier_types")
}

// List carrier types.
func (c *Client) ShowCarrierTypes(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowCarrierTypesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCarrierTypesRequest create the request corresponding to the show action endpoint of the carrier_types resource.
func (c *Client) NewShowCarrierTypesRequest(ctx context.Context, path string) (*http.Request, error) {
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
