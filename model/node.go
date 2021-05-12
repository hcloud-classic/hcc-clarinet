package model

import (
	"time"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

// Node - cgs
type Node struct {
	UUID        string            `json:"uuid"`
	ServerUUID  string            `json:"server_uuid"`
	BmcMacAddr  string            `json:"bmc_mac_addr"`
	BmcIP       string            `json:"bmc_ip"`
	PXEMacAddr  string            `json:"pxe_mac_addr"`
	Status      string            `json:"status"`
	CPUCores    int               `json:"cpu_cores"`
	Memory      int               `json:"memory"`
	Description string            `json:"description"`
	CreatedAt   time.Time         `json:"created_at"`
	Active      int               `json:"active"`
	ForceOff    bool              `json:"force_off"`
	Errors      []errors.HccError `json:"errors"`
}

type NodeDetail struct {
	NodeUUID   string            `json:"node_uuid"`
	NodeDetail string            `json:"node_detail_data"`
	NicDetail  string            `json:"nic_detail_data"`
	Errors     []errors.HccError `json:"errors"`
}

type NodeDetailData struct {
	CPUs     []CPU    `json:"cpus"`
	Memories []Memory `json:"memories"`
	NICs     []NIC    `json:"nics"`
	Errors   []errors.HccError
}

type CPU struct {
	Cores       int    `json:"cores"`
	ID          string `json:"id"`
	Manufacture string `json:"manufacture"`
	MaxSpeed    int    `json:"max_speed_mhz"`
	Model       string `json:"model"`
	Socket      string `json:"socket"`
	Threads     int    `json:"threads"`
	Status      Status `json:"status"`
}

type Memory struct {
	CapacityMB    int    `json:"capacity_mb"`
	DeviceLocator string `json:"device_locator"`
	ID            string `json:"id"`
	Manufacture   string `json:"manufacture"`
	PartNumber    string `json:"part_number"`
	SerialNumber  string `json:"serial_number"`
	Speed         int    `json:"speed_mhz"`
	Status        Status `json:"status"`
}

type NIC struct {
	ID    string `json:"id"`
	Mac   string `json:"mac"`
	Model string `json:"model"`
	Speed string `json:"speed"`
	Type  string `json:"type"`
}

type Status struct {
	Health       string `json:"health"`
	HealthRollup string `json:"health_rollup"`
	State        string `json:"state"`
}

// Nodes - cgs
type Nodes struct {
	Nodes  []Node            `json:"node_list"`
	Errors []errors.HccError `json:"errors"`
}

// NodeNum - cgs
type NodeNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}

// PowerState - younseok.shim
type PowerState int

const (
	On PowerState = 1 + iota
	Off
	Restart
)

type PowerStateNode struct {
	State  string            `json:"power_state"`
	Errors []errors.HccError `json:"errors"`
}
