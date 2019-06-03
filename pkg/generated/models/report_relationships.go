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

// ReportRelationships report relationships
// swagger:model ReportRelationships
type ReportRelationships struct {

	// report admission
	ReportAdmission *ReportRelationshipsReportAdmission `json:"report_admission,omitempty"`
}

func ReportRelationshipsWithDefaults(defaults client.Defaults) *ReportRelationships {
	return &ReportRelationships{

		ReportAdmission: ReportRelationshipsReportAdmissionWithDefaults(defaults),
	}
}

func (m *ReportRelationships) WithReportAdmission(reportAdmission ReportRelationshipsReportAdmission) *ReportRelationships {

	m.ReportAdmission = &reportAdmission

	return m
}

func (m *ReportRelationships) WithoutReportAdmission() *ReportRelationships {
	m.ReportAdmission = nil
	return m
}

// Validate validates this report relationships
func (m *ReportRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRelationships) validateReportAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportAdmission) { // not required
		return nil
	}

	if m.ReportAdmission != nil {
		if err := m.ReportAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_admission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRelationships) UnmarshalBinary(b []byte) error {
	var res ReportRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReportRelationshipsReportAdmission report relationships report admission
// swagger:model ReportRelationshipsReportAdmission
type ReportRelationshipsReportAdmission struct {

	// data
	Data []*ReportAdmission `json:"data"`
}

func ReportRelationshipsReportAdmissionWithDefaults(defaults client.Defaults) *ReportRelationshipsReportAdmission {
	return &ReportRelationshipsReportAdmission{

		Data: make([]*ReportAdmission, 0),
	}
}

func (m *ReportRelationshipsReportAdmission) WithData(data []*ReportAdmission) *ReportRelationshipsReportAdmission {

	m.Data = data

	return m
}

// Validate validates this report relationships report admission
func (m *ReportRelationshipsReportAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRelationshipsReportAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("report_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRelationshipsReportAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRelationshipsReportAdmission) UnmarshalBinary(b []byte) error {
	var res ReportRelationshipsReportAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRelationshipsReportAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}