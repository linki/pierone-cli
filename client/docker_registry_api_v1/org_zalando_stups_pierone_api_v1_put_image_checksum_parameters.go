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

// NewOrgZalandoStupsPieroneAPIV1PutImageChecksumParams creates a new OrgZalandoStupsPieroneAPIV1PutImageChecksumParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV1PutImageChecksumParams() *OrgZalandoStupsPieroneAPIV1PutImageChecksumParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1PutImageChecksumParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV1PutImageChecksumParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV1PutImageChecksumParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV1PutImageChecksumParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV1PutImageChecksumParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1PutImageChecksumParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV1PutImageChecksumParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v1 put image checksum operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV1PutImageChecksumParams struct {

	/*Image*/
	Image string

	timeout time.Duration
}

// WithImage adds the image to the org zalando stups pierone api v1 put image checksum params
func (o *OrgZalandoStupsPieroneAPIV1PutImageChecksumParams) WithImage(image string) *OrgZalandoStupsPieroneAPIV1PutImageChecksumParams {
	o.Image = image
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV1PutImageChecksumParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param image
	if err := r.SetPathParam("image", o.Image); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
