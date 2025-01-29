// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// PatchApplicationDefinitionReader is a Reader for the PatchApplicationDefinition structure.
type PatchApplicationDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchApplicationDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApplicationDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchApplicationDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchApplicationDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchApplicationDefinitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchApplicationDefinitionOK creates a PatchApplicationDefinitionOK with default headers values
func NewPatchApplicationDefinitionOK() *PatchApplicationDefinitionOK {
	return &PatchApplicationDefinitionOK{}
}

/*
PatchApplicationDefinitionOK describes a response with status code 200, with default header values.

ApplicationDefinition
*/
type PatchApplicationDefinitionOK struct {
	Payload *models.ApplicationDefinition
}

// IsSuccess returns true when this patch application definition o k response has a 2xx status code
func (o *PatchApplicationDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch application definition o k response has a 3xx status code
func (o *PatchApplicationDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch application definition o k response has a 4xx status code
func (o *PatchApplicationDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch application definition o k response has a 5xx status code
func (o *PatchApplicationDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch application definition o k response a status code equal to that given
func (o *PatchApplicationDefinitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch application definition o k response
func (o *PatchApplicationDefinitionOK) Code() int {
	return 200
}

func (o *PatchApplicationDefinitionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinitionOK %s", 200, payload)
}

func (o *PatchApplicationDefinitionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinitionOK %s", 200, payload)
}

func (o *PatchApplicationDefinitionOK) GetPayload() *models.ApplicationDefinition {
	return o.Payload
}

func (o *PatchApplicationDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApplicationDefinitionUnauthorized creates a PatchApplicationDefinitionUnauthorized with default headers values
func NewPatchApplicationDefinitionUnauthorized() *PatchApplicationDefinitionUnauthorized {
	return &PatchApplicationDefinitionUnauthorized{}
}

/*
PatchApplicationDefinitionUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchApplicationDefinitionUnauthorized struct {
}

// IsSuccess returns true when this patch application definition unauthorized response has a 2xx status code
func (o *PatchApplicationDefinitionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch application definition unauthorized response has a 3xx status code
func (o *PatchApplicationDefinitionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch application definition unauthorized response has a 4xx status code
func (o *PatchApplicationDefinitionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch application definition unauthorized response has a 5xx status code
func (o *PatchApplicationDefinitionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch application definition unauthorized response a status code equal to that given
func (o *PatchApplicationDefinitionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch application definition unauthorized response
func (o *PatchApplicationDefinitionUnauthorized) Code() int {
	return 401
}

func (o *PatchApplicationDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinitionUnauthorized", 401)
}

func (o *PatchApplicationDefinitionUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinitionUnauthorized", 401)
}

func (o *PatchApplicationDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchApplicationDefinitionForbidden creates a PatchApplicationDefinitionForbidden with default headers values
func NewPatchApplicationDefinitionForbidden() *PatchApplicationDefinitionForbidden {
	return &PatchApplicationDefinitionForbidden{}
}

/*
PatchApplicationDefinitionForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchApplicationDefinitionForbidden struct {
}

// IsSuccess returns true when this patch application definition forbidden response has a 2xx status code
func (o *PatchApplicationDefinitionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch application definition forbidden response has a 3xx status code
func (o *PatchApplicationDefinitionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch application definition forbidden response has a 4xx status code
func (o *PatchApplicationDefinitionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch application definition forbidden response has a 5xx status code
func (o *PatchApplicationDefinitionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch application definition forbidden response a status code equal to that given
func (o *PatchApplicationDefinitionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch application definition forbidden response
func (o *PatchApplicationDefinitionForbidden) Code() int {
	return 403
}

func (o *PatchApplicationDefinitionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinitionForbidden", 403)
}

func (o *PatchApplicationDefinitionForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinitionForbidden", 403)
}

func (o *PatchApplicationDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchApplicationDefinitionDefault creates a PatchApplicationDefinitionDefault with default headers values
func NewPatchApplicationDefinitionDefault(code int) *PatchApplicationDefinitionDefault {
	return &PatchApplicationDefinitionDefault{
		_statusCode: code,
	}
}

/*
PatchApplicationDefinitionDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchApplicationDefinitionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this patch application definition default response has a 2xx status code
func (o *PatchApplicationDefinitionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch application definition default response has a 3xx status code
func (o *PatchApplicationDefinitionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch application definition default response has a 4xx status code
func (o *PatchApplicationDefinitionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch application definition default response has a 5xx status code
func (o *PatchApplicationDefinitionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch application definition default response a status code equal to that given
func (o *PatchApplicationDefinitionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch application definition default response
func (o *PatchApplicationDefinitionDefault) Code() int {
	return o._statusCode
}

func (o *PatchApplicationDefinitionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinition default %s", o._statusCode, payload)
}

func (o *PatchApplicationDefinitionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/applicationdefinitions/{appdef_name}][%d] patchApplicationDefinition default %s", o._statusCode, payload)
}

func (o *PatchApplicationDefinitionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchApplicationDefinitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
