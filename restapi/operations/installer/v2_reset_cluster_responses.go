// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ResetClusterAcceptedCode is the HTTP code returned for type V2ResetClusterAccepted
const V2ResetClusterAcceptedCode int = 202

/*V2ResetClusterAccepted Success.

swagger:response v2ResetClusterAccepted
*/
type V2ResetClusterAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewV2ResetClusterAccepted creates V2ResetClusterAccepted with default headers values
func NewV2ResetClusterAccepted() *V2ResetClusterAccepted {

	return &V2ResetClusterAccepted{}
}

// WithPayload adds the payload to the v2 reset cluster accepted response
func (o *V2ResetClusterAccepted) WithPayload(payload *models.Cluster) *V2ResetClusterAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster accepted response
func (o *V2ResetClusterAccepted) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ResetClusterUnauthorizedCode is the HTTP code returned for type V2ResetClusterUnauthorized
const V2ResetClusterUnauthorizedCode int = 401

/*V2ResetClusterUnauthorized Unauthorized.

swagger:response v2ResetClusterUnauthorized
*/
type V2ResetClusterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ResetClusterUnauthorized creates V2ResetClusterUnauthorized with default headers values
func NewV2ResetClusterUnauthorized() *V2ResetClusterUnauthorized {

	return &V2ResetClusterUnauthorized{}
}

// WithPayload adds the payload to the v2 reset cluster unauthorized response
func (o *V2ResetClusterUnauthorized) WithPayload(payload *models.InfraError) *V2ResetClusterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster unauthorized response
func (o *V2ResetClusterUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ResetClusterForbiddenCode is the HTTP code returned for type V2ResetClusterForbidden
const V2ResetClusterForbiddenCode int = 403

/*V2ResetClusterForbidden Forbidden.

swagger:response v2ResetClusterForbidden
*/
type V2ResetClusterForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ResetClusterForbidden creates V2ResetClusterForbidden with default headers values
func NewV2ResetClusterForbidden() *V2ResetClusterForbidden {

	return &V2ResetClusterForbidden{}
}

// WithPayload adds the payload to the v2 reset cluster forbidden response
func (o *V2ResetClusterForbidden) WithPayload(payload *models.InfraError) *V2ResetClusterForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster forbidden response
func (o *V2ResetClusterForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ResetClusterNotFoundCode is the HTTP code returned for type V2ResetClusterNotFound
const V2ResetClusterNotFoundCode int = 404

/*V2ResetClusterNotFound Error.

swagger:response v2ResetClusterNotFound
*/
type V2ResetClusterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ResetClusterNotFound creates V2ResetClusterNotFound with default headers values
func NewV2ResetClusterNotFound() *V2ResetClusterNotFound {

	return &V2ResetClusterNotFound{}
}

// WithPayload adds the payload to the v2 reset cluster not found response
func (o *V2ResetClusterNotFound) WithPayload(payload *models.Error) *V2ResetClusterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster not found response
func (o *V2ResetClusterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ResetClusterMethodNotAllowedCode is the HTTP code returned for type V2ResetClusterMethodNotAllowed
const V2ResetClusterMethodNotAllowedCode int = 405

/*V2ResetClusterMethodNotAllowed Method Not Allowed.

swagger:response v2ResetClusterMethodNotAllowed
*/
type V2ResetClusterMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ResetClusterMethodNotAllowed creates V2ResetClusterMethodNotAllowed with default headers values
func NewV2ResetClusterMethodNotAllowed() *V2ResetClusterMethodNotAllowed {

	return &V2ResetClusterMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 reset cluster method not allowed response
func (o *V2ResetClusterMethodNotAllowed) WithPayload(payload *models.Error) *V2ResetClusterMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster method not allowed response
func (o *V2ResetClusterMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ResetClusterConflictCode is the HTTP code returned for type V2ResetClusterConflict
const V2ResetClusterConflictCode int = 409

/*V2ResetClusterConflict Error.

swagger:response v2ResetClusterConflict
*/
type V2ResetClusterConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ResetClusterConflict creates V2ResetClusterConflict with default headers values
func NewV2ResetClusterConflict() *V2ResetClusterConflict {

	return &V2ResetClusterConflict{}
}

// WithPayload adds the payload to the v2 reset cluster conflict response
func (o *V2ResetClusterConflict) WithPayload(payload *models.Error) *V2ResetClusterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster conflict response
func (o *V2ResetClusterConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ResetClusterInternalServerErrorCode is the HTTP code returned for type V2ResetClusterInternalServerError
const V2ResetClusterInternalServerErrorCode int = 500

/*V2ResetClusterInternalServerError Error.

swagger:response v2ResetClusterInternalServerError
*/
type V2ResetClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ResetClusterInternalServerError creates V2ResetClusterInternalServerError with default headers values
func NewV2ResetClusterInternalServerError() *V2ResetClusterInternalServerError {

	return &V2ResetClusterInternalServerError{}
}

// WithPayload adds the payload to the v2 reset cluster internal server error response
func (o *V2ResetClusterInternalServerError) WithPayload(payload *models.Error) *V2ResetClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 reset cluster internal server error response
func (o *V2ResetClusterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ResetClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
