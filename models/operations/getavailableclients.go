// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type Server struct {
	Name                 *string  `json:"name,omitempty"`
	Host                 *string  `json:"host,omitempty"`
	Address              *string  `json:"address,omitempty"`
	Port                 *float64 `json:"port,omitempty"`
	MachineIdentifier    *string  `json:"machineIdentifier,omitempty"`
	Version              *string  `json:"version,omitempty"`
	Protocol             *string  `json:"protocol,omitempty"`
	Product              *string  `json:"product,omitempty"`
	DeviceClass          *string  `json:"deviceClass,omitempty"`
	ProtocolVersion      *float64 `json:"protocolVersion,omitempty"`
	ProtocolCapabilities *string  `json:"protocolCapabilities,omitempty"`
}

func (o *Server) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Server) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *Server) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Server) GetPort() *float64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Server) GetMachineIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.MachineIdentifier
}

func (o *Server) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *Server) GetProtocol() *string {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *Server) GetProduct() *string {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *Server) GetDeviceClass() *string {
	if o == nil {
		return nil
	}
	return o.DeviceClass
}

func (o *Server) GetProtocolVersion() *float64 {
	if o == nil {
		return nil
	}
	return o.ProtocolVersion
}

func (o *Server) GetProtocolCapabilities() *string {
	if o == nil {
		return nil
	}
	return o.ProtocolCapabilities
}

type GetAvailableClientsMediaContainer struct {
	Size   *float64 `json:"size,omitempty"`
	Server []Server `json:"Server,omitempty"`
}

func (o *GetAvailableClientsMediaContainer) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetAvailableClientsMediaContainer) GetServer() []Server {
	if o == nil {
		return nil
	}
	return o.Server
}

type ResponseBody struct {
	MediaContainer *GetAvailableClientsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *ResponseBody) GetMediaContainer() *GetAvailableClientsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetAvailableClientsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Available Clients
	ResponseBodies []ResponseBody
}

func (o *GetAvailableClientsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAvailableClientsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAvailableClientsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAvailableClientsResponse) GetResponseBodies() []ResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
