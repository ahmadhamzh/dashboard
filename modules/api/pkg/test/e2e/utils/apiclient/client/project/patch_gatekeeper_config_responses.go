// Code generated by go-swagger; DO NOT EDIT.

package project

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

// PatchGatekeeperConfigReader is a Reader for the PatchGatekeeperConfig structure.
type PatchGatekeeperConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGatekeeperConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGatekeeperConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchGatekeeperConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchGatekeeperConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchGatekeeperConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGatekeeperConfigOK creates a PatchGatekeeperConfigOK with default headers values
func NewPatchGatekeeperConfigOK() *PatchGatekeeperConfigOK {
	return &PatchGatekeeperConfigOK{}
}

/*
PatchGatekeeperConfigOK describes a response with status code 200, with default header values.

GatekeeperConfig
*/
type PatchGatekeeperConfigOK struct {
	Payload *models.GatekeeperConfig
}

// IsSuccess returns true when this patch gatekeeper config o k response has a 2xx status code
func (o *PatchGatekeeperConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch gatekeeper config o k response has a 3xx status code
func (o *PatchGatekeeperConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gatekeeper config o k response has a 4xx status code
func (o *PatchGatekeeperConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch gatekeeper config o k response has a 5xx status code
func (o *PatchGatekeeperConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gatekeeper config o k response a status code equal to that given
func (o *PatchGatekeeperConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch gatekeeper config o k response
func (o *PatchGatekeeperConfigOK) Code() int {
	return 200
}

func (o *PatchGatekeeperConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigOK %s", 200, payload)
}

func (o *PatchGatekeeperConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigOK %s", 200, payload)
}

func (o *PatchGatekeeperConfigOK) GetPayload() *models.GatekeeperConfig {
	return o.Payload
}

func (o *PatchGatekeeperConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatekeeperConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGatekeeperConfigUnauthorized creates a PatchGatekeeperConfigUnauthorized with default headers values
func NewPatchGatekeeperConfigUnauthorized() *PatchGatekeeperConfigUnauthorized {
	return &PatchGatekeeperConfigUnauthorized{}
}

/*
PatchGatekeeperConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchGatekeeperConfigUnauthorized struct {
}

// IsSuccess returns true when this patch gatekeeper config unauthorized response has a 2xx status code
func (o *PatchGatekeeperConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gatekeeper config unauthorized response has a 3xx status code
func (o *PatchGatekeeperConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gatekeeper config unauthorized response has a 4xx status code
func (o *PatchGatekeeperConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gatekeeper config unauthorized response has a 5xx status code
func (o *PatchGatekeeperConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gatekeeper config unauthorized response a status code equal to that given
func (o *PatchGatekeeperConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch gatekeeper config unauthorized response
func (o *PatchGatekeeperConfigUnauthorized) Code() int {
	return 401
}

func (o *PatchGatekeeperConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigUnauthorized", 401)
}

func (o *PatchGatekeeperConfigUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigUnauthorized", 401)
}

func (o *PatchGatekeeperConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGatekeeperConfigForbidden creates a PatchGatekeeperConfigForbidden with default headers values
func NewPatchGatekeeperConfigForbidden() *PatchGatekeeperConfigForbidden {
	return &PatchGatekeeperConfigForbidden{}
}

/*
PatchGatekeeperConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchGatekeeperConfigForbidden struct {
}

// IsSuccess returns true when this patch gatekeeper config forbidden response has a 2xx status code
func (o *PatchGatekeeperConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gatekeeper config forbidden response has a 3xx status code
func (o *PatchGatekeeperConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gatekeeper config forbidden response has a 4xx status code
func (o *PatchGatekeeperConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gatekeeper config forbidden response has a 5xx status code
func (o *PatchGatekeeperConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gatekeeper config forbidden response a status code equal to that given
func (o *PatchGatekeeperConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch gatekeeper config forbidden response
func (o *PatchGatekeeperConfigForbidden) Code() int {
	return 403
}

func (o *PatchGatekeeperConfigForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigForbidden", 403)
}

func (o *PatchGatekeeperConfigForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigForbidden", 403)
}

func (o *PatchGatekeeperConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGatekeeperConfigDefault creates a PatchGatekeeperConfigDefault with default headers values
func NewPatchGatekeeperConfigDefault(code int) *PatchGatekeeperConfigDefault {
	return &PatchGatekeeperConfigDefault{
		_statusCode: code,
	}
}

/*
PatchGatekeeperConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchGatekeeperConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this patch gatekeeper config default response has a 2xx status code
func (o *PatchGatekeeperConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch gatekeeper config default response has a 3xx status code
func (o *PatchGatekeeperConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch gatekeeper config default response has a 4xx status code
func (o *PatchGatekeeperConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch gatekeeper config default response has a 5xx status code
func (o *PatchGatekeeperConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch gatekeeper config default response a status code equal to that given
func (o *PatchGatekeeperConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch gatekeeper config default response
func (o *PatchGatekeeperConfigDefault) Code() int {
	return o._statusCode
}

func (o *PatchGatekeeperConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfig default %s", o._statusCode, payload)
}

func (o *PatchGatekeeperConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfig default %s", o._statusCode, payload)
}

func (o *PatchGatekeeperConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchGatekeeperConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
