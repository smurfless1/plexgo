// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetLibraryItemsRequest struct {
	// the Id of the library to query
	SectionID float64 `pathParam:"style=simple,explode=false,name=sectionId"`
	// item type
	Type *float64 `queryParam:"style=form,explode=true,name=type"`
	// the filter parameter
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetLibraryItemsRequest) GetSectionID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SectionID
}

func (o *GetLibraryItemsRequest) GetType() *float64 {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetLibraryItemsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetLibraryItemsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetLibraryItemsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLibraryItemsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLibraryItemsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
