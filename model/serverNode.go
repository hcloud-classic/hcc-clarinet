package model

import (
	errors "innogrid.com/hcloud-classic/hcc_errors"
	"time"
)

// ServerNode - cgs
type ServerNode struct {
	NodeUUID   string            `json:"node_uuid"`
	CPUModel   string            `json:"cpu_model"`
	CPUSocket  int               `json:"cpu_processors"`
	CPUCores   int               `json:"cpu_cores"`
	CPUThreads int               `json:"cpu_threads"`
	Memory     int               `json:"memory"`
	CreatedAt  time.Time         `json:"created_at"`
	Errors     []errors.HccError `json:"errors"`
}

// ServerNodes - cgs
type ServerNodes struct {
	NodeList []ServerNode      `json:"server_node_list"`
	Errors   []errors.HccError `json:"errors"`
}

// ServerNodeNum - ish
type ServerNodeNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}
