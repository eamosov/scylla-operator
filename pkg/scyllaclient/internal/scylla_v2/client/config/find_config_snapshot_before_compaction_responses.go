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

// FindConfigSnapshotBeforeCompactionReader is a Reader for the FindConfigSnapshotBeforeCompaction structure.
type FindConfigSnapshotBeforeCompactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigSnapshotBeforeCompactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigSnapshotBeforeCompactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigSnapshotBeforeCompactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigSnapshotBeforeCompactionOK creates a FindConfigSnapshotBeforeCompactionOK with default headers values
func NewFindConfigSnapshotBeforeCompactionOK() *FindConfigSnapshotBeforeCompactionOK {
	return &FindConfigSnapshotBeforeCompactionOK{}
}

/*FindConfigSnapshotBeforeCompactionOK handles this case with default header values.

Config value
*/
type FindConfigSnapshotBeforeCompactionOK struct {
	Payload bool
}

func (o *FindConfigSnapshotBeforeCompactionOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigSnapshotBeforeCompactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigSnapshotBeforeCompactionDefault creates a FindConfigSnapshotBeforeCompactionDefault with default headers values
func NewFindConfigSnapshotBeforeCompactionDefault(code int) *FindConfigSnapshotBeforeCompactionDefault {
	return &FindConfigSnapshotBeforeCompactionDefault{
		_statusCode: code,
	}
}

/*FindConfigSnapshotBeforeCompactionDefault handles this case with default header values.

unexpected error
*/
type FindConfigSnapshotBeforeCompactionDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config snapshot before compaction default response
func (o *FindConfigSnapshotBeforeCompactionDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigSnapshotBeforeCompactionDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigSnapshotBeforeCompactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigSnapshotBeforeCompactionDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
