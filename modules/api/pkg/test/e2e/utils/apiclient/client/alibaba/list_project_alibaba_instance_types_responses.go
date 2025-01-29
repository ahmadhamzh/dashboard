// Code generated by go-swagger; DO NOT EDIT.

package alibaba

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

// ListProjectAlibabaInstanceTypesReader is a Reader for the ListProjectAlibabaInstanceTypes structure.
type ListProjectAlibabaInstanceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectAlibabaInstanceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectAlibabaInstanceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectAlibabaInstanceTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectAlibabaInstanceTypesOK creates a ListProjectAlibabaInstanceTypesOK with default headers values
func NewListProjectAlibabaInstanceTypesOK() *ListProjectAlibabaInstanceTypesOK {
	return &ListProjectAlibabaInstanceTypesOK{}
}

/*
ListProjectAlibabaInstanceTypesOK describes a response with status code 200, with default header values.

AlibabaInstanceTypeList
*/
type ListProjectAlibabaInstanceTypesOK struct {
	Payload models.AlibabaInstanceTypeList
}

// IsSuccess returns true when this list project alibaba instance types o k response has a 2xx status code
func (o *ListProjectAlibabaInstanceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project alibaba instance types o k response has a 3xx status code
func (o *ListProjectAlibabaInstanceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project alibaba instance types o k response has a 4xx status code
func (o *ListProjectAlibabaInstanceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project alibaba instance types o k response has a 5xx status code
func (o *ListProjectAlibabaInstanceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project alibaba instance types o k response a status code equal to that given
func (o *ListProjectAlibabaInstanceTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project alibaba instance types o k response
func (o *ListProjectAlibabaInstanceTypesOK) Code() int {
	return 200
}

func (o *ListProjectAlibabaInstanceTypesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/alibaba/instancetypes][%d] listProjectAlibabaInstanceTypesOK %s", 200, payload)
}

func (o *ListProjectAlibabaInstanceTypesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/alibaba/instancetypes][%d] listProjectAlibabaInstanceTypesOK %s", 200, payload)
}

func (o *ListProjectAlibabaInstanceTypesOK) GetPayload() models.AlibabaInstanceTypeList {
	return o.Payload
}

func (o *ListProjectAlibabaInstanceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectAlibabaInstanceTypesDefault creates a ListProjectAlibabaInstanceTypesDefault with default headers values
func NewListProjectAlibabaInstanceTypesDefault(code int) *ListProjectAlibabaInstanceTypesDefault {
	return &ListProjectAlibabaInstanceTypesDefault{
		_statusCode: code,
	}
}

/*
ListProjectAlibabaInstanceTypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectAlibabaInstanceTypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project alibaba instance types default response has a 2xx status code
func (o *ListProjectAlibabaInstanceTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project alibaba instance types default response has a 3xx status code
func (o *ListProjectAlibabaInstanceTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project alibaba instance types default response has a 4xx status code
func (o *ListProjectAlibabaInstanceTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project alibaba instance types default response has a 5xx status code
func (o *ListProjectAlibabaInstanceTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project alibaba instance types default response a status code equal to that given
func (o *ListProjectAlibabaInstanceTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project alibaba instance types default response
func (o *ListProjectAlibabaInstanceTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectAlibabaInstanceTypesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/alibaba/instancetypes][%d] listProjectAlibabaInstanceTypes default %s", o._statusCode, payload)
}

func (o *ListProjectAlibabaInstanceTypesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/alibaba/instancetypes][%d] listProjectAlibabaInstanceTypes default %s", o._statusCode, payload)
}

func (o *ListProjectAlibabaInstanceTypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectAlibabaInstanceTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
