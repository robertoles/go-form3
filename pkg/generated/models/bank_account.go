// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BankAccount bank account
// swagger:model BankAccount
type BankAccount struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// address
	Address []string `json:"address"`

	// bank id
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	BankIDCode string `json:"bank_id_code,omitempty"`

	// bank name
	BankName string `json:"bank_name,omitempty"`

	// bic
	Bic string `json:"bic,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// iban
	Iban string `json:"iban,omitempty"`

	// post code
	PostCode string `json:"post_code,omitempty"`
}

func BankAccountWithDefaults(defaults client.Defaults) *BankAccount {
	return &BankAccount{

		AccountNumber: defaults.GetString("BankAccount", "account_number"),

		Address: make([]string, 0),

		BankID: defaults.GetString("BankAccount", "bank_id"),

		BankIDCode: defaults.GetString("BankAccount", "bank_id_code"),

		BankName: defaults.GetString("BankAccount", "bank_name"),

		Bic: defaults.GetString("BankAccount", "bic"),

		City: defaults.GetString("BankAccount", "city"),

		Iban: defaults.GetString("BankAccount", "iban"),

		PostCode: defaults.GetString("BankAccount", "post_code"),
	}
}

func (m *BankAccount) WithAccountNumber(accountNumber string) *BankAccount {

	m.AccountNumber = accountNumber

	return m
}

func (m *BankAccount) WithAddress(address []string) *BankAccount {

	m.Address = address

	return m
}

func (m *BankAccount) WithBankID(bankID string) *BankAccount {

	m.BankID = bankID

	return m
}

func (m *BankAccount) WithBankIDCode(bankIDCode string) *BankAccount {

	m.BankIDCode = bankIDCode

	return m
}

func (m *BankAccount) WithBankName(bankName string) *BankAccount {

	m.BankName = bankName

	return m
}

func (m *BankAccount) WithBic(bic string) *BankAccount {

	m.Bic = bic

	return m
}

func (m *BankAccount) WithCity(city string) *BankAccount {

	m.City = city

	return m
}

func (m *BankAccount) WithIban(iban string) *BankAccount {

	m.Iban = iban

	return m
}

func (m *BankAccount) WithPostCode(postCode string) *BankAccount {

	m.PostCode = postCode

	return m
}

// Validate validates this bank account
func (m *BankAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BankAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankAccount) UnmarshalBinary(b []byte) error {
	var res BankAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BankAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
