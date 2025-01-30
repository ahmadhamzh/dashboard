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

// GetPolicyBindingReader is a Reader for the GetPolicyBinding structure.
type GetPolicyBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPolicyBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicyBindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPolicyBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPolicyBindingOK creates a GetPolicyBindingOK with default headers values
func NewGetPolicyBindingOK() *GetPolicyBindingOK {
	return &GetPolicyBindingOK{}
}

/*
GetPolicyBindingOK describes a response with status code 200, with default header values.

PolicyBinding
*/
type GetPolicyBindingOK struct {
	Payload *models.PolicyBinding
}

// IsSuccess returns true when this get policy binding o k response has a 2xx status code
func (o *GetPolicyBindingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policy binding o k response has a 3xx status code
func (o *GetPolicyBindingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy binding o k response has a 4xx status code
func (o *GetPolicyBindingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policy binding o k response has a 5xx status code
func (o *GetPolicyBindingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy binding o k response a status code equal to that given
func (o *GetPolicyBindingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get policy binding o k response
func (o *GetPolicyBindingOK) Code() int {
	return 200
}

func (o *GetPolicyBindingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBindingOK %s", 200, payload)
}

func (o *GetPolicyBindingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBindingOK %s", 200, payload)
}

func (o *GetPolicyBindingOK) GetPayload() *models.PolicyBinding {
	return o.Payload
}

func (o *GetPolicyBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyBindingUnauthorized creates a GetPolicyBindingUnauthorized with default headers values
func NewGetPolicyBindingUnauthorized() *GetPolicyBindingUnauthorized {
	return &GetPolicyBindingUnauthorized{}
}

/*
GetPolicyBindingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetPolicyBindingUnauthorized struct {
}

// IsSuccess returns true when this get policy binding unauthorized response has a 2xx status code
func (o *GetPolicyBindingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy binding unauthorized response has a 3xx status code
func (o *GetPolicyBindingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy binding unauthorized response has a 4xx status code
func (o *GetPolicyBindingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy binding unauthorized response has a 5xx status code
func (o *GetPolicyBindingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy binding unauthorized response a status code equal to that given
func (o *GetPolicyBindingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get policy binding unauthorized response
func (o *GetPolicyBindingUnauthorized) Code() int {
	return 401
}

func (o *GetPolicyBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBindingUnauthorized", 401)
}

func (o *GetPolicyBindingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBindingUnauthorized", 401)
}

func (o *GetPolicyBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyBindingForbidden creates a GetPolicyBindingForbidden with default headers values
func NewGetPolicyBindingForbidden() *GetPolicyBindingForbidden {
	return &GetPolicyBindingForbidden{}
}

/*
GetPolicyBindingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetPolicyBindingForbidden struct {
}

// IsSuccess returns true when this get policy binding forbidden response has a 2xx status code
func (o *GetPolicyBindingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy binding forbidden response has a 3xx status code
func (o *GetPolicyBindingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy binding forbidden response has a 4xx status code
func (o *GetPolicyBindingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy binding forbidden response has a 5xx status code
func (o *GetPolicyBindingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy binding forbidden response a status code equal to that given
func (o *GetPolicyBindingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get policy binding forbidden response
func (o *GetPolicyBindingForbidden) Code() int {
	return 403
}

func (o *GetPolicyBindingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBindingForbidden", 403)
}

func (o *GetPolicyBindingForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBindingForbidden", 403)
}

func (o *GetPolicyBindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyBindingDefault creates a GetPolicyBindingDefault with default headers values
func NewGetPolicyBindingDefault(code int) *GetPolicyBindingDefault {
	return &GetPolicyBindingDefault{
		_statusCode: code,
	}
}

/*
GetPolicyBindingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetPolicyBindingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get policy binding default response has a 2xx status code
func (o *GetPolicyBindingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get policy binding default response has a 3xx status code
func (o *GetPolicyBindingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get policy binding default response has a 4xx status code
func (o *GetPolicyBindingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get policy binding default response has a 5xx status code
func (o *GetPolicyBindingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get policy binding default response a status code equal to that given
func (o *GetPolicyBindingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get policy binding default response
func (o *GetPolicyBindingDefault) Code() int {
	return o._statusCode
}

func (o *GetPolicyBindingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBinding default %s", o._statusCode, payload)
}

func (o *GetPolicyBindingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/policybinding/{binding_name}][%d] getPolicyBinding default %s", o._statusCode, payload)
}

func (o *GetPolicyBindingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPolicyBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
