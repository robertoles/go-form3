// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetPaymentReturnReversalReader is a Reader for the GetPaymentReturnReversal structure.
type GetPaymentReturnReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReturnReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetPaymentReturnReversalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReturnReversalOK creates a GetPaymentReturnReversalOK with default headers values
func NewGetPaymentReturnReversalOK() *GetPaymentReturnReversalOK {
	return &GetPaymentReturnReversalOK{}
}

/*GetPaymentReturnReversalOK handles this case with default header values.

Return reversal details
*/
type GetPaymentReturnReversalOK struct {

	//Payload

	// isStream: false
	*models.ReturnReversalDetailsResponse
}

func (o *GetPaymentReturnReversalOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/returns/{returnId}/reversals/{reversalId}][%d] getPaymentReturnReversalOK  %+v", 200, o)
}

func (o *GetPaymentReturnReversalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnReversalDetailsResponse = new(models.ReturnReversalDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnReversalDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
