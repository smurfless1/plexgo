// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLibraryRequest struct {
	// the Id of the library to query
	SectionID float64 `pathParam:"style=simple,explode=false,name=sectionId"`
}

func (o *DeleteLibraryRequest) GetSectionID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SectionID
}

type DeleteLibraryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteLibraryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLibraryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLibraryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
