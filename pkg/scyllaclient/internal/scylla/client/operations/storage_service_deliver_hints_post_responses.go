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

// StorageServiceDeliverHintsPostReader is a Reader for the StorageServiceDeliverHintsPost structure.
type StorageServiceDeliverHintsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceDeliverHintsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceDeliverHintsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceDeliverHintsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceDeliverHintsPostOK creates a StorageServiceDeliverHintsPostOK with default headers values
func NewStorageServiceDeliverHintsPostOK() *StorageServiceDeliverHintsPostOK {
	return &StorageServiceDeliverHintsPostOK{}
}

/*StorageServiceDeliverHintsPostOK handles this case with default header values.

StorageServiceDeliverHintsPostOK storage service deliver hints post o k
*/
type StorageServiceDeliverHintsPostOK struct {
}

func (o *StorageServiceDeliverHintsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceDeliverHintsPostDefault creates a StorageServiceDeliverHintsPostDefault with default headers values
func NewStorageServiceDeliverHintsPostDefault(code int) *StorageServiceDeliverHintsPostDefault {
	return &StorageServiceDeliverHintsPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceDeliverHintsPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceDeliverHintsPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service deliver hints post default response
func (o *StorageServiceDeliverHintsPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceDeliverHintsPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceDeliverHintsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceDeliverHintsPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
