// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DirectDebitSubmissionRelationships direct debit submission relationships
// swagger:model DirectDebitSubmissionRelationships
type DirectDebitSubmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitSubmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`
}

func DirectDebitSubmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitSubmissionRelationships {
	return &DirectDebitSubmissionRelationships{

		DirectDebit: DirectDebitSubmissionRelationshipsDirectDebitWithDefaults(defaults),
	}
}

func (m *DirectDebitSubmissionRelationships) WithDirectDebit(directDebit DirectDebitSubmissionRelationshipsDirectDebit) *DirectDebitSubmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitSubmissionRelationships) WithoutDirectDebit() *DirectDebitSubmissionRelationships {
	m.DirectDebit = nil
	return m
}

// Validate validates this direct debit submission relationships
func (m *DirectDebitSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitSubmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitSubmissionRelationshipsDirectDebit direct debit submission relationships direct debit
// swagger:model DirectDebitSubmissionRelationshipsDirectDebit
type DirectDebitSubmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitSubmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitSubmissionRelationshipsDirectDebit {
	return &DirectDebitSubmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitSubmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitSubmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit submission relationships direct debit
func (m *DirectDebitSubmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitSubmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitSubmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitSubmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitSubmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitSubmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
