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

// FundingRequestRelationships funding request relationships
// swagger:model FundingRequestRelationships
type FundingRequestRelationships struct {

	// contact account
	ContactAccount *RelationshipsContactAccount `json:"contact_account,omitempty"`

	// funding request admission
	FundingRequestAdmission *RelationshipFundingRequestAdmission `json:"funding_request_admission,omitempty"`

	// fx deal
	FxDeal *RelationshipsFxDeal `json:"fx_deal,omitempty"`

	// party
	// Required: true
	Party *RelationshipsParty `json:"party"`
}

func FundingRequestRelationshipsWithDefaults(defaults client.Defaults) *FundingRequestRelationships {
	return &FundingRequestRelationships{

		ContactAccount: RelationshipsContactAccountWithDefaults(defaults),

		FundingRequestAdmission: RelationshipFundingRequestAdmissionWithDefaults(defaults),

		FxDeal: RelationshipsFxDealWithDefaults(defaults),

		Party: RelationshipsPartyWithDefaults(defaults),
	}
}

func (m *FundingRequestRelationships) WithContactAccount(contactAccount RelationshipsContactAccount) *FundingRequestRelationships {

	m.ContactAccount = &contactAccount

	return m
}

func (m *FundingRequestRelationships) WithoutContactAccount() *FundingRequestRelationships {
	m.ContactAccount = nil
	return m
}

func (m *FundingRequestRelationships) WithFundingRequestAdmission(fundingRequestAdmission RelationshipFundingRequestAdmission) *FundingRequestRelationships {

	m.FundingRequestAdmission = &fundingRequestAdmission

	return m
}

func (m *FundingRequestRelationships) WithoutFundingRequestAdmission() *FundingRequestRelationships {
	m.FundingRequestAdmission = nil
	return m
}

func (m *FundingRequestRelationships) WithFxDeal(fxDeal RelationshipsFxDeal) *FundingRequestRelationships {

	m.FxDeal = &fxDeal

	return m
}

func (m *FundingRequestRelationships) WithoutFxDeal() *FundingRequestRelationships {
	m.FxDeal = nil
	return m
}

func (m *FundingRequestRelationships) WithParty(party RelationshipsParty) *FundingRequestRelationships {

	m.Party = &party

	return m
}

func (m *FundingRequestRelationships) WithoutParty() *FundingRequestRelationships {
	m.Party = nil
	return m
}

// Validate validates this funding request relationships
func (m *FundingRequestRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFundingRequestAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFxDeal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FundingRequestRelationships) validateContactAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactAccount) { // not required
		return nil
	}

	if m.ContactAccount != nil {
		if err := m.ContactAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact_account")
			}
			return err
		}
	}

	return nil
}

func (m *FundingRequestRelationships) validateFundingRequestAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.FundingRequestAdmission) { // not required
		return nil
	}

	if m.FundingRequestAdmission != nil {
		if err := m.FundingRequestAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("funding_request_admission")
			}
			return err
		}
	}

	return nil
}

func (m *FundingRequestRelationships) validateFxDeal(formats strfmt.Registry) error {

	if swag.IsZero(m.FxDeal) { // not required
		return nil
	}

	if m.FxDeal != nil {
		if err := m.FxDeal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fx_deal")
			}
			return err
		}
	}

	return nil
}

func (m *FundingRequestRelationships) validateParty(formats strfmt.Registry) error {

	if err := validate.Required("party", "body", m.Party); err != nil {
		return err
	}

	if m.Party != nil {
		if err := m.Party.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FundingRequestRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FundingRequestRelationships) UnmarshalBinary(b []byte) error {
	var res FundingRequestRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FundingRequestRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}