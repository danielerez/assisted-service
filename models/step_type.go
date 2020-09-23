// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StepType step type
//
// swagger:model step-type
type StepType string

const (

	// StepTypeConnectivityCheck captures enum value "connectivity-check"
	StepTypeConnectivityCheck StepType = "connectivity-check"

	// StepTypeExecute captures enum value "execute"
	StepTypeExecute StepType = "execute"

	// StepTypeInventory captures enum value "inventory"
	StepTypeInventory StepType = "inventory"

	// StepTypeInstall captures enum value "install"
	StepTypeInstall StepType = "install"

	// StepTypeFreeNetworkAddresses captures enum value "free-network-addresses"
	StepTypeFreeNetworkAddresses StepType = "free-network-addresses"

	// StepTypeResetInstallation captures enum value "reset-installation"
	StepTypeResetInstallation StepType = "reset-installation"

	// StepTypeDhcpLeaseAllocate captures enum value "dhcp-lease-allocate"
	StepTypeDhcpLeaseAllocate StepType = "dhcp-lease-allocate"

	// StepTypeAPIVipConnectivityCheck captures enum value "api-vip-connectivity-check"
	StepTypeAPIVipConnectivityCheck StepType = "api-vip-connectivity-check"
)

// for schema
var stepTypeEnum []interface{}

func init() {
	var res []StepType
	if err := json.Unmarshal([]byte(`["connectivity-check","execute","inventory","install","free-network-addresses","reset-installation","dhcp-lease-allocate","api-vip-connectivity-check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stepTypeEnum = append(stepTypeEnum, v)
	}
}

func (m StepType) validateStepTypeEnum(path, location string, value StepType) error {
	if err := validate.EnumCase(path, location, value, stepTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this step type
func (m StepType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStepTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
