// Code generated by go-swagger; DO NOT EDIT.

package object_details

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
	"github.com/go-openapi/validate"
)

// GetHistogramReader is a Reader for the GetHistogram structure.
type GetHistogramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistogramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHistogramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHistogramDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHistogramOK creates a GetHistogramOK with default headers values
func NewGetHistogramOK() *GetHistogramOK {
	return &GetHistogramOK{}
}

/*GetHistogramOK handles this case with default header values.

A successful response.
*/
type GetHistogramOK struct {
	Payload *GetHistogramOKBody
}

func (o *GetHistogramOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetHistogram][%d] getHistogramOk  %+v", 200, o.Payload)
}

func (o *GetHistogramOK) GetPayload() *GetHistogramOKBody {
	return o.Payload
}

func (o *GetHistogramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHistogramOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHistogramDefault creates a GetHistogramDefault with default headers values
func NewGetHistogramDefault(code int) *GetHistogramDefault {
	return &GetHistogramDefault{
		_statusCode: code,
	}
}

/*GetHistogramDefault handles this case with default header values.

An unexpected error response.
*/
type GetHistogramDefault struct {
	_statusCode int

	Payload *GetHistogramDefaultBody
}

// Code gets the status code for the get histogram default response
func (o *GetHistogramDefault) Code() int {
	return o._statusCode
}

func (o *GetHistogramDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetHistogram][%d] GetHistogram default  %+v", o._statusCode, o.Payload)
}

func (o *GetHistogramDefault) GetPayload() *GetHistogramDefaultBody {
	return o.Payload
}

func (o *GetHistogramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHistogramDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

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

/*GetHistogramBody HistogramRequest defines filtering by time range, labels and queryid.
swagger:model GetHistogramBody
*/
type GetHistogramBody struct {

	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// labels
	Labels []*LabelsItems0 `json:"labels"`

	// queryid
	Queryid string `json:"queryid,omitempty"`
}

// Validate validates this get histogram body
func (o *GetHistogramBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHistogramBody) validatePeriodStartFrom(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetHistogramBody) validatePeriodStartTo(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetHistogramBody) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for i := 0; i < len(o.Labels); i++ {
		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if o.Labels[i] != nil {
			if err := o.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHistogramBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistogramBody) UnmarshalBinary(b []byte) error {
	var res GetHistogramBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetHistogramDefaultBody get histogram default body
swagger:model GetHistogramDefaultBody
*/
type GetHistogramDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get histogram default body
func (o *GetHistogramDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHistogramDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetHistogram default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHistogramDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistogramDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetHistogramDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetHistogramOKBody HistogramReply is histogram items as a list.
swagger:model GetHistogramOKBody
*/
type GetHistogramOKBody struct {

	// histogram items
	HistogramItems []*HistogramItemsItems0 `json:"histogram_items"`
}

// Validate validates this get histogram OK body
func (o *GetHistogramOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHistogramItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHistogramOKBody) validateHistogramItems(formats strfmt.Registry) error {

	if swag.IsZero(o.HistogramItems) { // not required
		return nil
	}

	for i := 0; i < len(o.HistogramItems); i++ {
		if swag.IsZero(o.HistogramItems[i]) { // not required
			continue
		}

		if o.HistogramItems[i] != nil {
			if err := o.HistogramItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getHistogramOk" + "." + "histogram_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHistogramOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistogramOKBody) UnmarshalBinary(b []byte) error {
	var res GetHistogramOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*HistogramItemsItems0 HistogramItem represents one item in histogram.
swagger:model HistogramItemsItems0
*/
type HistogramItemsItems0 struct {

	// range
	Range string `json:"range,omitempty"`

	// frequency
	Frequency int64 `json:"frequency,omitempty"`
}

// Validate validates this histogram items items0
func (o *HistogramItemsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HistogramItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HistogramItemsItems0) UnmarshalBinary(b []byte) error {
	var res HistogramItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LabelsItems0 MapFieldEntry allows to pass labels/dimensions in form like {"server": ["db1", "db2"...]}.
swagger:model LabelsItems0
*/
type LabelsItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this labels items0
func (o *LabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LabelsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LabelsItems0) UnmarshalBinary(b []byte) error {
	var res LabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
