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

// User user
// swagger:model User
type User struct {

	// attributes
	Attributes *UserAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// Name of the resource type
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func UserWithDefaults(defaults client.Defaults) *User {
	return &User{

		Attributes: UserAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("User", "id"),

		OrganisationID: defaults.GetStrfmtUUID("User", "organisation_id"),

		Type: defaults.GetString("User", "type"),

		Version: defaults.GetInt64Ptr("User", "version"),
	}
}

func (m *User) WithAttributes(attributes UserAttributes) *User {

	m.Attributes = &attributes

	return m
}

func (m *User) WithoutAttributes() *User {
	m.Attributes = nil
	return m
}

func (m *User) WithID(id strfmt.UUID) *User {

	m.ID = id

	return m
}

func (m *User) WithOrganisationID(organisationID strfmt.UUID) *User {

	m.OrganisationID = organisationID

	return m
}

func (m *User) WithType(typeVar string) *User {

	m.Type = typeVar

	return m
}

func (m *User) WithVersion(version int64) *User {

	m.Version = &version

	return m
}

func (m *User) WithoutVersion() *User {
	m.Version = nil
	return m
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
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

func (m *User) validateAttributes(formats strfmt.Registry) error {

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

func (m *User) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *User) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// UserAttributes user attributes
// swagger:model UserAttributes
type UserAttributes struct {

	// client credential ids
	ClientCredentialIds []string `json:"client_credential_ids"`

	// Email address
	Email string `json:"email,omitempty"`

	// public key ids
	PublicKeyIds []strfmt.UUID `json:"public_key_ids"`

	// List of roles that this user belongs to
	RoleIds []strfmt.UUID `json:"role_ids"`

	// User name
	Username string `json:"username,omitempty"`
}

func UserAttributesWithDefaults(defaults client.Defaults) *UserAttributes {
	return &UserAttributes{

		ClientCredentialIds: make([]string, 0),

		Email: defaults.GetString("UserAttributes", "email"),

		PublicKeyIds: make([]strfmt.UUID, 0),

		RoleIds: make([]strfmt.UUID, 0),

		Username: defaults.GetString("UserAttributes", "username"),
	}
}

func (m *UserAttributes) WithClientCredentialIds(clientCredentialIds []string) *UserAttributes {

	m.ClientCredentialIds = clientCredentialIds

	return m
}

func (m *UserAttributes) WithEmail(email string) *UserAttributes {

	m.Email = email

	return m
}

func (m *UserAttributes) WithPublicKeyIds(publicKeyIds []strfmt.UUID) *UserAttributes {

	m.PublicKeyIds = publicKeyIds

	return m
}

func (m *UserAttributes) WithRoleIds(roleIds []strfmt.UUID) *UserAttributes {

	m.RoleIds = roleIds

	return m
}

func (m *UserAttributes) WithUsername(username string) *UserAttributes {

	m.Username = username

	return m
}

// Validate validates this user attributes
func (m *UserAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKeyIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAttributes) validatePublicKeyIds(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicKeyIds) { // not required
		return nil
	}

	for i := 0; i < len(m.PublicKeyIds); i++ {

		if err := validate.FormatOf("attributes"+"."+"public_key_ids"+"."+strconv.Itoa(i), "body", "uuid", m.PublicKeyIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UserAttributes) validateRoleIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleIds) { // not required
		return nil
	}

	for i := 0; i < len(m.RoleIds); i++ {

		if err := validate.FormatOf("attributes"+"."+"role_ids"+"."+strconv.Itoa(i), "body", "uuid", m.RoleIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAttributes) UnmarshalBinary(b []byte) error {
	var res UserAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *UserAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
