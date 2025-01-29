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

// DeleteMachineDeploymentReader is a Reader for the DeleteMachineDeployment structure.
type DeleteMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMachineDeploymentOK creates a DeleteMachineDeploymentOK with default headers values
func NewDeleteMachineDeploymentOK() *DeleteMachineDeploymentOK {
	return &DeleteMachineDeploymentOK{}
}

/*
DeleteMachineDeploymentOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteMachineDeploymentOK struct {
}

// IsSuccess returns true when this delete machine deployment o k response has a 2xx status code
func (o *DeleteMachineDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete machine deployment o k response has a 3xx status code
func (o *DeleteMachineDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete machine deployment o k response has a 4xx status code
func (o *DeleteMachineDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete machine deployment o k response has a 5xx status code
func (o *DeleteMachineDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete machine deployment o k response a status code equal to that given
func (o *DeleteMachineDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete machine deployment o k response
func (o *DeleteMachineDeploymentOK) Code() int {
	return 200
}

func (o *DeleteMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentOK", 200)
}

func (o *DeleteMachineDeploymentOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentOK", 200)
}

func (o *DeleteMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDeploymentUnauthorized creates a DeleteMachineDeploymentUnauthorized with default headers values
func NewDeleteMachineDeploymentUnauthorized() *DeleteMachineDeploymentUnauthorized {
	return &DeleteMachineDeploymentUnauthorized{}
}

/*
DeleteMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteMachineDeploymentUnauthorized struct {
}

// IsSuccess returns true when this delete machine deployment unauthorized response has a 2xx status code
func (o *DeleteMachineDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete machine deployment unauthorized response has a 3xx status code
func (o *DeleteMachineDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete machine deployment unauthorized response has a 4xx status code
func (o *DeleteMachineDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete machine deployment unauthorized response has a 5xx status code
func (o *DeleteMachineDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete machine deployment unauthorized response a status code equal to that given
func (o *DeleteMachineDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete machine deployment unauthorized response
func (o *DeleteMachineDeploymentUnauthorized) Code() int {
	return 401
}

func (o *DeleteMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentUnauthorized", 401)
}

func (o *DeleteMachineDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentUnauthorized", 401)
}

func (o *DeleteMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDeploymentForbidden creates a DeleteMachineDeploymentForbidden with default headers values
func NewDeleteMachineDeploymentForbidden() *DeleteMachineDeploymentForbidden {
	return &DeleteMachineDeploymentForbidden{}
}

/*
DeleteMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteMachineDeploymentForbidden struct {
}

// IsSuccess returns true when this delete machine deployment forbidden response has a 2xx status code
func (o *DeleteMachineDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete machine deployment forbidden response has a 3xx status code
func (o *DeleteMachineDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete machine deployment forbidden response has a 4xx status code
func (o *DeleteMachineDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete machine deployment forbidden response has a 5xx status code
func (o *DeleteMachineDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete machine deployment forbidden response a status code equal to that given
func (o *DeleteMachineDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete machine deployment forbidden response
func (o *DeleteMachineDeploymentForbidden) Code() int {
	return 403
}

func (o *DeleteMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentForbidden", 403)
}

func (o *DeleteMachineDeploymentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentForbidden", 403)
}

func (o *DeleteMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDeploymentDefault creates a DeleteMachineDeploymentDefault with default headers values
func NewDeleteMachineDeploymentDefault(code int) *DeleteMachineDeploymentDefault {
	return &DeleteMachineDeploymentDefault{
		_statusCode: code,
	}
}

/*
DeleteMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete machine deployment default response has a 2xx status code
func (o *DeleteMachineDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete machine deployment default response has a 3xx status code
func (o *DeleteMachineDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete machine deployment default response has a 4xx status code
func (o *DeleteMachineDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete machine deployment default response has a 5xx status code
func (o *DeleteMachineDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete machine deployment default response a status code equal to that given
func (o *DeleteMachineDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete machine deployment default response
func (o *DeleteMachineDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMachineDeploymentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeployment default %s", o._statusCode, payload)
}

func (o *DeleteMachineDeploymentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeployment default %s", o._statusCode, payload)
}

func (o *DeleteMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
