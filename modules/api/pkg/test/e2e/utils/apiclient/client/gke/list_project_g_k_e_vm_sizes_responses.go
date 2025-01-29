// Code generated by go-swagger; DO NOT EDIT.

package gke

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

// ListProjectGKEVMSizesReader is a Reader for the ListProjectGKEVMSizes structure.
type ListProjectGKEVMSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectGKEVMSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectGKEVMSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectGKEVMSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectGKEVMSizesOK creates a ListProjectGKEVMSizesOK with default headers values
func NewListProjectGKEVMSizesOK() *ListProjectGKEVMSizesOK {
	return &ListProjectGKEVMSizesOK{}
}

/*
ListProjectGKEVMSizesOK describes a response with status code 200, with default header values.

GCPMachineSizeList
*/
type ListProjectGKEVMSizesOK struct {
	Payload models.GCPMachineSizeList
}

// IsSuccess returns true when this list project g k e Vm sizes o k response has a 2xx status code
func (o *ListProjectGKEVMSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project g k e Vm sizes o k response has a 3xx status code
func (o *ListProjectGKEVMSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project g k e Vm sizes o k response has a 4xx status code
func (o *ListProjectGKEVMSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project g k e Vm sizes o k response has a 5xx status code
func (o *ListProjectGKEVMSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project g k e Vm sizes o k response a status code equal to that given
func (o *ListProjectGKEVMSizesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project g k e Vm sizes o k response
func (o *ListProjectGKEVMSizesOK) Code() int {
	return 200
}

func (o *ListProjectGKEVMSizesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/vmsizes][%d] listProjectGKEVmSizesOK %s", 200, payload)
}

func (o *ListProjectGKEVMSizesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/vmsizes][%d] listProjectGKEVmSizesOK %s", 200, payload)
}

func (o *ListProjectGKEVMSizesOK) GetPayload() models.GCPMachineSizeList {
	return o.Payload
}

func (o *ListProjectGKEVMSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectGKEVMSizesDefault creates a ListProjectGKEVMSizesDefault with default headers values
func NewListProjectGKEVMSizesDefault(code int) *ListProjectGKEVMSizesDefault {
	return &ListProjectGKEVMSizesDefault{
		_statusCode: code,
	}
}

/*
ListProjectGKEVMSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectGKEVMSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project g k e VM sizes default response has a 2xx status code
func (o *ListProjectGKEVMSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project g k e VM sizes default response has a 3xx status code
func (o *ListProjectGKEVMSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project g k e VM sizes default response has a 4xx status code
func (o *ListProjectGKEVMSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project g k e VM sizes default response has a 5xx status code
func (o *ListProjectGKEVMSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project g k e VM sizes default response a status code equal to that given
func (o *ListProjectGKEVMSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project g k e VM sizes default response
func (o *ListProjectGKEVMSizesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectGKEVMSizesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/vmsizes][%d] listProjectGKEVMSizes default %s", o._statusCode, payload)
}

func (o *ListProjectGKEVMSizesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/vmsizes][%d] listProjectGKEVMSizes default %s", o._statusCode, payload)
}

func (o *ListProjectGKEVMSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectGKEVMSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
