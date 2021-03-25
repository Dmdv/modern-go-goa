// Code generated by goa v3.2.6, DO NOT EDIT.
//
// client HTTP client encoders and decoders
//
// Command:
// $ goa gen clients/design

package client

import (
	"bytes"
	client "clients/gen/client"
	clientviews "clients/gen/client/views"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "client" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		clientID string
	)
	{
		p, ok := v.(*client.AddPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("client", "add", "*client.AddPayload", v)
		}
		clientID = p.ClientID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddClientPath(clientID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("client", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the client add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*client.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("client", "add", "*client.AddPayload", v)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("client", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the client add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("client", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "client" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		clientID string
	)
	{
		p, ok := v.(*client.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("client", "get", "*client.GetPayload", v)
		}
		clientID = p.ClientID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetClientPath(clientID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("client", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the client get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("client", "get", err)
			}
			p := NewGetClientManagementOK(&body)
			view := "default"
			vres := &clientviews.ClientManagement{Projected: p, View: view}
			if err = clientviews.ValidateClientManagement(vres); err != nil {
				return nil, goahttp.ErrValidationError("client", "get", err)
			}
			res := client.NewClientManagement(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("client", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "client" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowClientPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("client", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the client
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("client", "show", err)
			}
			p := NewShowClientManagementCollectionOK(body)
			view := "default"
			vres := clientviews.ClientManagementCollection{Projected: p, View: view}
			if err = clientviews.ValidateClientManagementCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("client", "show", err)
			}
			res := client.NewClientManagementCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("client", "show", resp.StatusCode, string(body))
		}
	}
}

// unmarshalClientManagementResponseToClientviewsClientManagementView builds a
// value of type *clientviews.ClientManagementView from a value of type
// *ClientManagementResponse.
func unmarshalClientManagementResponseToClientviewsClientManagementView(v *ClientManagementResponse) *clientviews.ClientManagementView {
	res := &clientviews.ClientManagementView{
		ClientID:   v.ClientID,
		ClientName: v.ClientName,
	}

	return res
}