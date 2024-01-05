// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PlaylistType - limit to a type of playlist.
type PlaylistType string

const (
	PlaylistTypeAudio PlaylistType = "audio"
	PlaylistTypeVideo PlaylistType = "video"
	PlaylistTypePhoto PlaylistType = "photo"
)

func (e PlaylistType) ToPointer() *PlaylistType {
	return &e
}

func (e *PlaylistType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audio":
		fallthrough
	case "video":
		fallthrough
	case "photo":
		*e = PlaylistType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaylistType: %v", v)
	}
}

// QueryParamSmart - type of playlists to return (default is all).
type QueryParamSmart int64

const (
	QueryParamSmartZero QueryParamSmart = 0
	QueryParamSmartOne  QueryParamSmart = 1
)

func (e QueryParamSmart) ToPointer() *QueryParamSmart {
	return &e
}

func (e *QueryParamSmart) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = QueryParamSmart(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSmart: %v", v)
	}
}

type GetPlaylistsRequest struct {
	// limit to a type of playlist.
	PlaylistType *PlaylistType `queryParam:"style=form,explode=true,name=playlistType"`
	// type of playlists to return (default is all).
	Smart *QueryParamSmart `queryParam:"style=form,explode=true,name=smart"`
}

func (o *GetPlaylistsRequest) GetPlaylistType() *PlaylistType {
	if o == nil {
		return nil
	}
	return o.PlaylistType
}

func (o *GetPlaylistsRequest) GetSmart() *QueryParamSmart {
	if o == nil {
		return nil
	}
	return o.Smart
}

type GetPlaylistsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPlaylistsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPlaylistsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPlaylistsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
