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

// ListExternalClusterNodesReader is a Reader for the ListExternalClusterNodes structure.
type ListExternalClusterNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterNodesOK creates a ListExternalClusterNodesOK with default headers values
func NewListExternalClusterNodesOK() *ListExternalClusterNodesOK {
	return &ListExternalClusterNodesOK{}
}

/*
ListExternalClusterNodesOK describes a response with status code 200, with default header values.

ExternalClusterNode
*/
type ListExternalClusterNodesOK struct {
	Payload []*models.ExternalClusterNode
}

// IsSuccess returns true when this list external cluster nodes o k response has a 2xx status code
func (o *ListExternalClusterNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list external cluster nodes o k response has a 3xx status code
func (o *ListExternalClusterNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster nodes o k response has a 4xx status code
func (o *ListExternalClusterNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list external cluster nodes o k response has a 5xx status code
func (o *ListExternalClusterNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster nodes o k response a status code equal to that given
func (o *ListExternalClusterNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list external cluster nodes o k response
func (o *ListExternalClusterNodesOK) Code() int {
	return 200
}

func (o *ListExternalClusterNodesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesOK %s", 200, payload)
}

func (o *ListExternalClusterNodesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesOK %s", 200, payload)
}

func (o *ListExternalClusterNodesOK) GetPayload() []*models.ExternalClusterNode {
	return o.Payload
}

func (o *ListExternalClusterNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterNodesUnauthorized creates a ListExternalClusterNodesUnauthorized with default headers values
func NewListExternalClusterNodesUnauthorized() *ListExternalClusterNodesUnauthorized {
	return &ListExternalClusterNodesUnauthorized{}
}

/*
ListExternalClusterNodesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterNodesUnauthorized struct {
}

// IsSuccess returns true when this list external cluster nodes unauthorized response has a 2xx status code
func (o *ListExternalClusterNodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster nodes unauthorized response has a 3xx status code
func (o *ListExternalClusterNodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster nodes unauthorized response has a 4xx status code
func (o *ListExternalClusterNodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster nodes unauthorized response has a 5xx status code
func (o *ListExternalClusterNodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster nodes unauthorized response a status code equal to that given
func (o *ListExternalClusterNodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list external cluster nodes unauthorized response
func (o *ListExternalClusterNodesUnauthorized) Code() int {
	return 401
}

func (o *ListExternalClusterNodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesUnauthorized", 401)
}

func (o *ListExternalClusterNodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesUnauthorized", 401)
}

func (o *ListExternalClusterNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterNodesForbidden creates a ListExternalClusterNodesForbidden with default headers values
func NewListExternalClusterNodesForbidden() *ListExternalClusterNodesForbidden {
	return &ListExternalClusterNodesForbidden{}
}

/*
ListExternalClusterNodesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterNodesForbidden struct {
}

// IsSuccess returns true when this list external cluster nodes forbidden response has a 2xx status code
func (o *ListExternalClusterNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster nodes forbidden response has a 3xx status code
func (o *ListExternalClusterNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster nodes forbidden response has a 4xx status code
func (o *ListExternalClusterNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster nodes forbidden response has a 5xx status code
func (o *ListExternalClusterNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster nodes forbidden response a status code equal to that given
func (o *ListExternalClusterNodesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list external cluster nodes forbidden response
func (o *ListExternalClusterNodesForbidden) Code() int {
	return 403
}

func (o *ListExternalClusterNodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesForbidden", 403)
}

func (o *ListExternalClusterNodesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesForbidden", 403)
}

func (o *ListExternalClusterNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterNodesDefault creates a ListExternalClusterNodesDefault with default headers values
func NewListExternalClusterNodesDefault(code int) *ListExternalClusterNodesDefault {
	return &ListExternalClusterNodesDefault{
		_statusCode: code,
	}
}

/*
ListExternalClusterNodesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListExternalClusterNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list external cluster nodes default response has a 2xx status code
func (o *ListExternalClusterNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list external cluster nodes default response has a 3xx status code
func (o *ListExternalClusterNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list external cluster nodes default response has a 4xx status code
func (o *ListExternalClusterNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list external cluster nodes default response has a 5xx status code
func (o *ListExternalClusterNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list external cluster nodes default response a status code equal to that given
func (o *ListExternalClusterNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list external cluster nodes default response
func (o *ListExternalClusterNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListExternalClusterNodesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodes default %s", o._statusCode, payload)
}

func (o *ListExternalClusterNodesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodes default %s", o._statusCode, payload)
}

func (o *ListExternalClusterNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
