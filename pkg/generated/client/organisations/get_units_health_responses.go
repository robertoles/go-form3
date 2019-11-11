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

// GetUnitsHealthReader is a Reader for the GetUnitsHealth structure.
type GetUnitsHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnitsHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetUnitsHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUnitsHealthOK creates a GetUnitsHealthOK with default headers values
func NewGetUnitsHealthOK() *GetUnitsHealthOK {
	return &GetUnitsHealthOK{}
}

/*GetUnitsHealthOK handles this case with default header values.

Organisation service health
*/
type GetUnitsHealthOK struct {

	//Payload

	// isStream: false
	*models.Health
}

func (o *GetUnitsHealthOK) Error() string {
	return fmt.Sprintf("[GET /organisation/units/health][%d] getUnitsHealthOK  %+v", 200, o)
}

func (o *GetUnitsHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Health = new(models.Health)

	// response payload

	if err := consumer.Consume(response.Body(), o.Health); err != nil && err != io.EOF {
		return err
	}

	return nil
}
