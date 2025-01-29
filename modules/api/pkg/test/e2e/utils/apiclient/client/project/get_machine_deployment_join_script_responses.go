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

// GetMachineDeploymentJoinScriptReader is a Reader for the GetMachineDeploymentJoinScript structure.
type GetMachineDeploymentJoinScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineDeploymentJoinScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineDeploymentJoinScriptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMachineDeploymentJoinScriptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMachineDeploymentJoinScriptForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMachineDeploymentJoinScriptDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachineDeploymentJoinScriptOK creates a GetMachineDeploymentJoinScriptOK with default headers values
func NewGetMachineDeploymentJoinScriptOK() *GetMachineDeploymentJoinScriptOK {
	return &GetMachineDeploymentJoinScriptOK{}
}

/*
GetMachineDeploymentJoinScriptOK describes a response with status code 200, with default header values.

JoiningScript
*/
type GetMachineDeploymentJoinScriptOK struct {
	Payload models.JoiningScript
}

// IsSuccess returns true when this get machine deployment join script o k response has a 2xx status code
func (o *GetMachineDeploymentJoinScriptOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get machine deployment join script o k response has a 3xx status code
func (o *GetMachineDeploymentJoinScriptOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine deployment join script o k response has a 4xx status code
func (o *GetMachineDeploymentJoinScriptOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get machine deployment join script o k response has a 5xx status code
func (o *GetMachineDeploymentJoinScriptOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine deployment join script o k response a status code equal to that given
func (o *GetMachineDeploymentJoinScriptOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get machine deployment join script o k response
func (o *GetMachineDeploymentJoinScriptOK) Code() int {
	return 200
}

func (o *GetMachineDeploymentJoinScriptOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScriptOK %s", 200, payload)
}

func (o *GetMachineDeploymentJoinScriptOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScriptOK %s", 200, payload)
}

func (o *GetMachineDeploymentJoinScriptOK) GetPayload() models.JoiningScript {
	return o.Payload
}

func (o *GetMachineDeploymentJoinScriptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineDeploymentJoinScriptUnauthorized creates a GetMachineDeploymentJoinScriptUnauthorized with default headers values
func NewGetMachineDeploymentJoinScriptUnauthorized() *GetMachineDeploymentJoinScriptUnauthorized {
	return &GetMachineDeploymentJoinScriptUnauthorized{}
}

/*
GetMachineDeploymentJoinScriptUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetMachineDeploymentJoinScriptUnauthorized struct {
}

// IsSuccess returns true when this get machine deployment join script unauthorized response has a 2xx status code
func (o *GetMachineDeploymentJoinScriptUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get machine deployment join script unauthorized response has a 3xx status code
func (o *GetMachineDeploymentJoinScriptUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine deployment join script unauthorized response has a 4xx status code
func (o *GetMachineDeploymentJoinScriptUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get machine deployment join script unauthorized response has a 5xx status code
func (o *GetMachineDeploymentJoinScriptUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine deployment join script unauthorized response a status code equal to that given
func (o *GetMachineDeploymentJoinScriptUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get machine deployment join script unauthorized response
func (o *GetMachineDeploymentJoinScriptUnauthorized) Code() int {
	return 401
}

func (o *GetMachineDeploymentJoinScriptUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScriptUnauthorized", 401)
}

func (o *GetMachineDeploymentJoinScriptUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScriptUnauthorized", 401)
}

func (o *GetMachineDeploymentJoinScriptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineDeploymentJoinScriptForbidden creates a GetMachineDeploymentJoinScriptForbidden with default headers values
func NewGetMachineDeploymentJoinScriptForbidden() *GetMachineDeploymentJoinScriptForbidden {
	return &GetMachineDeploymentJoinScriptForbidden{}
}

/*
GetMachineDeploymentJoinScriptForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetMachineDeploymentJoinScriptForbidden struct {
}

// IsSuccess returns true when this get machine deployment join script forbidden response has a 2xx status code
func (o *GetMachineDeploymentJoinScriptForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get machine deployment join script forbidden response has a 3xx status code
func (o *GetMachineDeploymentJoinScriptForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine deployment join script forbidden response has a 4xx status code
func (o *GetMachineDeploymentJoinScriptForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get machine deployment join script forbidden response has a 5xx status code
func (o *GetMachineDeploymentJoinScriptForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine deployment join script forbidden response a status code equal to that given
func (o *GetMachineDeploymentJoinScriptForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get machine deployment join script forbidden response
func (o *GetMachineDeploymentJoinScriptForbidden) Code() int {
	return 403
}

func (o *GetMachineDeploymentJoinScriptForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScriptForbidden", 403)
}

func (o *GetMachineDeploymentJoinScriptForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScriptForbidden", 403)
}

func (o *GetMachineDeploymentJoinScriptForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineDeploymentJoinScriptDefault creates a GetMachineDeploymentJoinScriptDefault with default headers values
func NewGetMachineDeploymentJoinScriptDefault(code int) *GetMachineDeploymentJoinScriptDefault {
	return &GetMachineDeploymentJoinScriptDefault{
		_statusCode: code,
	}
}

/*
GetMachineDeploymentJoinScriptDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetMachineDeploymentJoinScriptDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get machine deployment join script default response has a 2xx status code
func (o *GetMachineDeploymentJoinScriptDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get machine deployment join script default response has a 3xx status code
func (o *GetMachineDeploymentJoinScriptDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get machine deployment join script default response has a 4xx status code
func (o *GetMachineDeploymentJoinScriptDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get machine deployment join script default response has a 5xx status code
func (o *GetMachineDeploymentJoinScriptDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get machine deployment join script default response a status code equal to that given
func (o *GetMachineDeploymentJoinScriptDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get machine deployment join script default response
func (o *GetMachineDeploymentJoinScriptDefault) Code() int {
	return o._statusCode
}

func (o *GetMachineDeploymentJoinScriptDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScript default %s", o._statusCode, payload)
}

func (o *GetMachineDeploymentJoinScriptDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/joiningscript][%d] getMachineDeploymentJoinScript default %s", o._statusCode, payload)
}

func (o *GetMachineDeploymentJoinScriptDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetMachineDeploymentJoinScriptDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
