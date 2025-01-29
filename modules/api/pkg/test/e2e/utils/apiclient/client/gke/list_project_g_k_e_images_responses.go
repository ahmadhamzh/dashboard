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

// ListProjectGKEImagesReader is a Reader for the ListProjectGKEImages structure.
type ListProjectGKEImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectGKEImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectGKEImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectGKEImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectGKEImagesOK creates a ListProjectGKEImagesOK with default headers values
func NewListProjectGKEImagesOK() *ListProjectGKEImagesOK {
	return &ListProjectGKEImagesOK{}
}

/*
ListProjectGKEImagesOK describes a response with status code 200, with default header values.

GKEImageList
*/
type ListProjectGKEImagesOK struct {
	Payload models.GKEImageList
}

// IsSuccess returns true when this list project g k e images o k response has a 2xx status code
func (o *ListProjectGKEImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project g k e images o k response has a 3xx status code
func (o *ListProjectGKEImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project g k e images o k response has a 4xx status code
func (o *ListProjectGKEImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project g k e images o k response has a 5xx status code
func (o *ListProjectGKEImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project g k e images o k response a status code equal to that given
func (o *ListProjectGKEImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project g k e images o k response
func (o *ListProjectGKEImagesOK) Code() int {
	return 200
}

func (o *ListProjectGKEImagesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/images][%d] listProjectGKEImagesOK %s", 200, payload)
}

func (o *ListProjectGKEImagesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/images][%d] listProjectGKEImagesOK %s", 200, payload)
}

func (o *ListProjectGKEImagesOK) GetPayload() models.GKEImageList {
	return o.Payload
}

func (o *ListProjectGKEImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectGKEImagesDefault creates a ListProjectGKEImagesDefault with default headers values
func NewListProjectGKEImagesDefault(code int) *ListProjectGKEImagesDefault {
	return &ListProjectGKEImagesDefault{
		_statusCode: code,
	}
}

/*
ListProjectGKEImagesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectGKEImagesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project g k e images default response has a 2xx status code
func (o *ListProjectGKEImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project g k e images default response has a 3xx status code
func (o *ListProjectGKEImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project g k e images default response has a 4xx status code
func (o *ListProjectGKEImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project g k e images default response has a 5xx status code
func (o *ListProjectGKEImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project g k e images default response a status code equal to that given
func (o *ListProjectGKEImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project g k e images default response
func (o *ListProjectGKEImagesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectGKEImagesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/images][%d] listProjectGKEImages default %s", o._statusCode, payload)
}

func (o *ListProjectGKEImagesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gke/images][%d] listProjectGKEImages default %s", o._statusCode, payload)
}

func (o *ListProjectGKEImagesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectGKEImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
