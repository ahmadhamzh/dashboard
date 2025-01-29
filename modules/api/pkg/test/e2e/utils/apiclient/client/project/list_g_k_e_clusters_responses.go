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

// ListGKEClustersReader is a Reader for the ListGKEClusters structure.
type ListGKEClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGKEClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGKEClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGKEClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGKEClustersOK creates a ListGKEClustersOK with default headers values
func NewListGKEClustersOK() *ListGKEClustersOK {
	return &ListGKEClustersOK{}
}

/*
ListGKEClustersOK describes a response with status code 200, with default header values.

GKEClusterList
*/
type ListGKEClustersOK struct {
	Payload models.GKEClusterList
}

// IsSuccess returns true when this list g k e clusters o k response has a 2xx status code
func (o *ListGKEClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g k e clusters o k response has a 3xx status code
func (o *ListGKEClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g k e clusters o k response has a 4xx status code
func (o *ListGKEClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g k e clusters o k response has a 5xx status code
func (o *ListGKEClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g k e clusters o k response a status code equal to that given
func (o *ListGKEClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list g k e clusters o k response
func (o *ListGKEClustersOK) Code() int {
	return 200
}

func (o *ListGKEClustersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/clusters][%d] listGKEClustersOK %s", 200, payload)
}

func (o *ListGKEClustersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/clusters][%d] listGKEClustersOK %s", 200, payload)
}

func (o *ListGKEClustersOK) GetPayload() models.GKEClusterList {
	return o.Payload
}

func (o *ListGKEClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGKEClustersDefault creates a ListGKEClustersDefault with default headers values
func NewListGKEClustersDefault(code int) *ListGKEClustersDefault {
	return &ListGKEClustersDefault{
		_statusCode: code,
	}
}

/*
ListGKEClustersDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGKEClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list g k e clusters default response has a 2xx status code
func (o *ListGKEClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g k e clusters default response has a 3xx status code
func (o *ListGKEClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g k e clusters default response has a 4xx status code
func (o *ListGKEClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g k e clusters default response has a 5xx status code
func (o *ListGKEClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g k e clusters default response a status code equal to that given
func (o *ListGKEClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list g k e clusters default response
func (o *ListGKEClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListGKEClustersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/clusters][%d] listGKEClusters default %s", o._statusCode, payload)
}

func (o *ListGKEClustersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/clusters][%d] listGKEClusters default %s", o._statusCode, payload)
}

func (o *ListGKEClustersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGKEClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
