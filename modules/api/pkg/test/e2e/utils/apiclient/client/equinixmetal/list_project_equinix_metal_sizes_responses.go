// Code generated by go-swagger; DO NOT EDIT.

package equinixmetal

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

// ListProjectEquinixMetalSizesReader is a Reader for the ListProjectEquinixMetalSizes structure.
type ListProjectEquinixMetalSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectEquinixMetalSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectEquinixMetalSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectEquinixMetalSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectEquinixMetalSizesOK creates a ListProjectEquinixMetalSizesOK with default headers values
func NewListProjectEquinixMetalSizesOK() *ListProjectEquinixMetalSizesOK {
	return &ListProjectEquinixMetalSizesOK{}
}

/*
ListProjectEquinixMetalSizesOK describes a response with status code 200, with default header values.

PacketSizeList
*/
type ListProjectEquinixMetalSizesOK struct {
	Payload []models.PacketSizeList
}

// IsSuccess returns true when this list project equinix metal sizes o k response has a 2xx status code
func (o *ListProjectEquinixMetalSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project equinix metal sizes o k response has a 3xx status code
func (o *ListProjectEquinixMetalSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project equinix metal sizes o k response has a 4xx status code
func (o *ListProjectEquinixMetalSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project equinix metal sizes o k response has a 5xx status code
func (o *ListProjectEquinixMetalSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project equinix metal sizes o k response a status code equal to that given
func (o *ListProjectEquinixMetalSizesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project equinix metal sizes o k response
func (o *ListProjectEquinixMetalSizesOK) Code() int {
	return 200
}

func (o *ListProjectEquinixMetalSizesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/equinixmetal/sizes][%d] listProjectEquinixMetalSizesOK %s", 200, payload)
}

func (o *ListProjectEquinixMetalSizesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/equinixmetal/sizes][%d] listProjectEquinixMetalSizesOK %s", 200, payload)
}

func (o *ListProjectEquinixMetalSizesOK) GetPayload() []models.PacketSizeList {
	return o.Payload
}

func (o *ListProjectEquinixMetalSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectEquinixMetalSizesDefault creates a ListProjectEquinixMetalSizesDefault with default headers values
func NewListProjectEquinixMetalSizesDefault(code int) *ListProjectEquinixMetalSizesDefault {
	return &ListProjectEquinixMetalSizesDefault{
		_statusCode: code,
	}
}

/*
ListProjectEquinixMetalSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectEquinixMetalSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project equinix metal sizes default response has a 2xx status code
func (o *ListProjectEquinixMetalSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project equinix metal sizes default response has a 3xx status code
func (o *ListProjectEquinixMetalSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project equinix metal sizes default response has a 4xx status code
func (o *ListProjectEquinixMetalSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project equinix metal sizes default response has a 5xx status code
func (o *ListProjectEquinixMetalSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project equinix metal sizes default response a status code equal to that given
func (o *ListProjectEquinixMetalSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project equinix metal sizes default response
func (o *ListProjectEquinixMetalSizesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectEquinixMetalSizesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/equinixmetal/sizes][%d] listProjectEquinixMetalSizes default %s", o._statusCode, payload)
}

func (o *ListProjectEquinixMetalSizesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/equinixmetal/sizes][%d] listProjectEquinixMetalSizes default %s", o._statusCode, payload)
}

func (o *ListProjectEquinixMetalSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectEquinixMetalSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
