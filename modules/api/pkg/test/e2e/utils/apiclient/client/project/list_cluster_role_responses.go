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

// ListClusterRoleReader is a Reader for the ListClusterRole structure.
type ListClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterRoleOK creates a ListClusterRoleOK with default headers values
func NewListClusterRoleOK() *ListClusterRoleOK {
	return &ListClusterRoleOK{}
}

/*
ListClusterRoleOK describes a response with status code 200, with default header values.

ClusterRole
*/
type ListClusterRoleOK struct {
	Payload []*models.ClusterRole
}

// IsSuccess returns true when this list cluster role o k response has a 2xx status code
func (o *ListClusterRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cluster role o k response has a 3xx status code
func (o *ListClusterRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role o k response has a 4xx status code
func (o *ListClusterRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cluster role o k response has a 5xx status code
func (o *ListClusterRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role o k response a status code equal to that given
func (o *ListClusterRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list cluster role o k response
func (o *ListClusterRoleOK) Code() int {
	return 200
}

func (o *ListClusterRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleOK %s", 200, payload)
}

func (o *ListClusterRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleOK %s", 200, payload)
}

func (o *ListClusterRoleOK) GetPayload() []*models.ClusterRole {
	return o.Payload
}

func (o *ListClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterRoleUnauthorized creates a ListClusterRoleUnauthorized with default headers values
func NewListClusterRoleUnauthorized() *ListClusterRoleUnauthorized {
	return &ListClusterRoleUnauthorized{}
}

/*
ListClusterRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleUnauthorized struct {
}

// IsSuccess returns true when this list cluster role unauthorized response has a 2xx status code
func (o *ListClusterRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster role unauthorized response has a 3xx status code
func (o *ListClusterRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role unauthorized response has a 4xx status code
func (o *ListClusterRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster role unauthorized response has a 5xx status code
func (o *ListClusterRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role unauthorized response a status code equal to that given
func (o *ListClusterRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list cluster role unauthorized response
func (o *ListClusterRoleUnauthorized) Code() int {
	return 401
}

func (o *ListClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleUnauthorized", 401)
}

func (o *ListClusterRoleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleUnauthorized", 401)
}

func (o *ListClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleForbidden creates a ListClusterRoleForbidden with default headers values
func NewListClusterRoleForbidden() *ListClusterRoleForbidden {
	return &ListClusterRoleForbidden{}
}

/*
ListClusterRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleForbidden struct {
}

// IsSuccess returns true when this list cluster role forbidden response has a 2xx status code
func (o *ListClusterRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster role forbidden response has a 3xx status code
func (o *ListClusterRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role forbidden response has a 4xx status code
func (o *ListClusterRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster role forbidden response has a 5xx status code
func (o *ListClusterRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role forbidden response a status code equal to that given
func (o *ListClusterRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list cluster role forbidden response
func (o *ListClusterRoleForbidden) Code() int {
	return 403
}

func (o *ListClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleForbidden", 403)
}

func (o *ListClusterRoleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleForbidden", 403)
}

func (o *ListClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleDefault creates a ListClusterRoleDefault with default headers values
func NewListClusterRoleDefault(code int) *ListClusterRoleDefault {
	return &ListClusterRoleDefault{
		_statusCode: code,
	}
}

/*
ListClusterRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list cluster role default response has a 2xx status code
func (o *ListClusterRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list cluster role default response has a 3xx status code
func (o *ListClusterRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list cluster role default response has a 4xx status code
func (o *ListClusterRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list cluster role default response has a 5xx status code
func (o *ListClusterRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list cluster role default response a status code equal to that given
func (o *ListClusterRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list cluster role default response
func (o *ListClusterRoleDefault) Code() int {
	return o._statusCode
}

func (o *ListClusterRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRole default %s", o._statusCode, payload)
}

func (o *ListClusterRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRole default %s", o._statusCode, payload)
}

func (o *ListClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
