// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectEKSVPCsReader is a Reader for the ListProjectEKSVPCs structure.
type ListProjectEKSVPCsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectEKSVPCsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectEKSVPCsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectEKSVPCsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectEKSVPCsOK creates a ListProjectEKSVPCsOK with default headers values
func NewListProjectEKSVPCsOK() *ListProjectEKSVPCsOK {
	return &ListProjectEKSVPCsOK{}
}

/*
ListProjectEKSVPCsOK describes a response with status code 200, with default header values.

EKSVPCList
*/
type ListProjectEKSVPCsOK struct {
	Payload models.EKSVPCList
}

// IsSuccess returns true when this list project e k s v p cs o k response has a 2xx status code
func (o *ListProjectEKSVPCsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project e k s v p cs o k response has a 3xx status code
func (o *ListProjectEKSVPCsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project e k s v p cs o k response has a 4xx status code
func (o *ListProjectEKSVPCsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project e k s v p cs o k response has a 5xx status code
func (o *ListProjectEKSVPCsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project e k s v p cs o k response a status code equal to that given
func (o *ListProjectEKSVPCsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectEKSVPCsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/vpcs][%d] listProjectEKSVPCsOK  %+v", 200, o.Payload)
}

func (o *ListProjectEKSVPCsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/vpcs][%d] listProjectEKSVPCsOK  %+v", 200, o.Payload)
}

func (o *ListProjectEKSVPCsOK) GetPayload() models.EKSVPCList {
	return o.Payload
}

func (o *ListProjectEKSVPCsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectEKSVPCsDefault creates a ListProjectEKSVPCsDefault with default headers values
func NewListProjectEKSVPCsDefault(code int) *ListProjectEKSVPCsDefault {
	return &ListProjectEKSVPCsDefault{
		_statusCode: code,
	}
}

/*
ListProjectEKSVPCsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectEKSVPCsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project e k s v p cs default response
func (o *ListProjectEKSVPCsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project e k s v p cs default response has a 2xx status code
func (o *ListProjectEKSVPCsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project e k s v p cs default response has a 3xx status code
func (o *ListProjectEKSVPCsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project e k s v p cs default response has a 4xx status code
func (o *ListProjectEKSVPCsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project e k s v p cs default response has a 5xx status code
func (o *ListProjectEKSVPCsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project e k s v p cs default response a status code equal to that given
func (o *ListProjectEKSVPCsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectEKSVPCsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/vpcs][%d] listProjectEKSVPCs default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectEKSVPCsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/vpcs][%d] listProjectEKSVPCs default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectEKSVPCsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectEKSVPCsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
