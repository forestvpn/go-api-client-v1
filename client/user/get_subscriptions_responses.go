// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/forestvpn/go-api-client-v1/models"
)

// GetSubscriptionsReader is a Reader for the GetSubscriptions structure.
type GetSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionsOK creates a GetSubscriptionsOK with default headers values
func NewGetSubscriptionsOK() *GetSubscriptionsOK {
	return &GetSubscriptionsOK{}
}

/* GetSubscriptionsOK describes a response with status code 200, with default header values.

OK
*/
type GetSubscriptionsOK struct {
	Payload []*models.Subscription
}

func (o *GetSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /user/subscriptions/][%d] getSubscriptionsOK  %+v", 200, o.Payload)
}
func (o *GetSubscriptionsOK) GetPayload() []*models.Subscription {
	return o.Payload
}

func (o *GetSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsDefault creates a GetSubscriptionsDefault with default headers values
func NewGetSubscriptionsDefault(code int) *GetSubscriptionsDefault {
	return &GetSubscriptionsDefault{
		_statusCode: code,
	}
}

/* GetSubscriptionsDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GetSubscriptionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get subscriptions default response
func (o *GetSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[GET /user/subscriptions/][%d] GetSubscriptions default  %+v", o._statusCode, o.Payload)
}
func (o *GetSubscriptionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSubscriptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
