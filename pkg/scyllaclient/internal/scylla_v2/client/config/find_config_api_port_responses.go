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

// FindConfigAPIPortReader is a Reader for the FindConfigAPIPort structure.
type FindConfigAPIPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAPIPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAPIPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAPIPortDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAPIPortOK creates a FindConfigAPIPortOK with default headers values
func NewFindConfigAPIPortOK() *FindConfigAPIPortOK {
	return &FindConfigAPIPortOK{}
}

/*FindConfigAPIPortOK handles this case with default header values.

Config value
*/
type FindConfigAPIPortOK struct {
	Payload int64
}

func (o *FindConfigAPIPortOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigAPIPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAPIPortDefault creates a FindConfigAPIPortDefault with default headers values
func NewFindConfigAPIPortDefault(code int) *FindConfigAPIPortDefault {
	return &FindConfigAPIPortDefault{
		_statusCode: code,
	}
}

/*FindConfigAPIPortDefault handles this case with default header values.

unexpected error
*/
type FindConfigAPIPortDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config api port default response
func (o *FindConfigAPIPortDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAPIPortDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAPIPortDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAPIPortDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
