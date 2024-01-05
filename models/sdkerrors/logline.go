// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type LogLineErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *LogLineErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *LogLineErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *LogLineErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// LogLineResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type LogLineResponseBody struct {
	Errors []LogLineErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &LogLineResponseBody{}

func (e *LogLineResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
