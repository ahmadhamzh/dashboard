// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeletePolicyTemplateReader is a Reader for the DeletePolicyTemplate structure.
type DeletePolicyTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePolicyTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePolicyTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePolicyTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePolicyTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePolicyTemplateOK creates a DeletePolicyTemplateOK with default headers values
func NewDeletePolicyTemplateOK() *DeletePolicyTemplateOK {
	return &DeletePolicyTemplateOK{}
}

/*
DeletePolicyTemplateOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeletePolicyTemplateOK struct {
}

// IsSuccess returns true when this delete policy template o k response has a 2xx status code
func (o *DeletePolicyTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete policy template o k response has a 3xx status code
func (o *DeletePolicyTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy template o k response has a 4xx status code
func (o *DeletePolicyTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete policy template o k response has a 5xx status code
func (o *DeletePolicyTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy template o k response a status code equal to that given
func (o *DeletePolicyTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeletePolicyTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplateOK ", 200)
}

func (o *DeletePolicyTemplateOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplateOK ", 200)
}

func (o *DeletePolicyTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyTemplateUnauthorized creates a DeletePolicyTemplateUnauthorized with default headers values
func NewDeletePolicyTemplateUnauthorized() *DeletePolicyTemplateUnauthorized {
	return &DeletePolicyTemplateUnauthorized{}
}

/*
DeletePolicyTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeletePolicyTemplateUnauthorized struct {
}

// IsSuccess returns true when this delete policy template unauthorized response has a 2xx status code
func (o *DeletePolicyTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy template unauthorized response has a 3xx status code
func (o *DeletePolicyTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy template unauthorized response has a 4xx status code
func (o *DeletePolicyTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete policy template unauthorized response has a 5xx status code
func (o *DeletePolicyTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy template unauthorized response a status code equal to that given
func (o *DeletePolicyTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeletePolicyTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplateUnauthorized ", 401)
}

func (o *DeletePolicyTemplateUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplateUnauthorized ", 401)
}

func (o *DeletePolicyTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyTemplateForbidden creates a DeletePolicyTemplateForbidden with default headers values
func NewDeletePolicyTemplateForbidden() *DeletePolicyTemplateForbidden {
	return &DeletePolicyTemplateForbidden{}
}

/*
DeletePolicyTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeletePolicyTemplateForbidden struct {
}

// IsSuccess returns true when this delete policy template forbidden response has a 2xx status code
func (o *DeletePolicyTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete policy template forbidden response has a 3xx status code
func (o *DeletePolicyTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy template forbidden response has a 4xx status code
func (o *DeletePolicyTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete policy template forbidden response has a 5xx status code
func (o *DeletePolicyTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy template forbidden response a status code equal to that given
func (o *DeletePolicyTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeletePolicyTemplateForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplateForbidden ", 403)
}

func (o *DeletePolicyTemplateForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplateForbidden ", 403)
}

func (o *DeletePolicyTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyTemplateDefault creates a DeletePolicyTemplateDefault with default headers values
func NewDeletePolicyTemplateDefault(code int) *DeletePolicyTemplateDefault {
	return &DeletePolicyTemplateDefault{
		_statusCode: code,
	}
}

/*
DeletePolicyTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeletePolicyTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete policy template default response
func (o *DeletePolicyTemplateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete policy template default response has a 2xx status code
func (o *DeletePolicyTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete policy template default response has a 3xx status code
func (o *DeletePolicyTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete policy template default response has a 4xx status code
func (o *DeletePolicyTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete policy template default response has a 5xx status code
func (o *DeletePolicyTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete policy template default response a status code equal to that given
func (o *DeletePolicyTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeletePolicyTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePolicyTemplateDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/policytemplates/{template_name}][%d] deletePolicyTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePolicyTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePolicyTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
