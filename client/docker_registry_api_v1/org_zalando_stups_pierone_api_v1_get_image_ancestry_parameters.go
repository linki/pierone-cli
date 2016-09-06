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

// NewOrgZalandoStupsPieroneAPIV1GetImageAncestryParams creates a new OrgZalandoStupsPieroneAPIV1GetImageAncestryParams object
// with the default values initialized.
func NewOrgZalandoStupsPieroneAPIV1GetImageAncestryParams() *OrgZalandoStupsPieroneAPIV1GetImageAncestryParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1GetImageAncestryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgZalandoStupsPieroneAPIV1GetImageAncestryParamsWithTimeout creates a new OrgZalandoStupsPieroneAPIV1GetImageAncestryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgZalandoStupsPieroneAPIV1GetImageAncestryParamsWithTimeout(timeout time.Duration) *OrgZalandoStupsPieroneAPIV1GetImageAncestryParams {
	var ()
	return &OrgZalandoStupsPieroneAPIV1GetImageAncestryParams{

		timeout: timeout,
	}
}

/*OrgZalandoStupsPieroneAPIV1GetImageAncestryParams contains all the parameters to send to the API endpoint
for the org zalando stups pierone api v1 get image ancestry operation typically these are written to a http.Request
*/
type OrgZalandoStupsPieroneAPIV1GetImageAncestryParams struct {

	/*Image*/
	Image string

	timeout time.Duration
}

// WithImage adds the image to the org zalando stups pierone api v1 get image ancestry params
func (o *OrgZalandoStupsPieroneAPIV1GetImageAncestryParams) WithImage(image string) *OrgZalandoStupsPieroneAPIV1GetImageAncestryParams {
	o.Image = image
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *OrgZalandoStupsPieroneAPIV1GetImageAncestryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
