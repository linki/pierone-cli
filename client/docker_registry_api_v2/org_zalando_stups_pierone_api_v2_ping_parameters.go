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

// NewOrgZalandoStupsPieroneAPIV2PingParams creates a new OrgZalandoStupsPieroneAPIV2PingParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV2PingParams() *OrgZalandoStupsPieroneAPIV2PingParams {

	return &OrgZalandoStupsPieroneAPIV2PingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV2PingParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV2PingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV2PingParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV2PingParams {

	return &OrgZalandoStupsPieroneAPIV2PingParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV2PingParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v2 ping operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV2PingParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV2PingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
