// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// FXDealSubmissionResourceType f x deal submission resource type
// swagger:model FXDealSubmissionResourceType
type FXDealSubmissionResourceType string

const (

	// FXDealSubmissionResourceTypeFxDealSubmissions captures enum value "fx_deal_submissions"
	FXDealSubmissionResourceTypeFxDealSubmissions FXDealSubmissionResourceType = "fx_deal_submissions"
)

// for schema
var fXDealSubmissionResourceTypeEnum []interface{}

func init() {
	var res []FXDealSubmissionResourceType
	if err := json.Unmarshal([]byte(`["fx_deal_submissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fXDealSubmissionResourceTypeEnum = append(fXDealSubmissionResourceTypeEnum, v)
	}
}

func (m FXDealSubmissionResourceType) validateFXDealSubmissionResourceTypeEnum(path, location string, value FXDealSubmissionResourceType) error {
	if err := validate.Enum(path, location, value, fXDealSubmissionResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this f x deal submission resource type
func (m FXDealSubmissionResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFXDealSubmissionResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealSubmissionResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
