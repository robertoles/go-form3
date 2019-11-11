// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserCredentialReader is a Reader for the DeleteUserCredential structure.
type DeleteUserCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:

		result := NewDeleteUserCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialNoContent creates a DeleteUserCredentialNoContent with default headers values
func NewDeleteUserCredentialNoContent() *DeleteUserCredentialNoContent {
	return &DeleteUserCredentialNoContent{}
}

/*DeleteUserCredentialNoContent handles this case with default header values.

Credential deleted
*/
type DeleteUserCredentialNoContent struct {
}

func (o *DeleteUserCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security/users/{user_id}/credentials/{client_id}][%d] deleteUserCredentialNoContent ", 204)
}

func (o *DeleteUserCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
