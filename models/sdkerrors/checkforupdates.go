// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type CheckForUpdatesErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *CheckForUpdatesErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *CheckForUpdatesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CheckForUpdatesErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// CheckForUpdatesResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type CheckForUpdatesResponseBody struct {
	Errors []CheckForUpdatesErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &CheckForUpdatesResponseBody{}

func (e *CheckForUpdatesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
