// Code generated by go-swagger; DO NOT EDIT.

package fundingrequest_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDReader is a Reader for the GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionID structure.
type GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK creates a GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK with default headers values
func NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK() *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK {
	return &GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK{}
}

/*GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK handles this case with default header values.

funding request admissionresponse
*/
type GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK struct {

	//Payload

	// isStream: false
	*models.FundingRequestAdmissionResponse
}

func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/fundingrequests/{funding_request_id}/admissions/{funding_request_admission_id}][%d] getTransactionFundingrequestsFundingRequestIdAdmissionsFundingRequestAdmissionIdOK", 200)
}

func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.FundingRequestAdmissionResponse = new(models.FundingRequestAdmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.FundingRequestAdmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest creates a GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest with default headers values
func NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest() *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest {
	return &GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest{}
}

/*GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest handles this case with default header values.

bad request
*/
type GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/fundingrequests/{funding_request_id}/admissions/{funding_request_admission_id}][%d] getTransactionFundingrequestsFundingRequestIdAdmissionsFundingRequestAdmissionIdBadRequest", 400)
}

func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden creates a GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden with default headers values
func NewGetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden() *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden {
	return &GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden{}
}

/*GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden handles this case with default header values.

forbidden
*/
type GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /transaction/fundingrequests/{funding_request_id}/admissions/{funding_request_admission_id}][%d] getTransactionFundingrequestsFundingRequestIdAdmissionsFundingRequestAdmissionIdForbidden", 403)
}

func (o *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
