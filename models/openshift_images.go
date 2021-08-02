// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenshiftImages openshift images
//
// swagger:model openshift-images
type OpenshiftImages struct {

	// The installation image of the OpenShift cluster.
	// Required: true
	ReleaseImage *string `json:"release_image"`

	// OCP version from the release metadata.
	// Required: true
	ReleaseVersion *string `json:"release_version"`

	// The base RHCOS image used for the discovery iso.
	// Required: true
	RhcosImage *string `json:"rhcos_image"`

	// The RHCOS rootfs url.
	// Required: true
	RhcosRootfs *string `json:"rhcos_rootfs"`

	// Build ID of the RHCOS image.
	// Required: true
	RhcosVersion *string `json:"rhcos_version"`
}

// Validate validates this openshift images
func (m *OpenshiftImages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleaseImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRhcosImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRhcosRootfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRhcosVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenshiftImages) validateReleaseImage(formats strfmt.Registry) error {

	if err := validate.Required("release_image", "body", m.ReleaseImage); err != nil {
		return err
	}

	return nil
}

func (m *OpenshiftImages) validateReleaseVersion(formats strfmt.Registry) error {

	if err := validate.Required("release_version", "body", m.ReleaseVersion); err != nil {
		return err
	}

	return nil
}

func (m *OpenshiftImages) validateRhcosImage(formats strfmt.Registry) error {

	if err := validate.Required("rhcos_image", "body", m.RhcosImage); err != nil {
		return err
	}

	return nil
}

func (m *OpenshiftImages) validateRhcosRootfs(formats strfmt.Registry) error {

	if err := validate.Required("rhcos_rootfs", "body", m.RhcosRootfs); err != nil {
		return err
	}

	return nil
}

func (m *OpenshiftImages) validateRhcosVersion(formats strfmt.Registry) error {

	if err := validate.Required("rhcos_version", "body", m.RhcosVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenshiftImages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenshiftImages) UnmarshalBinary(b []byte) error {
	var res OpenshiftImages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
