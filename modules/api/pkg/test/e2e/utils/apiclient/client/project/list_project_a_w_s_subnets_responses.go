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

// ListProjectAWSSubnetsReader is a Reader for the ListProjectAWSSubnets structure.
type ListProjectAWSSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectAWSSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectAWSSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectAWSSubnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectAWSSubnetsOK creates a ListProjectAWSSubnetsOK with default headers values
func NewListProjectAWSSubnetsOK() *ListProjectAWSSubnetsOK {
	return &ListProjectAWSSubnetsOK{}
}

/*
ListProjectAWSSubnetsOK describes a response with status code 200, with default header values.

AWSSubnetList
*/
type ListProjectAWSSubnetsOK struct {
	Payload models.AWSSubnetList
}

// IsSuccess returns true when this list project a w s subnets o k response has a 2xx status code
func (o *ListProjectAWSSubnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project a w s subnets o k response has a 3xx status code
func (o *ListProjectAWSSubnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project a w s subnets o k response has a 4xx status code
func (o *ListProjectAWSSubnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project a w s subnets o k response has a 5xx status code
func (o *ListProjectAWSSubnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project a w s subnets o k response a status code equal to that given
func (o *ListProjectAWSSubnetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project a w s subnets o k response
func (o *ListProjectAWSSubnetsOK) Code() int {
	return 200
}

func (o *ListProjectAWSSubnetsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/subnets][%d] listProjectAWSSubnetsOK %s", 200, payload)
}

func (o *ListProjectAWSSubnetsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/subnets][%d] listProjectAWSSubnetsOK %s", 200, payload)
}

func (o *ListProjectAWSSubnetsOK) GetPayload() models.AWSSubnetList {
	return o.Payload
}

func (o *ListProjectAWSSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectAWSSubnetsDefault creates a ListProjectAWSSubnetsDefault with default headers values
func NewListProjectAWSSubnetsDefault(code int) *ListProjectAWSSubnetsDefault {
	return &ListProjectAWSSubnetsDefault{
		_statusCode: code,
	}
}

/*
ListProjectAWSSubnetsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectAWSSubnetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project a w s subnets default response has a 2xx status code
func (o *ListProjectAWSSubnetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project a w s subnets default response has a 3xx status code
func (o *ListProjectAWSSubnetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project a w s subnets default response has a 4xx status code
func (o *ListProjectAWSSubnetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project a w s subnets default response has a 5xx status code
func (o *ListProjectAWSSubnetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project a w s subnets default response a status code equal to that given
func (o *ListProjectAWSSubnetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project a w s subnets default response
func (o *ListProjectAWSSubnetsDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectAWSSubnetsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/subnets][%d] listProjectAWSSubnets default %s", o._statusCode, payload)
}

func (o *ListProjectAWSSubnetsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/subnets][%d] listProjectAWSSubnets default %s", o._statusCode, payload)
}

func (o *ListProjectAWSSubnetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectAWSSubnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
