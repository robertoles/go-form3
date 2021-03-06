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

// GetDealReader is a Reader for the GetDeal structure.
type GetDealReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDealReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDealOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDealBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDealForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDealNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetDealBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDealDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDealOK creates a GetDealOK with default headers values
func NewGetDealOK() *GetDealOK {
	return &GetDealOK{}
}

/*GetDealOK handles this case with default header values.

fx deal response
*/
type GetDealOK struct {

	//Payload

	// isStream: false
	*models.FXDealResponse
}

func (o *GetDealOK) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealOK", 200)
}

func (o *GetDealOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FXDealResponse = new(models.FXDealResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FXDealResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealBadRequest creates a GetDealBadRequest with default headers values
func NewGetDealBadRequest() *GetDealBadRequest {
	return &GetDealBadRequest{}
}

/*GetDealBadRequest handles this case with default header values.

Bad Request
*/
type GetDealBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetDealBadRequest) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealBadRequest", 400)
}

func (o *GetDealBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealForbidden creates a GetDealForbidden with default headers values
func NewGetDealForbidden() *GetDealForbidden {
	return &GetDealForbidden{}
}

/*GetDealForbidden handles this case with default header values.

Action Forbidden
*/
type GetDealForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetDealForbidden) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealForbidden", 403)
}

func (o *GetDealForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealNotFound creates a GetDealNotFound with default headers values
func NewGetDealNotFound() *GetDealNotFound {
	return &GetDealNotFound{}
}

/*GetDealNotFound handles this case with default header values.

Not Found
*/
type GetDealNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetDealNotFound) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealNotFound", 404)
}

func (o *GetDealNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealBadGateway creates a GetDealBadGateway with default headers values
func NewGetDealBadGateway() *GetDealBadGateway {
	return &GetDealBadGateway{}
}

/*GetDealBadGateway handles this case with default header values.

Bad Gateway
*/
type GetDealBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetDealBadGateway) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] getDealBadGateway", 502)
}

func (o *GetDealBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealDefault creates a GetDealDefault with default headers values
func NewGetDealDefault(code int) *GetDealDefault {
	return &GetDealDefault{
		_statusCode: code,
	}
}

/*GetDealDefault handles this case with default header values.

Unexpected Error
*/
type GetDealDefault struct {
	_statusCode int

	//Payload

	// isStream: false
	*models.APIError
}

// Code gets the status code for the get deal default response
func (o *GetDealDefault) Code() int {
	return o._statusCode
}

func (o *GetDealDefault) Error() string {
	return fmt.Sprintf("[GET /fx/deals/{fx_deal_id}][%d] GetDeal default", o._statusCode)
}

func (o *GetDealDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
