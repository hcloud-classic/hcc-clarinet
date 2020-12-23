package model

import errors "github.com/hcloud-classic/hcc_errors"

// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	//AdaptiveIPUUID string               `json:"adaptiveip_uuid"`
	ServerUUID     string               `json:"server_uuid"`
	PublicIP       string               `json:"public_ip"`
	PrivateIP      string               `json:"private_ip"`
	PrivateGateway string               `json:"private_gateway"`
	CreatedAt      string               `json:"created_at"`
	Errors         errors.HccErrorStack `json:"errors"`
}

// AdaptiveIPServers - ish
type AdaptiveIPServers struct {
	AdaptiveIPServers []AdaptiveIPServer   `json:"adaptiveip_server_list"`
	Errors            errors.HccErrorStack `json:"errors"`
}

// AdaptiveIPServerNum - ish
type AdaptiveIPServerNum struct {
	Number int                  `json:"number"`
	Errors errors.HccErrorStack `json:"errors"`
}
