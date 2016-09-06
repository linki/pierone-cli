package pier_one_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/linki/pierone-cli/models"
)

// OrgZalandoStupsPieroneAPIReadTagsReader is a Reader for the OrgZalandoStupsPieroneAPIReadTags structure.
type OrgZalandoStupsPieroneAPIReadTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIReadTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIReadTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIReadTagsOK creates a OrgZalandoStupsPieroneAPIReadTagsOK with default headers values
func NewOrgZalandoStupsPieroneAPIReadTagsOK() *OrgZalandoStupsPieroneAPIReadTagsOK {
	return &OrgZalandoStupsPieroneAPIReadTagsOK{}
}

/*OrgZalandoStupsPieroneAPIReadTagsOK handles this case with default header values.

Return list of tags
*/
type OrgZalandoStupsPieroneAPIReadTagsOK struct {
	Payload []*models.TagSummary
}

func (o *OrgZalandoStupsPieroneAPIReadTagsOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team}/artifacts/{artifact}/tags][%d] orgZalandoStupsPieroneApiReadTagsOK  %+v", 200, o.Payload)
}

func (o *OrgZalandoStupsPieroneAPIReadTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}