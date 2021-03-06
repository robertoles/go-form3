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
)

// QueryResponseRelationships query response relationships
// swagger:model QueryResponseRelationships
type QueryResponseRelationships struct {

	// query
	Query *RelationshipsQuery `json:"query,omitempty"`

	// query submission
	QuerySubmission *RelationshipQuerySubmission `json:"query_submission,omitempty"`
}

func QueryResponseRelationshipsWithDefaults(defaults client.Defaults) *QueryResponseRelationships {
	return &QueryResponseRelationships{

		Query: RelationshipsQueryWithDefaults(defaults),

		QuerySubmission: RelationshipQuerySubmissionWithDefaults(defaults),
	}
}

func (m *QueryResponseRelationships) WithQuery(query RelationshipsQuery) *QueryResponseRelationships {

	m.Query = &query

	return m
}

func (m *QueryResponseRelationships) WithoutQuery() *QueryResponseRelationships {
	m.Query = nil
	return m
}

func (m *QueryResponseRelationships) WithQuerySubmission(querySubmission RelationshipQuerySubmission) *QueryResponseRelationships {

	m.QuerySubmission = &querySubmission

	return m
}

func (m *QueryResponseRelationships) WithoutQuerySubmission() *QueryResponseRelationships {
	m.QuerySubmission = nil
	return m
}

// Validate validates this query response relationships
func (m *QueryResponseRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuerySubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseRelationships) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResponseRelationships) validateQuerySubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.QuerySubmission) { // not required
		return nil
	}

	if m.QuerySubmission != nil {
		if err := m.QuerySubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseRelationships) UnmarshalBinary(b []byte) error {
	var res QueryResponseRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
