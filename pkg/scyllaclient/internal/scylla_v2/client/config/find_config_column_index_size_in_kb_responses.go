// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigColumnIndexSizeInKbReader is a Reader for the FindConfigColumnIndexSizeInKb structure.
type FindConfigColumnIndexSizeInKbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigColumnIndexSizeInKbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigColumnIndexSizeInKbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigColumnIndexSizeInKbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigColumnIndexSizeInKbOK creates a FindConfigColumnIndexSizeInKbOK with default headers values
func NewFindConfigColumnIndexSizeInKbOK() *FindConfigColumnIndexSizeInKbOK {
	return &FindConfigColumnIndexSizeInKbOK{}
}

/*FindConfigColumnIndexSizeInKbOK handles this case with default header values.

Config value
*/
type FindConfigColumnIndexSizeInKbOK struct {
	Payload int64
}

func (o *FindConfigColumnIndexSizeInKbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigColumnIndexSizeInKbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigColumnIndexSizeInKbDefault creates a FindConfigColumnIndexSizeInKbDefault with default headers values
func NewFindConfigColumnIndexSizeInKbDefault(code int) *FindConfigColumnIndexSizeInKbDefault {
	return &FindConfigColumnIndexSizeInKbDefault{
		_statusCode: code,
	}
}

/*FindConfigColumnIndexSizeInKbDefault handles this case with default header values.

unexpected error
*/
type FindConfigColumnIndexSizeInKbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config column index size in kb default response
func (o *FindConfigColumnIndexSizeInKbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigColumnIndexSizeInKbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigColumnIndexSizeInKbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigColumnIndexSizeInKbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
