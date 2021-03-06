// Code generated by go-swagger; DO NOT EDIT.

package fx_api

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

// Client.CreateDeal creates a new CreateDealRequest object
// with the default values initialized.
func (c *Client) CreateDeal() *CreateDealRequest {
	var ()
	return &CreateDealRequest{

		FXDealCreation: models.FXDealCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateDealRequest struct {

	/*CreationRequest*/

	*models.FXDealCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateDealRequest) FromJson(j string) *CreateDealRequest {

	var m models.FXDealCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.FXDealCreation = &m

	return o
}

func (o *CreateDealRequest) WithCreationRequest(creationRequest models.FXDealCreation) *CreateDealRequest {

	o.FXDealCreation = &creationRequest

	return o
}

func (o *CreateDealRequest) WithoutCreationRequest() *CreateDealRequest {

	o.FXDealCreation = &models.FXDealCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create deal Request
func (o *CreateDealRequest) WithContext(ctx context.Context) *CreateDealRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create deal Request
func (o *CreateDealRequest) WithHTTPClient(client *http.Client) *CreateDealRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateDealRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.FXDealCreation != nil {
		if err := r.SetBodyParam(o.FXDealCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
