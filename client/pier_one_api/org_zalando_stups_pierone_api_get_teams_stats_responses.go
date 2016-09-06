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

// OrgZalandoStupsPieroneAPIGetTeamsStatsReader is a Reader for the OrgZalandoStupsPieroneAPIGetTeamsStats structure.
type OrgZalandoStupsPieroneAPIGetTeamsStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *OrgZalandoStupsPieroneAPIGetTeamsStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgZalandoStupsPieroneAPIGetTeamsStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgZalandoStupsPieroneAPIGetTeamsStatsOK creates a OrgZalandoStupsPieroneAPIGetTeamsStatsOK with default headers values
func NewOrgZalandoStupsPieroneAPIGetTeamsStatsOK() *OrgZalandoStupsPieroneAPIGetTeamsStatsOK {
	return &OrgZalandoStupsPieroneAPIGetTeamsStatsOK{}
}

/*OrgZalandoStupsPieroneAPIGetTeamsStatsOK handles this case with default header values.

Return team stats
*/
type OrgZalandoStupsPieroneAPIGetTeamsStatsOK struct {
	Payload *models.TeamStatistics
}

func (o *OrgZalandoStupsPieroneAPIGetTeamsStatsOK) Error() string {
	return fmt.Sprintf("[GET /stats/teams][%d] orgZalandoStupsPieroneApiGetTeamsStatsOK  %+v", 200, o.Payload)
}

func (o *OrgZalandoStupsPieroneAPIGetTeamsStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
