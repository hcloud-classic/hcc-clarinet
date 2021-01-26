package model

import (
	errors "github.com/hcloud-classic/hcc_errors"
	"time"
)

// Node - cgs
type Node struct {
	UUID        string               `json:"uuid"`
	ServerUUID  string               `json:"server_uuid"`
	BmcMacAddr  string               `json:"bmc_mac_addr"`
	BmcIP       string               `json:"bmc_ip"`
	PXEMacAddr  string               `json:"pxe_mac_addr"`
	Status      string               `json:"status"`
	CPUCores    int                  `json:"cpu_cores"`
	Memory      int                  `json:"memory"`
	Description string               `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	Active      int                  `json:"active"`
	ForceOff    bool                 `json:"force_off"`
	Errors      errors.HccErrorStack `json:"errors"`
}

type NodeDetail struct {
	NodeUUID      string               `json:"node_uuid"`
	CPUModel      string               `json:"cpu_model"`
	CPUProcessors int                  `json:"cpu_processors"`
	CPUThreads    int                  `json:"cpu_threads"`
	Errors        errors.HccErrorStack `json:"errors"`
}

// Nodes - cgs
type Nodes struct {
	Nodes  []Node               `json:"node_list"`
	Errors errors.HccErrorStack `json:"errors"`
}

// NodeNum - cgs
type NodeNum struct {
	Number int                  `json:"number"`
	Errors errors.HccErrorStack `json:"errors"`
}

// PowerState - younseok.shim
type PowerState int

const (
	On PowerState = 1 + iota
	Off
	Restart
)

type PowerStateNode struct {
	State  string               `json:"power_state"`
	Errors errors.HccErrorStack `json:"errors"`
}
