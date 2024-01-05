// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/LukeHagar/plexgo/internal/utils"
	"net/http"
)

type PerformSearchRequest struct {
	// The query term
	Query string `queryParam:"style=form,explode=true,name=query"`
	// This gives context to the search, and can result in re-ordering of search result hubs
	SectionID *float64 `queryParam:"style=form,explode=true,name=sectionId"`
	// The number of items to return per hub
	Limit *float64 `default:"3" queryParam:"style=form,explode=true,name=limit"`
}

func (p PerformSearchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PerformSearchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PerformSearchRequest) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *PerformSearchRequest) GetSectionID() *float64 {
	if o == nil {
		return nil
	}
	return o.SectionID
}

func (o *PerformSearchRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type PerformSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PerformSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PerformSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PerformSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
