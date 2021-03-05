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

// GetProfilesReader is a Reader for the GetProfiles structure.
type GetProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProfilesOK creates a GetProfilesOK with default headers values
func NewGetProfilesOK() *GetProfilesOK {
	return &GetProfilesOK{}
}

/* GetProfilesOK describes a response with status code 200, with default header values.

OK
*/
type GetProfilesOK struct {
	Payload []*models.WireGuardProfile
}

func (o *GetProfilesOK) Error() string {
	return fmt.Sprintf("[GET /wg/profiles/][%d] getProfilesOK  %+v", 200, o.Payload)
}
func (o *GetProfilesOK) GetPayload() []*models.WireGuardProfile {
	return o.Payload
}

func (o *GetProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesDefault creates a GetProfilesDefault with default headers values
func NewGetProfilesDefault(code int) *GetProfilesDefault {
	return &GetProfilesDefault{
		_statusCode: code,
	}
}

/* GetProfilesDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GetProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get profiles default response
func (o *GetProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /wg/profiles/][%d] GetProfiles default  %+v", o._statusCode, o.Payload)
}
func (o *GetProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}