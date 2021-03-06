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

// FXDealSubmissionUpdatedEvent f x deal submission updated event
// swagger:model FXDealSubmissionUpdatedEvent
type FXDealSubmissionUpdatedEvent struct {

	// data
	// Required: true
	Data *FXDealSubmissionUpdatedEventData `json:"data"`
}

func FXDealSubmissionUpdatedEventWithDefaults(defaults client.Defaults) *FXDealSubmissionUpdatedEvent {
	return &FXDealSubmissionUpdatedEvent{

		Data: FXDealSubmissionUpdatedEventDataWithDefaults(defaults),
	}
}

func (m *FXDealSubmissionUpdatedEvent) WithData(data FXDealSubmissionUpdatedEventData) *FXDealSubmissionUpdatedEvent {

	m.Data = &data

	return m
}

func (m *FXDealSubmissionUpdatedEvent) WithoutData() *FXDealSubmissionUpdatedEvent {
	m.Data = nil
	return m
}

// Validate validates this f x deal submission updated event
func (m *FXDealSubmissionUpdatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealSubmissionUpdatedEvent) validateData(formats strfmt.Registry) error {

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
func (m *FXDealSubmissionUpdatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealSubmissionUpdatedEvent) UnmarshalBinary(b []byte) error {
	var res FXDealSubmissionUpdatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealSubmissionUpdatedEvent) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// FXDealSubmissionUpdatedEventData f x deal submission updated event data
// swagger:model FXDealSubmissionUpdatedEventData
type FXDealSubmissionUpdatedEventData struct {

	// type
	// Required: true
	Type FXDealSubmissionResourceType `json:"type"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// attributes
	// Required: true
	Attributes *FXDealSubmissionAttributes `json:"attributes"`

	// relationships
	// Required: true
	Relationships *FXDealSubmissionRelationships `json:"relationships"`
}

func FXDealSubmissionUpdatedEventDataWithDefaults(defaults client.Defaults) *FXDealSubmissionUpdatedEventData {
	return &FXDealSubmissionUpdatedEventData{

		// TODO Type: FXDealSubmissionResourceType,

		ID: defaults.GetStrfmtUUIDPtr("FXDealSubmissionUpdatedEventData", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("FXDealSubmissionUpdatedEventData", "organisation_id"),

		Attributes: FXDealSubmissionAttributesWithDefaults(defaults),

		Relationships: FXDealSubmissionRelationshipsWithDefaults(defaults),
	}
}

func (m *FXDealSubmissionUpdatedEventData) WithType(typeVar FXDealSubmissionResourceType) *FXDealSubmissionUpdatedEventData {

	m.Type = typeVar

	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithID(id strfmt.UUID) *FXDealSubmissionUpdatedEventData {

	m.ID = &id

	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithoutID() *FXDealSubmissionUpdatedEventData {
	m.ID = nil
	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithOrganisationID(organisationID strfmt.UUID) *FXDealSubmissionUpdatedEventData {

	m.OrganisationID = &organisationID

	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithoutOrganisationID() *FXDealSubmissionUpdatedEventData {
	m.OrganisationID = nil
	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithAttributes(attributes FXDealSubmissionAttributes) *FXDealSubmissionUpdatedEventData {

	m.Attributes = &attributes

	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithoutAttributes() *FXDealSubmissionUpdatedEventData {
	m.Attributes = nil
	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithRelationships(relationships FXDealSubmissionRelationships) *FXDealSubmissionUpdatedEventData {

	m.Relationships = &relationships

	return m
}

func (m *FXDealSubmissionUpdatedEventData) WithoutRelationships() *FXDealSubmissionUpdatedEventData {
	m.Relationships = nil
	return m
}

// Validate validates this f x deal submission updated event data
func (m *FXDealSubmissionUpdatedEventData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealSubmissionUpdatedEventData) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "type")
		}
		return err
	}

	return nil
}

func (m *FXDealSubmissionUpdatedEventData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("data"+"."+"id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FXDealSubmissionUpdatedEventData) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("data"+"."+"organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FXDealSubmissionUpdatedEventData) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

func (m *FXDealSubmissionUpdatedEventData) validateRelationships(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"relationships", "body", m.Relationships); err != nil {
		return err
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "relationships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FXDealSubmissionUpdatedEventData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealSubmissionUpdatedEventData) UnmarshalBinary(b []byte) error {
	var res FXDealSubmissionUpdatedEventData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealSubmissionUpdatedEventData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
