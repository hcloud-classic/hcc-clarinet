package init

import "hcc/clarinet/lib/config"

// MainInit : Main initialization function
func MainInit() error {
	config.Parser()
	return cmdInit()
}
