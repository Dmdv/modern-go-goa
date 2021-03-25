// Code generated by goa v3.2.6, DO NOT EDIT.
//
// client HTTP server types
//
// Command:
// $ goa gen clients/design

package server

import (
	client "clients/gen/client"
	clientviews "clients/gen/client/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "client" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Client ID
	ClientName *string `form:"ClientName,omitempty" json:"ClientName,omitempty" xml:"ClientName,omitempty"`
}

// GetResponseBody is the type of the "client" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID is the unique id of the Client.
	ClientID string `form:"ClientID" json:"ClientID" xml:"ClientID"`
	// Name of the Client
	ClientName string `form:"ClientName" json:"ClientName" xml:"ClientName"`
}

// ClientManagementResponseCollection is the type of the "client" service
// "show" endpoint HTTP response body.
type ClientManagementResponseCollection []*ClientManagementResponse

// ClientManagementResponse is used to define fields on response body types.
type ClientManagementResponse struct {
	// ID is the unique id of the Client.
	ClientID string `form:"ClientID" json:"ClientID" xml:"ClientID"`
	// Name of the Client
	ClientName string `form:"ClientName" json:"ClientName" xml:"ClientName"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "client" service.
func NewGetResponseBody(res *clientviews.ClientManagementView) *GetResponseBody {
	body := &GetResponseBody{
		ClientID:   *res.ClientID,
		ClientName: *res.ClientName,
	}
	return body
}

// NewClientManagementResponseCollection builds the HTTP response body from the
// result of the "show" endpoint of the "client" service.
func NewClientManagementResponseCollection(res clientviews.ClientManagementCollectionView) ClientManagementResponseCollection {
	body := make([]*ClientManagementResponse, len(res))
	for i, val := range res {
		body[i] = marshalClientviewsClientManagementViewToClientManagementResponse(val)
	}
	return body
}

// NewAddPayload builds a client service add endpoint payload.
func NewAddPayload(body *AddRequestBody, clientID string) *client.AddPayload {
	v := &client.AddPayload{
		ClientName: *body.ClientName,
	}
	v.ClientID = clientID

	return v
}

// NewGetPayload builds a client service get endpoint payload.
func NewGetPayload(clientID string) *client.GetPayload {
	v := &client.GetPayload{}
	v.ClientID = clientID

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.ClientName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ClientName", "body"))
	}
	return
}
