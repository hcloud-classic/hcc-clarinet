package model

import errors "innogrid.com/hcloud-classic/hcc_errors"

type Login struct {
	Token  string            `json:"token"`
	Errors []errors.HccError `json:"errors"`
}

type Valid struct {
	Errors  []errors.HccError `json:"errors"`
	IsValid bool              `json:"isvalid"`
}
