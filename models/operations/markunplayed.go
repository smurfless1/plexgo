// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MarkUnplayedRequest struct {
	// The media key to mark as Unplayed
	Key float64 `queryParam:"style=form,explode=true,name=key"`
}

func (o *MarkUnplayedRequest) GetKey() float64 {
	if o == nil {
		return 0.0
	}
	return o.Key
}

type MarkUnplayedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *MarkUnplayedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MarkUnplayedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MarkUnplayedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
