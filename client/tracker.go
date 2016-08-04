package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateTrackerPayload is the tracker create action payload.
type CreateTrackerPayload struct {
	// The carrier associated with the tracking_code you provided. The carrier will get auto-detected if none is provided
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The tracking code associated with the package you'd like to track
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// CreateTrackerPath computes a request path to the create action of tracker.
func CreateTrackerPath() string {
	return fmt.Sprintf("/v2/trackers")
}

// A Tracker can be created with only a tracking_code. Optionally, you can provide the carrier parameter, which indicates the carrier the package was shipped with. If no carrier is provided, EasyPost will attempt to determine the carrier based on the tracking_code provided. Providing a carrier parameter is recommended, since some tracking_codes are ambiguous and may match with more than one carrier. In addition, not having to auto-match the carrier will significantly speed up the response time.
func (c *Client) CreateTracker(ctx context.Context, path string, payload *CreateTrackerPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateTrackerRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTrackerRequest create the request corresponding to the create action endpoint of the tracker resource.
func (c *Client) NewCreateTrackerRequest(ctx context.Context, path string, payload *CreateTrackerPayload, contentType string) (*http.Request, error) {
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

// ListTrackerPayload is the tracker list action payload.
type ListTrackerPayload struct {
	// Optional pagination parameter. Only trackers created after the given ID will be included. May not be used with before_id
	AfterID *string `form:"after_id,omitempty" json:"after_id,omitempty" xml:"after_id,omitempty"`
	// Optional pagination parameter. Only trackers created before the given ID will be included. May not be used with after_id
	BeforeID *string `form:"before_id,omitempty" json:"before_id,omitempty" xml:"before_id,omitempty"`
	// Only returns Trackers with the given carrier value
	Carrier *string `form:"carrier,omitempty" json:"carrier,omitempty" xml:"carrier,omitempty"`
	// Only return Trackers created before this timestamp. Defaults to end of the current day, or 1 month after a passed start_datetime
	EndDatetime *string `form:"end_datetime,omitempty" json:"end_datetime,omitempty" xml:"end_datetime,omitempty"`
	// The number of Trackers to return on each page. The maximum value is 100
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
	// Only return Trackers created after this timestamp. Defaults to 1 month ago, or 1 month before a passed end_datetime
	StartDatetime *string `form:"start_datetime,omitempty" json:"start_datetime,omitempty" xml:"start_datetime,omitempty"`
	// Only returns Trackers with the given tracking_code. Useful for retrieving an individual Tracker by tracking_code rather than by ID
	TrackingCode string `form:"tracking_code" json:"tracking_code" xml:"tracking_code"`
}

// ListTrackerPath computes a request path to the list action of tracker.
func ListTrackerPath() string {
	return fmt.Sprintf("/v2/trackers")
}

// The Tracker List is a paginated list of all Tracker objects associated with the given API Key. It accepts a variety of parameters which can be used to modify the scope. The has_more attribute indicates whether or not additional pages can be requested. The recommended way of paginating is to use either the before_id or after_id parameter to specify where the next page begins.
func (c *Client) ListTracker(ctx context.Context, path string, payload *ListTrackerPayload, contentType string) (*http.Response, error) {
	req, err := c.NewListTrackerRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTrackerRequest create the request corresponding to the list action endpoint of the tracker resource.
func (c *Client) NewListTrackerRequest(ctx context.Context, path string, payload *ListTrackerPayload, contentType string) (*http.Request, error) {
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

// ShowTrackerPath computes a request path to the show action of tracker.
func ShowTrackerPath(id string) string {
	return fmt.Sprintf("/v2/trackers/%v", id)
}

// Retrieve a Tracker by id.
func (c *Client) ShowTracker(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTrackerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTrackerRequest create the request corresponding to the show action endpoint of the tracker resource.
func (c *Client) NewShowTrackerRequest(ctx context.Context, path string) (*http.Request, error) {
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
