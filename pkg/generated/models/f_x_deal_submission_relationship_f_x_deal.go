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

// FXDealSubmissionRelationshipFXDeal f x deal submission relationship f x deal
// swagger:model FXDealSubmissionRelationshipFXDeal
type FXDealSubmissionRelationshipFXDeal struct {

	// data
	// Required: true
	Data *FXDealSubmissionRelationshipFXDealData `json:"data"`
}

func FXDealSubmissionRelationshipFXDealWithDefaults(defaults client.Defaults) *FXDealSubmissionRelationshipFXDeal {
	return &FXDealSubmissionRelationshipFXDeal{

		Data: FXDealSubmissionRelationshipFXDealDataWithDefaults(defaults),
	}
}

func (m *FXDealSubmissionRelationshipFXDeal) WithData(data FXDealSubmissionRelationshipFXDealData) *FXDealSubmissionRelationshipFXDeal {

	m.Data = &data

	return m
}

func (m *FXDealSubmissionRelationshipFXDeal) WithoutData() *FXDealSubmissionRelationshipFXDeal {
	m.Data = nil
	return m
}

// Validate validates this f x deal submission relationship f x deal
func (m *FXDealSubmissionRelationshipFXDeal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealSubmissionRelationshipFXDeal) validateData(formats strfmt.Registry) error {

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
func (m *FXDealSubmissionRelationshipFXDeal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealSubmissionRelationshipFXDeal) UnmarshalBinary(b []byte) error {
	var res FXDealSubmissionRelationshipFXDeal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealSubmissionRelationshipFXDeal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// FXDealSubmissionRelationshipFXDealData f x deal submission relationship f x deal data
// swagger:model FXDealSubmissionRelationshipFXDealData
type FXDealSubmissionRelationshipFXDealData struct {

	// The FX Deal that is being submitted
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type FXDealResourceType `json:"type"`
}

func FXDealSubmissionRelationshipFXDealDataWithDefaults(defaults client.Defaults) *FXDealSubmissionRelationshipFXDealData {
	return &FXDealSubmissionRelationshipFXDealData{

		ID: defaults.GetStrfmtUUIDPtr("FXDealSubmissionRelationshipFXDealData", "id"),

		// TODO Type: FXDealResourceType,

	}
}

func (m *FXDealSubmissionRelationshipFXDealData) WithID(id strfmt.UUID) *FXDealSubmissionRelationshipFXDealData {

	m.ID = &id

	return m
}

func (m *FXDealSubmissionRelationshipFXDealData) WithoutID() *FXDealSubmissionRelationshipFXDealData {
	m.ID = nil
	return m
}

func (m *FXDealSubmissionRelationshipFXDealData) WithType(typeVar FXDealResourceType) *FXDealSubmissionRelationshipFXDealData {

	m.Type = typeVar

	return m
}

// Validate validates this f x deal submission relationship f x deal data
func (m *FXDealSubmissionRelationshipFXDealData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealSubmissionRelationshipFXDealData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("data"+"."+"id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FXDealSubmissionRelationshipFXDealData) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FXDealSubmissionRelationshipFXDealData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealSubmissionRelationshipFXDealData) UnmarshalBinary(b []byte) error {
	var res FXDealSubmissionRelationshipFXDealData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealSubmissionRelationshipFXDealData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
