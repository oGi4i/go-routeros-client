package go_routeros_client

import (
	"context"
	"net/http"

	"github.com/ogi4i/go-routeros-client/model"
	"github.com/ogi4i/go-routeros-client/query"
)

type (
	// EthernetInterface represents an ethernet interface from `/interface/ethernet`
	EthernetInterface struct {
		ID                      model.ID                 `json:".id"`
		Advertise               model.Advertise          `json:"advertise"`
		ARPMode                 model.ARPMode            `json:"arp"`
		ARPTimeout              model.DurationWithAuto   `json:"arp-timeout"`
		AutoNegotiation         bool                     `json:"auto-negotiation,string"`
		CableSettings           model.CableSettings      `json:"cable-settings"`
		DefaultName             string                   `json:"default-name"`
		DisableRunningCheck     bool                     `json:"disable-running-check,string"`
		Disabled                bool                     `json:"disabled,string"`
		FullDuplex              bool                     `json:"full-duplex,string"`
		LoopProtect             model.ToggleWithDefault  `json:"loop-protect"`
		LoopProtectDisableTime  model.Duration           `json:"loop-protect-disable-time"`
		LoopProtectSendInterval model.Duration           `json:"loop-protect-send-interval"`
		LoopProtectStatus       model.Status             `json:"loop-protect-status"`
		MACAddress              model.MACAddress         `json:"mac-address"`
		MTU                     uint16                   `json:"mtu,string"`
		Name                    string                   `json:"name"`
		OriginalMACAddress      model.MACAddress         `json:"orig-mac-address"`
		Running                 bool                     `json:"running,string"`
		RXBytes                 uint64                   `json:"rx-bytes,string"`
		RXDrops                 uint64                   `json:"rx-drop,string"`
		RXPackets               uint64                   `json:"rx-packet,string"`
		Speed                   model.Speed              `json:"speed"`
		TXBytes                 uint64                   `json:"tx-bytes,string"`
		TXPackets               uint64                   `json:"tx-packet,string"`
		ComboMode               *model.ComboMode         `json:"combo-mode,omitempty"`
		FECMode                 *model.FECMode           `json:"fec-mode,omitempty"`
		L2MTU                   *uint16                  `json:"l2mtu,omitempty"`
		MDIXEnable              *bool                    `json:"mdix-enable,omitempty"`
		RXFlowControl           *model.ToggleWithAuto    `json:"rx-flow-control,omitempty"`
		SFPRateSelectMode       *model.SFPRateSelectMode `json:"sfp-rate-select,omitempty"`
		SFPShutdownTemperature  *uint8                   `json:"sfp-shutdown-temperature,omitempty"`
		TXFlowControl           *model.ToggleWithAuto    `json:"tx-flow-control,omitempty"`
		Comment                 *string                  `json:"comment,omitempty"`
	}

	getEthernetInterfaceParams struct{}

	// UpdateEthernetInterfaceParams represents parameters with which an existing ethernet interface can be updated
	UpdateEthernetInterfaceParams struct {
		Advertise               *model.Advertise         `json:"advertise,omitempty"`
		ARPMode                 *model.ARPMode           `json:"arp,omitempty"`
		ARPTimeout              *model.DurationWithAuto  `json:"arp-timeout,omitempty"`
		AutoNegotiation         *bool                    `json:"auto-negotiation,omitempty,string"`
		CableSettings           *model.CableSettings     `json:"cable-settings,omitempty"`
		DisableRunningCheck     *bool                    `json:"disable-running-check,omitempty,string"`
		Disabled                *bool                    `json:"disabled,omitempty,string"`
		LoopProtect             *model.ToggleWithDefault `json:"loop-protect,omitempty"`
		LoopProtectDisableTime  *model.Duration          `json:"loop-protect-disable-time,omitempty"`
		LoopProtectSendInterval *model.Duration          `json:"loop-protect-send-interval,omitempty"`
		MACAddress              *model.MACAddress        `json:"mac-address,omitempty"`
		MTU                     *uint16                  `json:"mtu,omitempty,string"`
		Name                    *string                  `json:"name,omitempty"`
		OriginalMACAddress      *model.MACAddress        `json:"orig-mac-address,omitempty"`
		ComboMode               *model.ComboMode         `json:"combo-mode,omitempty"`
		FECMode                 *model.FECMode           `json:"fec-mode,omitempty"`
		L2MTU                   *uint16                  `json:"l2mtu,omitempty"`
		MDIXEnable              *bool                    `json:"mdix-enable,omitempty"`
		RXFlowControl           *model.ToggleWithAuto    `json:"rx-flow-control,omitempty"`
		SFPRateSelectMode       *model.SFPRateSelectMode `json:"sfp-rate-select,omitempty"`
		SFPShutdownTemperature  *uint8                   `json:"sfp-shutdown-temperature,omitempty"`
		TXFlowControl           *model.ToggleWithAuto    `json:"tx-flow-control,omitempty"`
		Comment                 *string                  `json:"comment,omitempty"`
	}
)

// GetEthernetInterfaceParams creates a builder for query parameters
func GetEthernetInterfaceParams() getEthernetInterfaceParams {
	return getEthernetInterfaceParams{}
}

func (gip getEthernetInterfaceParams) Advertise() query.Param[model.Advertise] {
	return "advertise"
}

func (gip getEthernetInterfaceParams) ARPMode() query.Param[model.ARPMode] {
	return "arp"
}

func (gip getEthernetInterfaceParams) ARPTimeout() query.Param[model.DurationWithAuto] {
	return "arp-timeout"
}

func (gip getEthernetInterfaceParams) AutoNegotiation() query.Param[bool] {
	return "auto-negotiation"
}

func (gip getEthernetInterfaceParams) CableSettings() query.Param[model.CableSettings] {
	return "cable-settings"
}

func (gip getEthernetInterfaceParams) DefaultName() query.Param[string] {
	return "default-name"
}

func (gip getEthernetInterfaceParams) DisableRunningCheck() query.Param[bool] {
	return "disable-running-check"
}

func (gip getEthernetInterfaceParams) Disabled() query.Param[bool] {
	return "disabled"
}

func (gip getEthernetInterfaceParams) FullDuplex() query.Param[bool] {
	return "full-duplex"
}

func (gip getEthernetInterfaceParams) LoopProtect() query.Param[model.ToggleWithDefault] {
	return "loop-protect"
}

func (gip getEthernetInterfaceParams) LoopProtectDisableTime() query.Param[model.Duration] {
	return "loop-protect-disable-time"
}

func (gip getEthernetInterfaceParams) LoopProtectSendInterval() query.Param[model.Duration] {
	return "loop-protect-send-interval"
}

func (gip getEthernetInterfaceParams) LoopProtectStatus() query.Param[model.Status] {
	return "loop-protect-status"
}

func (gip getEthernetInterfaceParams) MACAddress() query.Param[model.MACAddress] {
	return "mac-address"
}

func (gip getEthernetInterfaceParams) MTU() query.Param[uint16] {
	return "mtu"
}

func (gip getEthernetInterfaceParams) Name() query.Param[string] {
	return "name"
}

func (gip getEthernetInterfaceParams) OriginalMACAddress() query.Param[model.MACAddress] {
	return "orig-mac-address"
}

func (gip getEthernetInterfaceParams) Running() query.Param[bool] {
	return "running"
}

func (gip getEthernetInterfaceParams) RXBytes() query.Param[uint64] {
	return "rx-bytes"
}

func (gip getEthernetInterfaceParams) RXDrops() query.Param[uint64] {
	return "rx-drop"
}

func (gip getEthernetInterfaceParams) RXPackets() query.Param[uint64] {
	return "rx-packet"
}

func (gip getEthernetInterfaceParams) Speed() query.Param[model.Speed] {
	return "speed"
}

func (gip getEthernetInterfaceParams) TXBytes() query.Param[uint64] {
	return "tx-bytes"
}

func (gip getEthernetInterfaceParams) TXPackets() query.Param[uint64] {
	return "tx-packet"
}

func (gip getEthernetInterfaceParams) ComboMode() query.Param[model.ComboMode] {
	return "combo-mode"
}

func (gip getEthernetInterfaceParams) FECMode() query.Param[model.FECMode] {
	return "fec-mode"
}

func (gip getEthernetInterfaceParams) L2MTU() query.Param[uint64] {
	return "l2mtu"
}

func (gip getEthernetInterfaceParams) MDIXEnable() query.Param[bool] {
	return "mdix-enable"
}

func (gip getEthernetInterfaceParams) RXFlowControl() query.Param[model.ToggleWithAuto] {
	return "rx-flow-control"
}

func (gip getEthernetInterfaceParams) SFPRateSelectMode() query.Param[model.SFPRateSelectMode] {
	return "sfp-rate-select"
}

func (gip getEthernetInterfaceParams) SFPShutdownTemperature() query.Param[uint8] {
	return "sfp-shutdown-temperature"
}

func (gip getEthernetInterfaceParams) TXFlowControl() query.Param[model.ToggleWithAuto] {
	return "tx-flow-control"
}

func (gip getEthernetInterfaceParams) Comment() query.Param[string] {
	return "comment"
}

// ListEthernetInterface returns a list of all ethernet interfaces
// Same as FilterEthernetInterfaces without any filters
func (c *Client) ListEthernetInterface(ctx context.Context) ([]EthernetInterface, error) {
	var ethernetInterfaces []EthernetInterface
	err := c.do(ctx, http.MethodGet, c.baseURL+"/interface/ethernet", http.NoBody, &ethernetInterfaces)
	return ethernetInterfaces, err
}

// FilterEthernetInterfaces returns a list of ethernet interfaces which pass a given list of Filters
// Without any Filters is the same as ListEthernetInterface
func (c *Client) FilterEthernetInterfaces(ctx context.Context, filters ...query.Filter) ([]EthernetInterface, error) {
	var ethernetInterfaces []EthernetInterface
	err := c.do(ctx, http.MethodPost, c.baseURL+"/interface/ethernet/print", prepareQuery(filters...), &ethernetInterfaces)
	return ethernetInterfaces, err
}

// GetEthernetInterfaceByID returns a single ethernet interface by its identifier
// The ID must be previously obtained through other methods
func (c *Client) GetEthernetInterfaceByID(ctx context.Context, id model.ID) (EthernetInterface, error) {
	var ethernetInterface EthernetInterface
	err := c.do(ctx, http.MethodGet, c.baseURL+"/interface/ethernet/"+id.String(), http.NoBody, &ethernetInterface)
	return ethernetInterface, err
}

// UpdateEthernetInterface updates an existing ethernet interface and returns its new state
// The ID must be previously obtained through other methods
func (c *Client) UpdateEthernetInterface(ctx context.Context, id model.ID, params UpdateEthernetInterfaceParams) (EthernetInterface, error) {
	var ethernetInterface EthernetInterface
	err := c.do(ctx, http.MethodPatch, c.baseURL+"/interface/ethernet/"+id.String(), params, &ethernetInterface)
	return ethernetInterface, err
}
