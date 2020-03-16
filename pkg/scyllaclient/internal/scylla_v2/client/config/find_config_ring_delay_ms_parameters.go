// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigRingDelayMsParams creates a new FindConfigRingDelayMsParams object
// with the default values initialized.
func NewFindConfigRingDelayMsParams() *FindConfigRingDelayMsParams {

	return &FindConfigRingDelayMsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigRingDelayMsParamsWithTimeout creates a new FindConfigRingDelayMsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigRingDelayMsParamsWithTimeout(timeout time.Duration) *FindConfigRingDelayMsParams {

	return &FindConfigRingDelayMsParams{

		timeout: timeout,
	}
}

// NewFindConfigRingDelayMsParamsWithContext creates a new FindConfigRingDelayMsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigRingDelayMsParamsWithContext(ctx context.Context) *FindConfigRingDelayMsParams {

	return &FindConfigRingDelayMsParams{

		Context: ctx,
	}
}

// NewFindConfigRingDelayMsParamsWithHTTPClient creates a new FindConfigRingDelayMsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigRingDelayMsParamsWithHTTPClient(client *http.Client) *FindConfigRingDelayMsParams {

	return &FindConfigRingDelayMsParams{
		HTTPClient: client,
	}
}

/*FindConfigRingDelayMsParams contains all the parameters to send to the API endpoint
for the find config ring delay ms operation typically these are written to a http.Request
*/
type FindConfigRingDelayMsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config ring delay ms params
func (o *FindConfigRingDelayMsParams) WithTimeout(timeout time.Duration) *FindConfigRingDelayMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config ring delay ms params
func (o *FindConfigRingDelayMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config ring delay ms params
func (o *FindConfigRingDelayMsParams) WithContext(ctx context.Context) *FindConfigRingDelayMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config ring delay ms params
func (o *FindConfigRingDelayMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config ring delay ms params
func (o *FindConfigRingDelayMsParams) WithHTTPClient(client *http.Client) *FindConfigRingDelayMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config ring delay ms params
func (o *FindConfigRingDelayMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigRingDelayMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
