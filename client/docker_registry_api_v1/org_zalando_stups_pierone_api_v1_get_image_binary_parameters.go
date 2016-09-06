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

// NewOrgZalandoStupsPieroneAPIV1GetImageBinaryParams creates a new OrgZalandoStupsPieroneAPIV1GetImageBinaryParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV1GetImageBinaryParams() *OrgZalandoStupsPieroneAPIV1GetImageBinaryParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1GetImageBinaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV1GetImageBinaryParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV1GetImageBinaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV1GetImageBinaryParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV1GetImageBinaryParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1GetImageBinaryParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV1GetImageBinaryParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v1 get image binary operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV1GetImageBinaryParams struct {

	/*Image*/
	Image string

	timeout time.Duration
}

// WithImage adds the image to the org zalando stups pierone api v1 get image binary params
func (o *OrgZalandoStupsPieroneAPIV1GetImageBinaryParams) WithImage(image string) *OrgZalandoStupsPieroneAPIV1GetImageBinaryParams {
	o.Image = image
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV1GetImageBinaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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