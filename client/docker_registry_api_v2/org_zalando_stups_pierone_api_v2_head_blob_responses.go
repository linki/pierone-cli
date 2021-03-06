package docker_registry_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgZalandoStupsPieroneAPIV2HeadBlobReader is a Reader for the OrgZalandoStupsPieroneAPIV2HeadBlob structure.
type OrgZalandoStupsPieroneAPIV2HeadBlobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIV2HeadBlobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIV2HeadBlobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewOrgZalandoStupsPieroneAPIV2HeadBlobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIV2HeadBlobOK creates a OrgZalandoStupsPieroneAPIV2HeadBlobOK with default headers values
func NewOrgZalandoStupsPieroneAPIV2HeadBlobOK() *OrgZalandoStupsPieroneAPIV2HeadBlobOK {
	return &OrgZalandoStupsPieroneAPIV2HeadBlobOK{}
}

/*OrgZalandoStupsPieroneAPIV2HeadBlobOK handles this case with default header values.

exists
*/
type OrgZalandoStupsPieroneAPIV2HeadBlobOK struct {
}

func (o *OrgZalandoStupsPieroneAPIV2HeadBlobOK) Error() string {
	return fmt.Sprintf("[HEAD /v2/{team}/{artifact}/blobs/{digest}][%d] orgZalandoStupsPieroneApiV2HeadBlobOK ", 200)
}

func (o *OrgZalandoStupsPieroneAPIV2HeadBlobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgZalandoStupsPieroneAPIV2HeadBlobNotFound creates a OrgZalandoStupsPieroneAPIV2HeadBlobNotFound with default headers values
func NewOrgZalandoStupsPieroneAPIV2HeadBlobNotFound() *OrgZalandoStupsPieroneAPIV2HeadBlobNotFound {
	return &OrgZalandoStupsPieroneAPIV2HeadBlobNotFound{}
}

/*OrgZalandoStupsPieroneAPIV2HeadBlobNotFound handles this case with default header values.

FS layer blob does not exist
*/
type OrgZalandoStupsPieroneAPIV2HeadBlobNotFound struct {
}

func (o *OrgZalandoStupsPieroneAPIV2HeadBlobNotFound) Error() string {
	return fmt.Sprintf("[HEAD /v2/{team}/{artifact}/blobs/{digest}][%d] orgZalandoStupsPieroneApiV2HeadBlobNotFound ", 404)
}

func (o *OrgZalandoStupsPieroneAPIV2HeadBlobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
