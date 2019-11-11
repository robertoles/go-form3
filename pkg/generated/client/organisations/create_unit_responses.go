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

// CreateUnitReader is a Reader for the CreateUnit structure.
type CreateUnitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUnitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:

		result := NewCreateUnitCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUnitCreated creates a CreateUnitCreated with default headers values
func NewCreateUnitCreated() *CreateUnitCreated {
	return &CreateUnitCreated{}
}

/*CreateUnitCreated handles this case with default header values.

Organisation creation response
*/
type CreateUnitCreated struct {

	//Payload

	// isStream: false
	*models.OrganisationCreationResponse
}

func (o *CreateUnitCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/units][%d] createUnitCreated  %+v", 201, o)
}

func (o *CreateUnitCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.OrganisationCreationResponse = new(models.OrganisationCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.OrganisationCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
