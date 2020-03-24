// Code generated by go-swagger; DO NOT EDIT.

package addressbook

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

// Client.PostAddressbookPartiesIDContactsContactIDAccounts creates a new PostAddressbookPartiesIDContactsContactIDAccountsRequest object
// with the default values initialized.
func (c *Client) PostAddressbookPartiesIDContactsContactIDAccounts() *PostAddressbookPartiesIDContactsContactIDAccountsRequest {
	var ()
	return &PostAddressbookPartiesIDContactsContactIDAccountsRequest{

		ContactID: c.Defaults.GetStrfmtUUID("PostAddressbookPartiesIDContactsContactIDAccounts", "contact_id"),

		ContactAccountCreation: models.ContactAccountCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("PostAddressbookPartiesIDContactsContactIDAccounts", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostAddressbookPartiesIDContactsContactIDAccountsRequest struct {

	/*ContactID      Id of Contact      */

	ContactID strfmt.UUID

	/*CreationRequest*/

	*models.ContactAccountCreation

	/*ID      Id of party      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) FromJson(j string) *PostAddressbookPartiesIDContactsContactIDAccountsRequest {

	var m models.ContactAccountCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.ContactAccountCreation = &m

	return o
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WithContactID(contactID strfmt.UUID) *PostAddressbookPartiesIDContactsContactIDAccountsRequest {

	o.ContactID = contactID

	return o
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WithCreationRequest(creationRequest models.ContactAccountCreation) *PostAddressbookPartiesIDContactsContactIDAccountsRequest {

	o.ContactAccountCreation = &creationRequest

	return o
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WithoutCreationRequest() *PostAddressbookPartiesIDContactsContactIDAccountsRequest {

	o.ContactAccountCreation = &models.ContactAccountCreation{}

	return o
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WithID(id strfmt.UUID) *PostAddressbookPartiesIDContactsContactIDAccountsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the post addressbook parties ID contacts contact ID accounts Request
func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WithContext(ctx context.Context) *PostAddressbookPartiesIDContactsContactIDAccountsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post addressbook parties ID contacts contact ID accounts Request
func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WithHTTPClient(client *http.Client) *PostAddressbookPartiesIDContactsContactIDAccountsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostAddressbookPartiesIDContactsContactIDAccountsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID.String()); err != nil {
		return err
	}

	// ISBODYPARAM
	if o.ContactAccountCreation != nil {
		if err := r.SetBodyParam(o.ContactAccountCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
