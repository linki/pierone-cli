package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*ScmSourceInformation scm source information

swagger:model ScmSourceInformation
*/
type ScmSourceInformation struct {

	/* author
	 */
	Author string `json:"author,omitempty"`

	/* revision
	 */
	Revision string `json:"revision,omitempty"`

	/* status
	 */
	Status string `json:"status,omitempty"`

	/* url
	 */
	URL string `json:"url,omitempty"`
}

// Validate validates this scm source information
func (m *ScmSourceInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}