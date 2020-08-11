package model

import "time"

type AdaptiveIP struct {
	UUID           string    `json:"uuid"`
	NetworkAddress string    `json:"network_address"`
	Netmask        string    `json:"netmask"`
	Gateway        string    `json:"gateway"`
	StartIPAddress string    `json:"start_ip_address"`
	EndIPAddress   string    `json:"end_ip_address"`
	CreatedAt      time.Time `json:"created_at"`
}

type AdaptiveIPs struct {
	AdaptiveIP []Subnet `json:"adaptiveip"`
}

type AdaptiveIPNum struct {
	Number int `json:"number"`
}
