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

// MandateAttributes mandate attributes
// swagger:model MandateAttributes
type MandateAttributes struct {

	// amount
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// beneficiary party
	BeneficiaryParty *MandateAttributesBeneficiaryParty `json:"beneficiary_party,omitempty"`

	// clearing id
	ClearingID string `json:"clearing_id,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// debtor party
	DebtorParty *MandateAttributesDebtorParty `json:"debtor_party,omitempty"`

	// frequency
	Frequency MandateFrequency `json:"frequency,omitempty"`

	// numeric reference
	NumericReference string `json:"numeric_reference,omitempty"`

	// payment scheme
	PaymentScheme string `json:"payment_scheme,omitempty"`

	// processing date
	// Format: date
	ProcessingDate *strfmt.Date `json:"processing_date,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`

	// scheme payment type
	SchemePaymentType string `json:"scheme_payment_type,omitempty"`

	// scheme processing date
	// Format: date
	SchemeProcessingDate *strfmt.Date `json:"scheme_processing_date,omitempty"`

	// unique scheme id
	UniqueSchemeID string `json:"unique_scheme_id,omitempty"`
}

func MandateAttributesWithDefaults(defaults client.Defaults) *MandateAttributes {
	return &MandateAttributes{

		Amount: defaults.GetString("MandateAttributes", "amount"),

		BeneficiaryParty: MandateAttributesBeneficiaryPartyWithDefaults(defaults),

		ClearingID: defaults.GetString("MandateAttributes", "clearing_id"),

		Currency: defaults.GetString("MandateAttributes", "currency"),

		DebtorParty: MandateAttributesDebtorPartyWithDefaults(defaults),

		// TODO Frequency: MandateFrequency,

		NumericReference: defaults.GetString("MandateAttributes", "numeric_reference"),

		PaymentScheme: defaults.GetString("MandateAttributes", "payment_scheme"),

		ProcessingDate: defaults.GetStrfmtDatePtr("MandateAttributes", "processing_date"),

		Reference: defaults.GetString("MandateAttributes", "reference"),

		SchemePaymentType: defaults.GetString("MandateAttributes", "scheme_payment_type"),

		SchemeProcessingDate: defaults.GetStrfmtDatePtr("MandateAttributes", "scheme_processing_date"),

		UniqueSchemeID: defaults.GetString("MandateAttributes", "unique_scheme_id"),
	}
}

func (m *MandateAttributes) WithAmount(amount string) *MandateAttributes {

	m.Amount = amount

	return m
}

func (m *MandateAttributes) WithBeneficiaryParty(beneficiaryParty MandateAttributesBeneficiaryParty) *MandateAttributes {

	m.BeneficiaryParty = &beneficiaryParty

	return m
}

func (m *MandateAttributes) WithoutBeneficiaryParty() *MandateAttributes {
	m.BeneficiaryParty = nil
	return m
}

func (m *MandateAttributes) WithClearingID(clearingID string) *MandateAttributes {

	m.ClearingID = clearingID

	return m
}

func (m *MandateAttributes) WithCurrency(currency string) *MandateAttributes {

	m.Currency = currency

	return m
}

func (m *MandateAttributes) WithDebtorParty(debtorParty MandateAttributesDebtorParty) *MandateAttributes {

	m.DebtorParty = &debtorParty

	return m
}

func (m *MandateAttributes) WithoutDebtorParty() *MandateAttributes {
	m.DebtorParty = nil
	return m
}

func (m *MandateAttributes) WithFrequency(frequency MandateFrequency) *MandateAttributes {

	m.Frequency = frequency

	return m
}

func (m *MandateAttributes) WithNumericReference(numericReference string) *MandateAttributes {

	m.NumericReference = numericReference

	return m
}

func (m *MandateAttributes) WithPaymentScheme(paymentScheme string) *MandateAttributes {

	m.PaymentScheme = paymentScheme

	return m
}

func (m *MandateAttributes) WithProcessingDate(processingDate strfmt.Date) *MandateAttributes {

	m.ProcessingDate = &processingDate

	return m
}

func (m *MandateAttributes) WithoutProcessingDate() *MandateAttributes {
	m.ProcessingDate = nil
	return m
}

func (m *MandateAttributes) WithReference(reference string) *MandateAttributes {

	m.Reference = reference

	return m
}

func (m *MandateAttributes) WithSchemePaymentType(schemePaymentType string) *MandateAttributes {

	m.SchemePaymentType = schemePaymentType

	return m
}

func (m *MandateAttributes) WithSchemeProcessingDate(schemeProcessingDate strfmt.Date) *MandateAttributes {

	m.SchemeProcessingDate = &schemeProcessingDate

	return m
}

func (m *MandateAttributes) WithoutSchemeProcessingDate() *MandateAttributes {
	m.SchemeProcessingDate = nil
	return m
}

func (m *MandateAttributes) WithUniqueSchemeID(uniqueSchemeID string) *MandateAttributes {

	m.UniqueSchemeID = uniqueSchemeID

	return m
}

// Validate validates this mandate attributes
func (m *MandateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAttributes) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryParty) { // not required
		return nil
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary_party")
			}
			return err
		}
	}

	return nil
}

func (m *MandateAttributes) validateDebtorParty(formats strfmt.Registry) error {

	if swag.IsZero(m.DebtorParty) { // not required
		return nil
	}

	if m.DebtorParty != nil {
		if err := m.DebtorParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor_party")
			}
			return err
		}
	}

	return nil
}

func (m *MandateAttributes) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := m.Frequency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency")
		}
		return err
	}

	return nil
}

func (m *MandateAttributes) validateProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("processing_date", "body", "date", m.ProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateAttributes) validateSchemeProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("scheme_processing_date", "body", "date", m.SchemeProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAttributes) UnmarshalBinary(b []byte) error {
	var res MandateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
