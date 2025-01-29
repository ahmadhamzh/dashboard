// Code generated by go-swagger; DO NOT EDIT.

package nutanix

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

// ListProjectNutanixCategoriesReader is a Reader for the ListProjectNutanixCategories structure.
type ListProjectNutanixCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectNutanixCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectNutanixCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectNutanixCategoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectNutanixCategoriesOK creates a ListProjectNutanixCategoriesOK with default headers values
func NewListProjectNutanixCategoriesOK() *ListProjectNutanixCategoriesOK {
	return &ListProjectNutanixCategoriesOK{}
}

/*
ListProjectNutanixCategoriesOK describes a response with status code 200, with default header values.

NutanixCategoryList
*/
type ListProjectNutanixCategoriesOK struct {
	Payload models.NutanixCategoryList
}

// IsSuccess returns true when this list project nutanix categories o k response has a 2xx status code
func (o *ListProjectNutanixCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project nutanix categories o k response has a 3xx status code
func (o *ListProjectNutanixCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project nutanix categories o k response has a 4xx status code
func (o *ListProjectNutanixCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project nutanix categories o k response has a 5xx status code
func (o *ListProjectNutanixCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project nutanix categories o k response a status code equal to that given
func (o *ListProjectNutanixCategoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project nutanix categories o k response
func (o *ListProjectNutanixCategoriesOK) Code() int {
	return 200
}

func (o *ListProjectNutanixCategoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/nutanix/{dc}/categories][%d] listProjectNutanixCategoriesOK %s", 200, payload)
}

func (o *ListProjectNutanixCategoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/nutanix/{dc}/categories][%d] listProjectNutanixCategoriesOK %s", 200, payload)
}

func (o *ListProjectNutanixCategoriesOK) GetPayload() models.NutanixCategoryList {
	return o.Payload
}

func (o *ListProjectNutanixCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectNutanixCategoriesDefault creates a ListProjectNutanixCategoriesDefault with default headers values
func NewListProjectNutanixCategoriesDefault(code int) *ListProjectNutanixCategoriesDefault {
	return &ListProjectNutanixCategoriesDefault{
		_statusCode: code,
	}
}

/*
ListProjectNutanixCategoriesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectNutanixCategoriesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project nutanix categories default response has a 2xx status code
func (o *ListProjectNutanixCategoriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project nutanix categories default response has a 3xx status code
func (o *ListProjectNutanixCategoriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project nutanix categories default response has a 4xx status code
func (o *ListProjectNutanixCategoriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project nutanix categories default response has a 5xx status code
func (o *ListProjectNutanixCategoriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project nutanix categories default response a status code equal to that given
func (o *ListProjectNutanixCategoriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project nutanix categories default response
func (o *ListProjectNutanixCategoriesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectNutanixCategoriesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/nutanix/{dc}/categories][%d] listProjectNutanixCategories default %s", o._statusCode, payload)
}

func (o *ListProjectNutanixCategoriesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/nutanix/{dc}/categories][%d] listProjectNutanixCategories default %s", o._statusCode, payload)
}

func (o *ListProjectNutanixCategoriesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectNutanixCategoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
