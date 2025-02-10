// Code generated by go-swagger; DO NOT EDIT.

package nutanix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListNutanixCategoriesReader is a Reader for the ListNutanixCategories structure.
type ListNutanixCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNutanixCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNutanixCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListNutanixCategoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNutanixCategoriesOK creates a ListNutanixCategoriesOK with default headers values
func NewListNutanixCategoriesOK() *ListNutanixCategoriesOK {
	return &ListNutanixCategoriesOK{}
}

/*
ListNutanixCategoriesOK describes a response with status code 200, with default header values.

NutanixCategoryList
*/
type ListNutanixCategoriesOK struct {
	Payload models.NutanixCategoryList
}

// IsSuccess returns true when this list nutanix categories o k response has a 2xx status code
func (o *ListNutanixCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list nutanix categories o k response has a 3xx status code
func (o *ListNutanixCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list nutanix categories o k response has a 4xx status code
func (o *ListNutanixCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list nutanix categories o k response has a 5xx status code
func (o *ListNutanixCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list nutanix categories o k response a status code equal to that given
func (o *ListNutanixCategoriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListNutanixCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/categories][%d] listNutanixCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListNutanixCategoriesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/categories][%d] listNutanixCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListNutanixCategoriesOK) GetPayload() models.NutanixCategoryList {
	return o.Payload
}

func (o *ListNutanixCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNutanixCategoriesDefault creates a ListNutanixCategoriesDefault with default headers values
func NewListNutanixCategoriesDefault(code int) *ListNutanixCategoriesDefault {
	return &ListNutanixCategoriesDefault{
		_statusCode: code,
	}
}

/*
ListNutanixCategoriesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNutanixCategoriesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list nutanix categories default response
func (o *ListNutanixCategoriesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list nutanix categories default response has a 2xx status code
func (o *ListNutanixCategoriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list nutanix categories default response has a 3xx status code
func (o *ListNutanixCategoriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list nutanix categories default response has a 4xx status code
func (o *ListNutanixCategoriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list nutanix categories default response has a 5xx status code
func (o *ListNutanixCategoriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list nutanix categories default response a status code equal to that given
func (o *ListNutanixCategoriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListNutanixCategoriesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/categories][%d] listNutanixCategories default  %+v", o._statusCode, o.Payload)
}

func (o *ListNutanixCategoriesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/nutanix/{dc}/categories][%d] listNutanixCategories default  %+v", o._statusCode, o.Payload)
}

func (o *ListNutanixCategoriesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNutanixCategoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
