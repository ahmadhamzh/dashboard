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

// UpgradeClusterNodeDeploymentsV2Reader is a Reader for the UpgradeClusterNodeDeploymentsV2 structure.
type UpgradeClusterNodeDeploymentsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeClusterNodeDeploymentsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpgradeClusterNodeDeploymentsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpgradeClusterNodeDeploymentsV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpgradeClusterNodeDeploymentsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpgradeClusterNodeDeploymentsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpgradeClusterNodeDeploymentsV2OK creates a UpgradeClusterNodeDeploymentsV2OK with default headers values
func NewUpgradeClusterNodeDeploymentsV2OK() *UpgradeClusterNodeDeploymentsV2OK {
	return &UpgradeClusterNodeDeploymentsV2OK{}
}

/*
UpgradeClusterNodeDeploymentsV2OK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type UpgradeClusterNodeDeploymentsV2OK struct {
}

// IsSuccess returns true when this upgrade cluster node deployments v2 o k response has a 2xx status code
func (o *UpgradeClusterNodeDeploymentsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upgrade cluster node deployments v2 o k response has a 3xx status code
func (o *UpgradeClusterNodeDeploymentsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upgrade cluster node deployments v2 o k response has a 4xx status code
func (o *UpgradeClusterNodeDeploymentsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upgrade cluster node deployments v2 o k response has a 5xx status code
func (o *UpgradeClusterNodeDeploymentsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this upgrade cluster node deployments v2 o k response a status code equal to that given
func (o *UpgradeClusterNodeDeploymentsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upgrade cluster node deployments v2 o k response
func (o *UpgradeClusterNodeDeploymentsV2OK) Code() int {
	return 200
}

func (o *UpgradeClusterNodeDeploymentsV2OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2OK", 200)
}

func (o *UpgradeClusterNodeDeploymentsV2OK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2OK", 200)
}

func (o *UpgradeClusterNodeDeploymentsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpgradeClusterNodeDeploymentsV2Unauthorized creates a UpgradeClusterNodeDeploymentsV2Unauthorized with default headers values
func NewUpgradeClusterNodeDeploymentsV2Unauthorized() *UpgradeClusterNodeDeploymentsV2Unauthorized {
	return &UpgradeClusterNodeDeploymentsV2Unauthorized{}
}

/*
UpgradeClusterNodeDeploymentsV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpgradeClusterNodeDeploymentsV2Unauthorized struct {
}

// IsSuccess returns true when this upgrade cluster node deployments v2 unauthorized response has a 2xx status code
func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upgrade cluster node deployments v2 unauthorized response has a 3xx status code
func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upgrade cluster node deployments v2 unauthorized response has a 4xx status code
func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upgrade cluster node deployments v2 unauthorized response has a 5xx status code
func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upgrade cluster node deployments v2 unauthorized response a status code equal to that given
func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the upgrade cluster node deployments v2 unauthorized response
func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) Code() int {
	return 401
}

func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2Unauthorized", 401)
}

func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2Unauthorized", 401)
}

func (o *UpgradeClusterNodeDeploymentsV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpgradeClusterNodeDeploymentsV2Forbidden creates a UpgradeClusterNodeDeploymentsV2Forbidden with default headers values
func NewUpgradeClusterNodeDeploymentsV2Forbidden() *UpgradeClusterNodeDeploymentsV2Forbidden {
	return &UpgradeClusterNodeDeploymentsV2Forbidden{}
}

/*
UpgradeClusterNodeDeploymentsV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpgradeClusterNodeDeploymentsV2Forbidden struct {
}

// IsSuccess returns true when this upgrade cluster node deployments v2 forbidden response has a 2xx status code
func (o *UpgradeClusterNodeDeploymentsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upgrade cluster node deployments v2 forbidden response has a 3xx status code
func (o *UpgradeClusterNodeDeploymentsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upgrade cluster node deployments v2 forbidden response has a 4xx status code
func (o *UpgradeClusterNodeDeploymentsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upgrade cluster node deployments v2 forbidden response has a 5xx status code
func (o *UpgradeClusterNodeDeploymentsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upgrade cluster node deployments v2 forbidden response a status code equal to that given
func (o *UpgradeClusterNodeDeploymentsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upgrade cluster node deployments v2 forbidden response
func (o *UpgradeClusterNodeDeploymentsV2Forbidden) Code() int {
	return 403
}

func (o *UpgradeClusterNodeDeploymentsV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2Forbidden", 403)
}

func (o *UpgradeClusterNodeDeploymentsV2Forbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2Forbidden", 403)
}

func (o *UpgradeClusterNodeDeploymentsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpgradeClusterNodeDeploymentsV2Default creates a UpgradeClusterNodeDeploymentsV2Default with default headers values
func NewUpgradeClusterNodeDeploymentsV2Default(code int) *UpgradeClusterNodeDeploymentsV2Default {
	return &UpgradeClusterNodeDeploymentsV2Default{
		_statusCode: code,
	}
}

/*
UpgradeClusterNodeDeploymentsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type UpgradeClusterNodeDeploymentsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this upgrade cluster node deployments v2 default response has a 2xx status code
func (o *UpgradeClusterNodeDeploymentsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upgrade cluster node deployments v2 default response has a 3xx status code
func (o *UpgradeClusterNodeDeploymentsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upgrade cluster node deployments v2 default response has a 4xx status code
func (o *UpgradeClusterNodeDeploymentsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upgrade cluster node deployments v2 default response has a 5xx status code
func (o *UpgradeClusterNodeDeploymentsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upgrade cluster node deployments v2 default response a status code equal to that given
func (o *UpgradeClusterNodeDeploymentsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the upgrade cluster node deployments v2 default response
func (o *UpgradeClusterNodeDeploymentsV2Default) Code() int {
	return o._statusCode
}

func (o *UpgradeClusterNodeDeploymentsV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2 default %s", o._statusCode, payload)
}

func (o *UpgradeClusterNodeDeploymentsV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades][%d] upgradeClusterNodeDeploymentsV2 default %s", o._statusCode, payload)
}

func (o *UpgradeClusterNodeDeploymentsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpgradeClusterNodeDeploymentsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
