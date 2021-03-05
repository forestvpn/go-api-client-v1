// Code generated by go-swagger; DO NOT EDIT.

package usage

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
)

// NewGetTrafficParams creates a new GetTrafficParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrafficParams() *GetTrafficParams {
	return &GetTrafficParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrafficParamsWithTimeout creates a new GetTrafficParams object
// with the ability to set a timeout on a request.
func NewGetTrafficParamsWithTimeout(timeout time.Duration) *GetTrafficParams {
	return &GetTrafficParams{
		timeout: timeout,
	}
}

// NewGetTrafficParamsWithContext creates a new GetTrafficParams object
// with the ability to set a context for a request.
func NewGetTrafficParamsWithContext(ctx context.Context) *GetTrafficParams {
	return &GetTrafficParams{
		Context: ctx,
	}
}

// NewGetTrafficParamsWithHTTPClient creates a new GetTrafficParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrafficParamsWithHTTPClient(client *http.Client) *GetTrafficParams {
	return &GetTrafficParams{
		HTTPClient: client,
	}
}

/* GetTrafficParams contains all the parameters to send to the API endpoint
   for the get traffic operation.

   Typically these are written to a http.Request.
*/
type GetTrafficParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get traffic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrafficParams) WithDefaults() *GetTrafficParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get traffic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrafficParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get traffic params
func (o *GetTrafficParams) WithTimeout(timeout time.Duration) *GetTrafficParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get traffic params
func (o *GetTrafficParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get traffic params
func (o *GetTrafficParams) WithContext(ctx context.Context) *GetTrafficParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get traffic params
func (o *GetTrafficParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get traffic params
func (o *GetTrafficParams) WithHTTPClient(client *http.Client) *GetTrafficParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get traffic params
func (o *GetTrafficParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrafficParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}