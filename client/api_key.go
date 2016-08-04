package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ListAPIKeyPath computes a request path to the list action of api_key.
func ListAPIKeyPath() string {
	return fmt.Sprintf("/v2/api_keys")
}

// Both production and test keys will be returned a User and all of its children. If the request is authenticated as a Child only the API Keys for that Child will be returned.
func (c *Client) ListAPIKey(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAPIKeyRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAPIKeyRequest create the request corresponding to the list action endpoint of the api_key resource.
func (c *Client) NewListAPIKeyRequest(ctx context.Context, path string) (*http.Request, error) {
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
