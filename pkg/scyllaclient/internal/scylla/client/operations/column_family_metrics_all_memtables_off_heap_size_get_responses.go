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

// ColumnFamilyMetricsAllMemtablesOffHeapSizeGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesOffHeapSizeGet structure.
type ColumnFamilyMetricsAllMemtablesOffHeapSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK creates a ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK() *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK {
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK{}
}

/*ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK column family metrics all memtables off heap size get o k
*/
type ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault creates a ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault with default headers values
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault(code int) *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault {
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics all memtables off heap size get default response
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
