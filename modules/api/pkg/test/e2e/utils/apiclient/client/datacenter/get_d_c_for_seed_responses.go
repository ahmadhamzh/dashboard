// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetDCForSeedReader is a Reader for the GetDCForSeed structure.
type GetDCForSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDCForSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDCForSeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDCForSeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDCForSeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDCForSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDCForSeedOK creates a GetDCForSeedOK with default headers values
func NewGetDCForSeedOK() *GetDCForSeedOK {
	return &GetDCForSeedOK{}
}

/*
GetDCForSeedOK describes a response with status code 200, with default header values.

Datacenter
*/
type GetDCForSeedOK struct {
	Payload *models.Datacenter
}

// IsSuccess returns true when this get d c for seed o k response has a 2xx status code
func (o *GetDCForSeedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d c for seed o k response has a 3xx status code
func (o *GetDCForSeedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d c for seed o k response has a 4xx status code
func (o *GetDCForSeedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d c for seed o k response has a 5xx status code
func (o *GetDCForSeedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d c for seed o k response a status code equal to that given
func (o *GetDCForSeedOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDCForSeedOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeedOK  %+v", 200, o.Payload)
}

func (o *GetDCForSeedOK) String() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeedOK  %+v", 200, o.Payload)
}

func (o *GetDCForSeedOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *GetDCForSeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDCForSeedUnauthorized creates a GetDCForSeedUnauthorized with default headers values
func NewGetDCForSeedUnauthorized() *GetDCForSeedUnauthorized {
	return &GetDCForSeedUnauthorized{}
}

/*
GetDCForSeedUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetDCForSeedUnauthorized struct {
}

// IsSuccess returns true when this get d c for seed unauthorized response has a 2xx status code
func (o *GetDCForSeedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d c for seed unauthorized response has a 3xx status code
func (o *GetDCForSeedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d c for seed unauthorized response has a 4xx status code
func (o *GetDCForSeedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d c for seed unauthorized response has a 5xx status code
func (o *GetDCForSeedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get d c for seed unauthorized response a status code equal to that given
func (o *GetDCForSeedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDCForSeedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeedUnauthorized ", 401)
}

func (o *GetDCForSeedUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeedUnauthorized ", 401)
}

func (o *GetDCForSeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDCForSeedForbidden creates a GetDCForSeedForbidden with default headers values
func NewGetDCForSeedForbidden() *GetDCForSeedForbidden {
	return &GetDCForSeedForbidden{}
}

/*
GetDCForSeedForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetDCForSeedForbidden struct {
}

// IsSuccess returns true when this get d c for seed forbidden response has a 2xx status code
func (o *GetDCForSeedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d c for seed forbidden response has a 3xx status code
func (o *GetDCForSeedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d c for seed forbidden response has a 4xx status code
func (o *GetDCForSeedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d c for seed forbidden response has a 5xx status code
func (o *GetDCForSeedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get d c for seed forbidden response a status code equal to that given
func (o *GetDCForSeedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDCForSeedForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeedForbidden ", 403)
}

func (o *GetDCForSeedForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeedForbidden ", 403)
}

func (o *GetDCForSeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDCForSeedDefault creates a GetDCForSeedDefault with default headers values
func NewGetDCForSeedDefault(code int) *GetDCForSeedDefault {
	return &GetDCForSeedDefault{
		_statusCode: code,
	}
}

/*
GetDCForSeedDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetDCForSeedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get d c for seed default response
func (o *GetDCForSeedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get d c for seed default response has a 2xx status code
func (o *GetDCForSeedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get d c for seed default response has a 3xx status code
func (o *GetDCForSeedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get d c for seed default response has a 4xx status code
func (o *GetDCForSeedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get d c for seed default response has a 5xx status code
func (o *GetDCForSeedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get d c for seed default response a status code equal to that given
func (o *GetDCForSeedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDCForSeedDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeed default  %+v", o._statusCode, o.Payload)
}

func (o *GetDCForSeedDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/seed/{seed_name}/dc/{dc}][%d] getDCForSeed default  %+v", o._statusCode, o.Payload)
}

func (o *GetDCForSeedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDCForSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
