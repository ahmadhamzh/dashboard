// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

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

// ListProjectVMwareCloudDirectorCatalogsReader is a Reader for the ListProjectVMwareCloudDirectorCatalogs structure.
type ListProjectVMwareCloudDirectorCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectVMwareCloudDirectorCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectVMwareCloudDirectorCatalogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectVMwareCloudDirectorCatalogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectVMwareCloudDirectorCatalogsOK creates a ListProjectVMwareCloudDirectorCatalogsOK with default headers values
func NewListProjectVMwareCloudDirectorCatalogsOK() *ListProjectVMwareCloudDirectorCatalogsOK {
	return &ListProjectVMwareCloudDirectorCatalogsOK{}
}

/*
ListProjectVMwareCloudDirectorCatalogsOK describes a response with status code 200, with default header values.

VMwareCloudDirectorCatalogList
*/
type ListProjectVMwareCloudDirectorCatalogsOK struct {
	Payload models.VMwareCloudDirectorCatalogList
}

// IsSuccess returns true when this list project v mware cloud director catalogs o k response has a 2xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project v mware cloud director catalogs o k response has a 3xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project v mware cloud director catalogs o k response has a 4xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project v mware cloud director catalogs o k response has a 5xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project v mware cloud director catalogs o k response a status code equal to that given
func (o *ListProjectVMwareCloudDirectorCatalogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project v mware cloud director catalogs o k response
func (o *ListProjectVMwareCloudDirectorCatalogsOK) Code() int {
	return 200
}

func (o *ListProjectVMwareCloudDirectorCatalogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/catalogs][%d] listProjectVMwareCloudDirectorCatalogsOK %s", 200, payload)
}

func (o *ListProjectVMwareCloudDirectorCatalogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/catalogs][%d] listProjectVMwareCloudDirectorCatalogsOK %s", 200, payload)
}

func (o *ListProjectVMwareCloudDirectorCatalogsOK) GetPayload() models.VMwareCloudDirectorCatalogList {
	return o.Payload
}

func (o *ListProjectVMwareCloudDirectorCatalogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectVMwareCloudDirectorCatalogsDefault creates a ListProjectVMwareCloudDirectorCatalogsDefault with default headers values
func NewListProjectVMwareCloudDirectorCatalogsDefault(code int) *ListProjectVMwareCloudDirectorCatalogsDefault {
	return &ListProjectVMwareCloudDirectorCatalogsDefault{
		_statusCode: code,
	}
}

/*
ListProjectVMwareCloudDirectorCatalogsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectVMwareCloudDirectorCatalogsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project v mware cloud director catalogs default response has a 2xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project v mware cloud director catalogs default response has a 3xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project v mware cloud director catalogs default response has a 4xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project v mware cloud director catalogs default response has a 5xx status code
func (o *ListProjectVMwareCloudDirectorCatalogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project v mware cloud director catalogs default response a status code equal to that given
func (o *ListProjectVMwareCloudDirectorCatalogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project v mware cloud director catalogs default response
func (o *ListProjectVMwareCloudDirectorCatalogsDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectVMwareCloudDirectorCatalogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/catalogs][%d] listProjectVMwareCloudDirectorCatalogs default %s", o._statusCode, payload)
}

func (o *ListProjectVMwareCloudDirectorCatalogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/vmwareclouddirector/{dc}/catalogs][%d] listProjectVMwareCloudDirectorCatalogs default %s", o._statusCode, payload)
}

func (o *ListProjectVMwareCloudDirectorCatalogsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectVMwareCloudDirectorCatalogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
