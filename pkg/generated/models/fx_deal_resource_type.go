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

// FxDealResourceType fx deal resource type
// swagger:model FxDealResourceType
type FxDealResourceType string

const (

	// FxDealResourceTypeFxDeals captures enum value "fx_deals"
	FxDealResourceTypeFxDeals FxDealResourceType = "fx_deals"
)

// for schema
var fxDealResourceTypeEnum []interface{}

func init() {
	var res []FxDealResourceType
	if err := json.Unmarshal([]byte(`["fx_deals"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fxDealResourceTypeEnum = append(fxDealResourceTypeEnum, v)
	}
}

func (m FxDealResourceType) validateFxDealResourceTypeEnum(path, location string, value FxDealResourceType) error {
	if err := validate.Enum(path, location, value, fxDealResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this fx deal resource type
func (m FxDealResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFxDealResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FxDealResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
