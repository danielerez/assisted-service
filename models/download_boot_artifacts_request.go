// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DownloadBootArtifactsRequest Information sent to the agent for downloading artifacts to boot a host into discovery.
//
// swagger:model download_boot_artifacts_request
type DownloadBootArtifactsRequest struct {

	// The base directory on the host that contains the /boot folder. The host will download boot
	// artifacts into a folder in this directory.
	// Required: true
	HostFsMountDir *string `json:"host_fs_mount_dir"`

	// URL address to download the initrd.
	// Required: true
	InitrdURL *string `json:"initrd_url"`

	// URL address to download the kernel.
	// Required: true
	KernelURL *string `json:"kernel_url"`

	// URL address to download the rootfs.
	// Required: true
	RootfsURL *string `json:"rootfs_url"`
}

// Validate validates this download boot artifacts request
func (m *DownloadBootArtifactsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostFsMountDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitrdURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernelURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootfsURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DownloadBootArtifactsRequest) validateHostFsMountDir(formats strfmt.Registry) error {

	if err := validate.Required("host_fs_mount_dir", "body", m.HostFsMountDir); err != nil {
		return err
	}

	return nil
}

func (m *DownloadBootArtifactsRequest) validateInitrdURL(formats strfmt.Registry) error {

	if err := validate.Required("initrd_url", "body", m.InitrdURL); err != nil {
		return err
	}

	return nil
}

func (m *DownloadBootArtifactsRequest) validateKernelURL(formats strfmt.Registry) error {

	if err := validate.Required("kernel_url", "body", m.KernelURL); err != nil {
		return err
	}

	return nil
}

func (m *DownloadBootArtifactsRequest) validateRootfsURL(formats strfmt.Registry) error {

	if err := validate.Required("rootfs_url", "body", m.RootfsURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this download boot artifacts request based on context it is used
func (m *DownloadBootArtifactsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DownloadBootArtifactsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadBootArtifactsRequest) UnmarshalBinary(b []byte) error {
	var res DownloadBootArtifactsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
