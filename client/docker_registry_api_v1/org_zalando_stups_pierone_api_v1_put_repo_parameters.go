package docker_registry_api_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrgZalandoStupsPieroneAPIV1PutRepoParams creates a new OrgZalandoStupsPieroneAPIV1PutRepoParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV1PutRepoParams() *OrgZalandoStupsPieroneAPIV1PutRepoParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1PutRepoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV1PutRepoParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV1PutRepoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV1PutRepoParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV1PutRepoParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1PutRepoParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV1PutRepoParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v1 put repo operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV1PutRepoParams struct {

	/*Artifact*/
	Artifact string
	/*Team*/
	Team string

	timeout time.Duration
}

// WithArtifact adds the artifact to the org zalando stups pierone api v1 put repo params
func (o *OrgZalandoStupsPieroneAPIV1PutRepoParams) WithArtifact(artifact string) *OrgZalandoStupsPieroneAPIV1PutRepoParams {
	o.Artifact = artifact
	return o
}

// WithTeam adds the team to the org zalando stups pierone api v1 put repo params
func (o *OrgZalandoStupsPieroneAPIV1PutRepoParams) WithTeam(team string) *OrgZalandoStupsPieroneAPIV1PutRepoParams {
	o.Team = team
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV1PutRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param artifact
	if err := r.SetPathParam("artifact", o.Artifact); err != nil {
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
