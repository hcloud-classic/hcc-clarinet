package model

import (
	"hcc/clarinet/lib/errors"
	"time"
)

// AdaptiveIP - ish
type AdaptiveIP struct {
	UUID           string               `json:"uuid"`
	NetworkAddress string               `json:"network_address"`
	Netmask        string               `json:"netmask"`
	Gateway        string               `json:"gateway"`
	StartIPAddress string               `json:"start_ip_address"`
	EndIPAddress   string               `json:"end_ip_address"`
	CreatedAt      time.Time            `json:"created_at"`
	Errors         errors.HccErrorStack `json:"errors"`
}

// AdaptiveIPs - ish
type AdaptiveIPs struct {
	AdaptiveIP []AdaptiveIP         `json:"adaptiveip"`
	Errors     errors.HccErrorStack `json:"errors"`
}

type AvailableIPList struct {
	AvailableIPs []string             `json:"available_ip_list"`
	Errors       errors.HccErrorStack `json:"errors"`
}

// AdaptiveIPNum - ish
type AdaptiveIPNum struct {
	Number int                  `json:"number"`
	Errors errors.HccErrorStack `json:"errors"`
}
