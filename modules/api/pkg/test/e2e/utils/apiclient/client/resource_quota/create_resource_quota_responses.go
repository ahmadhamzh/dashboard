// Code generated by go-swagger; DO NOT EDIT.

package resource_quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateResourceQuotaReader is a Reader for the CreateResourceQuota structure.
type CreateResourceQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateResourceQuotaCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateResourceQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateResourceQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateResourceQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateResourceQuotaCreated creates a CreateResourceQuotaCreated with default headers values
func NewCreateResourceQuotaCreated() *CreateResourceQuotaCreated {
	return &CreateResourceQuotaCreated{}
}

/*
CreateResourceQuotaCreated describes a response with status code 201, with default header values.

EmptyResponse is a empty response
*/
type CreateResourceQuotaCreated struct {
}

// IsSuccess returns true when this create resource quota created response has a 2xx status code
func (o *CreateResourceQuotaCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create resource quota created response has a 3xx status code
func (o *CreateResourceQuotaCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource quota created response has a 4xx status code
func (o *CreateResourceQuotaCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create resource quota created response has a 5xx status code
func (o *CreateResourceQuotaCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource quota created response a status code equal to that given
func (o *CreateResourceQuotaCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create resource quota created response
func (o *CreateResourceQuotaCreated) Code() int {
	return 201
}

func (o *CreateResourceQuotaCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuotaCreated", 201)
}

func (o *CreateResourceQuotaCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuotaCreated", 201)
}

func (o *CreateResourceQuotaCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateResourceQuotaUnauthorized creates a CreateResourceQuotaUnauthorized with default headers values
func NewCreateResourceQuotaUnauthorized() *CreateResourceQuotaUnauthorized {
	return &CreateResourceQuotaUnauthorized{}
}

/*
CreateResourceQuotaUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateResourceQuotaUnauthorized struct {
}

// IsSuccess returns true when this create resource quota unauthorized response has a 2xx status code
func (o *CreateResourceQuotaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource quota unauthorized response has a 3xx status code
func (o *CreateResourceQuotaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource quota unauthorized response has a 4xx status code
func (o *CreateResourceQuotaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource quota unauthorized response has a 5xx status code
func (o *CreateResourceQuotaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource quota unauthorized response a status code equal to that given
func (o *CreateResourceQuotaUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create resource quota unauthorized response
func (o *CreateResourceQuotaUnauthorized) Code() int {
	return 401
}

func (o *CreateResourceQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuotaUnauthorized", 401)
}

func (o *CreateResourceQuotaUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuotaUnauthorized", 401)
}

func (o *CreateResourceQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateResourceQuotaForbidden creates a CreateResourceQuotaForbidden with default headers values
func NewCreateResourceQuotaForbidden() *CreateResourceQuotaForbidden {
	return &CreateResourceQuotaForbidden{}
}

/*
CreateResourceQuotaForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateResourceQuotaForbidden struct {
}

// IsSuccess returns true when this create resource quota forbidden response has a 2xx status code
func (o *CreateResourceQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource quota forbidden response has a 3xx status code
func (o *CreateResourceQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource quota forbidden response has a 4xx status code
func (o *CreateResourceQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource quota forbidden response has a 5xx status code
func (o *CreateResourceQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource quota forbidden response a status code equal to that given
func (o *CreateResourceQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create resource quota forbidden response
func (o *CreateResourceQuotaForbidden) Code() int {
	return 403
}

func (o *CreateResourceQuotaForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuotaForbidden", 403)
}

func (o *CreateResourceQuotaForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuotaForbidden", 403)
}

func (o *CreateResourceQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateResourceQuotaDefault creates a CreateResourceQuotaDefault with default headers values
func NewCreateResourceQuotaDefault(code int) *CreateResourceQuotaDefault {
	return &CreateResourceQuotaDefault{
		_statusCode: code,
	}
}

/*
CreateResourceQuotaDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateResourceQuotaDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create resource quota default response has a 2xx status code
func (o *CreateResourceQuotaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create resource quota default response has a 3xx status code
func (o *CreateResourceQuotaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create resource quota default response has a 4xx status code
func (o *CreateResourceQuotaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create resource quota default response has a 5xx status code
func (o *CreateResourceQuotaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create resource quota default response a status code equal to that given
func (o *CreateResourceQuotaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create resource quota default response
func (o *CreateResourceQuotaDefault) Code() int {
	return o._statusCode
}

func (o *CreateResourceQuotaDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuota default %s", o._statusCode, payload)
}

func (o *CreateResourceQuotaDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/quotas][%d] createResourceQuota default %s", o._statusCode, payload)
}

func (o *CreateResourceQuotaDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateResourceQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateResourceQuotaBody create resource quota body
swagger:model CreateResourceQuotaBody
*/
type CreateResourceQuotaBody struct {

	// subject kind
	SubjectKind string `json:"subjectKind,omitempty"`

	// subject name
	SubjectName string `json:"subjectName,omitempty"`

	// quota
	Quota *models.Quota `json:"quota,omitempty"`
}

// Validate validates this create resource quota body
func (o *CreateResourceQuotaBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateResourceQuotaBody) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(o.Quota) { // not required
		return nil
	}

	if o.Quota != nil {
		if err := o.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "quota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create resource quota body based on the context it is used
func (o *CreateResourceQuotaBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateResourceQuotaBody) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if o.Quota != nil {

		if swag.IsZero(o.Quota) { // not required
			return nil
		}

		if err := o.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "quota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateResourceQuotaBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateResourceQuotaBody) UnmarshalBinary(b []byte) error {
	var res CreateResourceQuotaBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
