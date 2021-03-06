package docker_registry_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrgZalandoStupsPieroneAPIV2PatchUploadParams creates a new OrgZalandoStupsPieroneAPIV2PatchUploadParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV2PatchUploadParams() *OrgZalandoStupsPieroneAPIV2PatchUploadParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV2PatchUploadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV2PatchUploadParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV2PatchUploadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV2PatchUploadParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV2PatchUploadParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV2PatchUploadParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV2PatchUploadParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v2 patch upload operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV2PatchUploadParams struct {

	/*Artifact*/
	Artifact string
	/*Data*/
	Data interface{}
	/*Team*/
	Team string
	/*UUID
	  Upload UUID as described on https://docs.docker.com/registry/spec/api/#pushing-an-image

	*/
	UUID string

	timeout time.Duration
}

// WithArtifact adds the artifact to the org zalando stups pierone api v2 patch upload params
func (o *OrgZalandoStupsPieroneAPIV2PatchUploadParams) WithArtifact(artifact string) *OrgZalandoStupsPieroneAPIV2PatchUploadParams {
	o.Artifact = artifact
	return o
}

// WithData adds the data to the org zalando stups pierone api v2 patch upload params
func (o *OrgZalandoStupsPieroneAPIV2PatchUploadParams) WithData(data interface{}) *OrgZalandoStupsPieroneAPIV2PatchUploadParams {
	o.Data = data
	return o
}

// WithTeam adds the team to the org zalando stups pierone api v2 patch upload params
func (o *OrgZalandoStupsPieroneAPIV2PatchUploadParams) WithTeam(team string) *OrgZalandoStupsPieroneAPIV2PatchUploadParams {
	o.Team = team
	return o
}

// WithUUID adds the uuid to the org zalando stups pierone api v2 patch upload params
func (o *OrgZalandoStupsPieroneAPIV2PatchUploadParams) WithUUID(uuid string) *OrgZalandoStupsPieroneAPIV2PatchUploadParams {
	o.UUID = uuid
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV2PatchUploadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param artifact
	if err := r.SetPathParam("artifact", o.Artifact); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
