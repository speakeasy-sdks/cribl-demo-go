// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl-demo/pkg/models/shared"
	"net/http"
)

// ListCertificates200ApplicationJSON - a list of Certificate objects
type ListCertificates200ApplicationJSON struct {
	// number of items present in the items array
	Count *int64               `json:"count,omitempty"`
	Items []shared.Certificate `json:"items,omitempty"`
}

type ListCertificatesResponse struct {
	ContentType string
	// Unexpected error
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// a list of Certificate objects
	ListCertificates200ApplicationJSONObject *ListCertificates200ApplicationJSON
}