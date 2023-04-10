package go_routeros_client

import (
	"context"
	"net/http"

	"github.com/ogi4i/go-routeros-client/model"
	"github.com/ogi4i/go-routeros-client/query"
)

type (
	// IPRoute represents an IP route entry from `/ip/route`
	IPRoute struct {
		ID                 model.ID                `json:".id"`
		Active             bool                    `json:"active,string"`
		Distance           uint8                   `json:"distance,string"`
		DestinationAddress model.IPPrefix          `json:"dst-address"`
		Disabled           bool                    `json:"disabled,string"`
		Dynamic            bool                    `json:"dynamic,string"`
		ECMP               bool                    `json:"ecmp,string"`
		Gateway            model.Gateway           `json:"gateway"`
		HWOffloaded        bool                    `json:"hw-offloaded,string"`
		ImmediateGateway   model.Gateway           `json:"immediate-gw"`
		Inactive           bool                    `json:"inactive,string"`
		OSPF               bool                    `json:"ospf,string"`
		PreferredSource    model.IPAddr            `json:"pref-src"`
		RoutingTable       string                  `json:"routing-table"`
		Scope              uint8                   `json:"scope,string"`
		SuppressHWOffload  bool                    `json:"suppress-hw-offload,string"`
		TargetScope        uint8                   `json:"target-scope,string"`
		Comment            *string                 `json:"comment,omitempty"`
		Blackhole          *bool                   `json:"blackhole,string,omitempty"`
		CheckGateway       *model.CheckGatewayMode `json:"check-gateway,omitempty"`
		VRFInterface       *string                 `json:"vrf-interface,omitempty"`
	}

	getIPRouteParams struct{}

	// CreateIPRouteParams represents parameters with which a new IP route can be created
	CreateIPRouteParams struct {
		DestinationAddress model.IPPrefix          `json:"dst-address"`
		Gateway            model.Gateway           `json:"gateway"`
		Disabled           *bool                   `json:"disabled,string"`
		Distance           *uint8                  `json:"distance,string,omitempty"`
		PreferredSource    *model.IPAddr           `json:"pref-src,omitempty"`
		RoutingTable       *string                 `json:"routing-table,omitempty"`
		Scope              *uint8                  `json:"scope,string,string,omitempty"`
		SuppressHWOffload  *bool                   `json:"suppress-hw-offload,string,omitempty"`
		TargetScope        *uint8                  `json:"target-scope,string,omitempty"`
		Comment            *string                 `json:"comment,omitempty"`
		Blackhole          *bool                   `json:"blackhole,string,omitempty"`
		CheckGateway       *model.CheckGatewayMode `json:"check-gateway,omitempty"`
		VRFInterface       *string                 `json:"vrf-interface,omitempty"`
	}

	// UpdateIPRouteParams represents parameters with which an existing IP route can be updated
	UpdateIPRouteParams struct {
		DestinationAddress *model.IPPrefix         `json:"dst-address"`
		Gateway            *model.Gateway          `json:"gateway"`
		Disabled           *bool                   `json:"disabled,string"`
		Distance           *uint8                  `json:"distance,string,omitempty"`
		PreferredSource    *model.IPAddr           `json:"pref-src,omitempty"`
		RoutingTable       *string                 `json:"routing-table,omitempty"`
		Scope              *uint8                  `json:"scope,string,string,omitempty"`
		SuppressHWOffload  *bool                   `json:"suppress-hw-offload,string,omitempty"`
		TargetScope        *uint8                  `json:"target-scope,string,omitempty"`
		Comment            *string                 `json:"comment,omitempty"`
		CheckGateway       *model.CheckGatewayMode `json:"check-gateway,omitempty"`
		VRFInterface       *string                 `json:"vrf-interface,omitempty"`
	}
)

// GetIPRouteParams creates a builder for query parameters
func GetIPRouteParams() getIPRouteParams {
	return getIPRouteParams{}
}

func (gip getIPRouteParams) Active() query.Param[bool] {
	return "active"
}

func (gip getIPRouteParams) Distance() query.Param[uint8] {
	return "distance"
}

func (gip getIPRouteParams) DestinationAddress() query.Param[model.IPPrefix] {
	return "dst-address"
}

func (gip getIPRouteParams) Disabled() query.Param[bool] {
	return "disabled"
}

func (gip getIPRouteParams) Dynamic() query.Param[bool] {
	return "dynamic"
}

func (gip getIPRouteParams) ECMP() query.Param[bool] {
	return "ecmp"
}

func (gip getIPRouteParams) Gateway() query.Param[model.Gateway] {
	return "gateway"
}

func (gip getIPRouteParams) HWOffloaded() query.Param[bool] {
	return "hw-offloaded"
}

func (gip getIPRouteParams) ImmediateGateway() query.Param[model.Gateway] {
	return "immediate-gw"
}

func (gip getIPRouteParams) Inactive() query.Param[bool] {
	return "inactive"
}

func (gip getIPRouteParams) OSPF() query.Param[bool] {
	return "ospf"
}

func (gip getIPRouteParams) PreferredSource() query.Param[model.IPAddr] {
	return "pref-src"
}

func (gip getIPRouteParams) RoutingTable() query.Param[string] {
	return "routing-table"
}

func (gip getIPRouteParams) Scope() query.Param[uint8] {
	return "scope"
}

func (gip getIPRouteParams) SuppressHWOffload() query.Param[bool] {
	return "suppress-hw-offload"
}

func (gip getIPRouteParams) TargetScope() query.Param[uint8] {
	return "target-scope"
}

func (gip getIPRouteParams) Comment() query.Param[string] {
	return "comment"
}

func (gip getIPRouteParams) Blackhole() query.Param[bool] {
	return "blackhole"
}

func (gip getIPRouteParams) CheckGateway() query.Param[model.CheckGatewayMode] {
	return "check-gateway"
}

func (gip getIPRouteParams) VRFInterface() query.Param[string] {
	return "vfr-interface"
}

// ListIPRoutes returns a list of all IP routes
// Same as FilterIPRoutes without any filters
func (c *Client) ListIPRoutes(ctx context.Context) ([]IPRoute, error) {
	var ipRoutes []IPRoute
	err := c.do(ctx, http.MethodGet, c.baseURL+"/ip/route", http.NoBody, &ipRoutes)
	return ipRoutes, err
}

// FilterARPEntries returns a list of ARP entries which pass a given list of Filters
// Without any Filters is the same as ListARPEntries
func (c *Client) FilterIPRoutes(ctx context.Context, filters ...query.Filter) ([]IPRoute, error) {
	var ipRoutes []IPRoute
	err := c.do(ctx, http.MethodPost, c.baseURL+"/ip/route/print", prepareQuery(filters...), &ipRoutes)
	return ipRoutes, err
}

// GetIPRouteByID returns a single IP route by its identifier
// The ID must be previously obtained through other methods
func (c *Client) GetIPRouteByID(ctx context.Context, id model.ID) (IPRoute, error) {
	var ipRoute IPRoute
	err := c.do(ctx, http.MethodGet, c.baseURL+"/ip/route/"+id.String(), http.NoBody, &ipRoute)
	return ipRoute, err
}

// CreateIPRoute creates a new IP route and returns it
func (c *Client) CreateIPRoute(ctx context.Context, params CreateIPRouteParams) (IPRoute, error) {
	var ipRoute IPRoute
	err := c.do(ctx, http.MethodPut, c.baseURL+"/ip/route", params, &ipRoute)
	return ipRoute, err
}

// UpdateIPRoute updates an existing IP route and returns its new state
// The ID must be previously obtained through other methods
func (c *Client) UpdateIPRoute(ctx context.Context, id model.ID, params UpdateIPRouteParams) (IPRoute, error) {
	var ipRoute IPRoute
	err := c.do(ctx, http.MethodPatch, c.baseURL+"/ip/route/"+id.String(), params, &ipRoute)
	return ipRoute, err
}

// DeleteIPRoute deletes an existing IP route
// The ID must be previously obtained through other methods
func (c *Client) DeleteIPRoute(ctx context.Context, id model.ID) error {
	return c.do(ctx, http.MethodDelete, c.baseURL+"/ip/route/"+id.String(), http.NoBody, http.NoBody)
}
