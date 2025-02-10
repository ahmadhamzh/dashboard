// Code generated by go-swagger; DO NOT EDIT.

package preset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectPresetsReader is a Reader for the ListProjectPresets structure.
type ListProjectPresetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectPresetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectPresetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListProjectPresetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectPresetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectPresetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectPresetsOK creates a ListProjectPresetsOK with default headers values
func NewListProjectPresetsOK() *ListProjectPresetsOK {
	return &ListProjectPresetsOK{}
}

/*
ListProjectPresetsOK describes a response with status code 200, with default header values.

PresetList
*/
type ListProjectPresetsOK struct {
	Payload *models.PresetList
}

// IsSuccess returns true when this list project presets o k response has a 2xx status code
func (o *ListProjectPresetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project presets o k response has a 3xx status code
func (o *ListProjectPresetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project presets o k response has a 4xx status code
func (o *ListProjectPresetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project presets o k response has a 5xx status code
func (o *ListProjectPresetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project presets o k response a status code equal to that given
func (o *ListProjectPresetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectPresetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresetsOK  %+v", 200, o.Payload)
}

func (o *ListProjectPresetsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresetsOK  %+v", 200, o.Payload)
}

func (o *ListProjectPresetsOK) GetPayload() *models.PresetList {
	return o.Payload
}

func (o *ListProjectPresetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PresetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectPresetsUnauthorized creates a ListProjectPresetsUnauthorized with default headers values
func NewListProjectPresetsUnauthorized() *ListProjectPresetsUnauthorized {
	return &ListProjectPresetsUnauthorized{}
}

/*
ListProjectPresetsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListProjectPresetsUnauthorized struct {
}

// IsSuccess returns true when this list project presets unauthorized response has a 2xx status code
func (o *ListProjectPresetsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project presets unauthorized response has a 3xx status code
func (o *ListProjectPresetsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project presets unauthorized response has a 4xx status code
func (o *ListProjectPresetsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project presets unauthorized response has a 5xx status code
func (o *ListProjectPresetsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project presets unauthorized response a status code equal to that given
func (o *ListProjectPresetsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListProjectPresetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresetsUnauthorized ", 401)
}

func (o *ListProjectPresetsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresetsUnauthorized ", 401)
}

func (o *ListProjectPresetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectPresetsForbidden creates a ListProjectPresetsForbidden with default headers values
func NewListProjectPresetsForbidden() *ListProjectPresetsForbidden {
	return &ListProjectPresetsForbidden{}
}

/*
ListProjectPresetsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListProjectPresetsForbidden struct {
}

// IsSuccess returns true when this list project presets forbidden response has a 2xx status code
func (o *ListProjectPresetsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project presets forbidden response has a 3xx status code
func (o *ListProjectPresetsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project presets forbidden response has a 4xx status code
func (o *ListProjectPresetsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project presets forbidden response has a 5xx status code
func (o *ListProjectPresetsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project presets forbidden response a status code equal to that given
func (o *ListProjectPresetsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListProjectPresetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresetsForbidden ", 403)
}

func (o *ListProjectPresetsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresetsForbidden ", 403)
}

func (o *ListProjectPresetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectPresetsDefault creates a ListProjectPresetsDefault with default headers values
func NewListProjectPresetsDefault(code int) *ListProjectPresetsDefault {
	return &ListProjectPresetsDefault{
		_statusCode: code,
	}
}

/*
ListProjectPresetsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectPresetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project presets default response
func (o *ListProjectPresetsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project presets default response has a 2xx status code
func (o *ListProjectPresetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project presets default response has a 3xx status code
func (o *ListProjectPresetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project presets default response has a 4xx status code
func (o *ListProjectPresetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project presets default response has a 5xx status code
func (o *ListProjectPresetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project presets default response a status code equal to that given
func (o *ListProjectPresetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectPresetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresets default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectPresetsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/presets][%d] listProjectPresets default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectPresetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectPresetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
