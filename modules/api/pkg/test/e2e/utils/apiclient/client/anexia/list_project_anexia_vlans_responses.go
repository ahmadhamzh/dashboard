// Code generated by go-swagger; DO NOT EDIT.

package anexia

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

// ListProjectAnexiaVlansReader is a Reader for the ListProjectAnexiaVlans structure.
type ListProjectAnexiaVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectAnexiaVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectAnexiaVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectAnexiaVlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectAnexiaVlansOK creates a ListProjectAnexiaVlansOK with default headers values
func NewListProjectAnexiaVlansOK() *ListProjectAnexiaVlansOK {
	return &ListProjectAnexiaVlansOK{}
}

/*
ListProjectAnexiaVlansOK describes a response with status code 200, with default header values.

AnexiaVlanList
*/
type ListProjectAnexiaVlansOK struct {
	Payload models.AnexiaVlanList
}

// IsSuccess returns true when this list project anexia vlans o k response has a 2xx status code
func (o *ListProjectAnexiaVlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project anexia vlans o k response has a 3xx status code
func (o *ListProjectAnexiaVlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project anexia vlans o k response has a 4xx status code
func (o *ListProjectAnexiaVlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project anexia vlans o k response has a 5xx status code
func (o *ListProjectAnexiaVlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project anexia vlans o k response a status code equal to that given
func (o *ListProjectAnexiaVlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project anexia vlans o k response
func (o *ListProjectAnexiaVlansOK) Code() int {
	return 200
}

func (o *ListProjectAnexiaVlansOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/anexia/vlans][%d] listProjectAnexiaVlansOK %s", 200, payload)
}

func (o *ListProjectAnexiaVlansOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/anexia/vlans][%d] listProjectAnexiaVlansOK %s", 200, payload)
}

func (o *ListProjectAnexiaVlansOK) GetPayload() models.AnexiaVlanList {
	return o.Payload
}

func (o *ListProjectAnexiaVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectAnexiaVlansDefault creates a ListProjectAnexiaVlansDefault with default headers values
func NewListProjectAnexiaVlansDefault(code int) *ListProjectAnexiaVlansDefault {
	return &ListProjectAnexiaVlansDefault{
		_statusCode: code,
	}
}

/*
ListProjectAnexiaVlansDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectAnexiaVlansDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project anexia vlans default response has a 2xx status code
func (o *ListProjectAnexiaVlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project anexia vlans default response has a 3xx status code
func (o *ListProjectAnexiaVlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project anexia vlans default response has a 4xx status code
func (o *ListProjectAnexiaVlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project anexia vlans default response has a 5xx status code
func (o *ListProjectAnexiaVlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project anexia vlans default response a status code equal to that given
func (o *ListProjectAnexiaVlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project anexia vlans default response
func (o *ListProjectAnexiaVlansDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectAnexiaVlansDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/anexia/vlans][%d] listProjectAnexiaVlans default %s", o._statusCode, payload)
}

func (o *ListProjectAnexiaVlansDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/anexia/vlans][%d] listProjectAnexiaVlans default %s", o._statusCode, payload)
}

func (o *ListProjectAnexiaVlansDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectAnexiaVlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
