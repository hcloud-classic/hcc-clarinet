package model

import errors "innogrid.com/hcloud-classic/hcc_errors"

// Login : Login structure for use clarinet
type Login struct {
	Token  string            `json:"token"`
	Errors []errors.HccError `json:"errors"`
}

// Valid : Valid structure for validate user login in clarinet
type Valid struct {
	Errors  []errors.HccError `json:"errors"`
	IsValid bool              `json:"isvalid"`
}
