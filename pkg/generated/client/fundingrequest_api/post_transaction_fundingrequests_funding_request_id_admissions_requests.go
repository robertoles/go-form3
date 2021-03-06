// Code generated by go-swagger; DO NOT EDIT.

package fundingrequest_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// Client.PostTransactionFundingrequestsFundingRequestIDAdmissions creates a new PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest object
// with the default values initialized.
func (c *Client) PostTransactionFundingrequestsFundingRequestIDAdmissions() *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {
	var ()
	return &PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest{

		FundingRequestAdmissionCreation: models.FundingRequestAdmissionCreationWithDefaults(c.Defaults),

		FundingRequestID: c.Defaults.GetStrfmtUUID("PostTransactionFundingrequestsFundingRequestIDAdmissions", "funding_request_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest struct {

	/*CreationRequest*/

	*models.FundingRequestAdmissionCreation

	/*FundingRequestID      Funding Request ID      */

	FundingRequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) FromJson(j string) *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {

	var m models.FundingRequestAdmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.FundingRequestAdmissionCreation = &m

	return o
}

func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) WithCreationRequest(creationRequest models.FundingRequestAdmissionCreation) *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {

	o.FundingRequestAdmissionCreation = &creationRequest

	return o
}

func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) WithoutCreationRequest() *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {

	o.FundingRequestAdmissionCreation = &models.FundingRequestAdmissionCreation{}

	return o
}

func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) WithFundingRequestID(fundingRequestID strfmt.UUID) *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {

	o.FundingRequestID = fundingRequestID

	return o
}

//////////////////
// WithContext adds the context to the post transaction fundingrequests funding request ID admissions Request
func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) WithContext(ctx context.Context) *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post transaction fundingrequests funding request ID admissions Request
func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) WithHTTPClient(client *http.Client) *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.FundingRequestAdmissionCreation != nil {
		if err := r.SetBodyParam(o.FundingRequestAdmissionCreation); err != nil {
			return err
		}
	}

	// path param funding_request_id
	if err := r.SetPathParam("funding_request_id", o.FundingRequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
