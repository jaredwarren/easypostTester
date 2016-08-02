package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/jaredwarren/easypostTester/app"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// ShowCarrierTypesOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCarrierTypesOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CarrierTypesController) (http.ResponseWriter, app.EasypostCarrierTypesCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v2/carrier_types"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CarrierTypesTest"), rw, req, prms)
	showCtx, err := app.NewShowCarrierTypesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.EasypostCarrierTypesCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.EasypostCarrierTypesCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.EasypostCarrierTypesCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
