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

// CreateClusterServiceAccountReader is a Reader for the CreateClusterServiceAccount structure.
type CreateClusterServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterServiceAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateClusterServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateClusterServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateClusterServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterServiceAccountCreated creates a CreateClusterServiceAccountCreated with default headers values
func NewCreateClusterServiceAccountCreated() *CreateClusterServiceAccountCreated {
	return &CreateClusterServiceAccountCreated{}
}

/*
CreateClusterServiceAccountCreated describes a response with status code 201, with default header values.

ClusterServiceAccount
*/
type CreateClusterServiceAccountCreated struct {
	Payload *models.ClusterServiceAccount
}

// IsSuccess returns true when this create cluster service account created response has a 2xx status code
func (o *CreateClusterServiceAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster service account created response has a 3xx status code
func (o *CreateClusterServiceAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster service account created response has a 4xx status code
func (o *CreateClusterServiceAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster service account created response has a 5xx status code
func (o *CreateClusterServiceAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster service account created response a status code equal to that given
func (o *CreateClusterServiceAccountCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create cluster service account created response
func (o *CreateClusterServiceAccountCreated) Code() int {
	return 201
}

func (o *CreateClusterServiceAccountCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccountCreated %s", 201, payload)
}

func (o *CreateClusterServiceAccountCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccountCreated %s", 201, payload)
}

func (o *CreateClusterServiceAccountCreated) GetPayload() *models.ClusterServiceAccount {
	return o.Payload
}

func (o *CreateClusterServiceAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterServiceAccountUnauthorized creates a CreateClusterServiceAccountUnauthorized with default headers values
func NewCreateClusterServiceAccountUnauthorized() *CreateClusterServiceAccountUnauthorized {
	return &CreateClusterServiceAccountUnauthorized{}
}

/*
CreateClusterServiceAccountUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterServiceAccountUnauthorized struct {
}

// IsSuccess returns true when this create cluster service account unauthorized response has a 2xx status code
func (o *CreateClusterServiceAccountUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster service account unauthorized response has a 3xx status code
func (o *CreateClusterServiceAccountUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster service account unauthorized response has a 4xx status code
func (o *CreateClusterServiceAccountUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster service account unauthorized response has a 5xx status code
func (o *CreateClusterServiceAccountUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster service account unauthorized response a status code equal to that given
func (o *CreateClusterServiceAccountUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create cluster service account unauthorized response
func (o *CreateClusterServiceAccountUnauthorized) Code() int {
	return 401
}

func (o *CreateClusterServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccountUnauthorized", 401)
}

func (o *CreateClusterServiceAccountUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccountUnauthorized", 401)
}

func (o *CreateClusterServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterServiceAccountForbidden creates a CreateClusterServiceAccountForbidden with default headers values
func NewCreateClusterServiceAccountForbidden() *CreateClusterServiceAccountForbidden {
	return &CreateClusterServiceAccountForbidden{}
}

/*
CreateClusterServiceAccountForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterServiceAccountForbidden struct {
}

// IsSuccess returns true when this create cluster service account forbidden response has a 2xx status code
func (o *CreateClusterServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster service account forbidden response has a 3xx status code
func (o *CreateClusterServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster service account forbidden response has a 4xx status code
func (o *CreateClusterServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster service account forbidden response has a 5xx status code
func (o *CreateClusterServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster service account forbidden response a status code equal to that given
func (o *CreateClusterServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create cluster service account forbidden response
func (o *CreateClusterServiceAccountForbidden) Code() int {
	return 403
}

func (o *CreateClusterServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccountForbidden", 403)
}

func (o *CreateClusterServiceAccountForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccountForbidden", 403)
}

func (o *CreateClusterServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterServiceAccountDefault creates a CreateClusterServiceAccountDefault with default headers values
func NewCreateClusterServiceAccountDefault(code int) *CreateClusterServiceAccountDefault {
	return &CreateClusterServiceAccountDefault{
		_statusCode: code,
	}
}

/*
CreateClusterServiceAccountDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateClusterServiceAccountDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create cluster service account default response has a 2xx status code
func (o *CreateClusterServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cluster service account default response has a 3xx status code
func (o *CreateClusterServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cluster service account default response has a 4xx status code
func (o *CreateClusterServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cluster service account default response has a 5xx status code
func (o *CreateClusterServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cluster service account default response a status code equal to that given
func (o *CreateClusterServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create cluster service account default response
func (o *CreateClusterServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterServiceAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccount default %s", o._statusCode, payload)
}

func (o *CreateClusterServiceAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] createClusterServiceAccount default %s", o._statusCode, payload)
}

func (o *CreateClusterServiceAccountDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateClusterServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
