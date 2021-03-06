// Code generated by goa v3.2.6, DO NOT EDIT.
//
// client HTTP client types
//
// Command:
// $ goa gen clients/design

package client

import (
	client "clients/gen/client"
	clientviews "clients/gen/client/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "client" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Client ID
	ClientName string `form:"ClientName" json:"ClientName" xml:"ClientName"`
}

// GetResponseBody is the type of the "client" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID is the unique id of the Client.
	ClientID *string `form:"ClientID,omitempty" json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// Name of the Client
	ClientName *string `form:"ClientName,omitempty" json:"ClientName,omitempty" xml:"ClientName,omitempty"`
}

// ShowResponseBody is the type of the "client" service "show" endpoint HTTP
// response body.
type ShowResponseBody []*ClientManagementResponse

// ClientManagementResponse is used to define fields on response body types.
type ClientManagementResponse struct {
	// ID is the unique id of the Client.
	ClientID *string `form:"ClientID,omitempty" json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// Name of the Client
	ClientName *string `form:"ClientName,omitempty" json:"ClientName,omitempty" xml:"ClientName,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "client" service.
func NewAddRequestBody(p *client.AddPayload) *AddRequestBody {
	body := &AddRequestBody{
		ClientName: p.ClientName,
	}
	return body
}

// NewGetClientManagementOK builds a "client" service "get" endpoint result
// from a HTTP "OK" response.
func NewGetClientManagementOK(body *GetResponseBody) *clientviews.ClientManagementView {
	v := &clientviews.ClientManagementView{
		ClientID:   body.ClientID,
		ClientName: body.ClientName,
	}

	return v
}

// NewShowClientManagementCollectionOK builds a "client" service "show"
// endpoint result from a HTTP "OK" response.
func NewShowClientManagementCollectionOK(body ShowResponseBody) clientviews.ClientManagementCollectionView {
	v := make([]*clientviews.ClientManagementView, len(body))
	for i, val := range body {
		v[i] = unmarshalClientManagementResponseToClientviewsClientManagementView(val)
	}
	return v
}

// ValidateClientManagementResponse runs the validations defined on
// ClientManagementResponse
func ValidateClientManagementResponse(body *ClientManagementResponse) (err error) {
	if body.ClientID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ClientID", "body"))
	}
	if body.ClientName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ClientName", "body"))
	}
	return
}
