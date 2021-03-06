package docker_registry_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgZalandoStupsPieroneAPIV2ListTagsReader is a Reader for the OrgZalandoStupsPieroneAPIV2ListTags structure.
type OrgZalandoStupsPieroneAPIV2ListTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIV2ListTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIV2ListTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIV2ListTagsOK creates a OrgZalandoStupsPieroneAPIV2ListTagsOK with default headers values
func NewOrgZalandoStupsPieroneAPIV2ListTagsOK() *OrgZalandoStupsPieroneAPIV2ListTagsOK {
	return &OrgZalandoStupsPieroneAPIV2ListTagsOK{}
}

/*OrgZalandoStupsPieroneAPIV2ListTagsOK handles this case with default header values.

list of tags
*/
type OrgZalandoStupsPieroneAPIV2ListTagsOK struct {
	Payload OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody
}

func (o *OrgZalandoStupsPieroneAPIV2ListTagsOK) Error() string {
	return fmt.Sprintf("[GET /v2/{team}/{artifact}/tags/list][%d] orgZalandoStupsPieroneApiV2ListTagsOK  %+v", 200, o.Payload)
}

func (o *OrgZalandoStupsPieroneAPIV2ListTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody org zalando stups pierone API v2 list tags o k body body

swagger:model OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody
*/
type OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody struct {

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* tags

	Required: true
	*/
	Tags []string `json:"tags"`
}

// Validate validates this org zalando stups pierone API v2 list tags o k body body
func (o *OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("orgZalandoStupsPieroneApiV2ListTagsOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *OrgZalandoStupsPieroneAPIV2ListTagsOKBodyBody) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("orgZalandoStupsPieroneApiV2ListTagsOK"+"."+"tags", "body", o.Tags); err != nil {
		return err
	}

	return nil
}
