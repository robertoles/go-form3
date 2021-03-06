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

// FundingRequestAdmissionWithoutRelationships funding request admission without relationships
// swagger:model FundingRequestAdmissionWithoutRelationships
type FundingRequestAdmissionWithoutRelationships struct {

	// type
	// Required: true
	Type FundingRequestAdmissionResourceType `json:"type"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// version
	// Required: true
	Version *float64 `json:"version"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// attributes
	// Required: true
	Attributes *FundingRequestAdmissionAttributes `json:"attributes"`
}

func FundingRequestAdmissionWithoutRelationshipsWithDefaults(defaults client.Defaults) *FundingRequestAdmissionWithoutRelationships {
	return &FundingRequestAdmissionWithoutRelationships{

		// TODO Type: FundingRequestAdmissionResourceType,

		ID: defaults.GetStrfmtUUIDPtr("FundingRequestAdmissionWithoutRelationships", "id"),

		Version: defaults.GetFloat64Ptr("FundingRequestAdmissionWithoutRelationships", "version"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("FundingRequestAdmissionWithoutRelationships", "organisation_id"),

		CreatedOn: defaults.GetStrfmtDateTime("FundingRequestAdmissionWithoutRelationships", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("FundingRequestAdmissionWithoutRelationships", "modified_on"),

		Attributes: FundingRequestAdmissionAttributesWithDefaults(defaults),
	}
}

func (m *FundingRequestAdmissionWithoutRelationships) WithType(typeVar FundingRequestAdmissionResourceType) *FundingRequestAdmissionWithoutRelationships {

	m.Type = typeVar

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithID(id strfmt.UUID) *FundingRequestAdmissionWithoutRelationships {

	m.ID = &id

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithoutID() *FundingRequestAdmissionWithoutRelationships {
	m.ID = nil
	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithVersion(version float64) *FundingRequestAdmissionWithoutRelationships {

	m.Version = &version

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithoutVersion() *FundingRequestAdmissionWithoutRelationships {
	m.Version = nil
	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithOrganisationID(organisationID strfmt.UUID) *FundingRequestAdmissionWithoutRelationships {

	m.OrganisationID = &organisationID

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithoutOrganisationID() *FundingRequestAdmissionWithoutRelationships {
	m.OrganisationID = nil
	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithCreatedOn(createdOn strfmt.DateTime) *FundingRequestAdmissionWithoutRelationships {

	m.CreatedOn = createdOn

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithModifiedOn(modifiedOn strfmt.DateTime) *FundingRequestAdmissionWithoutRelationships {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithAttributes(attributes FundingRequestAdmissionAttributes) *FundingRequestAdmissionWithoutRelationships {

	m.Attributes = &attributes

	return m
}

func (m *FundingRequestAdmissionWithoutRelationships) WithoutAttributes() *FundingRequestAdmissionWithoutRelationships {
	m.Attributes = nil
	return m
}

// Validate validates this funding request admission without relationships
func (m *FundingRequestAdmissionWithoutRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FundingRequestAdmissionWithoutRelationships) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FundingRequestAdmissionWithoutRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FundingRequestAdmissionWithoutRelationships) UnmarshalBinary(b []byte) error {
	var res FundingRequestAdmissionWithoutRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FundingRequestAdmissionWithoutRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
