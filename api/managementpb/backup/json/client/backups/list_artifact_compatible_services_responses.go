// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// ListArtifactCompatibleServicesReader is a Reader for the ListArtifactCompatibleServices structure.
type ListArtifactCompatibleServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArtifactCompatibleServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArtifactCompatibleServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListArtifactCompatibleServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArtifactCompatibleServicesOK creates a ListArtifactCompatibleServicesOK with default headers values
func NewListArtifactCompatibleServicesOK() *ListArtifactCompatibleServicesOK {
	return &ListArtifactCompatibleServicesOK{}
}

/*ListArtifactCompatibleServicesOK handles this case with default header values.

A successful response.
*/
type ListArtifactCompatibleServicesOK struct {
	Payload *ListArtifactCompatibleServicesOKBody
}

func (o *ListArtifactCompatibleServicesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ListArtifactCompatibleServices][%d] listArtifactCompatibleServicesOk  %+v", 200, o.Payload)
}

func (o *ListArtifactCompatibleServicesOK) GetPayload() *ListArtifactCompatibleServicesOKBody {
	return o.Payload
}

func (o *ListArtifactCompatibleServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListArtifactCompatibleServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactCompatibleServicesDefault creates a ListArtifactCompatibleServicesDefault with default headers values
func NewListArtifactCompatibleServicesDefault(code int) *ListArtifactCompatibleServicesDefault {
	return &ListArtifactCompatibleServicesDefault{
		_statusCode: code,
	}
}

/*ListArtifactCompatibleServicesDefault handles this case with default header values.

An unexpected error response.
*/
type ListArtifactCompatibleServicesDefault struct {
	_statusCode int

	Payload *ListArtifactCompatibleServicesDefaultBody
}

// Code gets the status code for the list artifact compatible services default response
func (o *ListArtifactCompatibleServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListArtifactCompatibleServicesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ListArtifactCompatibleServices][%d] ListArtifactCompatibleServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListArtifactCompatibleServicesDefault) GetPayload() *ListArtifactCompatibleServicesDefaultBody {
	return o.Payload
}

func (o *ListArtifactCompatibleServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListArtifactCompatibleServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListArtifactCompatibleServicesBody list artifact compatible services body
swagger:model ListArtifactCompatibleServicesBody
*/
type ListArtifactCompatibleServicesBody struct {

	// Artifact id used to determine restore compatibility.
	ArtifactID string `json:"artifact_id,omitempty"`
}

// Validate validates this list artifact compatible services body
func (o *ListArtifactCompatibleServicesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListArtifactCompatibleServicesDefaultBody list artifact compatible services default body
swagger:model ListArtifactCompatibleServicesDefaultBody
*/
type ListArtifactCompatibleServicesDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list artifact compatible services default body
func (o *ListArtifactCompatibleServicesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactCompatibleServicesDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListArtifactCompatibleServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListArtifactCompatibleServicesOKBody list artifact compatible services OK body
swagger:model ListArtifactCompatibleServicesOKBody
*/
type ListArtifactCompatibleServicesOKBody struct {

	// mysql
	Mysql []*MysqlItems0 `json:"mysql"`

	// mongodb
	Mongodb []*MongodbItems0 `json:"mongodb"`
}

// Validate validates this list artifact compatible services OK body
func (o *ListArtifactCompatibleServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactCompatibleServicesOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	for i := 0; i < len(o.Mysql); i++ {
		if swag.IsZero(o.Mysql[i]) { // not required
			continue
		}

		if o.Mysql[i] != nil {
			if err := o.Mysql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactCompatibleServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListArtifactCompatibleServicesOKBody) validateMongodb(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	for i := 0; i < len(o.Mongodb); i++ {
		if swag.IsZero(o.Mongodb[i]) { // not required
			continue
		}

		if o.Mongodb[i] != nil {
			if err := o.Mongodb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactCompatibleServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactCompatibleServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactCompatibleServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MongodbItems0 MongoDBService represents a generic MongoDB instance.
swagger:model MongodbItems0
*/
type MongodbItems0 struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this mongodb items0
func (o *MongodbItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MongodbItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MongodbItems0) UnmarshalBinary(b []byte) error {
	var res MongodbItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MysqlItems0 MySQLService represents a generic MySQL instance.
swagger:model MysqlItems0
*/
type MysqlItems0 struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this mysql items0
func (o *MysqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MysqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MysqlItems0) UnmarshalBinary(b []byte) error {
	var res MysqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
