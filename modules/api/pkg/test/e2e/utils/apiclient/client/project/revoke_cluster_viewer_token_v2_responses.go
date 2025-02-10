// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// RevokeClusterViewerTokenV2Reader is a Reader for the RevokeClusterViewerTokenV2 structure.
type RevokeClusterViewerTokenV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeClusterViewerTokenV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeClusterViewerTokenV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeClusterViewerTokenV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeClusterViewerTokenV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRevokeClusterViewerTokenV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeClusterViewerTokenV2OK creates a RevokeClusterViewerTokenV2OK with default headers values
func NewRevokeClusterViewerTokenV2OK() *RevokeClusterViewerTokenV2OK {
	return &RevokeClusterViewerTokenV2OK{}
}

/*
RevokeClusterViewerTokenV2OK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterViewerTokenV2OK struct {
}

// IsSuccess returns true when this revoke cluster viewer token v2 o k response has a 2xx status code
func (o *RevokeClusterViewerTokenV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revoke cluster viewer token v2 o k response has a 3xx status code
func (o *RevokeClusterViewerTokenV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster viewer token v2 o k response has a 4xx status code
func (o *RevokeClusterViewerTokenV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revoke cluster viewer token v2 o k response has a 5xx status code
func (o *RevokeClusterViewerTokenV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster viewer token v2 o k response a status code equal to that given
func (o *RevokeClusterViewerTokenV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *RevokeClusterViewerTokenV2OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2OK ", 200)
}

func (o *RevokeClusterViewerTokenV2OK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2OK ", 200)
}

func (o *RevokeClusterViewerTokenV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterViewerTokenV2Unauthorized creates a RevokeClusterViewerTokenV2Unauthorized with default headers values
func NewRevokeClusterViewerTokenV2Unauthorized() *RevokeClusterViewerTokenV2Unauthorized {
	return &RevokeClusterViewerTokenV2Unauthorized{}
}

/*
RevokeClusterViewerTokenV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterViewerTokenV2Unauthorized struct {
}

// IsSuccess returns true when this revoke cluster viewer token v2 unauthorized response has a 2xx status code
func (o *RevokeClusterViewerTokenV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke cluster viewer token v2 unauthorized response has a 3xx status code
func (o *RevokeClusterViewerTokenV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster viewer token v2 unauthorized response has a 4xx status code
func (o *RevokeClusterViewerTokenV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke cluster viewer token v2 unauthorized response has a 5xx status code
func (o *RevokeClusterViewerTokenV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster viewer token v2 unauthorized response a status code equal to that given
func (o *RevokeClusterViewerTokenV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RevokeClusterViewerTokenV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2Unauthorized ", 401)
}

func (o *RevokeClusterViewerTokenV2Unauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2Unauthorized ", 401)
}

func (o *RevokeClusterViewerTokenV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterViewerTokenV2Forbidden creates a RevokeClusterViewerTokenV2Forbidden with default headers values
func NewRevokeClusterViewerTokenV2Forbidden() *RevokeClusterViewerTokenV2Forbidden {
	return &RevokeClusterViewerTokenV2Forbidden{}
}

/*
RevokeClusterViewerTokenV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterViewerTokenV2Forbidden struct {
}

// IsSuccess returns true when this revoke cluster viewer token v2 forbidden response has a 2xx status code
func (o *RevokeClusterViewerTokenV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke cluster viewer token v2 forbidden response has a 3xx status code
func (o *RevokeClusterViewerTokenV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke cluster viewer token v2 forbidden response has a 4xx status code
func (o *RevokeClusterViewerTokenV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke cluster viewer token v2 forbidden response has a 5xx status code
func (o *RevokeClusterViewerTokenV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke cluster viewer token v2 forbidden response a status code equal to that given
func (o *RevokeClusterViewerTokenV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RevokeClusterViewerTokenV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2Forbidden ", 403)
}

func (o *RevokeClusterViewerTokenV2Forbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2Forbidden ", 403)
}

func (o *RevokeClusterViewerTokenV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterViewerTokenV2Default creates a RevokeClusterViewerTokenV2Default with default headers values
func NewRevokeClusterViewerTokenV2Default(code int) *RevokeClusterViewerTokenV2Default {
	return &RevokeClusterViewerTokenV2Default{
		_statusCode: code,
	}
}

/*
RevokeClusterViewerTokenV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type RevokeClusterViewerTokenV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the revoke cluster viewer token v2 default response
func (o *RevokeClusterViewerTokenV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this revoke cluster viewer token v2 default response has a 2xx status code
func (o *RevokeClusterViewerTokenV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this revoke cluster viewer token v2 default response has a 3xx status code
func (o *RevokeClusterViewerTokenV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this revoke cluster viewer token v2 default response has a 4xx status code
func (o *RevokeClusterViewerTokenV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this revoke cluster viewer token v2 default response has a 5xx status code
func (o *RevokeClusterViewerTokenV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this revoke cluster viewer token v2 default response a status code equal to that given
func (o *RevokeClusterViewerTokenV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RevokeClusterViewerTokenV2Default) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2 default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeClusterViewerTokenV2Default) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenV2 default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeClusterViewerTokenV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RevokeClusterViewerTokenV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
