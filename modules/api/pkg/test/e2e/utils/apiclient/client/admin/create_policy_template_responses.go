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

// CreatePolicyTemplateReader is a Reader for the CreatePolicyTemplate structure.
type CreatePolicyTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePolicyTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePolicyTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePolicyTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePolicyTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePolicyTemplateOK creates a CreatePolicyTemplateOK with default headers values
func NewCreatePolicyTemplateOK() *CreatePolicyTemplateOK {
	return &CreatePolicyTemplateOK{}
}

/*
CreatePolicyTemplateOK describes a response with status code 200, with default header values.

PolicyTemplate
*/
type CreatePolicyTemplateOK struct {
	Payload *models.PolicyTemplate
}

// IsSuccess returns true when this create policy template o k response has a 2xx status code
func (o *CreatePolicyTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create policy template o k response has a 3xx status code
func (o *CreatePolicyTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy template o k response has a 4xx status code
func (o *CreatePolicyTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create policy template o k response has a 5xx status code
func (o *CreatePolicyTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy template o k response a status code equal to that given
func (o *CreatePolicyTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create policy template o k response
func (o *CreatePolicyTemplateOK) Code() int {
	return 200
}

func (o *CreatePolicyTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplateOK %s", 200, payload)
}

func (o *CreatePolicyTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplateOK %s", 200, payload)
}

func (o *CreatePolicyTemplateOK) GetPayload() *models.PolicyTemplate {
	return o.Payload
}

func (o *CreatePolicyTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyTemplateUnauthorized creates a CreatePolicyTemplateUnauthorized with default headers values
func NewCreatePolicyTemplateUnauthorized() *CreatePolicyTemplateUnauthorized {
	return &CreatePolicyTemplateUnauthorized{}
}

/*
CreatePolicyTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreatePolicyTemplateUnauthorized struct {
}

// IsSuccess returns true when this create policy template unauthorized response has a 2xx status code
func (o *CreatePolicyTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy template unauthorized response has a 3xx status code
func (o *CreatePolicyTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy template unauthorized response has a 4xx status code
func (o *CreatePolicyTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy template unauthorized response has a 5xx status code
func (o *CreatePolicyTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy template unauthorized response a status code equal to that given
func (o *CreatePolicyTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create policy template unauthorized response
func (o *CreatePolicyTemplateUnauthorized) Code() int {
	return 401
}

func (o *CreatePolicyTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplateUnauthorized", 401)
}

func (o *CreatePolicyTemplateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplateUnauthorized", 401)
}

func (o *CreatePolicyTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePolicyTemplateForbidden creates a CreatePolicyTemplateForbidden with default headers values
func NewCreatePolicyTemplateForbidden() *CreatePolicyTemplateForbidden {
	return &CreatePolicyTemplateForbidden{}
}

/*
CreatePolicyTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreatePolicyTemplateForbidden struct {
}

// IsSuccess returns true when this create policy template forbidden response has a 2xx status code
func (o *CreatePolicyTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy template forbidden response has a 3xx status code
func (o *CreatePolicyTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy template forbidden response has a 4xx status code
func (o *CreatePolicyTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy template forbidden response has a 5xx status code
func (o *CreatePolicyTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy template forbidden response a status code equal to that given
func (o *CreatePolicyTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create policy template forbidden response
func (o *CreatePolicyTemplateForbidden) Code() int {
	return 403
}

func (o *CreatePolicyTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplateForbidden", 403)
}

func (o *CreatePolicyTemplateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplateForbidden", 403)
}

func (o *CreatePolicyTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePolicyTemplateDefault creates a CreatePolicyTemplateDefault with default headers values
func NewCreatePolicyTemplateDefault(code int) *CreatePolicyTemplateDefault {
	return &CreatePolicyTemplateDefault{
		_statusCode: code,
	}
}

/*
CreatePolicyTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreatePolicyTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create policy template default response has a 2xx status code
func (o *CreatePolicyTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create policy template default response has a 3xx status code
func (o *CreatePolicyTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create policy template default response has a 4xx status code
func (o *CreatePolicyTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create policy template default response has a 5xx status code
func (o *CreatePolicyTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create policy template default response a status code equal to that given
func (o *CreatePolicyTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create policy template default response
func (o *CreatePolicyTemplateDefault) Code() int {
	return o._statusCode
}

func (o *CreatePolicyTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplate default %s", o._statusCode, payload)
}

func (o *CreatePolicyTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/policytemplate][%d] createPolicyTemplate default %s", o._statusCode, payload)
}

func (o *CreatePolicyTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreatePolicyTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
