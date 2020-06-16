// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new audit API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for audit API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get audit entry API
*/
func (a *GetAuditEntryRequest) Do() (*GetAuditEntryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAuditEntry",
		Method:             "GET",
		PathPattern:        "/audit/entries/{record_type}/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAuditEntryReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditEntryOK), nil

}

func (a *GetAuditEntryRequest) MustDo() *GetAuditEntryOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list audit entries API
*/
func (a *ListAuditEntriesRequest) Do() (*ListAuditEntriesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAuditEntries",
		Method:             "GET",
		PathPattern:        "/audit/entries/{record_type}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListAuditEntriesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAuditEntriesOK), nil

}

func (a *ListAuditEntriesRequest) MustDo() *ListAuditEntriesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
