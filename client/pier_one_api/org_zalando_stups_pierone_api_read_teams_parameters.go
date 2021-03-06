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

// NewOrgZalandoStupsPieroneAPIReadTeamsParams creates a new OrgZalandoStupsPieroneAPIReadTeamsParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIReadTeamsParams() *OrgZalandoStupsPieroneAPIReadTeamsParams {

	return &OrgZalandoStupsPieroneAPIReadTeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIReadTeamsParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIReadTeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIReadTeamsParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIReadTeamsParams {

	return &OrgZalandoStupsPieroneAPIReadTeamsParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIReadTeamsParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api read teams operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIReadTeamsParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIReadTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
