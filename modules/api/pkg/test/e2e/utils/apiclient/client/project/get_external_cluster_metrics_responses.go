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

// GetExternalClusterMetricsReader is a Reader for the GetExternalClusterMetrics structure.
type GetExternalClusterMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalClusterMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalClusterMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExternalClusterMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalClusterMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetExternalClusterMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalClusterMetricsOK creates a GetExternalClusterMetricsOK with default headers values
func NewGetExternalClusterMetricsOK() *GetExternalClusterMetricsOK {
	return &GetExternalClusterMetricsOK{}
}

/*
GetExternalClusterMetricsOK describes a response with status code 200, with default header values.

ClusterMetrics
*/
type GetExternalClusterMetricsOK struct {
	Payload *models.ClusterMetrics
}

// IsSuccess returns true when this get external cluster metrics o k response has a 2xx status code
func (o *GetExternalClusterMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external cluster metrics o k response has a 3xx status code
func (o *GetExternalClusterMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster metrics o k response has a 4xx status code
func (o *GetExternalClusterMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external cluster metrics o k response has a 5xx status code
func (o *GetExternalClusterMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster metrics o k response a status code equal to that given
func (o *GetExternalClusterMetricsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get external cluster metrics o k response
func (o *GetExternalClusterMetricsOK) Code() int {
	return 200
}

func (o *GetExternalClusterMetricsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetricsOK %s", 200, payload)
}

func (o *GetExternalClusterMetricsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetricsOK %s", 200, payload)
}

func (o *GetExternalClusterMetricsOK) GetPayload() *models.ClusterMetrics {
	return o.Payload
}

func (o *GetExternalClusterMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalClusterMetricsUnauthorized creates a GetExternalClusterMetricsUnauthorized with default headers values
func NewGetExternalClusterMetricsUnauthorized() *GetExternalClusterMetricsUnauthorized {
	return &GetExternalClusterMetricsUnauthorized{}
}

/*
GetExternalClusterMetricsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterMetricsUnauthorized struct {
}

// IsSuccess returns true when this get external cluster metrics unauthorized response has a 2xx status code
func (o *GetExternalClusterMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external cluster metrics unauthorized response has a 3xx status code
func (o *GetExternalClusterMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster metrics unauthorized response has a 4xx status code
func (o *GetExternalClusterMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external cluster metrics unauthorized response has a 5xx status code
func (o *GetExternalClusterMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster metrics unauthorized response a status code equal to that given
func (o *GetExternalClusterMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get external cluster metrics unauthorized response
func (o *GetExternalClusterMetricsUnauthorized) Code() int {
	return 401
}

func (o *GetExternalClusterMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetricsUnauthorized", 401)
}

func (o *GetExternalClusterMetricsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetricsUnauthorized", 401)
}

func (o *GetExternalClusterMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterMetricsForbidden creates a GetExternalClusterMetricsForbidden with default headers values
func NewGetExternalClusterMetricsForbidden() *GetExternalClusterMetricsForbidden {
	return &GetExternalClusterMetricsForbidden{}
}

/*
GetExternalClusterMetricsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetExternalClusterMetricsForbidden struct {
}

// IsSuccess returns true when this get external cluster metrics forbidden response has a 2xx status code
func (o *GetExternalClusterMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external cluster metrics forbidden response has a 3xx status code
func (o *GetExternalClusterMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external cluster metrics forbidden response has a 4xx status code
func (o *GetExternalClusterMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external cluster metrics forbidden response has a 5xx status code
func (o *GetExternalClusterMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get external cluster metrics forbidden response a status code equal to that given
func (o *GetExternalClusterMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get external cluster metrics forbidden response
func (o *GetExternalClusterMetricsForbidden) Code() int {
	return 403
}

func (o *GetExternalClusterMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetricsForbidden", 403)
}

func (o *GetExternalClusterMetricsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetricsForbidden", 403)
}

func (o *GetExternalClusterMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExternalClusterMetricsDefault creates a GetExternalClusterMetricsDefault with default headers values
func NewGetExternalClusterMetricsDefault(code int) *GetExternalClusterMetricsDefault {
	return &GetExternalClusterMetricsDefault{
		_statusCode: code,
	}
}

/*
GetExternalClusterMetricsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetExternalClusterMetricsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get external cluster metrics default response has a 2xx status code
func (o *GetExternalClusterMetricsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get external cluster metrics default response has a 3xx status code
func (o *GetExternalClusterMetricsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get external cluster metrics default response has a 4xx status code
func (o *GetExternalClusterMetricsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get external cluster metrics default response has a 5xx status code
func (o *GetExternalClusterMetricsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get external cluster metrics default response a status code equal to that given
func (o *GetExternalClusterMetricsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get external cluster metrics default response
func (o *GetExternalClusterMetricsDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalClusterMetricsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetrics default %s", o._statusCode, payload)
}

func (o *GetExternalClusterMetricsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics][%d] getExternalClusterMetrics default %s", o._statusCode, payload)
}

func (o *GetExternalClusterMetricsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetExternalClusterMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
