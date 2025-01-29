// Code generated by go-swagger; DO NOT EDIT.

package anexia

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

// ListAnexiaTemplatesNoCredentialsV2Reader is a Reader for the ListAnexiaTemplatesNoCredentialsV2 structure.
type ListAnexiaTemplatesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAnexiaTemplatesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAnexiaTemplatesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAnexiaTemplatesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAnexiaTemplatesNoCredentialsV2OK creates a ListAnexiaTemplatesNoCredentialsV2OK with default headers values
func NewListAnexiaTemplatesNoCredentialsV2OK() *ListAnexiaTemplatesNoCredentialsV2OK {
	return &ListAnexiaTemplatesNoCredentialsV2OK{}
}

/*
ListAnexiaTemplatesNoCredentialsV2OK describes a response with status code 200, with default header values.

AnexiaTemplateList
*/
type ListAnexiaTemplatesNoCredentialsV2OK struct {
	Payload models.AnexiaTemplateList
}

// IsSuccess returns true when this list anexia templates no credentials v2 o k response has a 2xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list anexia templates no credentials v2 o k response has a 3xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list anexia templates no credentials v2 o k response has a 4xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list anexia templates no credentials v2 o k response has a 5xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list anexia templates no credentials v2 o k response a status code equal to that given
func (o *ListAnexiaTemplatesNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list anexia templates no credentials v2 o k response
func (o *ListAnexiaTemplatesNoCredentialsV2OK) Code() int {
	return 200
}

func (o *ListAnexiaTemplatesNoCredentialsV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/templates][%d] listAnexiaTemplatesNoCredentialsV2OK %s", 200, payload)
}

func (o *ListAnexiaTemplatesNoCredentialsV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/templates][%d] listAnexiaTemplatesNoCredentialsV2OK %s", 200, payload)
}

func (o *ListAnexiaTemplatesNoCredentialsV2OK) GetPayload() models.AnexiaTemplateList {
	return o.Payload
}

func (o *ListAnexiaTemplatesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAnexiaTemplatesNoCredentialsV2Default creates a ListAnexiaTemplatesNoCredentialsV2Default with default headers values
func NewListAnexiaTemplatesNoCredentialsV2Default(code int) *ListAnexiaTemplatesNoCredentialsV2Default {
	return &ListAnexiaTemplatesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*
ListAnexiaTemplatesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAnexiaTemplatesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list anexia templates no credentials v2 default response has a 2xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list anexia templates no credentials v2 default response has a 3xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list anexia templates no credentials v2 default response has a 4xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list anexia templates no credentials v2 default response has a 5xx status code
func (o *ListAnexiaTemplatesNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list anexia templates no credentials v2 default response a status code equal to that given
func (o *ListAnexiaTemplatesNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list anexia templates no credentials v2 default response
func (o *ListAnexiaTemplatesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAnexiaTemplatesNoCredentialsV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/templates][%d] listAnexiaTemplatesNoCredentialsV2 default %s", o._statusCode, payload)
}

func (o *ListAnexiaTemplatesNoCredentialsV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/templates][%d] listAnexiaTemplatesNoCredentialsV2 default %s", o._statusCode, payload)
}

func (o *ListAnexiaTemplatesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAnexiaTemplatesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
