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

// StorageServiceOwnershipByKeyspaceGetReader is a Reader for the StorageServiceOwnershipByKeyspaceGet structure.
type StorageServiceOwnershipByKeyspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceOwnershipByKeyspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceOwnershipByKeyspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceOwnershipByKeyspaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceOwnershipByKeyspaceGetOK creates a StorageServiceOwnershipByKeyspaceGetOK with default headers values
func NewStorageServiceOwnershipByKeyspaceGetOK() *StorageServiceOwnershipByKeyspaceGetOK {
	return &StorageServiceOwnershipByKeyspaceGetOK{}
}

/*StorageServiceOwnershipByKeyspaceGetOK handles this case with default header values.

StorageServiceOwnershipByKeyspaceGetOK storage service ownership by keyspace get o k
*/
type StorageServiceOwnershipByKeyspaceGetOK struct {
	Payload []*models.Mapper
}

func (o *StorageServiceOwnershipByKeyspaceGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *StorageServiceOwnershipByKeyspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceOwnershipByKeyspaceGetDefault creates a StorageServiceOwnershipByKeyspaceGetDefault with default headers values
func NewStorageServiceOwnershipByKeyspaceGetDefault(code int) *StorageServiceOwnershipByKeyspaceGetDefault {
	return &StorageServiceOwnershipByKeyspaceGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceOwnershipByKeyspaceGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceOwnershipByKeyspaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service ownership by keyspace get default response
func (o *StorageServiceOwnershipByKeyspaceGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceOwnershipByKeyspaceGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceOwnershipByKeyspaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceOwnershipByKeyspaceGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
