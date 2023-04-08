package go_routeros_client

import (
	"context"
	"net/http"

	"github.com/ogi4i/go-routeros-client/model"
	"github.com/ogi4i/go-routeros-client/query"
)

type (
	// ARPEntry represents an ARP entry from `/ip/arp`
	ARPEntry struct {
		ID         model.ID         `json:".id"`
		DHCP       bool             `json:"DHCP,string"`
		Address    model.IPAddr     `json:"address"`
		Complete   bool             `json:"complete,string"`
		Disabled   bool             `json:"disabled,string"`
		Dynamic    bool             `json:"dynamic,string"`
		Interface  string           `json:"interface"`
		Invalid    bool             `json:"invalid,string"`
		MACAddress model.MACAddress `json:"mac-address"`
		Published  bool             `json:"published,string"`
		Comment    *string          `json:"comment,omitempty"`
	}

	getARPEntryParams struct{}

	// CreateARPEntryParams represents parameters with which a new IP ARP entry can be created
	CreateARPEntryParams struct {
		Address    model.IPAddr     `json:"address"`
		Disabled   bool             `json:"disabled,string"`
		Interface  string           `json:"interface"`
		MACAddress model.MACAddress `json:"mac-address"`
		Published  bool             `json:"published,string"`
		Comment    *string          `json:"comment,omitempty"`
	}

	// UpdateARPEntryParams represents parameters with which an existing IP ARP entry can be updated
	UpdateARPEntryParams struct {
		Address    *model.IPAddr     `json:"address,omitempty"`
		Disabled   *bool             `json:"disabled,omitempty,string"`
		Interface  *string           `json:"interface,omitempty"`
		MACAddress *model.MACAddress `json:"mac-address,omitempty"`
		Published  *bool             `json:"published,omitempty,string"`
		Comment    *string           `json:"comment,omitempty"`
	}
)

// GetARPEntryParams creates a builder for query parameters
func GetARPEntryParams() getARPEntryParams {
	return getARPEntryParams{}
}

func (gip getARPEntryParams) DHCP() query.Param[string] {
	return "DHCP"
}

func (gip getARPEntryParams) Address() query.Param[model.IPAddr] {
	return "address"
}

func (gip getARPEntryParams) Complete() query.Param[bool] {
	return "complete"
}

func (gip getARPEntryParams) Disabled() query.Param[bool] {
	return "disabled"
}

func (gip getARPEntryParams) Dynamic() query.Param[bool] {
	return "dynamic"
}

func (gip getARPEntryParams) Interface() query.Param[string] {
	return "interface"
}

func (gip getARPEntryParams) Invalid() query.Param[bool] {
	return "invalid"
}

func (gip getARPEntryParams) MACAddress() query.Param[model.MACAddress] {
	return "mac-address"
}

func (gip getARPEntryParams) Published() query.Param[bool] {
	return "published"
}

func (gip getARPEntryParams) Comment() query.Param[string] {
	return "comment"
}

// ListARPEntries returns a list of all ARP entries
// Same as FilterARPEntries without any filters
func (c *Client) ListARPEntries(ctx context.Context) ([]ARPEntry, error) {
	var arpEntries []ARPEntry
	err := c.do(ctx, http.MethodGet, c.baseURL+"/ip/arp", http.NoBody, &arpEntries)
	return arpEntries, err
}

// FilterARPEntries returns a list of ARP entries which pass a given list of Filters
// Without any Filters is the same as ListARPEntries
func (c *Client) FilterARPEntries(ctx context.Context, filters ...query.Filter) ([]ARPEntry, error) {
	var arpEntries []ARPEntry
	err := c.do(ctx, http.MethodPost, c.baseURL+"/ip/arp/print", prepareQuery(filters...), &arpEntries)
	return arpEntries, err
}

// GetARPEntryByID returns a single ARP entry by its identifier
// The ID must be previously obtained through other methods
func (c *Client) GetARPEntryByID(ctx context.Context, id model.ID) (ARPEntry, error) {
	var arpEntry ARPEntry
	err := c.do(ctx, http.MethodGet, c.baseURL+"/ip/arp/"+id.String(), http.NoBody, &arpEntry)
	return arpEntry, err
}

// CreateARPEntry creates a new ARP entry and returns it
func (c *Client) CreateARPEntry(ctx context.Context, params CreateARPEntryParams) (ARPEntry, error) {
	var arpEntry ARPEntry
	err := c.do(ctx, http.MethodPut, c.baseURL+"/ip/arp", params, &arpEntry)
	return arpEntry, err
}

// UpdateARPEntry updates an existing ARP entry and returns its new state
// The ID must be previously obtained through other methods
func (c *Client) UpdateARPEntry(ctx context.Context, id model.ID, params UpdateARPEntryParams) (ARPEntry, error) {
	var arpEntry ARPEntry
	err := c.do(ctx, http.MethodPatch, c.baseURL+"/ip/arp/"+id.String(), params, &arpEntry)
	return arpEntry, err
}

// DeleteARPEntry deletes an existing ARP entry
// The ID must be previously obtained through other methods
func (c *Client) DeleteARPEntry(ctx context.Context, id model.ID) error {
	return c.do(ctx, http.MethodDelete, c.baseURL+"/ip/arp/"+id.String(), http.NoBody, http.NoBody)
}
