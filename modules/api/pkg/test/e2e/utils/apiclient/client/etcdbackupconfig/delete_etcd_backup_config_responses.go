// Code generated by go-swagger; DO NOT EDIT.

package etcdbackupconfig

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

// DeleteEtcdBackupConfigReader is a Reader for the DeleteEtcdBackupConfig structure.
type DeleteEtcdBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEtcdBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEtcdBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEtcdBackupConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEtcdBackupConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteEtcdBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEtcdBackupConfigOK creates a DeleteEtcdBackupConfigOK with default headers values
func NewDeleteEtcdBackupConfigOK() *DeleteEtcdBackupConfigOK {
	return &DeleteEtcdBackupConfigOK{}
}

/*
DeleteEtcdBackupConfigOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdBackupConfigOK struct {
}

// IsSuccess returns true when this delete etcd backup config o k response has a 2xx status code
func (o *DeleteEtcdBackupConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete etcd backup config o k response has a 3xx status code
func (o *DeleteEtcdBackupConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd backup config o k response has a 4xx status code
func (o *DeleteEtcdBackupConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete etcd backup config o k response has a 5xx status code
func (o *DeleteEtcdBackupConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd backup config o k response a status code equal to that given
func (o *DeleteEtcdBackupConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete etcd backup config o k response
func (o *DeleteEtcdBackupConfigOK) Code() int {
	return 200
}

func (o *DeleteEtcdBackupConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigOK", 200)
}

func (o *DeleteEtcdBackupConfigOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigOK", 200)
}

func (o *DeleteEtcdBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdBackupConfigUnauthorized creates a DeleteEtcdBackupConfigUnauthorized with default headers values
func NewDeleteEtcdBackupConfigUnauthorized() *DeleteEtcdBackupConfigUnauthorized {
	return &DeleteEtcdBackupConfigUnauthorized{}
}

/*
DeleteEtcdBackupConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdBackupConfigUnauthorized struct {
}

// IsSuccess returns true when this delete etcd backup config unauthorized response has a 2xx status code
func (o *DeleteEtcdBackupConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete etcd backup config unauthorized response has a 3xx status code
func (o *DeleteEtcdBackupConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd backup config unauthorized response has a 4xx status code
func (o *DeleteEtcdBackupConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete etcd backup config unauthorized response has a 5xx status code
func (o *DeleteEtcdBackupConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd backup config unauthorized response a status code equal to that given
func (o *DeleteEtcdBackupConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete etcd backup config unauthorized response
func (o *DeleteEtcdBackupConfigUnauthorized) Code() int {
	return 401
}

func (o *DeleteEtcdBackupConfigUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigUnauthorized", 401)
}

func (o *DeleteEtcdBackupConfigUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigUnauthorized", 401)
}

func (o *DeleteEtcdBackupConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdBackupConfigForbidden creates a DeleteEtcdBackupConfigForbidden with default headers values
func NewDeleteEtcdBackupConfigForbidden() *DeleteEtcdBackupConfigForbidden {
	return &DeleteEtcdBackupConfigForbidden{}
}

/*
DeleteEtcdBackupConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdBackupConfigForbidden struct {
}

// IsSuccess returns true when this delete etcd backup config forbidden response has a 2xx status code
func (o *DeleteEtcdBackupConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete etcd backup config forbidden response has a 3xx status code
func (o *DeleteEtcdBackupConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd backup config forbidden response has a 4xx status code
func (o *DeleteEtcdBackupConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete etcd backup config forbidden response has a 5xx status code
func (o *DeleteEtcdBackupConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd backup config forbidden response a status code equal to that given
func (o *DeleteEtcdBackupConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete etcd backup config forbidden response
func (o *DeleteEtcdBackupConfigForbidden) Code() int {
	return 403
}

func (o *DeleteEtcdBackupConfigForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigForbidden", 403)
}

func (o *DeleteEtcdBackupConfigForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfigForbidden", 403)
}

func (o *DeleteEtcdBackupConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdBackupConfigDefault creates a DeleteEtcdBackupConfigDefault with default headers values
func NewDeleteEtcdBackupConfigDefault(code int) *DeleteEtcdBackupConfigDefault {
	return &DeleteEtcdBackupConfigDefault{
		_statusCode: code,
	}
}

/*
DeleteEtcdBackupConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteEtcdBackupConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete etcd backup config default response has a 2xx status code
func (o *DeleteEtcdBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete etcd backup config default response has a 3xx status code
func (o *DeleteEtcdBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete etcd backup config default response has a 4xx status code
func (o *DeleteEtcdBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete etcd backup config default response has a 5xx status code
func (o *DeleteEtcdBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete etcd backup config default response a status code equal to that given
func (o *DeleteEtcdBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete etcd backup config default response
func (o *DeleteEtcdBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEtcdBackupConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *DeleteEtcdBackupConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] deleteEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *DeleteEtcdBackupConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteEtcdBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
