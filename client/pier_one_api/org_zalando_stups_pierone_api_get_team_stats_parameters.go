package pier_one_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrgZalandoStupsPieroneAPIGetTeamStatsParams creates a new OrgZalandoStupsPieroneAPIGetTeamStatsParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIGetTeamStatsParams() *OrgZalandoStupsPieroneAPIGetTeamStatsParams {
	var ()
	return &OrgZalandoStupsPieroneAPIGetTeamStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIGetTeamStatsParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIGetTeamStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIGetTeamStatsParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIGetTeamStatsParams {
	var ()
	return &OrgZalandoStupsPieroneAPIGetTeamStatsParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIGetTeamStatsParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api get team stats operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIGetTeamStatsParams struct {

	/*Team*/
	Team string

	timeout time.Duration
}

// WithTeam adds the team to the org zalando stups pierone api get team stats params
func (o *OrgZalandoStupsPieroneAPIGetTeamStatsParams) WithTeam(team string) *OrgZalandoStupsPieroneAPIGetTeamStatsParams {
	o.Team = team
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIGetTeamStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
