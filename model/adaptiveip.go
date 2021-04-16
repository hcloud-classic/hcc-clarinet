package model

import errors "innogrid.com/hcloud-classic/hcc_errors"

// AdaptiveIP - ish
type AdaptiveIP struct {
	UUID            string            `json:"uuid"`
	ExtIfaceAddress string            `json:"ext_ifaceip_address"`
	Netmask         string            `json:"netmask"`
	Gateway         string            `json:"gateway_address"`
	StartIPAddress  string            `json:"start_ip_address"`
	EndIPAddress    string            `json:"end_ip_address"`
	Errors          []errors.HccError `json:"errors"`
}

// AdaptiveIPs - ish
type AdaptiveIPs struct {
	AdaptiveIP []AdaptiveIP      `json:"adaptiveip"`
	Errors     []errors.HccError `json:"errors"`
}

type AvailableIPList struct {
	AvailableIPs []string          `json:"available_ip_list"`
	Errors       []errors.HccError `json:"errors"`
}

// AdaptiveIPNum - ish
type AdaptiveIPNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}
