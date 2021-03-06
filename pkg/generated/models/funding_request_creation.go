// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FundingRequestCreation funding request creation
// swagger:model FundingRequestCreation
type FundingRequestCreation struct {

	// data
	// Required: true
	Data *NewFundingRequest `json:"data"`
}

func FundingRequestCreationWithDefaults(defaults client.Defaults) *FundingRequestCreation {
	return &FundingRequestCreation{

		Data: NewFundingRequestWithDefaults(defaults),
	}
}

func (m *FundingRequestCreation) WithData(data NewFundingRequest) *FundingRequestCreation {

	m.Data = &data

	return m
}

func (m *FundingRequestCreation) WithoutData() *FundingRequestCreation {
	m.Data = nil
	return m
}

// Validate validates this funding request creation
func (m *FundingRequestCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FundingRequestCreation) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FundingRequestCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FundingRequestCreation) UnmarshalBinary(b []byte) error {
	var res FundingRequestCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FundingRequestCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
