// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVMwareCloudDirectorNetworksNoCredentialsReader is a Reader for the ListVMwareCloudDirectorNetworksNoCredentials structure.
type ListVMwareCloudDirectorNetworksNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorNetworksNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorNetworksNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorNetworksNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorNetworksNoCredentialsOK creates a ListVMwareCloudDirectorNetworksNoCredentialsOK with default headers values
func NewListVMwareCloudDirectorNetworksNoCredentialsOK() *ListVMwareCloudDirectorNetworksNoCredentialsOK {
	return &ListVMwareCloudDirectorNetworksNoCredentialsOK{}
}

/*
ListVMwareCloudDirectorNetworksNoCredentialsOK describes a response with status code 200, with default header values.

VMwareCloudDirectorNetworkList
*/
type ListVMwareCloudDirectorNetworksNoCredentialsOK struct {
	Payload models.VMwareCloudDirectorNetworkList
}

// IsSuccess returns true when this list v mware cloud director networks no credentials o k response has a 2xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list v mware cloud director networks no credentials o k response has a 3xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list v mware cloud director networks no credentials o k response has a 4xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list v mware cloud director networks no credentials o k response has a 5xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list v mware cloud director networks no credentials o k response a status code equal to that given
func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/networks][%d] listVMwareCloudDirectorNetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/networks][%d] listVMwareCloudDirectorNetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) GetPayload() models.VMwareCloudDirectorNetworkList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorNetworksNoCredentialsDefault creates a ListVMwareCloudDirectorNetworksNoCredentialsDefault with default headers values
func NewListVMwareCloudDirectorNetworksNoCredentialsDefault(code int) *ListVMwareCloudDirectorNetworksNoCredentialsDefault {
	return &ListVMwareCloudDirectorNetworksNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListVMwareCloudDirectorNetworksNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorNetworksNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director networks no credentials default response
func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list v mware cloud director networks no credentials default response has a 2xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list v mware cloud director networks no credentials default response has a 3xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list v mware cloud director networks no credentials default response has a 4xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list v mware cloud director networks no credentials default response has a 5xx status code
func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list v mware cloud director networks no credentials default response a status code equal to that given
func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/networks][%d] listVMwareCloudDirectorNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vmwareclouddirector/networks][%d] listVMwareCloudDirectorNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorNetworksNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
