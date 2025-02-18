// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckForOperatorUpdateReader is a Reader for the CheckForOperatorUpdate structure.
type CheckForOperatorUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckForOperatorUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckForOperatorUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckForOperatorUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckForOperatorUpdateOK creates a CheckForOperatorUpdateOK with default headers values
func NewCheckForOperatorUpdateOK() *CheckForOperatorUpdateOK {
	return &CheckForOperatorUpdateOK{}
}

/*CheckForOperatorUpdateOK handles this case with default header values.

A successful response.
*/
type CheckForOperatorUpdateOK struct {
	Payload *CheckForOperatorUpdateOKBody
}

func (o *CheckForOperatorUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/CheckForOperatorUpdate][%d] checkForOperatorUpdateOk  %+v", 200, o.Payload)
}

func (o *CheckForOperatorUpdateOK) GetPayload() *CheckForOperatorUpdateOKBody {
	return o.Payload
}

func (o *CheckForOperatorUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckForOperatorUpdateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckForOperatorUpdateDefault creates a CheckForOperatorUpdateDefault with default headers values
func NewCheckForOperatorUpdateDefault(code int) *CheckForOperatorUpdateDefault {
	return &CheckForOperatorUpdateDefault{
		_statusCode: code,
	}
}

/*CheckForOperatorUpdateDefault handles this case with default header values.

An unexpected error response.
*/
type CheckForOperatorUpdateDefault struct {
	_statusCode int

	Payload *CheckForOperatorUpdateDefaultBody
}

// Code gets the status code for the check for operator update default response
func (o *CheckForOperatorUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CheckForOperatorUpdateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/CheckForOperatorUpdate][%d] CheckForOperatorUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *CheckForOperatorUpdateDefault) GetPayload() *CheckForOperatorUpdateDefaultBody {
	return o.Payload
}

func (o *CheckForOperatorUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckForOperatorUpdateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckForOperatorUpdateDefaultBody check for operator update default body
swagger:model CheckForOperatorUpdateDefaultBody
*/
type CheckForOperatorUpdateDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this check for operator update default body
func (o *CheckForOperatorUpdateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckForOperatorUpdateDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CheckForOperatorUpdate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckForOperatorUpdateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckForOperatorUpdateDefaultBody) UnmarshalBinary(b []byte) error {
	var res CheckForOperatorUpdateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckForOperatorUpdateOKBody check for operator update OK body
swagger:model CheckForOperatorUpdateOKBody
*/
type CheckForOperatorUpdateOKBody struct {

	// The cluster name is used as a key for this map, value contains components and their inforamtion about update.
	ClusterToComponents map[string]ClusterToComponentsAnon `json:"cluster_to_components,omitempty"`
}

// Validate validates this check for operator update OK body
func (o *CheckForOperatorUpdateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusterToComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckForOperatorUpdateOKBody) validateClusterToComponents(formats strfmt.Registry) error {

	if swag.IsZero(o.ClusterToComponents) { // not required
		return nil
	}

	for k := range o.ClusterToComponents {

		if swag.IsZero(o.ClusterToComponents[k]) { // not required
			continue
		}
		if val, ok := o.ClusterToComponents[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckForOperatorUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckForOperatorUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res CheckForOperatorUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClusterToComponentsAnon ComponentsUpdateInformation contains info about components and their available latest versions.
swagger:model ClusterToComponentsAnon
*/
type ClusterToComponentsAnon struct {

	// component_to_update_information stores, under the name of the component, information about the update.
	// "pxc-operator", "psmdb-operator" are names used by backend for our operators.
	ComponentToUpdateInformation map[string]ClusterToComponentsAnonComponentToUpdateInformationAnon `json:"component_to_update_information,omitempty"`
}

// Validate validates this cluster to components anon
func (o *ClusterToComponentsAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComponentToUpdateInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClusterToComponentsAnon) validateComponentToUpdateInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ComponentToUpdateInformation) { // not required
		return nil
	}

	for k := range o.ComponentToUpdateInformation {

		if swag.IsZero(o.ComponentToUpdateInformation[k]) { // not required
			continue
		}
		if val, ok := o.ComponentToUpdateInformation[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClusterToComponentsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClusterToComponentsAnon) UnmarshalBinary(b []byte) error {
	var res ClusterToComponentsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClusterToComponentsAnonComponentToUpdateInformationAnon ComponentUpdateInformation contains version we can update to for certain component.
swagger:model ClusterToComponentsAnonComponentToUpdateInformationAnon
*/
type ClusterToComponentsAnonComponentToUpdateInformationAnon struct {

	// available version
	AvailableVersion string `json:"available_version,omitempty"`
}

// Validate validates this cluster to components anon component to update information anon
func (o *ClusterToComponentsAnonComponentToUpdateInformationAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClusterToComponentsAnonComponentToUpdateInformationAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClusterToComponentsAnonComponentToUpdateInformationAnon) UnmarshalBinary(b []byte) error {
	var res ClusterToComponentsAnonComponentToUpdateInformationAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
