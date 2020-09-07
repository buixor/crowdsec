// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDecisionForIPRangeResponse get decision for Ip range response
//
// swagger:model GetDecisionForIpRangeResponse
type GetDecisionForIPRangeResponse []*GetDecisionForIPRangeResponseItems0

// Validate validates this get decision for Ip range response
func (m GetDecisionForIPRangeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// GetDecisionForIPRangeResponseItems0 get decision for IP range response items0
//
// swagger:model GetDecisionForIPRangeResponseItems0
type GetDecisionForIPRangeResponseItems0 struct {

	// decision
	Decision *Decision `json:"decision,omitempty"`

	// decision id
	DecisionID string `json:"decision_id,omitempty"`
}

// Validate validates this get decision for IP range response items0
func (m *GetDecisionForIPRangeResponseItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDecisionForIPRangeResponseItems0) validateDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.Decision) { // not required
		return nil
	}

	if m.Decision != nil {
		if err := m.Decision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDecisionForIPRangeResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDecisionForIPRangeResponseItems0) UnmarshalBinary(b []byte) error {
	var res GetDecisionForIPRangeResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
