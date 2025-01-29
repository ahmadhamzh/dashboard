// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

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

// GetRuleGroupReader is a Reader for the GetRuleGroup structure.
type GetRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRuleGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRuleGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuleGroupOK creates a GetRuleGroupOK with default headers values
func NewGetRuleGroupOK() *GetRuleGroupOK {
	return &GetRuleGroupOK{}
}

/*
GetRuleGroupOK describes a response with status code 200, with default header values.

RuleGroup
*/
type GetRuleGroupOK struct {
	Payload *models.RuleGroup
}

// IsSuccess returns true when this get rule group o k response has a 2xx status code
func (o *GetRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rule group o k response has a 3xx status code
func (o *GetRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule group o k response has a 4xx status code
func (o *GetRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rule group o k response has a 5xx status code
func (o *GetRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule group o k response a status code equal to that given
func (o *GetRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rule group o k response
func (o *GetRuleGroupOK) Code() int {
	return 200
}

func (o *GetRuleGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroupOK %s", 200, payload)
}

func (o *GetRuleGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroupOK %s", 200, payload)
}

func (o *GetRuleGroupOK) GetPayload() *models.RuleGroup {
	return o.Payload
}

func (o *GetRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleGroupUnauthorized creates a GetRuleGroupUnauthorized with default headers values
func NewGetRuleGroupUnauthorized() *GetRuleGroupUnauthorized {
	return &GetRuleGroupUnauthorized{}
}

/*
GetRuleGroupUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetRuleGroupUnauthorized struct {
}

// IsSuccess returns true when this get rule group unauthorized response has a 2xx status code
func (o *GetRuleGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule group unauthorized response has a 3xx status code
func (o *GetRuleGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule group unauthorized response has a 4xx status code
func (o *GetRuleGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule group unauthorized response has a 5xx status code
func (o *GetRuleGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule group unauthorized response a status code equal to that given
func (o *GetRuleGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get rule group unauthorized response
func (o *GetRuleGroupUnauthorized) Code() int {
	return 401
}

func (o *GetRuleGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroupUnauthorized", 401)
}

func (o *GetRuleGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroupUnauthorized", 401)
}

func (o *GetRuleGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuleGroupForbidden creates a GetRuleGroupForbidden with default headers values
func NewGetRuleGroupForbidden() *GetRuleGroupForbidden {
	return &GetRuleGroupForbidden{}
}

/*
GetRuleGroupForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetRuleGroupForbidden struct {
}

// IsSuccess returns true when this get rule group forbidden response has a 2xx status code
func (o *GetRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule group forbidden response has a 3xx status code
func (o *GetRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule group forbidden response has a 4xx status code
func (o *GetRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule group forbidden response has a 5xx status code
func (o *GetRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule group forbidden response a status code equal to that given
func (o *GetRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get rule group forbidden response
func (o *GetRuleGroupForbidden) Code() int {
	return 403
}

func (o *GetRuleGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroupForbidden", 403)
}

func (o *GetRuleGroupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroupForbidden", 403)
}

func (o *GetRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuleGroupDefault creates a GetRuleGroupDefault with default headers values
func NewGetRuleGroupDefault(code int) *GetRuleGroupDefault {
	return &GetRuleGroupDefault{
		_statusCode: code,
	}
}

/*
GetRuleGroupDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetRuleGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get rule group default response has a 2xx status code
func (o *GetRuleGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get rule group default response has a 3xx status code
func (o *GetRuleGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get rule group default response has a 4xx status code
func (o *GetRuleGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get rule group default response has a 5xx status code
func (o *GetRuleGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get rule group default response a status code equal to that given
func (o *GetRuleGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get rule group default response
func (o *GetRuleGroupDefault) Code() int {
	return o._statusCode
}

func (o *GetRuleGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroup default %s", o._statusCode, payload)
}

func (o *GetRuleGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}][%d] getRuleGroup default %s", o._statusCode, payload)
}

func (o *GetRuleGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRuleGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
