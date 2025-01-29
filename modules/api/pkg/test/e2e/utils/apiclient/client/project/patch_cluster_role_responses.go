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

// PatchClusterRoleReader is a Reader for the PatchClusterRole structure.
type PatchClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchClusterRoleOK creates a PatchClusterRoleOK with default headers values
func NewPatchClusterRoleOK() *PatchClusterRoleOK {
	return &PatchClusterRoleOK{}
}

/*
PatchClusterRoleOK describes a response with status code 200, with default header values.

ClusterRole
*/
type PatchClusterRoleOK struct {
	Payload *models.ClusterRole
}

// IsSuccess returns true when this patch cluster role o k response has a 2xx status code
func (o *PatchClusterRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch cluster role o k response has a 3xx status code
func (o *PatchClusterRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch cluster role o k response has a 4xx status code
func (o *PatchClusterRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch cluster role o k response has a 5xx status code
func (o *PatchClusterRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch cluster role o k response a status code equal to that given
func (o *PatchClusterRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch cluster role o k response
func (o *PatchClusterRoleOK) Code() int {
	return 200
}

func (o *PatchClusterRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleOK %s", 200, payload)
}

func (o *PatchClusterRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleOK %s", 200, payload)
}

func (o *PatchClusterRoleOK) GetPayload() *models.ClusterRole {
	return o.Payload
}

func (o *PatchClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClusterRoleUnauthorized creates a PatchClusterRoleUnauthorized with default headers values
func NewPatchClusterRoleUnauthorized() *PatchClusterRoleUnauthorized {
	return &PatchClusterRoleUnauthorized{}
}

/*
PatchClusterRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterRoleUnauthorized struct {
}

// IsSuccess returns true when this patch cluster role unauthorized response has a 2xx status code
func (o *PatchClusterRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch cluster role unauthorized response has a 3xx status code
func (o *PatchClusterRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch cluster role unauthorized response has a 4xx status code
func (o *PatchClusterRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch cluster role unauthorized response has a 5xx status code
func (o *PatchClusterRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch cluster role unauthorized response a status code equal to that given
func (o *PatchClusterRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch cluster role unauthorized response
func (o *PatchClusterRoleUnauthorized) Code() int {
	return 401
}

func (o *PatchClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleUnauthorized", 401)
}

func (o *PatchClusterRoleUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleUnauthorized", 401)
}

func (o *PatchClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterRoleForbidden creates a PatchClusterRoleForbidden with default headers values
func NewPatchClusterRoleForbidden() *PatchClusterRoleForbidden {
	return &PatchClusterRoleForbidden{}
}

/*
PatchClusterRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterRoleForbidden struct {
}

// IsSuccess returns true when this patch cluster role forbidden response has a 2xx status code
func (o *PatchClusterRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch cluster role forbidden response has a 3xx status code
func (o *PatchClusterRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch cluster role forbidden response has a 4xx status code
func (o *PatchClusterRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch cluster role forbidden response has a 5xx status code
func (o *PatchClusterRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch cluster role forbidden response a status code equal to that given
func (o *PatchClusterRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch cluster role forbidden response
func (o *PatchClusterRoleForbidden) Code() int {
	return 403
}

func (o *PatchClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleForbidden", 403)
}

func (o *PatchClusterRoleForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleForbidden", 403)
}

func (o *PatchClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterRoleDefault creates a PatchClusterRoleDefault with default headers values
func NewPatchClusterRoleDefault(code int) *PatchClusterRoleDefault {
	return &PatchClusterRoleDefault{
		_statusCode: code,
	}
}

/*
PatchClusterRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this patch cluster role default response has a 2xx status code
func (o *PatchClusterRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch cluster role default response has a 3xx status code
func (o *PatchClusterRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch cluster role default response has a 4xx status code
func (o *PatchClusterRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch cluster role default response has a 5xx status code
func (o *PatchClusterRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch cluster role default response a status code equal to that given
func (o *PatchClusterRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch cluster role default response
func (o *PatchClusterRoleDefault) Code() int {
	return o._statusCode
}

func (o *PatchClusterRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRole default %s", o._statusCode, payload)
}

func (o *PatchClusterRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRole default %s", o._statusCode, payload)
}

func (o *PatchClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
