package docker_registry_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgZalandoStupsPieroneAPIV2PingReader is a Reader for the OrgZalandoStupsPieroneAPIV2Ping structure.
type OrgZalandoStupsPieroneAPIV2PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIV2PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIV2PingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIV2PingOK creates a OrgZalandoStupsPieroneAPIV2PingOK with default headers values
func NewOrgZalandoStupsPieroneAPIV2PingOK() *OrgZalandoStupsPieroneAPIV2PingOK {
	return &OrgZalandoStupsPieroneAPIV2PingOK{}
}

/*OrgZalandoStupsPieroneAPIV2PingOK handles this case with default header values.

v2 is available
*/
type OrgZalandoStupsPieroneAPIV2PingOK struct {
}

func (o *OrgZalandoStupsPieroneAPIV2PingOK) Error() string {
	return fmt.Sprintf("[GET /v2/][%d] orgZalandoStupsPieroneApiV2PingOK ", 200)
}

func (o *OrgZalandoStupsPieroneAPIV2PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
