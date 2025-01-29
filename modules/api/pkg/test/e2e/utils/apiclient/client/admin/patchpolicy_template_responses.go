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

// PatchpolicyTemplateReader is a Reader for the PatchpolicyTemplate structure.
type PatchpolicyTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchpolicyTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchpolicyTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchpolicyTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchpolicyTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchpolicyTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchpolicyTemplateOK creates a PatchpolicyTemplateOK with default headers values
func NewPatchpolicyTemplateOK() *PatchpolicyTemplateOK {
	return &PatchpolicyTemplateOK{}
}

/*
PatchpolicyTemplateOK describes a response with status code 200, with default header values.

PolicyTemplate
*/
type PatchpolicyTemplateOK struct {
	Payload *models.PolicyTemplate
}

// IsSuccess returns true when this patchpolicy template o k response has a 2xx status code
func (o *PatchpolicyTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patchpolicy template o k response has a 3xx status code
func (o *PatchpolicyTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchpolicy template o k response has a 4xx status code
func (o *PatchpolicyTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patchpolicy template o k response has a 5xx status code
func (o *PatchpolicyTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patchpolicy template o k response a status code equal to that given
func (o *PatchpolicyTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patchpolicy template o k response
func (o *PatchpolicyTemplateOK) Code() int {
	return 200
}

func (o *PatchpolicyTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplateOK %s", 200, payload)
}

func (o *PatchpolicyTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplateOK %s", 200, payload)
}

func (o *PatchpolicyTemplateOK) GetPayload() *models.PolicyTemplate {
	return o.Payload
}

func (o *PatchpolicyTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchpolicyTemplateUnauthorized creates a PatchpolicyTemplateUnauthorized with default headers values
func NewPatchpolicyTemplateUnauthorized() *PatchpolicyTemplateUnauthorized {
	return &PatchpolicyTemplateUnauthorized{}
}

/*
PatchpolicyTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchpolicyTemplateUnauthorized struct {
}

// IsSuccess returns true when this patchpolicy template unauthorized response has a 2xx status code
func (o *PatchpolicyTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patchpolicy template unauthorized response has a 3xx status code
func (o *PatchpolicyTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchpolicy template unauthorized response has a 4xx status code
func (o *PatchpolicyTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patchpolicy template unauthorized response has a 5xx status code
func (o *PatchpolicyTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patchpolicy template unauthorized response a status code equal to that given
func (o *PatchpolicyTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patchpolicy template unauthorized response
func (o *PatchpolicyTemplateUnauthorized) Code() int {
	return 401
}

func (o *PatchpolicyTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplateUnauthorized", 401)
}

func (o *PatchpolicyTemplateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplateUnauthorized", 401)
}

func (o *PatchpolicyTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchpolicyTemplateForbidden creates a PatchpolicyTemplateForbidden with default headers values
func NewPatchpolicyTemplateForbidden() *PatchpolicyTemplateForbidden {
	return &PatchpolicyTemplateForbidden{}
}

/*
PatchpolicyTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchpolicyTemplateForbidden struct {
}

// IsSuccess returns true when this patchpolicy template forbidden response has a 2xx status code
func (o *PatchpolicyTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patchpolicy template forbidden response has a 3xx status code
func (o *PatchpolicyTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchpolicy template forbidden response has a 4xx status code
func (o *PatchpolicyTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patchpolicy template forbidden response has a 5xx status code
func (o *PatchpolicyTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patchpolicy template forbidden response a status code equal to that given
func (o *PatchpolicyTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patchpolicy template forbidden response
func (o *PatchpolicyTemplateForbidden) Code() int {
	return 403
}

func (o *PatchpolicyTemplateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplateForbidden", 403)
}

func (o *PatchpolicyTemplateForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplateForbidden", 403)
}

func (o *PatchpolicyTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchpolicyTemplateDefault creates a PatchpolicyTemplateDefault with default headers values
func NewPatchpolicyTemplateDefault(code int) *PatchpolicyTemplateDefault {
	return &PatchpolicyTemplateDefault{
		_statusCode: code,
	}
}

/*
PatchpolicyTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchpolicyTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this patchpolicy template default response has a 2xx status code
func (o *PatchpolicyTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patchpolicy template default response has a 3xx status code
func (o *PatchpolicyTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patchpolicy template default response has a 4xx status code
func (o *PatchpolicyTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patchpolicy template default response has a 5xx status code
func (o *PatchpolicyTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patchpolicy template default response a status code equal to that given
func (o *PatchpolicyTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patchpolicy template default response
func (o *PatchpolicyTemplateDefault) Code() int {
	return o._statusCode
}

func (o *PatchpolicyTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplate default %s", o._statusCode, payload)
}

func (o *PatchpolicyTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/policytemplate/{template_name}][%d] patchpolicyTemplate default %s", o._statusCode, payload)
}

func (o *PatchpolicyTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchpolicyTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
