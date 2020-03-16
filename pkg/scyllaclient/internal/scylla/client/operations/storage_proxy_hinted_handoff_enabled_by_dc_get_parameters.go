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

// NewStorageProxyHintedHandoffEnabledByDcGetParams creates a new StorageProxyHintedHandoffEnabledByDcGetParams object
// with the default values initialized.
func NewStorageProxyHintedHandoffEnabledByDcGetParams() *StorageProxyHintedHandoffEnabledByDcGetParams {

	return &StorageProxyHintedHandoffEnabledByDcGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyHintedHandoffEnabledByDcGetParamsWithTimeout creates a new StorageProxyHintedHandoffEnabledByDcGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyHintedHandoffEnabledByDcGetParamsWithTimeout(timeout time.Duration) *StorageProxyHintedHandoffEnabledByDcGetParams {

	return &StorageProxyHintedHandoffEnabledByDcGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyHintedHandoffEnabledByDcGetParamsWithContext creates a new StorageProxyHintedHandoffEnabledByDcGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyHintedHandoffEnabledByDcGetParamsWithContext(ctx context.Context) *StorageProxyHintedHandoffEnabledByDcGetParams {

	return &StorageProxyHintedHandoffEnabledByDcGetParams{

		Context: ctx,
	}
}

// NewStorageProxyHintedHandoffEnabledByDcGetParamsWithHTTPClient creates a new StorageProxyHintedHandoffEnabledByDcGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyHintedHandoffEnabledByDcGetParamsWithHTTPClient(client *http.Client) *StorageProxyHintedHandoffEnabledByDcGetParams {

	return &StorageProxyHintedHandoffEnabledByDcGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyHintedHandoffEnabledByDcGetParams contains all the parameters to send to the API endpoint
for the storage proxy hinted handoff enabled by dc get operation typically these are written to a http.Request
*/
type StorageProxyHintedHandoffEnabledByDcGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy hinted handoff enabled by dc get params
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) WithTimeout(timeout time.Duration) *StorageProxyHintedHandoffEnabledByDcGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy hinted handoff enabled by dc get params
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy hinted handoff enabled by dc get params
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) WithContext(ctx context.Context) *StorageProxyHintedHandoffEnabledByDcGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy hinted handoff enabled by dc get params
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy hinted handoff enabled by dc get params
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) WithHTTPClient(client *http.Client) *StorageProxyHintedHandoffEnabledByDcGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy hinted handoff enabled by dc get params
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyHintedHandoffEnabledByDcGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
