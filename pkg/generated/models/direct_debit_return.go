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
	"github.com/go-openapi/validate"
)

// DirectDebitReturn direct debit return
// swagger:model DirectDebitReturn
type DirectDebitReturn struct {

	// attributes
	Attributes *DirectDebitReturnAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *DirectDebitReturnRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitReturn {
	return &DirectDebitReturn{

		Attributes: DirectDebitReturnAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturn", "created_on"),

		ID: defaults.GetStrfmtUUID("DirectDebitReturn", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturn", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("DirectDebitReturn", "organisation_id"),

		Relationships: DirectDebitReturnRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReturn", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReturn", "version"),
	}
}

func (m *DirectDebitReturn) WithAttributes(attributes DirectDebitReturnAttributes) *DirectDebitReturn {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReturn) WithoutAttributes() *DirectDebitReturn {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReturn) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReturn {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReturn) WithoutCreatedOn() *DirectDebitReturn {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReturn) WithID(id strfmt.UUID) *DirectDebitReturn {

	m.ID = id

	return m
}

func (m *DirectDebitReturn) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReturn {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReturn) WithoutModifiedOn() *DirectDebitReturn {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReturn) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReturn {

	m.OrganisationID = organisationID

	return m
}

func (m *DirectDebitReturn) WithRelationships(relationships DirectDebitReturnRelationships) *DirectDebitReturn {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReturn) WithoutRelationships() *DirectDebitReturn {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReturn) WithType(typeVar string) *DirectDebitReturn {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReturn) WithVersion(version int64) *DirectDebitReturn {

	m.Version = &version

	return m
}

func (m *DirectDebitReturn) WithoutVersion() *DirectDebitReturn {
	m.Version = nil
	return m
}

// Validate validates this direct debit return
func (m *DirectDebitReturn) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReturn) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReturn) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReturn) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnAttributes direct debit return attributes
// swagger:model DirectDebitReturnAttributes
type DirectDebitReturnAttributes struct {

	// return code
	ReturnCode string `json:"return_code,omitempty"`

	// scheme transaction id
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`
}

func DirectDebitReturnAttributesWithDefaults(defaults client.Defaults) *DirectDebitReturnAttributes {
	return &DirectDebitReturnAttributes{

		ReturnCode: defaults.GetString("DirectDebitReturnAttributes", "return_code"),

		SchemeTransactionID: defaults.GetString("DirectDebitReturnAttributes", "scheme_transaction_id"),
	}
}

func (m *DirectDebitReturnAttributes) WithReturnCode(returnCode string) *DirectDebitReturnAttributes {

	m.ReturnCode = returnCode

	return m
}

func (m *DirectDebitReturnAttributes) WithSchemeTransactionID(schemeTransactionID string) *DirectDebitReturnAttributes {

	m.SchemeTransactionID = schemeTransactionID

	return m
}

// Validate validates this direct debit return attributes
func (m *DirectDebitReturnAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationships direct debit return relationships
// swagger:model DirectDebitReturnRelationships
type DirectDebitReturnRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReturnRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit return submission
	DirectDebitReturnSubmission *DirectDebitReturnRelationshipsDirectDebitReturnSubmission `json:"direct_debit_return_submission,omitempty"`
}

func DirectDebitReturnRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationships {
	return &DirectDebitReturnRelationships{

		DirectDebit: DirectDebitReturnRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReturnSubmission: DirectDebitReturnRelationshipsDirectDebitReturnSubmissionWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnRelationships) WithDirectDebit(directDebit DirectDebitReturnRelationshipsDirectDebit) *DirectDebitReturnRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReturnRelationships) WithoutDirectDebit() *DirectDebitReturnRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReturnRelationships) WithDirectDebitReturnSubmission(directDebitReturnSubmission DirectDebitReturnRelationshipsDirectDebitReturnSubmission) *DirectDebitReturnRelationships {

	m.DirectDebitReturnSubmission = &directDebitReturnSubmission

	return m
}

func (m *DirectDebitReturnRelationships) WithoutDirectDebitReturnSubmission() *DirectDebitReturnRelationships {
	m.DirectDebitReturnSubmission = nil
	return m
}

// Validate validates this direct debit return relationships
func (m *DirectDebitReturnRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturnSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnRelationships) validateDirectDebitReturnSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturnSubmission) { // not required
		return nil
	}

	if m.DirectDebitReturnSubmission != nil {
		if err := m.DirectDebitReturnSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationshipsDirectDebit direct debit return relationships direct debit
// swagger:model DirectDebitReturnRelationshipsDirectDebit
type DirectDebitReturnRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReturnRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationshipsDirectDebit {
	return &DirectDebitReturnRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReturnRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReturnRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit return relationships direct debit
func (m *DirectDebitReturnRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationshipsDirectDebitReturnSubmission direct debit return relationships direct debit return submission
// swagger:model DirectDebitReturnRelationshipsDirectDebitReturnSubmission
type DirectDebitReturnRelationshipsDirectDebitReturnSubmission struct {

	// data
	Data []*DirectDebitReturnSubmission `json:"data"`
}

func DirectDebitReturnRelationshipsDirectDebitReturnSubmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationshipsDirectDebitReturnSubmission {
	return &DirectDebitReturnRelationshipsDirectDebitReturnSubmission{

		Data: make([]*DirectDebitReturnSubmission, 0),
	}
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) WithData(data []*DirectDebitReturnSubmission) *DirectDebitReturnRelationshipsDirectDebitReturnSubmission {

	m.Data = data

	return m
}

// Validate validates this direct debit return relationships direct debit return submission
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationshipsDirectDebitReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}