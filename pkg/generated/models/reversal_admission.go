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

// ReversalAdmission reversal admission
// swagger:model ReversalAdmission
type ReversalAdmission struct {

	// attributes
	Attributes *ReversalAdmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ReversalAdmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReversalAdmissionWithDefaults(defaults client.Defaults) *ReversalAdmission {
	return &ReversalAdmission{

		Attributes: ReversalAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReversalAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReversalAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReversalAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReversalAdmission", "organisation_id"),

		Relationships: ReversalAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReversalAdmission", "type"),

		Version: defaults.GetInt64Ptr("ReversalAdmission", "version"),
	}
}

func (m *ReversalAdmission) WithAttributes(attributes ReversalAdmissionAttributes) *ReversalAdmission {

	m.Attributes = &attributes

	return m
}

func (m *ReversalAdmission) WithoutAttributes() *ReversalAdmission {
	m.Attributes = nil
	return m
}

func (m *ReversalAdmission) WithCreatedOn(createdOn strfmt.DateTime) *ReversalAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReversalAdmission) WithoutCreatedOn() *ReversalAdmission {
	m.CreatedOn = nil
	return m
}

func (m *ReversalAdmission) WithID(id strfmt.UUID) *ReversalAdmission {

	m.ID = &id

	return m
}

func (m *ReversalAdmission) WithoutID() *ReversalAdmission {
	m.ID = nil
	return m
}

func (m *ReversalAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ReversalAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReversalAdmission) WithoutModifiedOn() *ReversalAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *ReversalAdmission) WithOrganisationID(organisationID strfmt.UUID) *ReversalAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReversalAdmission) WithoutOrganisationID() *ReversalAdmission {
	m.OrganisationID = nil
	return m
}

func (m *ReversalAdmission) WithRelationships(relationships ReversalAdmissionRelationships) *ReversalAdmission {

	m.Relationships = &relationships

	return m
}

func (m *ReversalAdmission) WithoutRelationships() *ReversalAdmission {
	m.Relationships = nil
	return m
}

func (m *ReversalAdmission) WithType(typeVar string) *ReversalAdmission {

	m.Type = typeVar

	return m
}

func (m *ReversalAdmission) WithVersion(version int64) *ReversalAdmission {

	m.Version = &version

	return m
}

func (m *ReversalAdmission) WithoutVersion() *ReversalAdmission {
	m.Version = nil
	return m
}

// Validate validates this reversal admission
func (m *ReversalAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
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

func (m *ReversalAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmission) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmission) UnmarshalBinary(b []byte) error {
	var res ReversalAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalAdmissionAttributes reversal admission attributes
// swagger:model ReversalAdmissionAttributes
type ReversalAdmissionAttributes struct {

	// Scheme-specific status code. Refer to scheme documentation where available.
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// Description of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`
}

func ReversalAdmissionAttributesWithDefaults(defaults client.Defaults) *ReversalAdmissionAttributes {
	return &ReversalAdmissionAttributes{

		SchemeStatusCode: defaults.GetString("ReversalAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReversalAdmissionAttributes", "scheme_status_code_description"),

		SourceGateway: defaults.GetString("ReversalAdmissionAttributes", "source_gateway"),
	}
}

func (m *ReversalAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReversalAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReversalAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReversalAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReversalAdmissionAttributes) WithSourceGateway(sourceGateway string) *ReversalAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

// Validate validates this reversal admission attributes
func (m *ReversalAdmissionAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalAdmissionRelationships reversal admission relationships
// swagger:model ReversalAdmissionRelationships
type ReversalAdmissionRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// reversal
	Reversal *RelationshipReversals `json:"reversal,omitempty"`
}

func ReversalAdmissionRelationshipsWithDefaults(defaults client.Defaults) *ReversalAdmissionRelationships {
	return &ReversalAdmissionRelationships{

		Payment: RelationshipPaymentsWithDefaults(defaults),

		Reversal: RelationshipReversalsWithDefaults(defaults),
	}
}

func (m *ReversalAdmissionRelationships) WithPayment(payment RelationshipPayments) *ReversalAdmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *ReversalAdmissionRelationships) WithoutPayment() *ReversalAdmissionRelationships {
	m.Payment = nil
	return m
}

func (m *ReversalAdmissionRelationships) WithReversal(reversal RelationshipReversals) *ReversalAdmissionRelationships {

	m.Reversal = &reversal

	return m
}

func (m *ReversalAdmissionRelationships) WithoutReversal() *ReversalAdmissionRelationships {
	m.Reversal = nil
	return m
}

// Validate validates this reversal admission relationships
func (m *ReversalAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalAdmissionRelationships) validateReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.Reversal) { // not required
		return nil
	}

	if m.Reversal != nil {
		if err := m.Reversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "reversal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
