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

// StreamManagerMetricsOutgoingGetReader is a Reader for the StreamManagerMetricsOutgoingGet structure.
type StreamManagerMetricsOutgoingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerMetricsOutgoingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamManagerMetricsOutgoingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStreamManagerMetricsOutgoingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStreamManagerMetricsOutgoingGetOK creates a StreamManagerMetricsOutgoingGetOK with default headers values
func NewStreamManagerMetricsOutgoingGetOK() *StreamManagerMetricsOutgoingGetOK {
	return &StreamManagerMetricsOutgoingGetOK{}
}

/*StreamManagerMetricsOutgoingGetOK handles this case with default header values.

StreamManagerMetricsOutgoingGetOK stream manager metrics outgoing get o k
*/
type StreamManagerMetricsOutgoingGetOK struct {
	Payload int32
}

func (o *StreamManagerMetricsOutgoingGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StreamManagerMetricsOutgoingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamManagerMetricsOutgoingGetDefault creates a StreamManagerMetricsOutgoingGetDefault with default headers values
func NewStreamManagerMetricsOutgoingGetDefault(code int) *StreamManagerMetricsOutgoingGetDefault {
	return &StreamManagerMetricsOutgoingGetDefault{
		_statusCode: code,
	}
}

/*StreamManagerMetricsOutgoingGetDefault handles this case with default header values.

internal server error
*/
type StreamManagerMetricsOutgoingGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the stream manager metrics outgoing get default response
func (o *StreamManagerMetricsOutgoingGetDefault) Code() int {
	return o._statusCode
}

func (o *StreamManagerMetricsOutgoingGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StreamManagerMetricsOutgoingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StreamManagerMetricsOutgoingGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
