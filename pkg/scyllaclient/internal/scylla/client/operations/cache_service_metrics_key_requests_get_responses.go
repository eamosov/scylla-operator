// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// CacheServiceMetricsKeyRequestsGetReader is a Reader for the CacheServiceMetricsKeyRequestsGet structure.
type CacheServiceMetricsKeyRequestsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsKeyRequestsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsKeyRequestsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsKeyRequestsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsKeyRequestsGetOK creates a CacheServiceMetricsKeyRequestsGetOK with default headers values
func NewCacheServiceMetricsKeyRequestsGetOK() *CacheServiceMetricsKeyRequestsGetOK {
	return &CacheServiceMetricsKeyRequestsGetOK{}
}

/*CacheServiceMetricsKeyRequestsGetOK handles this case with default header values.

CacheServiceMetricsKeyRequestsGetOK cache service metrics key requests get o k
*/
type CacheServiceMetricsKeyRequestsGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsKeyRequestsGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsKeyRequestsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsKeyRequestsGetDefault creates a CacheServiceMetricsKeyRequestsGetDefault with default headers values
func NewCacheServiceMetricsKeyRequestsGetDefault(code int) *CacheServiceMetricsKeyRequestsGetDefault {
	return &CacheServiceMetricsKeyRequestsGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsKeyRequestsGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsKeyRequestsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics key requests get default response
func (o *CacheServiceMetricsKeyRequestsGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsKeyRequestsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsKeyRequestsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsKeyRequestsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
