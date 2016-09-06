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

// NewOrgZalandoStupsPieroneAPIV2PutManifestParams creates a new OrgZalandoStupsPieroneAPIV2PutManifestParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV2PutManifestParams() *OrgZalandoStupsPieroneAPIV2PutManifestParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV2PutManifestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV2PutManifestParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV2PutManifestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV2PutManifestParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV2PutManifestParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV2PutManifestParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV2PutManifestParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v2 put manifest operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV2PutManifestParams struct {

	/*Artifact*/
	Artifact string
	/*Data*/
	Data string
	/*Name
	  Manifest name

	*/
	Name string
	/*Team*/
	Team string

	timeout time.Duration
}

// WithArtifact adds the artifact to the org zalando stups pierone api v2 put manifest params
func (o *OrgZalandoStupsPieroneAPIV2PutManifestParams) WithArtifact(artifact string) *OrgZalandoStupsPieroneAPIV2PutManifestParams {
	o.Artifact = artifact
	return o
}

// WithData adds the data to the org zalando stups pierone api v2 put manifest params
func (o *OrgZalandoStupsPieroneAPIV2PutManifestParams) WithData(data string) *OrgZalandoStupsPieroneAPIV2PutManifestParams {
	o.Data = data
	return o
}

// WithName adds the name to the org zalando stups pierone api v2 put manifest params
func (o *OrgZalandoStupsPieroneAPIV2PutManifestParams) WithName(name string) *OrgZalandoStupsPieroneAPIV2PutManifestParams {
	o.Name = name
	return o
}

// WithTeam adds the team to the org zalando stups pierone api v2 put manifest params
func (o *OrgZalandoStupsPieroneAPIV2PutManifestParams) WithTeam(team string) *OrgZalandoStupsPieroneAPIV2PutManifestParams {
	o.Team = team
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV2PutManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param artifact
	if err := r.SetPathParam("artifact", o.Artifact); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
