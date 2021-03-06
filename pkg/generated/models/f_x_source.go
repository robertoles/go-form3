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

// FXSource f x source
// swagger:model FXSource
type FXSource struct {

	// The amount of currency to sell
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// The currency the party wants to sell
	// Required: true
	Currency *string `json:"currency"`
}

func FXSourceWithDefaults(defaults client.Defaults) *FXSource {
	return &FXSource{

		Amount: defaults.GetString("FXSource", "amount"),

		Currency: defaults.GetStringPtr("FXSource", "currency"),
	}
}

func (m *FXSource) WithAmount(amount string) *FXSource {

	m.Amount = amount

	return m
}

func (m *FXSource) WithCurrency(currency string) *FXSource {

	m.Currency = &currency

	return m
}

func (m *FXSource) WithoutCurrency() *FXSource {
	m.Currency = nil
	return m
}

// Validate validates this f x source
func (m *FXSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXSource) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

func (m *FXSource) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FXSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXSource) UnmarshalBinary(b []byte) error {
	var res FXSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXSource) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
