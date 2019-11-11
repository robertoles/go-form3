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

// ClaimSubmissionRelationships claim submission relationships
// swagger:model ClaimSubmissionRelationships
type ClaimSubmissionRelationships struct {

	// claim
	Claim *ClaimSubmissionRelationshipsClaim `json:"claim,omitempty"`
}

func ClaimSubmissionRelationshipsWithDefaults(defaults client.Defaults) *ClaimSubmissionRelationships {
	return &ClaimSubmissionRelationships{

		Claim: ClaimSubmissionRelationshipsClaimWithDefaults(defaults),
	}
}

func (m *ClaimSubmissionRelationships) WithClaim(claim ClaimSubmissionRelationshipsClaim) *ClaimSubmissionRelationships {

	m.Claim = &claim

	return m
}

func (m *ClaimSubmissionRelationships) WithoutClaim() *ClaimSubmissionRelationships {
	m.Claim = nil
	return m
}

// Validate validates this claim submission relationships
func (m *ClaimSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaim(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimSubmissionRelationships) validateClaim(formats strfmt.Registry) error {

	if swag.IsZero(m.Claim) { // not required
		return nil
	}

	if m.Claim != nil {
		if err := m.Claim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claim")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ClaimSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimSubmissionRelationshipsClaim claim submission relationships claim
// swagger:model ClaimSubmissionRelationshipsClaim
type ClaimSubmissionRelationshipsClaim struct {

	// data
	Data []*Claim `json:"data"`
}

func ClaimSubmissionRelationshipsClaimWithDefaults(defaults client.Defaults) *ClaimSubmissionRelationshipsClaim {
	return &ClaimSubmissionRelationshipsClaim{

		Data: make([]*Claim, 0),
	}
}

func (m *ClaimSubmissionRelationshipsClaim) WithData(data []*Claim) *ClaimSubmissionRelationshipsClaim {

	m.Data = data

	return m
}

// Validate validates this claim submission relationships claim
func (m *ClaimSubmissionRelationshipsClaim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimSubmissionRelationshipsClaim) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("claim" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimSubmissionRelationshipsClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimSubmissionRelationshipsClaim) UnmarshalBinary(b []byte) error {
	var res ClaimSubmissionRelationshipsClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimSubmissionRelationshipsClaim) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
