package host

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/filanov/bm-inventory/models"
	logutil "github.com/filanov/bm-inventory/pkg/log"

	"github.com/sirupsen/logrus"

	"github.com/filanov/stateswitch"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type transitionHandler struct {
	db  *gorm.DB
	log logrus.FieldLogger
}

////////////////////////////////////////////////////////////////////////////
// RegisterHost
////////////////////////////////////////////////////////////////////////////

type TransitionArgsRegisterHost struct {
	ctx                   context.Context
	discoveryAgentVersion string
}

func (th *transitionHandler) PostRegisterHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	sHost, ok := sw.(*stateHost)
	if !ok {
		return errors.New("PostRegisterHost incompatible type of StateSwitch")
	}
	params, ok := args.(*TransitionArgsRegisterHost)
	if !ok {
		return errors.New("PostRegisterHost invalid argument")
	}

	host := models.Host{}
	log := logutil.FromContext(params.ctx, th.log)

	// If host already exists
	if err := th.db.First(&host, "id = ? and cluster_id = ?", sHost.host.ID, sHost.host.ClusterID).Error; err == nil {
		currentState := swag.StringValue(host.Status)
		host.Status = sHost.host.Status

		// The reason for the double register is unknown (HW might have changed) -
		// so we reset the hw info and start the discovery process again.
		return updateHostStateWithParams(log, currentState, statusInfoDiscovering, &host, th.db,
			"hardware_info", "", "discovery_agent_version", params.discoveryAgentVersion)
	}

	sHost.host.StatusUpdatedAt = strfmt.DateTime(time.Now())
	log.Infof("Register new host %s cluster %s", sHost.host.ID.String(), sHost.host.ClusterID)
	return th.db.Create(sHost.host).Error
}

func (th *transitionHandler) PostRegisterDuringInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	sHost, ok := sw.(*stateHost)
	if !ok {
		return errors.New("RegisterNewHost incompatible type of StateSwitch")
	}
	params, ok := args.(*TransitionArgsRegisterHost)
	if !ok {
		return errors.New("PostRegisterDuringInstallation invalid argument")
	}
	return updateHostStateWithParams(logutil.FromContext(params.ctx, th.log), sHost.srcState,
		"Tried to register during installation", sHost.host, th.db)
}

////////////////////////////////////////////////////////////////////////////
// Installation failure
////////////////////////////////////////////////////////////////////////////

type TransitionArgsHostInstallationFailed struct {
	ctx    context.Context
	reason string
}

func (th *transitionHandler) PostHostInstallationFailed(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	sHost, ok := sw.(*stateHost)
	if !ok {
		return errors.New("HostInstallationFailed incompatible type of StateSwitch")
	}
	params, ok := args.(*TransitionArgsHostInstallationFailed)
	if !ok {
		return errors.New("HostInstallationFailed invalid argument")
	}
	return updateHostStateWithParams(logutil.FromContext(params.ctx, th.log), sHost.srcState,
		params.reason, sHost.host, th.db)
}

////////////////////////////////////////////////////////////////////////////
// Cancel Installation
////////////////////////////////////////////////////////////////////////////

type TransitionArgsCancelInstallation struct {
	ctx    context.Context
	reason string
	db     *gorm.DB
}

func (th *transitionHandler) PostCancelInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	sHost, ok := sw.(*stateHost)
	if !ok {
		return errors.New("PostCancelInstallation incompatible type of StateSwitch")
	}
	params, ok := args.(*TransitionArgsCancelInstallation)
	if !ok {
		return errors.New("PostCancelInstallation invalid argument")
	}
	if sHost.srcState == HostStatusError {
		return nil
	}
	return updateHostStateWithParams(logutil.FromContext(params.ctx, th.log), sHost.srcState,
		params.reason, sHost.host, params.db)
}

////////////////////////////////////////////////////////////////////////////
// Reset Host
////////////////////////////////////////////////////////////////////////////

type TransitionArgsResetHost struct {
	ctx    context.Context
	reason string
	db     *gorm.DB
}

func (th *transitionHandler) PostResetHost(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	sHost, ok := sw.(*stateHost)
	if !ok {
		return errors.New("PostResetHost incompatible type of StateSwitch")
	}
	params, ok := args.(*TransitionArgsResetHost)
	if !ok {
		return errors.New("PostResetHost invalid argument")
	}
	return updateHostStateWithParams(logutil.FromContext(params.ctx, th.log), sHost.srcState,
		params.reason, sHost.host, params.db)
}
