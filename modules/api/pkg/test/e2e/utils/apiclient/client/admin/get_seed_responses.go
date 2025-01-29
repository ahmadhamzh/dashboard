// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetSeedReader is a Reader for the GetSeed structure.
type GetSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSeedOK creates a GetSeedOK with default headers values
func NewGetSeedOK() *GetSeedOK {
	return &GetSeedOK{}
}

/*
GetSeedOK describes a response with status code 200, with default header values.

Seed
*/
type GetSeedOK struct {
	Payload *models.Seed
}

// IsSuccess returns true when this get seed o k response has a 2xx status code
func (o *GetSeedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get seed o k response has a 3xx status code
func (o *GetSeedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get seed o k response has a 4xx status code
func (o *GetSeedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get seed o k response has a 5xx status code
func (o *GetSeedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get seed o k response a status code equal to that given
func (o *GetSeedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get seed o k response
func (o *GetSeedOK) Code() int {
	return 200
}

func (o *GetSeedOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedOK %s", 200, payload)
}

func (o *GetSeedOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedOK %s", 200, payload)
}

func (o *GetSeedOK) GetPayload() *models.Seed {
	return o.Payload
}

func (o *GetSeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Seed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeedUnauthorized creates a GetSeedUnauthorized with default headers values
func NewGetSeedUnauthorized() *GetSeedUnauthorized {
	return &GetSeedUnauthorized{}
}

/*
GetSeedUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetSeedUnauthorized struct {
}

// IsSuccess returns true when this get seed unauthorized response has a 2xx status code
func (o *GetSeedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get seed unauthorized response has a 3xx status code
func (o *GetSeedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get seed unauthorized response has a 4xx status code
func (o *GetSeedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get seed unauthorized response has a 5xx status code
func (o *GetSeedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get seed unauthorized response a status code equal to that given
func (o *GetSeedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get seed unauthorized response
func (o *GetSeedUnauthorized) Code() int {
	return 401
}

func (o *GetSeedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedUnauthorized", 401)
}

func (o *GetSeedUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedUnauthorized", 401)
}

func (o *GetSeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeedForbidden creates a GetSeedForbidden with default headers values
func NewGetSeedForbidden() *GetSeedForbidden {
	return &GetSeedForbidden{}
}

/*
GetSeedForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetSeedForbidden struct {
}

// IsSuccess returns true when this get seed forbidden response has a 2xx status code
func (o *GetSeedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get seed forbidden response has a 3xx status code
func (o *GetSeedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get seed forbidden response has a 4xx status code
func (o *GetSeedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get seed forbidden response has a 5xx status code
func (o *GetSeedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get seed forbidden response a status code equal to that given
func (o *GetSeedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get seed forbidden response
func (o *GetSeedForbidden) Code() int {
	return 403
}

func (o *GetSeedForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedForbidden", 403)
}

func (o *GetSeedForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedForbidden", 403)
}

func (o *GetSeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeedDefault creates a GetSeedDefault with default headers values
func NewGetSeedDefault(code int) *GetSeedDefault {
	return &GetSeedDefault{
		_statusCode: code,
	}
}

/*
GetSeedDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetSeedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get seed default response has a 2xx status code
func (o *GetSeedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get seed default response has a 3xx status code
func (o *GetSeedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get seed default response has a 4xx status code
func (o *GetSeedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get seed default response has a 5xx status code
func (o *GetSeedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get seed default response a status code equal to that given
func (o *GetSeedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get seed default response
func (o *GetSeedDefault) Code() int {
	return o._statusCode
}

func (o *GetSeedDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeed default %s", o._statusCode, payload)
}

func (o *GetSeedDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeed default %s", o._statusCode, payload)
}

func (o *GetSeedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
