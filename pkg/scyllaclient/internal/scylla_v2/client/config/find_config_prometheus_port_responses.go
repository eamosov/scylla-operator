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

// FindConfigPrometheusPortReader is a Reader for the FindConfigPrometheusPort structure.
type FindConfigPrometheusPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPrometheusPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPrometheusPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPrometheusPortDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPrometheusPortOK creates a FindConfigPrometheusPortOK with default headers values
func NewFindConfigPrometheusPortOK() *FindConfigPrometheusPortOK {
	return &FindConfigPrometheusPortOK{}
}

/*FindConfigPrometheusPortOK handles this case with default header values.

Config value
*/
type FindConfigPrometheusPortOK struct {
	Payload int64
}

func (o *FindConfigPrometheusPortOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigPrometheusPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPrometheusPortDefault creates a FindConfigPrometheusPortDefault with default headers values
func NewFindConfigPrometheusPortDefault(code int) *FindConfigPrometheusPortDefault {
	return &FindConfigPrometheusPortDefault{
		_statusCode: code,
	}
}

/*FindConfigPrometheusPortDefault handles this case with default header values.

unexpected error
*/
type FindConfigPrometheusPortDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config prometheus port default response
func (o *FindConfigPrometheusPortDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPrometheusPortDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPrometheusPortDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPrometheusPortDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
