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

// FundingRequestAdmissionResponse funding request admission response
// swagger:model FundingRequestAdmissionResponse
type FundingRequestAdmissionResponse struct {

	// data
	// Required: true
	Data *FundingRequestAdmission `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func FundingRequestAdmissionResponseWithDefaults(defaults client.Defaults) *FundingRequestAdmissionResponse {
	return &FundingRequestAdmissionResponse{

		Data: FundingRequestAdmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *FundingRequestAdmissionResponse) WithData(data FundingRequestAdmission) *FundingRequestAdmissionResponse {

	m.Data = &data

	return m
}

func (m *FundingRequestAdmissionResponse) WithoutData() *FundingRequestAdmissionResponse {
	m.Data = nil
	return m
}

func (m *FundingRequestAdmissionResponse) WithLinks(links Links) *FundingRequestAdmissionResponse {

	m.Links = &links

	return m
}

func (m *FundingRequestAdmissionResponse) WithoutLinks() *FundingRequestAdmissionResponse {
	m.Links = nil
	return m
}

// Validate validates this funding request admission response
func (m *FundingRequestAdmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FundingRequestAdmissionResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *FundingRequestAdmissionResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FundingRequestAdmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FundingRequestAdmissionResponse) UnmarshalBinary(b []byte) error {
	var res FundingRequestAdmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FundingRequestAdmissionResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
