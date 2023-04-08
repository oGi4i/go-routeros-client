package go_routeros_client

import (
	"context"
	"net/http"

	"github.com/ogi4i/go-routeros-client/model"
	"github.com/ogi4i/go-routeros-client/query"
)

type (
	// IPAddress represents an IP address from `/ip/address`
	IPAddress struct {
		ID              model.ID       `json:".id"`
		ActualInterface string         `json:"actual-interface"`
		Address         model.IPPrefix `json:"address"`
		Disabled        bool           `json:"disabled,string"`
		Dynamic         bool           `json:"dynamic,string"`
		Interface       string         `json:"interface"`
		Invalid         bool           `json:"invalid,string"`
		Network         model.IPAddr   `json:"network"`
		Comment         *string        `json:"comment,omitempty"`
	}

	getIPAddressParams struct{}

	// CreateIPAddressParams represents parameters with which a new IP address can be created
	CreateIPAddressParams struct {
		Address   model.IPPrefix `json:"address"`
		Disabled  bool           `json:"disabled,string"`
		Interface string         `json:"interface"`
		Comment   *string        `json:"comment,omitempty"`
	}

	// UpdateIPAddressParams represents parameters with which an existing IP address can be updated
	UpdateIPAddressParams struct {
		Address   *model.IPPrefix `json:"address,omitempty"`
		Disabled  *bool           `json:"disabled,omitempty,string"`
		Interface *string         `json:"interface,omitempty"`
		Comment   *string         `json:"comment,omitempty"`
	}
)

// GetIPAddressParams creates a builder for query parameters
func GetIPAddressParams() getIPAddressParams {
	return getIPAddressParams{}
}

func (gip getIPAddressParams) ActualInterface() query.Param[string] {
	return "actual-interface"
}

func (gip getIPAddressParams) Address() query.Param[model.IPAddr] {
	return "address"
}

func (gip getIPAddressParams) Disabled() query.Param[bool] {
	return "disabled"
}

func (gip getIPAddressParams) Dynamic() query.Param[bool] {
	return "dynamic"
}

func (gip getIPAddressParams) Interface() query.Param[string] {
	return "interface"
}

func (gip getIPAddressParams) Invalid() query.Param[bool] {
	return "invalid"
}

func (gip getIPAddressParams) Network() query.Param[model.IPAddr] {
	return "network"
}

func (gip getIPAddressParams) Comment() query.Param[string] {
	return "comment"
}

// ListIPAddresses returns a list of all IP addresses
// Same as FilterIPAddresses without any filters
func (c *Client) ListIPAddresses(ctx context.Context) ([]IPAddress, error) {
	var ipAddressList []IPAddress
	err := c.do(ctx, http.MethodGet, c.baseURL+"/ip/address", http.NoBody, &ipAddressList)
	return ipAddressList, err
}

// FilterIPAddresses returns a list of IP addresses which pass a given list of Filters
// Without any Filters is the same as ListIPAddresses
func (c *Client) FilterIPAddresses(ctx context.Context, filters ...query.Filter) ([]IPAddress, error) {
	var ipAddressList []IPAddress
	err := c.do(ctx, http.MethodPost, c.baseURL+"/ip/address/print", prepareQuery(filters...), &ipAddressList)
	return ipAddressList, err
}

// ListIPAddresses returns a single IP address by its identifier
// The ID must be previously obtained through other methods
func (c *Client) GetIPAddressByID(ctx context.Context, id model.ID) (IPAddress, error) {
	var ipAddress IPAddress
	err := c.do(ctx, http.MethodGet, c.baseURL+"/ip/address/"+id.String(), http.NoBody, &ipAddress)
	return ipAddress, err
}

// CreateIPAddress creates a new IP address and returns it
func (c *Client) CreateIPAddress(ctx context.Context, params CreateIPAddressParams) (IPAddress, error) {
	var ipAddress IPAddress
	err := c.do(ctx, http.MethodPut, c.baseURL+"/ip/address", params, &ipAddress)
	return ipAddress, err
}

// UpdateIPAddress updates an existing IP address and returns its new state
// The ID must be previously obtained through other methods
func (c *Client) UpdateIPAddress(ctx context.Context, id model.ID, params UpdateIPAddressParams) (IPAddress, error) {
	var ipAddress IPAddress
	err := c.do(ctx, http.MethodPatch, c.baseURL+"/ip/address/"+id.String(), params, &ipAddress)
	return ipAddress, err
}

// DeleteIPAddress deletes an existing IP address
// The ID must be previously obtained through other methods
func (c *Client) DeleteIPAddress(ctx context.Context, id model.ID) error {
	return c.do(ctx, http.MethodDelete, c.baseURL+"/ip/address/"+id.String(), http.NoBody, http.NoBody)
}
