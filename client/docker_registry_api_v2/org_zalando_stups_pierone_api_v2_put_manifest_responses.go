package docker_registry_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgZalandoStupsPieroneAPIV2PutManifestReader is a Reader for the OrgZalandoStupsPieroneAPIV2PutManifest structure.
type OrgZalandoStupsPieroneAPIV2PutManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIV2PutManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIV2PutManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIV2PutManifestOK creates a OrgZalandoStupsPieroneAPIV2PutManifestOK with default headers values
func NewOrgZalandoStupsPieroneAPIV2PutManifestOK() *OrgZalandoStupsPieroneAPIV2PutManifestOK {
	return &OrgZalandoStupsPieroneAPIV2PutManifestOK{}
}

/*OrgZalandoStupsPieroneAPIV2PutManifestOK handles this case with default header values.

OK
*/
type OrgZalandoStupsPieroneAPIV2PutManifestOK struct {
}

func (o *OrgZalandoStupsPieroneAPIV2PutManifestOK) Error() string {
	return fmt.Sprintf("[PUT /v2/{team}/{artifact}/manifests/{name}][%d] orgZalandoStupsPieroneApiV2PutManifestOK ", 200)
}

func (o *OrgZalandoStupsPieroneAPIV2PutManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
