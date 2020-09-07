package model

import (
	"hcc/clarinet/lib/errors"
	"time"
)

// ServerNode - cgs
type ServerNode struct {
	UUID       string               `json:"uuid"`
	ServerUUID string               `json:"server_uuid"`
	NodeUUID   string               `json:"node_uuid"`
	CreatedAt  time.Time            `json:"created_at"`
	Errors     errors.HccErrorStack `json:"errors"`
}

// ServerNodes - cgs
type ServerNodes struct {
	Server []Server             `json:"server_node"`
	Errors errors.HccErrorStack `json:"errors"`
}

// ServerNodeNum - ish
type ServerNodeNum struct {
	Number int                  `json:"number"`
	Errors errors.HccErrorStack `json:"errors"`
}
