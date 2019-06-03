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

// DirectDebitReturnSubmissionStatus direct debit return submission status
// swagger:model DirectDebitReturnSubmissionStatus
type DirectDebitReturnSubmissionStatus string

const (

	// DirectDebitReturnSubmissionStatusAccepted captures enum value "accepted"
	DirectDebitReturnSubmissionStatusAccepted DirectDebitReturnSubmissionStatus = "accepted"

	// DirectDebitReturnSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	DirectDebitReturnSubmissionStatusReleasedToGateway DirectDebitReturnSubmissionStatus = "released_to_gateway"

	// DirectDebitReturnSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	DirectDebitReturnSubmissionStatusDeliveryConfirmed DirectDebitReturnSubmissionStatus = "delivery_confirmed"

	// DirectDebitReturnSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	DirectDebitReturnSubmissionStatusDeliveryFailed DirectDebitReturnSubmissionStatus = "delivery_failed"

	// DirectDebitReturnSubmissionStatusSubmitted captures enum value "submitted"
	DirectDebitReturnSubmissionStatusSubmitted DirectDebitReturnSubmissionStatus = "submitted"

	// DirectDebitReturnSubmissionStatusValidationPending captures enum value "validation_pending"
	DirectDebitReturnSubmissionStatusValidationPending DirectDebitReturnSubmissionStatus = "validation_pending"

	// DirectDebitReturnSubmissionStatusValidationPassed captures enum value "validation_passed"
	DirectDebitReturnSubmissionStatusValidationPassed DirectDebitReturnSubmissionStatus = "validation_passed"

	// DirectDebitReturnSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	DirectDebitReturnSubmissionStatusQueuedForDelivery DirectDebitReturnSubmissionStatus = "queued_for_delivery"
)

// for schema
var directDebitReturnSubmissionStatusEnum []interface{}

func init() {
	var res []DirectDebitReturnSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","released_to_gateway","delivery_confirmed","delivery_failed","submitted","validation_pending","validation_passed","queued_for_delivery"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitReturnSubmissionStatusEnum = append(directDebitReturnSubmissionStatusEnum, v)
	}
}

func (m DirectDebitReturnSubmissionStatus) validateDirectDebitReturnSubmissionStatusEnum(path, location string, value DirectDebitReturnSubmissionStatus) error {
	if err := validate.Enum(path, location, value, directDebitReturnSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit return submission status
func (m DirectDebitReturnSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitReturnSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}