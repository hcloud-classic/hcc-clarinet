package config

import (
	"log"

	"github.com/Terry-Mao/goconf"
)

var conf = goconf.New()
var config = piccoloConfig{}
var err error

func parseFlute() {
	config.FluteConfig = conf.Get("flute")
	if config.FluteConfig == nil {
		log.Panic("no flute section")
	}

	Flute = flute{}
	Flute.ServerAddress, err = config.FluteConfig.String("flute_server_address")
	if err != nil {
		log.Panic(err)
	}

	Flute.ServerPort, err = config.FluteConfig.Int("flute_server_port")
	if err != nil {
		log.Panic(err)
	}

	Flute.RequestTimeoutMs, err = config.FluteConfig.Int("flute_request_timeout_ms")
	if err != nil {
		log.Panic(err)
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		log.Panic(err)
	}

	parseFlute()

}
