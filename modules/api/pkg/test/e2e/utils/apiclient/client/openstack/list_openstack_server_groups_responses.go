// Code generated by go-swagger; DO NOT EDIT.

package openstack

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

// ListOpenstackServerGroupsReader is a Reader for the ListOpenstackServerGroups structure.
type ListOpenstackServerGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackServerGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackServerGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackServerGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackServerGroupsOK creates a ListOpenstackServerGroupsOK with default headers values
func NewListOpenstackServerGroupsOK() *ListOpenstackServerGroupsOK {
	return &ListOpenstackServerGroupsOK{}
}

/*
ListOpenstackServerGroupsOK describes a response with status code 200, with default header values.

OpenstackServerGroup
*/
type ListOpenstackServerGroupsOK struct {
	Payload []*models.OpenstackServerGroup
}

// IsSuccess returns true when this list openstack server groups o k response has a 2xx status code
func (o *ListOpenstackServerGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list openstack server groups o k response has a 3xx status code
func (o *ListOpenstackServerGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list openstack server groups o k response has a 4xx status code
func (o *ListOpenstackServerGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list openstack server groups o k response has a 5xx status code
func (o *ListOpenstackServerGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list openstack server groups o k response a status code equal to that given
func (o *ListOpenstackServerGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list openstack server groups o k response
func (o *ListOpenstackServerGroupsOK) Code() int {
	return 200
}

func (o *ListOpenstackServerGroupsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/providers/openstack/servergroups][%d] listOpenstackServerGroupsOK %s", 200, payload)
}

func (o *ListOpenstackServerGroupsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/providers/openstack/servergroups][%d] listOpenstackServerGroupsOK %s", 200, payload)
}

func (o *ListOpenstackServerGroupsOK) GetPayload() []*models.OpenstackServerGroup {
	return o.Payload
}

func (o *ListOpenstackServerGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackServerGroupsDefault creates a ListOpenstackServerGroupsDefault with default headers values
func NewListOpenstackServerGroupsDefault(code int) *ListOpenstackServerGroupsDefault {
	return &ListOpenstackServerGroupsDefault{
		_statusCode: code,
	}
}

/*
ListOpenstackServerGroupsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackServerGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list openstack server groups default response has a 2xx status code
func (o *ListOpenstackServerGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list openstack server groups default response has a 3xx status code
func (o *ListOpenstackServerGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list openstack server groups default response has a 4xx status code
func (o *ListOpenstackServerGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list openstack server groups default response has a 5xx status code
func (o *ListOpenstackServerGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list openstack server groups default response a status code equal to that given
func (o *ListOpenstackServerGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list openstack server groups default response
func (o *ListOpenstackServerGroupsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackServerGroupsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/providers/openstack/servergroups][%d] listOpenstackServerGroups default %s", o._statusCode, payload)
}

func (o *ListOpenstackServerGroupsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/providers/openstack/servergroups][%d] listOpenstackServerGroups default %s", o._statusCode, payload)
}

func (o *ListOpenstackServerGroupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackServerGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
