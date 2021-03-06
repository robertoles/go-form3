// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DirectDebitReturnAdmissionStatus direct debit return admission status
// swagger:model DirectDebitReturnAdmissionStatus
type DirectDebitReturnAdmissionStatus string

const (

	// DirectDebitReturnAdmissionStatusConfirmed captures enum value "confirmed"
	DirectDebitReturnAdmissionStatusConfirmed DirectDebitReturnAdmissionStatus = "confirmed"

	// DirectDebitReturnAdmissionStatusFailed captures enum value "failed"
	DirectDebitReturnAdmissionStatusFailed DirectDebitReturnAdmissionStatus = "failed"
)

// for schema
var directDebitReturnAdmissionStatusEnum []interface{}

func init() {
	var res []DirectDebitReturnAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitReturnAdmissionStatusEnum = append(directDebitReturnAdmissionStatusEnum, v)
	}
}

func (m DirectDebitReturnAdmissionStatus) validateDirectDebitReturnAdmissionStatusEnum(path, location string, value DirectDebitReturnAdmissionStatus) error {
	if err := validate.Enum(path, location, value, directDebitReturnAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit return admission status
func (m DirectDebitReturnAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitReturnAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
