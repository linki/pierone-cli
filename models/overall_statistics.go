package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*OverallStatistics overall statistics

swagger:model OverallStatistics
*/
type OverallStatistics struct {

	/* storage
	 */
	Storage int64 `json:"storage,omitempty"`

	/* teams
	 */
	Teams int64 `json:"teams,omitempty"`
}

// Validate validates this overall statistics
func (m *OverallStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
