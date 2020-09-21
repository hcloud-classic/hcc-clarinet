package model

import "hcc/clarinet/lib/errors"

type Login struct {
	Token  string               `json:"token"`
	Errors errors.HccErrorStack `json:"errors"`
}

type Valid struct {
	IsValid bool                 `json:"isvalid"`
	Errors  errors.HccErrorStack `json:"errors"`
}
