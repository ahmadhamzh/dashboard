// Code generated by go-swagger; DO NOT EDIT.

package preset

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

// DeletePresetReader is a Reader for the DeletePreset structure.
type DeletePresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePresetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePresetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePresetOK creates a DeletePresetOK with default headers values
func NewDeletePresetOK() *DeletePresetOK {
	return &DeletePresetOK{}
}

/*
DeletePresetOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetOK struct {
}

// IsSuccess returns true when this delete preset o k response has a 2xx status code
func (o *DeletePresetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete preset o k response has a 3xx status code
func (o *DeletePresetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset o k response has a 4xx status code
func (o *DeletePresetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete preset o k response has a 5xx status code
func (o *DeletePresetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset o k response a status code equal to that given
func (o *DeletePresetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete preset o k response
func (o *DeletePresetOK) Code() int {
	return 200
}

func (o *DeletePresetOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetOK", 200)
}

func (o *DeletePresetOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetOK", 200)
}

func (o *DeletePresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetUnauthorized creates a DeletePresetUnauthorized with default headers values
func NewDeletePresetUnauthorized() *DeletePresetUnauthorized {
	return &DeletePresetUnauthorized{}
}

/*
DeletePresetUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetUnauthorized struct {
}

// IsSuccess returns true when this delete preset unauthorized response has a 2xx status code
func (o *DeletePresetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete preset unauthorized response has a 3xx status code
func (o *DeletePresetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset unauthorized response has a 4xx status code
func (o *DeletePresetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete preset unauthorized response has a 5xx status code
func (o *DeletePresetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset unauthorized response a status code equal to that given
func (o *DeletePresetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete preset unauthorized response
func (o *DeletePresetUnauthorized) Code() int {
	return 401
}

func (o *DeletePresetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetUnauthorized", 401)
}

func (o *DeletePresetUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetUnauthorized", 401)
}

func (o *DeletePresetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetForbidden creates a DeletePresetForbidden with default headers values
func NewDeletePresetForbidden() *DeletePresetForbidden {
	return &DeletePresetForbidden{}
}

/*
DeletePresetForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetForbidden struct {
}

// IsSuccess returns true when this delete preset forbidden response has a 2xx status code
func (o *DeletePresetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete preset forbidden response has a 3xx status code
func (o *DeletePresetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset forbidden response has a 4xx status code
func (o *DeletePresetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete preset forbidden response has a 5xx status code
func (o *DeletePresetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset forbidden response a status code equal to that given
func (o *DeletePresetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete preset forbidden response
func (o *DeletePresetForbidden) Code() int {
	return 403
}

func (o *DeletePresetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetForbidden", 403)
}

func (o *DeletePresetForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetForbidden", 403)
}

func (o *DeletePresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetNotFound creates a DeletePresetNotFound with default headers values
func NewDeletePresetNotFound() *DeletePresetNotFound {
	return &DeletePresetNotFound{}
}

/*
DeletePresetNotFound describes a response with status code 404, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetNotFound struct {
}

// IsSuccess returns true when this delete preset not found response has a 2xx status code
func (o *DeletePresetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete preset not found response has a 3xx status code
func (o *DeletePresetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset not found response has a 4xx status code
func (o *DeletePresetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete preset not found response has a 5xx status code
func (o *DeletePresetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset not found response a status code equal to that given
func (o *DeletePresetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete preset not found response
func (o *DeletePresetNotFound) Code() int {
	return 404
}

func (o *DeletePresetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetNotFound", 404)
}

func (o *DeletePresetNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePresetNotFound", 404)
}

func (o *DeletePresetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetDefault creates a DeletePresetDefault with default headers values
func NewDeletePresetDefault(code int) *DeletePresetDefault {
	return &DeletePresetDefault{
		_statusCode: code,
	}
}

/*
DeletePresetDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeletePresetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete preset default response has a 2xx status code
func (o *DeletePresetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete preset default response has a 3xx status code
func (o *DeletePresetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete preset default response has a 4xx status code
func (o *DeletePresetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete preset default response has a 5xx status code
func (o *DeletePresetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete preset default response a status code equal to that given
func (o *DeletePresetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete preset default response
func (o *DeletePresetDefault) Code() int {
	return o._statusCode
}

func (o *DeletePresetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePreset default %s", o._statusCode, payload)
}

func (o *DeletePresetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}][%d] deletePreset default %s", o._statusCode, payload)
}

func (o *DeletePresetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
