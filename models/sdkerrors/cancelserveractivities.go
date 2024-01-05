// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type CancelServerActivitiesErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *CancelServerActivitiesErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *CancelServerActivitiesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CancelServerActivitiesErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// CancelServerActivitiesResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type CancelServerActivitiesResponseBody struct {
	Errors []CancelServerActivitiesErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &CancelServerActivitiesResponseBody{}

func (e *CancelServerActivitiesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
