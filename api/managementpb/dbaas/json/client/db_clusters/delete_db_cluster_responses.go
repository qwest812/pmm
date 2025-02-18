// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteDBClusterReader is a Reader for the DeleteDBCluster structure.
type DeleteDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDBClusterOK creates a DeleteDBClusterOK with default headers values
func NewDeleteDBClusterOK() *DeleteDBClusterOK {
	return &DeleteDBClusterOK{}
}

/*DeleteDBClusterOK handles this case with default header values.

A successful response.
*/
type DeleteDBClusterOK struct {
	Payload interface{}
}

func (o *DeleteDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Delete][%d] deleteDbClusterOk  %+v", 200, o.Payload)
}

func (o *DeleteDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDBClusterDefault creates a DeleteDBClusterDefault with default headers values
func NewDeleteDBClusterDefault(code int) *DeleteDBClusterDefault {
	return &DeleteDBClusterDefault{
		_statusCode: code,
	}
}

/*DeleteDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteDBClusterDefault struct {
	_statusCode int

	Payload *DeleteDBClusterDefaultBody
}

// Code gets the status code for the delete DB cluster default response
func (o *DeleteDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Delete][%d] DeleteDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDBClusterDefault) GetPayload() *DeleteDBClusterDefaultBody {
	return o.Payload
}

func (o *DeleteDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteDBClusterBody delete DB cluster body
swagger:model DeleteDBClusterBody
*/
type DeleteDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// DB cluster name.
	Name string `json:"name,omitempty"`

	// DBClusterType represents database cluster type.
	//
	//  - DB_CLUSTER_TYPE_INVALID: DB_CLUSTER_TYPE_INVALID represents unknown cluster type.
	//  - DB_CLUSTER_TYPE_PXC: DB_CLUSTER_TYPE_PXC represents pxc cluster type.
	//  - DB_CLUSTER_TYPE_PSMDB: DB_CLUSTER_TYPE_PSMDB represents psmdb cluster type.
	// Enum: [DB_CLUSTER_TYPE_INVALID DB_CLUSTER_TYPE_PXC DB_CLUSTER_TYPE_PSMDB]
	ClusterType *string `json:"cluster_type,omitempty"`
}

// Validate validates this delete DB cluster body
func (o *DeleteDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteDbClusterBodyTypeClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DB_CLUSTER_TYPE_INVALID","DB_CLUSTER_TYPE_PXC","DB_CLUSTER_TYPE_PSMDB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteDbClusterBodyTypeClusterTypePropEnum = append(deleteDbClusterBodyTypeClusterTypePropEnum, v)
	}
}

const (

	// DeleteDBClusterBodyClusterTypeDBCLUSTERTYPEINVALID captures enum value "DB_CLUSTER_TYPE_INVALID"
	DeleteDBClusterBodyClusterTypeDBCLUSTERTYPEINVALID string = "DB_CLUSTER_TYPE_INVALID"

	// DeleteDBClusterBodyClusterTypeDBCLUSTERTYPEPXC captures enum value "DB_CLUSTER_TYPE_PXC"
	DeleteDBClusterBodyClusterTypeDBCLUSTERTYPEPXC string = "DB_CLUSTER_TYPE_PXC"

	// DeleteDBClusterBodyClusterTypeDBCLUSTERTYPEPSMDB captures enum value "DB_CLUSTER_TYPE_PSMDB"
	DeleteDBClusterBodyClusterTypeDBCLUSTERTYPEPSMDB string = "DB_CLUSTER_TYPE_PSMDB"
)

// prop value enum
func (o *DeleteDBClusterBody) validateClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deleteDbClusterBodyTypeClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *DeleteDBClusterBody) validateClusterType(formats strfmt.Registry) error {

	if swag.IsZero(o.ClusterType) { // not required
		return nil
	}

	// value enum
	if err := o.validateClusterTypeEnum("body"+"."+"cluster_type", "body", *o.ClusterType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDBClusterBody) UnmarshalBinary(b []byte) error {
	var res DeleteDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteDBClusterDefaultBody delete DB cluster default body
swagger:model DeleteDBClusterDefaultBody
*/
type DeleteDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this delete DB cluster default body
func (o *DeleteDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("DeleteDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
