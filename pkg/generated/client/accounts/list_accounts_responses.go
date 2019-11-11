// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// ListAccountsReader is a Reader for the ListAccounts structure.
type ListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountsOK creates a ListAccountsOK with default headers values
func NewListAccountsOK() *ListAccountsOK {
	return &ListAccountsOK{}
}

/*ListAccountsOK handles this case with default header values.

List of account details
*/
type ListAccountsOK struct {

	//Payload

	// isStream: false
	*models.AccountDetailsListResponse
}

func (o *ListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts][%d] listAccountsOK  %+v", 200, o)
}

func (o *ListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountDetailsListResponse = new(models.AccountDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
