// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// ListMandatesReader is a Reader for the ListMandates structure.
type ListMandatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMandatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewListMandatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListMandatesOK creates a ListMandatesOK with default headers values
func NewListMandatesOK() *ListMandatesOK {
	return &ListMandatesOK{}
}

/*ListMandatesOK handles this case with default header values.

List of mandates details
*/
type ListMandatesOK struct {

	//Payload

	// isStream: false
	*models.MandateDetailsListResponse
}

func (o *ListMandatesOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates][%d] listMandatesOK  %+v", 200, o)
}

func (o *ListMandatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateDetailsListResponse = new(models.MandateDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
