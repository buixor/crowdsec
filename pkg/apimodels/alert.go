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

// Alert Alert
//
// swagger:model Alert
type Alert struct {

	// capacity
	Capacity int32 `json:"capacity,omitempty"`

	// decisions
	Decisions []*Decision `json:"decisions"`

	// the Meta of the events leading to overflow
	Events []*Event `json:"events"`

	// events count
	EventsCount int32 `json:"events_count,omitempty"`

	// labels
	Labels []string `json:"labels"`

	// leakspeed
	Leakspeed string `json:"leakspeed,omitempty"`

	// only relevant for APIL->APIC, ignored for cscli->APIL and crowdsec->APIL
	MachineID string `json:"machine_id,omitempty"`

	// a human readable message
	Message string `json:"message,omitempty"`

	// scenario
	Scenario string `json:"scenario,omitempty"`

	// scenario hash
	ScenarioHash string `json:"scenario_hash,omitempty"`

	// scenario version
	ScenarioVersion string `json:"scenario_version,omitempty"`

	// simulated
	Simulated bool `json:"simulated,omitempty"`

	// source
	Source *Source `json:"source,omitempty"`

	// start at
	StartAt string `json:"start_at,omitempty"`

	// stop at
	StopAt string `json:"stop_at,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) validateDecisions(formats strfmt.Registry) error {

	if swag.IsZero(m.Decisions) { // not required
		return nil
	}

	for i := 0; i < len(m.Decisions); i++ {
		if swag.IsZero(m.Decisions[i]) { // not required
			continue
		}

		if m.Decisions[i] != nil {
			if err := m.Decisions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("decisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Alert) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Alert) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
