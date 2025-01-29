// Code generated by go-swagger; DO NOT EDIT.

package eks

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

// ValidateProjectEKSCredentialsReader is a Reader for the ValidateProjectEKSCredentials structure.
type ValidateProjectEKSCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateProjectEKSCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateProjectEKSCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateProjectEKSCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateProjectEKSCredentialsOK creates a ValidateProjectEKSCredentialsOK with default headers values
func NewValidateProjectEKSCredentialsOK() *ValidateProjectEKSCredentialsOK {
	return &ValidateProjectEKSCredentialsOK{}
}

/*
ValidateProjectEKSCredentialsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type ValidateProjectEKSCredentialsOK struct {
}

// IsSuccess returns true when this validate project e k s credentials o k response has a 2xx status code
func (o *ValidateProjectEKSCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate project e k s credentials o k response has a 3xx status code
func (o *ValidateProjectEKSCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate project e k s credentials o k response has a 4xx status code
func (o *ValidateProjectEKSCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate project e k s credentials o k response has a 5xx status code
func (o *ValidateProjectEKSCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate project e k s credentials o k response a status code equal to that given
func (o *ValidateProjectEKSCredentialsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate project e k s credentials o k response
func (o *ValidateProjectEKSCredentialsOK) Code() int {
	return 200
}

func (o *ValidateProjectEKSCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/validatecredentials][%d] validateProjectEKSCredentialsOK", 200)
}

func (o *ValidateProjectEKSCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/validatecredentials][%d] validateProjectEKSCredentialsOK", 200)
}

func (o *ValidateProjectEKSCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateProjectEKSCredentialsDefault creates a ValidateProjectEKSCredentialsDefault with default headers values
func NewValidateProjectEKSCredentialsDefault(code int) *ValidateProjectEKSCredentialsDefault {
	return &ValidateProjectEKSCredentialsDefault{
		_statusCode: code,
	}
}

/*
ValidateProjectEKSCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ValidateProjectEKSCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this validate project e k s credentials default response has a 2xx status code
func (o *ValidateProjectEKSCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate project e k s credentials default response has a 3xx status code
func (o *ValidateProjectEKSCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate project e k s credentials default response has a 4xx status code
func (o *ValidateProjectEKSCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate project e k s credentials default response has a 5xx status code
func (o *ValidateProjectEKSCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate project e k s credentials default response a status code equal to that given
func (o *ValidateProjectEKSCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the validate project e k s credentials default response
func (o *ValidateProjectEKSCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ValidateProjectEKSCredentialsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/validatecredentials][%d] validateProjectEKSCredentials default %s", o._statusCode, payload)
}

func (o *ValidateProjectEKSCredentialsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/validatecredentials][%d] validateProjectEKSCredentials default %s", o._statusCode, payload)
}

func (o *ValidateProjectEKSCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ValidateProjectEKSCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
