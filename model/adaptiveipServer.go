package model

import "hcc/clarinet/lib/errors"

// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	AdaptiveIPUUID string               `json:"adaptiveip_uuid"`
	ServerUUID     string               `json:"server_uuid"`
	PublicIP       string               `json:"public_ip"`
	PrivateIP      string               `json:"private_ip"`
	PrivateGateway string               `json:"private_gateway"`
	Errors         errors.HccErrorStack `json:"errors"`
}

// AdaptiveIPServers - ish
type AdaptiveIPServers struct {
	AdaptiveIPServers []AdaptiveIPServer   `json:"adaptiveip_server"`
	Errors            errors.HccErrorStack `json:"errors"`
}

// AdaptiveIPServerNum - ish
type AdaptiveIPServerNum struct {
	Number int                  `json:"number"`
	Errors errors.HccErrorStack `json:"errors"`
}
