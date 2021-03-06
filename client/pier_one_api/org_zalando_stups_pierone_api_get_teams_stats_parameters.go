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

// NewOrgZalandoStupsPieroneAPIGetTeamsStatsParams creates a new OrgZalandoStupsPieroneAPIGetTeamsStatsParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIGetTeamsStatsParams() *OrgZalandoStupsPieroneAPIGetTeamsStatsParams {

	return &OrgZalandoStupsPieroneAPIGetTeamsStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIGetTeamsStatsParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIGetTeamsStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIGetTeamsStatsParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIGetTeamsStatsParams {

	return &OrgZalandoStupsPieroneAPIGetTeamsStatsParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIGetTeamsStatsParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api get teams stats operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIGetTeamsStatsParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIGetTeamsStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
