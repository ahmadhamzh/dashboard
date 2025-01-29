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

// AssignSSHKeyToClusterV2Reader is a Reader for the AssignSSHKeyToClusterV2 structure.
type AssignSSHKeyToClusterV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignSSHKeyToClusterV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAssignSSHKeyToClusterV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssignSSHKeyToClusterV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignSSHKeyToClusterV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAssignSSHKeyToClusterV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignSSHKeyToClusterV2Created creates a AssignSSHKeyToClusterV2Created with default headers values
func NewAssignSSHKeyToClusterV2Created() *AssignSSHKeyToClusterV2Created {
	return &AssignSSHKeyToClusterV2Created{}
}

/*
AssignSSHKeyToClusterV2Created describes a response with status code 201, with default header values.

SSHKey
*/
type AssignSSHKeyToClusterV2Created struct {
	Payload *models.SSHKey
}

// IsSuccess returns true when this assign Ssh key to cluster v2 created response has a 2xx status code
func (o *AssignSSHKeyToClusterV2Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign Ssh key to cluster v2 created response has a 3xx status code
func (o *AssignSSHKeyToClusterV2Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign Ssh key to cluster v2 created response has a 4xx status code
func (o *AssignSSHKeyToClusterV2Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign Ssh key to cluster v2 created response has a 5xx status code
func (o *AssignSSHKeyToClusterV2Created) IsServerError() bool {
	return false
}

// IsCode returns true when this assign Ssh key to cluster v2 created response a status code equal to that given
func (o *AssignSSHKeyToClusterV2Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the assign Ssh key to cluster v2 created response
func (o *AssignSSHKeyToClusterV2Created) Code() int {
	return 201
}

func (o *AssignSSHKeyToClusterV2Created) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterV2Created %s", 201, payload)
}

func (o *AssignSSHKeyToClusterV2Created) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterV2Created %s", 201, payload)
}

func (o *AssignSSHKeyToClusterV2Created) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *AssignSSHKeyToClusterV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignSSHKeyToClusterV2Unauthorized creates a AssignSSHKeyToClusterV2Unauthorized with default headers values
func NewAssignSSHKeyToClusterV2Unauthorized() *AssignSSHKeyToClusterV2Unauthorized {
	return &AssignSSHKeyToClusterV2Unauthorized{}
}

/*
AssignSSHKeyToClusterV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type AssignSSHKeyToClusterV2Unauthorized struct {
}

// IsSuccess returns true when this assign Ssh key to cluster v2 unauthorized response has a 2xx status code
func (o *AssignSSHKeyToClusterV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign Ssh key to cluster v2 unauthorized response has a 3xx status code
func (o *AssignSSHKeyToClusterV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign Ssh key to cluster v2 unauthorized response has a 4xx status code
func (o *AssignSSHKeyToClusterV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign Ssh key to cluster v2 unauthorized response has a 5xx status code
func (o *AssignSSHKeyToClusterV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this assign Ssh key to cluster v2 unauthorized response a status code equal to that given
func (o *AssignSSHKeyToClusterV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the assign Ssh key to cluster v2 unauthorized response
func (o *AssignSSHKeyToClusterV2Unauthorized) Code() int {
	return 401
}

func (o *AssignSSHKeyToClusterV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterV2Unauthorized", 401)
}

func (o *AssignSSHKeyToClusterV2Unauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterV2Unauthorized", 401)
}

func (o *AssignSSHKeyToClusterV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSSHKeyToClusterV2Forbidden creates a AssignSSHKeyToClusterV2Forbidden with default headers values
func NewAssignSSHKeyToClusterV2Forbidden() *AssignSSHKeyToClusterV2Forbidden {
	return &AssignSSHKeyToClusterV2Forbidden{}
}

/*
AssignSSHKeyToClusterV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type AssignSSHKeyToClusterV2Forbidden struct {
}

// IsSuccess returns true when this assign Ssh key to cluster v2 forbidden response has a 2xx status code
func (o *AssignSSHKeyToClusterV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign Ssh key to cluster v2 forbidden response has a 3xx status code
func (o *AssignSSHKeyToClusterV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign Ssh key to cluster v2 forbidden response has a 4xx status code
func (o *AssignSSHKeyToClusterV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign Ssh key to cluster v2 forbidden response has a 5xx status code
func (o *AssignSSHKeyToClusterV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign Ssh key to cluster v2 forbidden response a status code equal to that given
func (o *AssignSSHKeyToClusterV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the assign Ssh key to cluster v2 forbidden response
func (o *AssignSSHKeyToClusterV2Forbidden) Code() int {
	return 403
}

func (o *AssignSSHKeyToClusterV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterV2Forbidden", 403)
}

func (o *AssignSSHKeyToClusterV2Forbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterV2Forbidden", 403)
}

func (o *AssignSSHKeyToClusterV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSSHKeyToClusterV2Default creates a AssignSSHKeyToClusterV2Default with default headers values
func NewAssignSSHKeyToClusterV2Default(code int) *AssignSSHKeyToClusterV2Default {
	return &AssignSSHKeyToClusterV2Default{
		_statusCode: code,
	}
}

/*
AssignSSHKeyToClusterV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type AssignSSHKeyToClusterV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign SSH key to cluster v2 default response has a 2xx status code
func (o *AssignSSHKeyToClusterV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign SSH key to cluster v2 default response has a 3xx status code
func (o *AssignSSHKeyToClusterV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign SSH key to cluster v2 default response has a 4xx status code
func (o *AssignSSHKeyToClusterV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign SSH key to cluster v2 default response has a 5xx status code
func (o *AssignSSHKeyToClusterV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign SSH key to cluster v2 default response a status code equal to that given
func (o *AssignSSHKeyToClusterV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign SSH key to cluster v2 default response
func (o *AssignSSHKeyToClusterV2Default) Code() int {
	return o._statusCode
}

func (o *AssignSSHKeyToClusterV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSSHKeyToClusterV2 default %s", o._statusCode, payload)
}

func (o *AssignSSHKeyToClusterV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSSHKeyToClusterV2 default %s", o._statusCode, payload)
}

func (o *AssignSSHKeyToClusterV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignSSHKeyToClusterV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
