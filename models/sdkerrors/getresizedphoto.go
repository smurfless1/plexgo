// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetResizedPhotoErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetResizedPhotoErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetResizedPhotoErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetResizedPhotoErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetResizedPhotoResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetResizedPhotoResponseBody struct {
	Errors []GetResizedPhotoErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetResizedPhotoResponseBody{}

func (e *GetResizedPhotoResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
