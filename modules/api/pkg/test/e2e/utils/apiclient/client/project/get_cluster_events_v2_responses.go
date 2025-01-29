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

// GetClusterEventsV2Reader is a Reader for the GetClusterEventsV2 structure.
type GetClusterEventsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterEventsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterEventsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterEventsV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterEventsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterEventsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterEventsV2OK creates a GetClusterEventsV2OK with default headers values
func NewGetClusterEventsV2OK() *GetClusterEventsV2OK {
	return &GetClusterEventsV2OK{}
}

/*
GetClusterEventsV2OK describes a response with status code 200, with default header values.

Event
*/
type GetClusterEventsV2OK struct {
	Payload []*models.Event
}

// IsSuccess returns true when this get cluster events v2 o k response has a 2xx status code
func (o *GetClusterEventsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster events v2 o k response has a 3xx status code
func (o *GetClusterEventsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster events v2 o k response has a 4xx status code
func (o *GetClusterEventsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster events v2 o k response has a 5xx status code
func (o *GetClusterEventsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster events v2 o k response a status code equal to that given
func (o *GetClusterEventsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster events v2 o k response
func (o *GetClusterEventsV2OK) Code() int {
	return 200
}

func (o *GetClusterEventsV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2OK %s", 200, payload)
}

func (o *GetClusterEventsV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2OK %s", 200, payload)
}

func (o *GetClusterEventsV2OK) GetPayload() []*models.Event {
	return o.Payload
}

func (o *GetClusterEventsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterEventsV2Unauthorized creates a GetClusterEventsV2Unauthorized with default headers values
func NewGetClusterEventsV2Unauthorized() *GetClusterEventsV2Unauthorized {
	return &GetClusterEventsV2Unauthorized{}
}

/*
GetClusterEventsV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterEventsV2Unauthorized struct {
}

// IsSuccess returns true when this get cluster events v2 unauthorized response has a 2xx status code
func (o *GetClusterEventsV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster events v2 unauthorized response has a 3xx status code
func (o *GetClusterEventsV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster events v2 unauthorized response has a 4xx status code
func (o *GetClusterEventsV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster events v2 unauthorized response has a 5xx status code
func (o *GetClusterEventsV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster events v2 unauthorized response a status code equal to that given
func (o *GetClusterEventsV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cluster events v2 unauthorized response
func (o *GetClusterEventsV2Unauthorized) Code() int {
	return 401
}

func (o *GetClusterEventsV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2Unauthorized", 401)
}

func (o *GetClusterEventsV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2Unauthorized", 401)
}

func (o *GetClusterEventsV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterEventsV2Forbidden creates a GetClusterEventsV2Forbidden with default headers values
func NewGetClusterEventsV2Forbidden() *GetClusterEventsV2Forbidden {
	return &GetClusterEventsV2Forbidden{}
}

/*
GetClusterEventsV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterEventsV2Forbidden struct {
}

// IsSuccess returns true when this get cluster events v2 forbidden response has a 2xx status code
func (o *GetClusterEventsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster events v2 forbidden response has a 3xx status code
func (o *GetClusterEventsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster events v2 forbidden response has a 4xx status code
func (o *GetClusterEventsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster events v2 forbidden response has a 5xx status code
func (o *GetClusterEventsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster events v2 forbidden response a status code equal to that given
func (o *GetClusterEventsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cluster events v2 forbidden response
func (o *GetClusterEventsV2Forbidden) Code() int {
	return 403
}

func (o *GetClusterEventsV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2Forbidden", 403)
}

func (o *GetClusterEventsV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2Forbidden", 403)
}

func (o *GetClusterEventsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterEventsV2Default creates a GetClusterEventsV2Default with default headers values
func NewGetClusterEventsV2Default(code int) *GetClusterEventsV2Default {
	return &GetClusterEventsV2Default{
		_statusCode: code,
	}
}

/*
GetClusterEventsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterEventsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get cluster events v2 default response has a 2xx status code
func (o *GetClusterEventsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster events v2 default response has a 3xx status code
func (o *GetClusterEventsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster events v2 default response has a 4xx status code
func (o *GetClusterEventsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster events v2 default response has a 5xx status code
func (o *GetClusterEventsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster events v2 default response a status code equal to that given
func (o *GetClusterEventsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cluster events v2 default response
func (o *GetClusterEventsV2Default) Code() int {
	return o._statusCode
}

func (o *GetClusterEventsV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2 default %s", o._statusCode, payload)
}

func (o *GetClusterEventsV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events][%d] getClusterEventsV2 default %s", o._statusCode, payload)
}

func (o *GetClusterEventsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterEventsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
