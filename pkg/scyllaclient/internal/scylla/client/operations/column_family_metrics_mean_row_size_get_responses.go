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

// ColumnFamilyMetricsMeanRowSizeGetReader is a Reader for the ColumnFamilyMetricsMeanRowSizeGet structure.
type ColumnFamilyMetricsMeanRowSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMeanRowSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMeanRowSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMeanRowSizeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMeanRowSizeGetOK creates a ColumnFamilyMetricsMeanRowSizeGetOK with default headers values
func NewColumnFamilyMetricsMeanRowSizeGetOK() *ColumnFamilyMetricsMeanRowSizeGetOK {
	return &ColumnFamilyMetricsMeanRowSizeGetOK{}
}

/*ColumnFamilyMetricsMeanRowSizeGetOK handles this case with default header values.

ColumnFamilyMetricsMeanRowSizeGetOK column family metrics mean row size get o k
*/
type ColumnFamilyMetricsMeanRowSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMeanRowSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMeanRowSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMeanRowSizeGetDefault creates a ColumnFamilyMetricsMeanRowSizeGetDefault with default headers values
func NewColumnFamilyMetricsMeanRowSizeGetDefault(code int) *ColumnFamilyMetricsMeanRowSizeGetDefault {
	return &ColumnFamilyMetricsMeanRowSizeGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMeanRowSizeGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMeanRowSizeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics mean row size get default response
func (o *ColumnFamilyMetricsMeanRowSizeGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMeanRowSizeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMeanRowSizeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMeanRowSizeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
