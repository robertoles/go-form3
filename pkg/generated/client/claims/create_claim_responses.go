// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreateClaimReader is a Reader for the CreateClaim structure.
type CreateClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:

		result := NewCreateClaimCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:

		result := NewCreateClaimBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateClaimCreated creates a CreateClaimCreated with default headers values
func NewCreateClaimCreated() *CreateClaimCreated {
	return &CreateClaimCreated{}
}

/*CreateClaimCreated handles this case with default header values.

Claim creation response
*/
type CreateClaimCreated struct {

	//Payload

	// isStream: false
	*models.ClaimDetailsResponse
}

func (o *CreateClaimCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/claims][%d] createClaimCreated  %+v", 201, o)
}

func (o *CreateClaimCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimDetailsResponse = new(models.ClaimDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClaimBadRequest creates a CreateClaimBadRequest with default headers values
func NewCreateClaimBadRequest() *CreateClaimBadRequest {
	return &CreateClaimBadRequest{}
}

/*CreateClaimBadRequest handles this case with default header values.

Claim creation error
*/
type CreateClaimBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateClaimBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/claims][%d] createClaimBadRequest  %+v", 400, o)
}

func (o *CreateClaimBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
