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

// RelationshipsFxDealProperties relationships fx deal properties
// swagger:model RelationshipsFxDealProperties
type RelationshipsFxDealProperties struct {

	// The fx_deal which initiated this funding request
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type FxDealResourceType `json:"type"`
}

func RelationshipsFxDealPropertiesWithDefaults(defaults client.Defaults) *RelationshipsFxDealProperties {
	return &RelationshipsFxDealProperties{

		ID: defaults.GetStrfmtUUIDPtr("RelationshipsFxDealProperties", "id"),

		// TODO Type: FxDealResourceType,

	}
}

func (m *RelationshipsFxDealProperties) WithID(id strfmt.UUID) *RelationshipsFxDealProperties {

	m.ID = &id

	return m
}

func (m *RelationshipsFxDealProperties) WithoutID() *RelationshipsFxDealProperties {
	m.ID = nil
	return m
}

func (m *RelationshipsFxDealProperties) WithType(typeVar FxDealResourceType) *RelationshipsFxDealProperties {

	m.Type = typeVar

	return m
}

// Validate validates this relationships fx deal properties
func (m *RelationshipsFxDealProperties) Validate(formats strfmt.Registry) error {
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

func (m *RelationshipsFxDealProperties) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationshipsFxDealProperties) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipsFxDealProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsFxDealProperties) UnmarshalBinary(b []byte) error {
	var res RelationshipsFxDealProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsFxDealProperties) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
