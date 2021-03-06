// Code generated by go-swagger; DO NOT EDIT.

package apple

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/forestvpn/go-api-client-v1/models"
)

// NewVerifyReceiptParams creates a new VerifyReceiptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyReceiptParams() *VerifyReceiptParams {
	return &VerifyReceiptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyReceiptParamsWithTimeout creates a new VerifyReceiptParams object
// with the ability to set a timeout on a request.
func NewVerifyReceiptParamsWithTimeout(timeout time.Duration) *VerifyReceiptParams {
	return &VerifyReceiptParams{
		timeout: timeout,
	}
}

// NewVerifyReceiptParamsWithContext creates a new VerifyReceiptParams object
// with the ability to set a context for a request.
func NewVerifyReceiptParamsWithContext(ctx context.Context) *VerifyReceiptParams {
	return &VerifyReceiptParams{
		Context: ctx,
	}
}

// NewVerifyReceiptParamsWithHTTPClient creates a new VerifyReceiptParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyReceiptParamsWithHTTPClient(client *http.Client) *VerifyReceiptParams {
	return &VerifyReceiptParams{
		HTTPClient: client,
	}
}

/* VerifyReceiptParams contains all the parameters to send to the API endpoint
   for the verify receipt operation.

   Typically these are written to a http.Request.
*/
type VerifyReceiptParams struct {

	// Request.
	Request *models.AppleVerifyReceiptRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify receipt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyReceiptParams) WithDefaults() *VerifyReceiptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify receipt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyReceiptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify receipt params
func (o *VerifyReceiptParams) WithTimeout(timeout time.Duration) *VerifyReceiptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify receipt params
func (o *VerifyReceiptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify receipt params
func (o *VerifyReceiptParams) WithContext(ctx context.Context) *VerifyReceiptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify receipt params
func (o *VerifyReceiptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify receipt params
func (o *VerifyReceiptParams) WithHTTPClient(client *http.Client) *VerifyReceiptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify receipt params
func (o *VerifyReceiptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the verify receipt params
func (o *VerifyReceiptParams) WithRequest(request *models.AppleVerifyReceiptRequest) *VerifyReceiptParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the verify receipt params
func (o *VerifyReceiptParams) SetRequest(request *models.AppleVerifyReceiptRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyReceiptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
