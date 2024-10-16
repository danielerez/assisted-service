// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadClusterISOOKCode is the HTTP code returned for type DownloadClusterISOOK
const DownloadClusterISOOKCode int = 200

/*DownloadClusterISOOK Success.

swagger:response downloadClusterISOOK
*/
type DownloadClusterISOOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadClusterISOOK creates DownloadClusterISOOK with default headers values
func NewDownloadClusterISOOK() *DownloadClusterISOOK {

	return &DownloadClusterISOOK{}
}

// WithPayload adds the payload to the download cluster i s o o k response
func (o *DownloadClusterISOOK) WithPayload(payload io.ReadCloser) *DownloadClusterISOOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o o k response
func (o *DownloadClusterISOOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadClusterISOMovedPermanentlyCode is the HTTP code returned for type DownloadClusterISOMovedPermanently
const DownloadClusterISOMovedPermanentlyCode int = 301

/*DownloadClusterISOMovedPermanently Redirect.

swagger:response downloadClusterISOMovedPermanently
*/
type DownloadClusterISOMovedPermanently struct {
	/*

	 */
	Location string `json:"Location"`
}

// NewDownloadClusterISOMovedPermanently creates DownloadClusterISOMovedPermanently with default headers values
func NewDownloadClusterISOMovedPermanently() *DownloadClusterISOMovedPermanently {

	return &DownloadClusterISOMovedPermanently{}
}

// WithLocation adds the location to the download cluster i s o moved permanently response
func (o *DownloadClusterISOMovedPermanently) WithLocation(location string) *DownloadClusterISOMovedPermanently {
	o.Location = location
	return o
}

// SetLocation sets the location to the download cluster i s o moved permanently response
func (o *DownloadClusterISOMovedPermanently) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *DownloadClusterISOMovedPermanently) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(301)
}

// DownloadClusterISOBadRequestCode is the HTTP code returned for type DownloadClusterISOBadRequest
const DownloadClusterISOBadRequestCode int = 400

/*DownloadClusterISOBadRequest Error.

swagger:response downloadClusterISOBadRequest
*/
type DownloadClusterISOBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOBadRequest creates DownloadClusterISOBadRequest with default headers values
func NewDownloadClusterISOBadRequest() *DownloadClusterISOBadRequest {

	return &DownloadClusterISOBadRequest{}
}

// WithPayload adds the payload to the download cluster i s o bad request response
func (o *DownloadClusterISOBadRequest) WithPayload(payload *models.Error) *DownloadClusterISOBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o bad request response
func (o *DownloadClusterISOBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOUnauthorizedCode is the HTTP code returned for type DownloadClusterISOUnauthorized
const DownloadClusterISOUnauthorizedCode int = 401

/*DownloadClusterISOUnauthorized Unauthorized.

swagger:response downloadClusterISOUnauthorized
*/
type DownloadClusterISOUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterISOUnauthorized creates DownloadClusterISOUnauthorized with default headers values
func NewDownloadClusterISOUnauthorized() *DownloadClusterISOUnauthorized {

	return &DownloadClusterISOUnauthorized{}
}

// WithPayload adds the payload to the download cluster i s o unauthorized response
func (o *DownloadClusterISOUnauthorized) WithPayload(payload *models.InfraError) *DownloadClusterISOUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o unauthorized response
func (o *DownloadClusterISOUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOForbiddenCode is the HTTP code returned for type DownloadClusterISOForbidden
const DownloadClusterISOForbiddenCode int = 403

/*DownloadClusterISOForbidden Forbidden.

swagger:response downloadClusterISOForbidden
*/
type DownloadClusterISOForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterISOForbidden creates DownloadClusterISOForbidden with default headers values
func NewDownloadClusterISOForbidden() *DownloadClusterISOForbidden {

	return &DownloadClusterISOForbidden{}
}

// WithPayload adds the payload to the download cluster i s o forbidden response
func (o *DownloadClusterISOForbidden) WithPayload(payload *models.InfraError) *DownloadClusterISOForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o forbidden response
func (o *DownloadClusterISOForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISONotFoundCode is the HTTP code returned for type DownloadClusterISONotFound
const DownloadClusterISONotFoundCode int = 404

/*DownloadClusterISONotFound Error.

swagger:response downloadClusterISONotFound
*/
type DownloadClusterISONotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISONotFound creates DownloadClusterISONotFound with default headers values
func NewDownloadClusterISONotFound() *DownloadClusterISONotFound {

	return &DownloadClusterISONotFound{}
}

// WithPayload adds the payload to the download cluster i s o not found response
func (o *DownloadClusterISONotFound) WithPayload(payload *models.Error) *DownloadClusterISONotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o not found response
func (o *DownloadClusterISONotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISONotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOMethodNotAllowedCode is the HTTP code returned for type DownloadClusterISOMethodNotAllowed
const DownloadClusterISOMethodNotAllowedCode int = 405

/*DownloadClusterISOMethodNotAllowed Method Not Allowed.

swagger:response downloadClusterISOMethodNotAllowed
*/
type DownloadClusterISOMethodNotAllowed struct {
}

// NewDownloadClusterISOMethodNotAllowed creates DownloadClusterISOMethodNotAllowed with default headers values
func NewDownloadClusterISOMethodNotAllowed() *DownloadClusterISOMethodNotAllowed {

	return &DownloadClusterISOMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DownloadClusterISOMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// DownloadClusterISOConflictCode is the HTTP code returned for type DownloadClusterISOConflict
const DownloadClusterISOConflictCode int = 409

/*DownloadClusterISOConflict Error.

swagger:response downloadClusterISOConflict
*/
type DownloadClusterISOConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOConflict creates DownloadClusterISOConflict with default headers values
func NewDownloadClusterISOConflict() *DownloadClusterISOConflict {

	return &DownloadClusterISOConflict{}
}

// WithPayload adds the payload to the download cluster i s o conflict response
func (o *DownloadClusterISOConflict) WithPayload(payload *models.Error) *DownloadClusterISOConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o conflict response
func (o *DownloadClusterISOConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOInternalServerErrorCode is the HTTP code returned for type DownloadClusterISOInternalServerError
const DownloadClusterISOInternalServerErrorCode int = 500

/*DownloadClusterISOInternalServerError Error.

swagger:response downloadClusterISOInternalServerError
*/
type DownloadClusterISOInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOInternalServerError creates DownloadClusterISOInternalServerError with default headers values
func NewDownloadClusterISOInternalServerError() *DownloadClusterISOInternalServerError {

	return &DownloadClusterISOInternalServerError{}
}

// WithPayload adds the payload to the download cluster i s o internal server error response
func (o *DownloadClusterISOInternalServerError) WithPayload(payload *models.Error) *DownloadClusterISOInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o internal server error response
func (o *DownloadClusterISOInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
