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

// CreateNodeDeploymentReader is a Reader for the CreateNodeDeployment structure.
type CreateNodeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNodeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNodeDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateNodeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNodeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNodeDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNodeDeploymentCreated creates a CreateNodeDeploymentCreated with default headers values
func NewCreateNodeDeploymentCreated() *CreateNodeDeploymentCreated {
	return &CreateNodeDeploymentCreated{}
}

/*
CreateNodeDeploymentCreated describes a response with status code 201, with default header values.

NodeDeployment
*/
type CreateNodeDeploymentCreated struct {
	Payload *models.NodeDeployment
}

// IsSuccess returns true when this create node deployment created response has a 2xx status code
func (o *CreateNodeDeploymentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create node deployment created response has a 3xx status code
func (o *CreateNodeDeploymentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create node deployment created response has a 4xx status code
func (o *CreateNodeDeploymentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create node deployment created response has a 5xx status code
func (o *CreateNodeDeploymentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create node deployment created response a status code equal to that given
func (o *CreateNodeDeploymentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create node deployment created response
func (o *CreateNodeDeploymentCreated) Code() int {
	return 201
}

func (o *CreateNodeDeploymentCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentCreated %s", 201, payload)
}

func (o *CreateNodeDeploymentCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentCreated %s", 201, payload)
}

func (o *CreateNodeDeploymentCreated) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *CreateNodeDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNodeDeploymentUnauthorized creates a CreateNodeDeploymentUnauthorized with default headers values
func NewCreateNodeDeploymentUnauthorized() *CreateNodeDeploymentUnauthorized {
	return &CreateNodeDeploymentUnauthorized{}
}

/*
CreateNodeDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateNodeDeploymentUnauthorized struct {
}

// IsSuccess returns true when this create node deployment unauthorized response has a 2xx status code
func (o *CreateNodeDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create node deployment unauthorized response has a 3xx status code
func (o *CreateNodeDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create node deployment unauthorized response has a 4xx status code
func (o *CreateNodeDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create node deployment unauthorized response has a 5xx status code
func (o *CreateNodeDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create node deployment unauthorized response a status code equal to that given
func (o *CreateNodeDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create node deployment unauthorized response
func (o *CreateNodeDeploymentUnauthorized) Code() int {
	return 401
}

func (o *CreateNodeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentUnauthorized", 401)
}

func (o *CreateNodeDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentUnauthorized", 401)
}

func (o *CreateNodeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodeDeploymentForbidden creates a CreateNodeDeploymentForbidden with default headers values
func NewCreateNodeDeploymentForbidden() *CreateNodeDeploymentForbidden {
	return &CreateNodeDeploymentForbidden{}
}

/*
CreateNodeDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateNodeDeploymentForbidden struct {
}

// IsSuccess returns true when this create node deployment forbidden response has a 2xx status code
func (o *CreateNodeDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create node deployment forbidden response has a 3xx status code
func (o *CreateNodeDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create node deployment forbidden response has a 4xx status code
func (o *CreateNodeDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create node deployment forbidden response has a 5xx status code
func (o *CreateNodeDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create node deployment forbidden response a status code equal to that given
func (o *CreateNodeDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create node deployment forbidden response
func (o *CreateNodeDeploymentForbidden) Code() int {
	return 403
}

func (o *CreateNodeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentForbidden", 403)
}

func (o *CreateNodeDeploymentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentForbidden", 403)
}

func (o *CreateNodeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodeDeploymentDefault creates a CreateNodeDeploymentDefault with default headers values
func NewCreateNodeDeploymentDefault(code int) *CreateNodeDeploymentDefault {
	return &CreateNodeDeploymentDefault{
		_statusCode: code,
	}
}

/*
CreateNodeDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateNodeDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create node deployment default response has a 2xx status code
func (o *CreateNodeDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create node deployment default response has a 3xx status code
func (o *CreateNodeDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create node deployment default response has a 4xx status code
func (o *CreateNodeDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create node deployment default response has a 5xx status code
func (o *CreateNodeDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create node deployment default response a status code equal to that given
func (o *CreateNodeDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create node deployment default response
func (o *CreateNodeDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *CreateNodeDeploymentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeployment default %s", o._statusCode, payload)
}

func (o *CreateNodeDeploymentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeployment default %s", o._statusCode, payload)
}

func (o *CreateNodeDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateNodeDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
