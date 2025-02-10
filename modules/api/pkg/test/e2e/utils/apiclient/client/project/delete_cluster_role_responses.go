// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteClusterRoleReader is a Reader for the DeleteClusterRole structure.
type DeleteClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterRoleOK creates a DeleteClusterRoleOK with default headers values
func NewDeleteClusterRoleOK() *DeleteClusterRoleOK {
	return &DeleteClusterRoleOK{}
}

/*
DeleteClusterRoleOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterRoleOK struct {
}

// IsSuccess returns true when this delete cluster role o k response has a 2xx status code
func (o *DeleteClusterRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cluster role o k response has a 3xx status code
func (o *DeleteClusterRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster role o k response has a 4xx status code
func (o *DeleteClusterRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cluster role o k response has a 5xx status code
func (o *DeleteClusterRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster role o k response a status code equal to that given
func (o *DeleteClusterRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteClusterRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleOK ", 200)
}

func (o *DeleteClusterRoleOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleOK ", 200)
}

func (o *DeleteClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterRoleUnauthorized creates a DeleteClusterRoleUnauthorized with default headers values
func NewDeleteClusterRoleUnauthorized() *DeleteClusterRoleUnauthorized {
	return &DeleteClusterRoleUnauthorized{}
}

/*
DeleteClusterRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterRoleUnauthorized struct {
}

// IsSuccess returns true when this delete cluster role unauthorized response has a 2xx status code
func (o *DeleteClusterRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cluster role unauthorized response has a 3xx status code
func (o *DeleteClusterRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster role unauthorized response has a 4xx status code
func (o *DeleteClusterRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cluster role unauthorized response has a 5xx status code
func (o *DeleteClusterRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster role unauthorized response a status code equal to that given
func (o *DeleteClusterRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleUnauthorized ", 401)
}

func (o *DeleteClusterRoleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleUnauthorized ", 401)
}

func (o *DeleteClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterRoleForbidden creates a DeleteClusterRoleForbidden with default headers values
func NewDeleteClusterRoleForbidden() *DeleteClusterRoleForbidden {
	return &DeleteClusterRoleForbidden{}
}

/*
DeleteClusterRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterRoleForbidden struct {
}

// IsSuccess returns true when this delete cluster role forbidden response has a 2xx status code
func (o *DeleteClusterRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cluster role forbidden response has a 3xx status code
func (o *DeleteClusterRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster role forbidden response has a 4xx status code
func (o *DeleteClusterRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cluster role forbidden response has a 5xx status code
func (o *DeleteClusterRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster role forbidden response a status code equal to that given
func (o *DeleteClusterRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleForbidden ", 403)
}

func (o *DeleteClusterRoleForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleForbidden ", 403)
}

func (o *DeleteClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterRoleDefault creates a DeleteClusterRoleDefault with default headers values
func NewDeleteClusterRoleDefault(code int) *DeleteClusterRoleDefault {
	return &DeleteClusterRoleDefault{
		_statusCode: code,
	}
}

/*
DeleteClusterRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete cluster role default response
func (o *DeleteClusterRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete cluster role default response has a 2xx status code
func (o *DeleteClusterRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete cluster role default response has a 3xx status code
func (o *DeleteClusterRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete cluster role default response has a 4xx status code
func (o *DeleteClusterRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete cluster role default response has a 5xx status code
func (o *DeleteClusterRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete cluster role default response a status code equal to that given
func (o *DeleteClusterRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteClusterRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterRoleDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
