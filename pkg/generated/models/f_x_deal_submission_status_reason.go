// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"
)

// FXDealSubmissionStatusReason Plain-text description of the status attribute. May provide additional information in case of a submission error.
// swagger:model FXDealSubmissionStatusReason
type FXDealSubmissionStatusReason string

// Validate validates this f x deal submission status reason
func (m FXDealSubmissionStatusReason) Validate(formats strfmt.Registry) error {
	return nil
}
func (m *FXDealSubmissionStatusReason) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
