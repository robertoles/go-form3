// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// ModifyUnitReader is a Reader for the ModifyUnit structure.
type ModifyUnitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyUnitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewModifyUnitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyUnitOK creates a ModifyUnitOK with default headers values
func NewModifyUnitOK() *ModifyUnitOK {
	return &ModifyUnitOK{}
}

/*ModifyUnitOK handles this case with default header values.

Organisation details
*/
type ModifyUnitOK struct {

	//Payload

	// isStream: false
	*models.OrganisationDetailsResponse
}

func (o *ModifyUnitOK) Error() string {
	return fmt.Sprintf("[PATCH /organisation/units/{id}][%d] modifyUnitOK  %+v", 200, o)
}

func (o *ModifyUnitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.OrganisationDetailsResponse = new(models.OrganisationDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.OrganisationDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
