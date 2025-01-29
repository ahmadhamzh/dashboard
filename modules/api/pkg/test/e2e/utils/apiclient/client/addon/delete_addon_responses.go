// Code generated by go-swagger; DO NOT EDIT.

package addon

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

// DeleteAddonReader is a Reader for the DeleteAddon structure.
type DeleteAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAddonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAddonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAddonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAddonOK creates a DeleteAddonOK with default headers values
func NewDeleteAddonOK() *DeleteAddonOK {
	return &DeleteAddonOK{}
}

/*
DeleteAddonOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteAddonOK struct {
}

// IsSuccess returns true when this delete addon o k response has a 2xx status code
func (o *DeleteAddonOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete addon o k response has a 3xx status code
func (o *DeleteAddonOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete addon o k response has a 4xx status code
func (o *DeleteAddonOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete addon o k response has a 5xx status code
func (o *DeleteAddonOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete addon o k response a status code equal to that given
func (o *DeleteAddonOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete addon o k response
func (o *DeleteAddonOK) Code() int {
	return 200
}

func (o *DeleteAddonOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonOK", 200)
}

func (o *DeleteAddonOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonOK", 200)
}

func (o *DeleteAddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAddonUnauthorized creates a DeleteAddonUnauthorized with default headers values
func NewDeleteAddonUnauthorized() *DeleteAddonUnauthorized {
	return &DeleteAddonUnauthorized{}
}

/*
DeleteAddonUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteAddonUnauthorized struct {
}

// IsSuccess returns true when this delete addon unauthorized response has a 2xx status code
func (o *DeleteAddonUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete addon unauthorized response has a 3xx status code
func (o *DeleteAddonUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete addon unauthorized response has a 4xx status code
func (o *DeleteAddonUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete addon unauthorized response has a 5xx status code
func (o *DeleteAddonUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete addon unauthorized response a status code equal to that given
func (o *DeleteAddonUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete addon unauthorized response
func (o *DeleteAddonUnauthorized) Code() int {
	return 401
}

func (o *DeleteAddonUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonUnauthorized", 401)
}

func (o *DeleteAddonUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonUnauthorized", 401)
}

func (o *DeleteAddonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAddonForbidden creates a DeleteAddonForbidden with default headers values
func NewDeleteAddonForbidden() *DeleteAddonForbidden {
	return &DeleteAddonForbidden{}
}

/*
DeleteAddonForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteAddonForbidden struct {
}

// IsSuccess returns true when this delete addon forbidden response has a 2xx status code
func (o *DeleteAddonForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete addon forbidden response has a 3xx status code
func (o *DeleteAddonForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete addon forbidden response has a 4xx status code
func (o *DeleteAddonForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete addon forbidden response has a 5xx status code
func (o *DeleteAddonForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete addon forbidden response a status code equal to that given
func (o *DeleteAddonForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete addon forbidden response
func (o *DeleteAddonForbidden) Code() int {
	return 403
}

func (o *DeleteAddonForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonForbidden", 403)
}

func (o *DeleteAddonForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonForbidden", 403)
}

func (o *DeleteAddonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAddonDefault creates a DeleteAddonDefault with default headers values
func NewDeleteAddonDefault(code int) *DeleteAddonDefault {
	return &DeleteAddonDefault{
		_statusCode: code,
	}
}

/*
DeleteAddonDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteAddonDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete addon default response has a 2xx status code
func (o *DeleteAddonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete addon default response has a 3xx status code
func (o *DeleteAddonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete addon default response has a 4xx status code
func (o *DeleteAddonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete addon default response has a 5xx status code
func (o *DeleteAddonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete addon default response a status code equal to that given
func (o *DeleteAddonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete addon default response
func (o *DeleteAddonDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAddonDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddon default %s", o._statusCode, payload)
}

func (o *DeleteAddonDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddon default %s", o._statusCode, payload)
}

func (o *DeleteAddonDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteAddonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
