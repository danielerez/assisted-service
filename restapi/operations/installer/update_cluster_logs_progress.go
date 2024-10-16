// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateClusterLogsProgressHandlerFunc turns a function with the right signature into a update cluster logs progress handler
type UpdateClusterLogsProgressHandlerFunc func(UpdateClusterLogsProgressParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateClusterLogsProgressHandlerFunc) Handle(params UpdateClusterLogsProgressParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateClusterLogsProgressHandler interface for that can handle valid update cluster logs progress params
type UpdateClusterLogsProgressHandler interface {
	Handle(UpdateClusterLogsProgressParams, interface{}) middleware.Responder
}

// NewUpdateClusterLogsProgress creates a new http.Handler for the update cluster logs progress operation
func NewUpdateClusterLogsProgress(ctx *middleware.Context, handler UpdateClusterLogsProgressHandler) *UpdateClusterLogsProgress {
	return &UpdateClusterLogsProgress{Context: ctx, Handler: handler}
}

/* UpdateClusterLogsProgress swagger:route PUT /v1/clusters/{cluster_id}/logs_progress installer updateClusterLogsProgress

Update log collection state and progress.

*/
type UpdateClusterLogsProgress struct {
	Context *middleware.Context
	Handler UpdateClusterLogsProgressHandler
}

func (o *UpdateClusterLogsProgress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateClusterLogsProgressParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
