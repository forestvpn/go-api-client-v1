// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/forestvpn/go-api-client-v1/models"
)

// LookupReader is a Reader for the Lookup structure.
type LookupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLookupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLookupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLookupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLookupOK creates a LookupOK with default headers values
func NewLookupOK() *LookupOK {
	return &LookupOK{}
}

/* LookupOK describes a response with status code 200, with default header values.

OK
*/
type LookupOK struct {
	Payload *models.Lookup
}

func (o *LookupOK) Error() string {
	return fmt.Sprintf("[POST /dns/lookup][%d] lookupOK  %+v", 200, o.Payload)
}
func (o *LookupOK) GetPayload() *models.Lookup {
	return o.Payload
}

func (o *LookupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Lookup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookupNotFound creates a LookupNotFound with default headers values
func NewLookupNotFound() *LookupNotFound {
	return &LookupNotFound{}
}

/* LookupNotFound describes a response with status code 404, with default header values.

Not found
*/
type LookupNotFound struct {
}

func (o *LookupNotFound) Error() string {
	return fmt.Sprintf("[POST /dns/lookup][%d] lookupNotFound ", 404)
}

func (o *LookupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLookupDefault creates a LookupDefault with default headers values
func NewLookupDefault(code int) *LookupDefault {
	return &LookupDefault{
		_statusCode: code,
	}
}

/* LookupDefault describes a response with status code -1, with default header values.

error
*/
type LookupDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the lookup default response
func (o *LookupDefault) Code() int {
	return o._statusCode
}

func (o *LookupDefault) Error() string {
	return fmt.Sprintf("[POST /dns/lookup][%d] Lookup default  %+v", o._statusCode, o.Payload)
}
func (o *LookupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *LookupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
