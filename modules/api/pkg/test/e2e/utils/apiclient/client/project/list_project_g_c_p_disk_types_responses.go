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

// ListProjectGCPDiskTypesReader is a Reader for the ListProjectGCPDiskTypes structure.
type ListProjectGCPDiskTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectGCPDiskTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectGCPDiskTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectGCPDiskTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectGCPDiskTypesOK creates a ListProjectGCPDiskTypesOK with default headers values
func NewListProjectGCPDiskTypesOK() *ListProjectGCPDiskTypesOK {
	return &ListProjectGCPDiskTypesOK{}
}

/*
ListProjectGCPDiskTypesOK describes a response with status code 200, with default header values.

GCPDiskTypeList
*/
type ListProjectGCPDiskTypesOK struct {
	Payload models.GCPDiskTypeList
}

// IsSuccess returns true when this list project g c p disk types o k response has a 2xx status code
func (o *ListProjectGCPDiskTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project g c p disk types o k response has a 3xx status code
func (o *ListProjectGCPDiskTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project g c p disk types o k response has a 4xx status code
func (o *ListProjectGCPDiskTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project g c p disk types o k response has a 5xx status code
func (o *ListProjectGCPDiskTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project g c p disk types o k response a status code equal to that given
func (o *ListProjectGCPDiskTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project g c p disk types o k response
func (o *ListProjectGCPDiskTypesOK) Code() int {
	return 200
}

func (o *ListProjectGCPDiskTypesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/disktypes][%d] listProjectGCPDiskTypesOK %s", 200, payload)
}

func (o *ListProjectGCPDiskTypesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/disktypes][%d] listProjectGCPDiskTypesOK %s", 200, payload)
}

func (o *ListProjectGCPDiskTypesOK) GetPayload() models.GCPDiskTypeList {
	return o.Payload
}

func (o *ListProjectGCPDiskTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectGCPDiskTypesDefault creates a ListProjectGCPDiskTypesDefault with default headers values
func NewListProjectGCPDiskTypesDefault(code int) *ListProjectGCPDiskTypesDefault {
	return &ListProjectGCPDiskTypesDefault{
		_statusCode: code,
	}
}

/*
ListProjectGCPDiskTypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectGCPDiskTypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project g c p disk types default response has a 2xx status code
func (o *ListProjectGCPDiskTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project g c p disk types default response has a 3xx status code
func (o *ListProjectGCPDiskTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project g c p disk types default response has a 4xx status code
func (o *ListProjectGCPDiskTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project g c p disk types default response has a 5xx status code
func (o *ListProjectGCPDiskTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project g c p disk types default response a status code equal to that given
func (o *ListProjectGCPDiskTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project g c p disk types default response
func (o *ListProjectGCPDiskTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectGCPDiskTypesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/disktypes][%d] listProjectGCPDiskTypes default %s", o._statusCode, payload)
}

func (o *ListProjectGCPDiskTypesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/disktypes][%d] listProjectGCPDiskTypes default %s", o._statusCode, payload)
}

func (o *ListProjectGCPDiskTypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectGCPDiskTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
