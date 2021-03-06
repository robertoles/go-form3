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

// FXRateResourceType f x rate resource type
// swagger:model FXRateResourceType
type FXRateResourceType string

const (

	// FXRateResourceTypeFxRates captures enum value "fx_rates"
	FXRateResourceTypeFxRates FXRateResourceType = "fx_rates"
)

// for schema
var fXRateResourceTypeEnum []interface{}

func init() {
	var res []FXRateResourceType
	if err := json.Unmarshal([]byte(`["fx_rates"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fXRateResourceTypeEnum = append(fXRateResourceTypeEnum, v)
	}
}

func (m FXRateResourceType) validateFXRateResourceTypeEnum(path, location string, value FXRateResourceType) error {
	if err := validate.Enum(path, location, value, fXRateResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this f x rate resource type
func (m FXRateResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFXRateResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXRateResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
