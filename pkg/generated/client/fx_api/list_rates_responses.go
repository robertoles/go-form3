// Code generated by go-swagger; DO NOT EDIT.

package fx_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// ListRatesReader is a Reader for the ListRates structure.
type ListRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListRatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListRatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewListRatesBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListRatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRatesOK creates a ListRatesOK with default headers values
func NewListRatesOK() *ListRatesOK {
	return &ListRatesOK{}
}

/*ListRatesOK handles this case with default header values.

FX rates response
*/
type ListRatesOK struct {

	//Payload

	// isStream: false
	*models.FXRatesListResponse
}

func (o *ListRatesOK) Error() string {
	return fmt.Sprintf("[GET /fx/rates][%d] listRatesOK", 200)
}

func (o *ListRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FXRatesListResponse = new(models.FXRatesListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FXRatesListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRatesBadRequest creates a ListRatesBadRequest with default headers values
func NewListRatesBadRequest() *ListRatesBadRequest {
	return &ListRatesBadRequest{}
}

/*ListRatesBadRequest handles this case with default header values.

Bad Request
*/
type ListRatesBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListRatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fx/rates][%d] listRatesBadRequest", 400)
}

func (o *ListRatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRatesForbidden creates a ListRatesForbidden with default headers values
func NewListRatesForbidden() *ListRatesForbidden {
	return &ListRatesForbidden{}
}

/*ListRatesForbidden handles this case with default header values.

Action Forbidden
*/
type ListRatesForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListRatesForbidden) Error() string {
	return fmt.Sprintf("[GET /fx/rates][%d] listRatesForbidden", 403)
}

func (o *ListRatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRatesBadGateway creates a ListRatesBadGateway with default headers values
func NewListRatesBadGateway() *ListRatesBadGateway {
	return &ListRatesBadGateway{}
}

/*ListRatesBadGateway handles this case with default header values.

Bad Gateway
*/
type ListRatesBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListRatesBadGateway) Error() string {
	return fmt.Sprintf("[GET /fx/rates][%d] listRatesBadGateway", 502)
}

func (o *ListRatesBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRatesDefault creates a ListRatesDefault with default headers values
func NewListRatesDefault(code int) *ListRatesDefault {
	return &ListRatesDefault{
		_statusCode: code,
	}
}

/*ListRatesDefault handles this case with default header values.

Unexpected Error
*/
type ListRatesDefault struct {
	_statusCode int

	//Payload

	// isStream: false
	*models.APIError
}

// Code gets the status code for the list rates default response
func (o *ListRatesDefault) Code() int {
	return o._statusCode
}

func (o *ListRatesDefault) Error() string {
	return fmt.Sprintf("[GET /fx/rates][%d] ListRates default", o._statusCode)
}

func (o *ListRatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
