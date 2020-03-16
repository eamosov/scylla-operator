// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewCompactionManagerCompactionHistoryGetParams creates a new CompactionManagerCompactionHistoryGetParams object
// with the default values initialized.
func NewCompactionManagerCompactionHistoryGetParams() *CompactionManagerCompactionHistoryGetParams {

	return &CompactionManagerCompactionHistoryGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompactionManagerCompactionHistoryGetParamsWithTimeout creates a new CompactionManagerCompactionHistoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompactionManagerCompactionHistoryGetParamsWithTimeout(timeout time.Duration) *CompactionManagerCompactionHistoryGetParams {

	return &CompactionManagerCompactionHistoryGetParams{

		timeout: timeout,
	}
}

// NewCompactionManagerCompactionHistoryGetParamsWithContext creates a new CompactionManagerCompactionHistoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompactionManagerCompactionHistoryGetParamsWithContext(ctx context.Context) *CompactionManagerCompactionHistoryGetParams {

	return &CompactionManagerCompactionHistoryGetParams{

		Context: ctx,
	}
}

// NewCompactionManagerCompactionHistoryGetParamsWithHTTPClient creates a new CompactionManagerCompactionHistoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompactionManagerCompactionHistoryGetParamsWithHTTPClient(client *http.Client) *CompactionManagerCompactionHistoryGetParams {

	return &CompactionManagerCompactionHistoryGetParams{
		HTTPClient: client,
	}
}

/*CompactionManagerCompactionHistoryGetParams contains all the parameters to send to the API endpoint
for the compaction manager compaction history get operation typically these are written to a http.Request
*/
type CompactionManagerCompactionHistoryGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the compaction manager compaction history get params
func (o *CompactionManagerCompactionHistoryGetParams) WithTimeout(timeout time.Duration) *CompactionManagerCompactionHistoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compaction manager compaction history get params
func (o *CompactionManagerCompactionHistoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the compaction manager compaction history get params
func (o *CompactionManagerCompactionHistoryGetParams) WithContext(ctx context.Context) *CompactionManagerCompactionHistoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compaction manager compaction history get params
func (o *CompactionManagerCompactionHistoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the compaction manager compaction history get params
func (o *CompactionManagerCompactionHistoryGetParams) WithHTTPClient(client *http.Client) *CompactionManagerCompactionHistoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the compaction manager compaction history get params
func (o *CompactionManagerCompactionHistoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CompactionManagerCompactionHistoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
