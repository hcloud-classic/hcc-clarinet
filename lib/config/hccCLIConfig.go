package config

import "github.com/Terry-Mao/goconf"

var configLocation = "/etc/hcc/clarinet/clarinet.conf"

type piccoloConfig struct {
	FluteConfig       *goconf.Section
	HarpConfig        *goconf.Section
	ViolinConfig      *goconf.Section
}
