// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListKubeVirtPreferencesNoCredentialsReader is a Reader for the ListKubeVirtPreferencesNoCredentials structure.
type ListKubeVirtPreferencesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubeVirtPreferencesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubeVirtPreferencesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubeVirtPreferencesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubeVirtPreferencesNoCredentialsOK creates a ListKubeVirtPreferencesNoCredentialsOK with default headers values
func NewListKubeVirtPreferencesNoCredentialsOK() *ListKubeVirtPreferencesNoCredentialsOK {
	return &ListKubeVirtPreferencesNoCredentialsOK{}
}

/*
ListKubeVirtPreferencesNoCredentialsOK describes a response with status code 200, with default header values.

VirtualMachinePreferenceList
*/
type ListKubeVirtPreferencesNoCredentialsOK struct {
	Payload *models.VirtualMachinePreferenceList
}

// IsSuccess returns true when this list kube virt preferences no credentials o k response has a 2xx status code
func (o *ListKubeVirtPreferencesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list kube virt preferences no credentials o k response has a 3xx status code
func (o *ListKubeVirtPreferencesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list kube virt preferences no credentials o k response has a 4xx status code
func (o *ListKubeVirtPreferencesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list kube virt preferences no credentials o k response has a 5xx status code
func (o *ListKubeVirtPreferencesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list kube virt preferences no credentials o k response a status code equal to that given
func (o *ListKubeVirtPreferencesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListKubeVirtPreferencesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/preferences][%d] listKubeVirtPreferencesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListKubeVirtPreferencesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/preferences][%d] listKubeVirtPreferencesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListKubeVirtPreferencesNoCredentialsOK) GetPayload() *models.VirtualMachinePreferenceList {
	return o.Payload
}

func (o *ListKubeVirtPreferencesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachinePreferenceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubeVirtPreferencesNoCredentialsDefault creates a ListKubeVirtPreferencesNoCredentialsDefault with default headers values
func NewListKubeVirtPreferencesNoCredentialsDefault(code int) *ListKubeVirtPreferencesNoCredentialsDefault {
	return &ListKubeVirtPreferencesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListKubeVirtPreferencesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListKubeVirtPreferencesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list kube virt preferences no credentials default response
func (o *ListKubeVirtPreferencesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list kube virt preferences no credentials default response has a 2xx status code
func (o *ListKubeVirtPreferencesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list kube virt preferences no credentials default response has a 3xx status code
func (o *ListKubeVirtPreferencesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list kube virt preferences no credentials default response has a 4xx status code
func (o *ListKubeVirtPreferencesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list kube virt preferences no credentials default response has a 5xx status code
func (o *ListKubeVirtPreferencesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list kube virt preferences no credentials default response a status code equal to that given
func (o *ListKubeVirtPreferencesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListKubeVirtPreferencesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/preferences][%d] listKubeVirtPreferencesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubeVirtPreferencesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/preferences][%d] listKubeVirtPreferencesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubeVirtPreferencesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListKubeVirtPreferencesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
