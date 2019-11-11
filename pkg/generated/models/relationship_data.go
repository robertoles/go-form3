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

// RelationshipData relationship data
// swagger:model RelationshipData
type RelationshipData struct {

	// ID of the referenced resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the referenced resource type
	Type string `json:"type,omitempty"`
}

func RelationshipDataWithDefaults(defaults client.Defaults) *RelationshipData {
	return &RelationshipData{

		ID: defaults.GetStrfmtUUID("RelationshipData", "id"),

		Type: defaults.GetString("RelationshipData", "type"),
	}
}

func (m *RelationshipData) WithID(id strfmt.UUID) *RelationshipData {

	m.ID = id

	return m
}

func (m *RelationshipData) WithType(typeVar string) *RelationshipData {

	m.Type = typeVar

	return m
}

// Validate validates this relationship data
func (m *RelationshipData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipData) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipData) UnmarshalBinary(b []byte) error {
	var res RelationshipData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
