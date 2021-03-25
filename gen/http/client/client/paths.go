// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the client service.
//
// Command:
// $ goa gen clients/design

package client

import (
	"fmt"
)

// AddClientPath returns the URL path to the client service add HTTP endpoint.
func AddClientPath(clientID string) string {
	return fmt.Sprintf("/api/v1/client/%v", clientID)
}

// GetClientPath returns the URL path to the client service get HTTP endpoint.
func GetClientPath(clientID string) string {
	return fmt.Sprintf("/api/v1/client/%v", clientID)
}

// ShowClientPath returns the URL path to the client service show HTTP endpoint.
func ShowClientPath() string {
	return "/api/v1/client"
}
