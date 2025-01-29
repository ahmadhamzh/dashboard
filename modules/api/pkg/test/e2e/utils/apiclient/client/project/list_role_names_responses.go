// Code generated by go-swagger; DO NOT EDIT.

package project

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

// ListRoleNamesReader is a Reader for the ListRoleNames structure.
type ListRoleNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRoleNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRoleNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRoleNamesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRoleNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRoleNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRoleNamesOK creates a ListRoleNamesOK with default headers values
func NewListRoleNamesOK() *ListRoleNamesOK {
	return &ListRoleNamesOK{}
}

/*
ListRoleNamesOK describes a response with status code 200, with default header values.

RoleName
*/
type ListRoleNamesOK struct {
	Payload []*models.RoleName
}

// IsSuccess returns true when this list role names o k response has a 2xx status code
func (o *ListRoleNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list role names o k response has a 3xx status code
func (o *ListRoleNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list role names o k response has a 4xx status code
func (o *ListRoleNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list role names o k response has a 5xx status code
func (o *ListRoleNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list role names o k response a status code equal to that given
func (o *ListRoleNamesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list role names o k response
func (o *ListRoleNamesOK) Code() int {
	return 200
}

func (o *ListRoleNamesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesOK %s", 200, payload)
}

func (o *ListRoleNamesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesOK %s", 200, payload)
}

func (o *ListRoleNamesOK) GetPayload() []*models.RoleName {
	return o.Payload
}

func (o *ListRoleNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRoleNamesUnauthorized creates a ListRoleNamesUnauthorized with default headers values
func NewListRoleNamesUnauthorized() *ListRoleNamesUnauthorized {
	return &ListRoleNamesUnauthorized{}
}

/*
ListRoleNamesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListRoleNamesUnauthorized struct {
}

// IsSuccess returns true when this list role names unauthorized response has a 2xx status code
func (o *ListRoleNamesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list role names unauthorized response has a 3xx status code
func (o *ListRoleNamesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list role names unauthorized response has a 4xx status code
func (o *ListRoleNamesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list role names unauthorized response has a 5xx status code
func (o *ListRoleNamesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list role names unauthorized response a status code equal to that given
func (o *ListRoleNamesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list role names unauthorized response
func (o *ListRoleNamesUnauthorized) Code() int {
	return 401
}

func (o *ListRoleNamesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesUnauthorized", 401)
}

func (o *ListRoleNamesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesUnauthorized", 401)
}

func (o *ListRoleNamesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleNamesForbidden creates a ListRoleNamesForbidden with default headers values
func NewListRoleNamesForbidden() *ListRoleNamesForbidden {
	return &ListRoleNamesForbidden{}
}

/*
ListRoleNamesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListRoleNamesForbidden struct {
}

// IsSuccess returns true when this list role names forbidden response has a 2xx status code
func (o *ListRoleNamesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list role names forbidden response has a 3xx status code
func (o *ListRoleNamesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list role names forbidden response has a 4xx status code
func (o *ListRoleNamesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list role names forbidden response has a 5xx status code
func (o *ListRoleNamesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list role names forbidden response a status code equal to that given
func (o *ListRoleNamesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list role names forbidden response
func (o *ListRoleNamesForbidden) Code() int {
	return 403
}

func (o *ListRoleNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesForbidden", 403)
}

func (o *ListRoleNamesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesForbidden", 403)
}

func (o *ListRoleNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleNamesDefault creates a ListRoleNamesDefault with default headers values
func NewListRoleNamesDefault(code int) *ListRoleNamesDefault {
	return &ListRoleNamesDefault{
		_statusCode: code,
	}
}

/*
ListRoleNamesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListRoleNamesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list role names default response has a 2xx status code
func (o *ListRoleNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list role names default response has a 3xx status code
func (o *ListRoleNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list role names default response has a 4xx status code
func (o *ListRoleNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list role names default response has a 5xx status code
func (o *ListRoleNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list role names default response a status code equal to that given
func (o *ListRoleNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list role names default response
func (o *ListRoleNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListRoleNamesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNames default %s", o._statusCode, payload)
}

func (o *ListRoleNamesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNames default %s", o._statusCode, payload)
}

func (o *ListRoleNamesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListRoleNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
