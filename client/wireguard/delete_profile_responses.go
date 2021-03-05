// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/forestvpn/go-api-client-v1/models"
)

// DeleteProfileReader is a Reader for the DeleteProfile structure.
type DeleteProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProfileNoContent creates a DeleteProfileNoContent with default headers values
func NewDeleteProfileNoContent() *DeleteProfileNoContent {
	return &DeleteProfileNoContent{}
}

/* DeleteProfileNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteProfileNoContent struct {
}

func (o *DeleteProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wg/profiles/{profileId}/][%d] deleteProfileNoContent ", 204)
}

func (o *DeleteProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProfileDefault creates a DeleteProfileDefault with default headers values
func NewDeleteProfileDefault(code int) *DeleteProfileDefault {
	return &DeleteProfileDefault{
		_statusCode: code,
	}
}

/* DeleteProfileDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type DeleteProfileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete profile default response
func (o *DeleteProfileDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProfileDefault) Error() string {
	return fmt.Sprintf("[DELETE /wg/profiles/{profileId}/][%d] DeleteProfile default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteProfileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}