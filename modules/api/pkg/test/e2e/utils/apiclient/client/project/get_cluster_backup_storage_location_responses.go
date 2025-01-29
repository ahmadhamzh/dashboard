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

// GetClusterBackupStorageLocationReader is a Reader for the GetClusterBackupStorageLocation structure.
type GetClusterBackupStorageLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterBackupStorageLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterBackupStorageLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterBackupStorageLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterBackupStorageLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterBackupStorageLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterBackupStorageLocationOK creates a GetClusterBackupStorageLocationOK with default headers values
func NewGetClusterBackupStorageLocationOK() *GetClusterBackupStorageLocationOK {
	return &GetClusterBackupStorageLocationOK{}
}

/*
GetClusterBackupStorageLocationOK describes a response with status code 200, with default header values.

ClusterBackupStorageLocation
*/
type GetClusterBackupStorageLocationOK struct {
	Payload *models.ClusterBackupStorageLocation
}

// IsSuccess returns true when this get cluster backup storage location o k response has a 2xx status code
func (o *GetClusterBackupStorageLocationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster backup storage location o k response has a 3xx status code
func (o *GetClusterBackupStorageLocationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster backup storage location o k response has a 4xx status code
func (o *GetClusterBackupStorageLocationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster backup storage location o k response has a 5xx status code
func (o *GetClusterBackupStorageLocationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster backup storage location o k response a status code equal to that given
func (o *GetClusterBackupStorageLocationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster backup storage location o k response
func (o *GetClusterBackupStorageLocationOK) Code() int {
	return 200
}

func (o *GetClusterBackupStorageLocationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocationOK  %+v", 200, o.Payload)
}

func (o *GetClusterBackupStorageLocationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocationOK  %+v", 200, o.Payload)
}

func (o *GetClusterBackupStorageLocationOK) GetPayload() *models.ClusterBackupStorageLocation {
	return o.Payload
}

func (o *GetClusterBackupStorageLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterBackupStorageLocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterBackupStorageLocationUnauthorized creates a GetClusterBackupStorageLocationUnauthorized with default headers values
func NewGetClusterBackupStorageLocationUnauthorized() *GetClusterBackupStorageLocationUnauthorized {
	return &GetClusterBackupStorageLocationUnauthorized{}
}

/*
GetClusterBackupStorageLocationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterBackupStorageLocationUnauthorized struct {
}

// IsSuccess returns true when this get cluster backup storage location unauthorized response has a 2xx status code
func (o *GetClusterBackupStorageLocationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster backup storage location unauthorized response has a 3xx status code
func (o *GetClusterBackupStorageLocationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster backup storage location unauthorized response has a 4xx status code
func (o *GetClusterBackupStorageLocationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster backup storage location unauthorized response has a 5xx status code
func (o *GetClusterBackupStorageLocationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster backup storage location unauthorized response a status code equal to that given
func (o *GetClusterBackupStorageLocationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cluster backup storage location unauthorized response
func (o *GetClusterBackupStorageLocationUnauthorized) Code() int {
	return 401
}

func (o *GetClusterBackupStorageLocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocationUnauthorized ", 401)
}

func (o *GetClusterBackupStorageLocationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocationUnauthorized ", 401)
}

func (o *GetClusterBackupStorageLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterBackupStorageLocationForbidden creates a GetClusterBackupStorageLocationForbidden with default headers values
func NewGetClusterBackupStorageLocationForbidden() *GetClusterBackupStorageLocationForbidden {
	return &GetClusterBackupStorageLocationForbidden{}
}

/*
GetClusterBackupStorageLocationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterBackupStorageLocationForbidden struct {
}

// IsSuccess returns true when this get cluster backup storage location forbidden response has a 2xx status code
func (o *GetClusterBackupStorageLocationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster backup storage location forbidden response has a 3xx status code
func (o *GetClusterBackupStorageLocationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster backup storage location forbidden response has a 4xx status code
func (o *GetClusterBackupStorageLocationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster backup storage location forbidden response has a 5xx status code
func (o *GetClusterBackupStorageLocationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster backup storage location forbidden response a status code equal to that given
func (o *GetClusterBackupStorageLocationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cluster backup storage location forbidden response
func (o *GetClusterBackupStorageLocationForbidden) Code() int {
	return 403
}

func (o *GetClusterBackupStorageLocationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocationForbidden ", 403)
}

func (o *GetClusterBackupStorageLocationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocationForbidden ", 403)
}

func (o *GetClusterBackupStorageLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterBackupStorageLocationDefault creates a GetClusterBackupStorageLocationDefault with default headers values
func NewGetClusterBackupStorageLocationDefault(code int) *GetClusterBackupStorageLocationDefault {
	return &GetClusterBackupStorageLocationDefault{
		_statusCode: code,
	}
}

/*
GetClusterBackupStorageLocationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterBackupStorageLocationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get cluster backup storage location default response has a 2xx status code
func (o *GetClusterBackupStorageLocationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster backup storage location default response has a 3xx status code
func (o *GetClusterBackupStorageLocationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster backup storage location default response has a 4xx status code
func (o *GetClusterBackupStorageLocationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster backup storage location default response has a 5xx status code
func (o *GetClusterBackupStorageLocationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster backup storage location default response a status code equal to that given
func (o *GetClusterBackupStorageLocationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cluster backup storage location default response
func (o *GetClusterBackupStorageLocationDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterBackupStorageLocationDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocation default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterBackupStorageLocationDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusterbackupstoragelocation/{cbsl_name}][%d] getClusterBackupStorageLocation default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterBackupStorageLocationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterBackupStorageLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
