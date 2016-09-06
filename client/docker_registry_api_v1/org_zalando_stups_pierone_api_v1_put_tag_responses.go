package docker_registry_api_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgZalandoStupsPieroneAPIV1PutTagReader is a Reader for the OrgZalandoStupsPieroneAPIV1PutTag structure.
type OrgZalandoStupsPieroneAPIV1PutTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIV1PutTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIV1PutTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewOrgZalandoStupsPieroneAPIV1PutTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIV1PutTagOK creates a OrgZalandoStupsPieroneAPIV1PutTagOK with default headers values
func NewOrgZalandoStupsPieroneAPIV1PutTagOK() *OrgZalandoStupsPieroneAPIV1PutTagOK {
	return &OrgZalandoStupsPieroneAPIV1PutTagOK{}
}

/*OrgZalandoStupsPieroneAPIV1PutTagOK handles this case with default header values.

Tag was stored successfully
*/
type OrgZalandoStupsPieroneAPIV1PutTagOK struct {
}

func (o *OrgZalandoStupsPieroneAPIV1PutTagOK) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/{team}/{artifact}/tags/{name}][%d] orgZalandoStupsPieroneApiV1PutTagOK ", 200)
}

func (o *OrgZalandoStupsPieroneAPIV1PutTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgZalandoStupsPieroneAPIV1PutTagConflict creates a OrgZalandoStupsPieroneAPIV1PutTagConflict with default headers values
func NewOrgZalandoStupsPieroneAPIV1PutTagConflict() *OrgZalandoStupsPieroneAPIV1PutTagConflict {
	return &OrgZalandoStupsPieroneAPIV1PutTagConflict{}
}

/*OrgZalandoStupsPieroneAPIV1PutTagConflict handles this case with default header values.

Tag already exists
*/
type OrgZalandoStupsPieroneAPIV1PutTagConflict struct {
}

func (o *OrgZalandoStupsPieroneAPIV1PutTagConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/{team}/{artifact}/tags/{name}][%d] orgZalandoStupsPieroneApiV1PutTagConflict ", 409)
}

func (o *OrgZalandoStupsPieroneAPIV1PutTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
