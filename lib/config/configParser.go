package config

import (
	"github.com/Terry-Mao/goconf"

	"log"
)

var conf = goconf.New()
var usrConf = goconf.New()
var config = ClarinetConfig{}
var err error

func parseUser() {
	User = user{}
	config.UserConfig = usrConf.Get("user")
	if config.UserConfig == nil {
		log.Panic("Cannot parsing user config")
	} else {
		User.Token, err = config.UserConfig.String("token")
		if err != nil {
			log.Panic("Cannot parsing user token")
		}
	}
}

func parsePiccolo() {
	config.PiccoloConfig = conf.Get("piccolo")
	if config.PiccoloConfig == nil {
		log.Panic("no piccolo section")
	}

	Piccolo = piccolo{}
	Piccolo.ServerAddress, err = config.PiccoloConfig.String("piccolo_server_address")
	if err != nil {
		log.Panic(err)
	}

	Piccolo.ServerPort, err = config.PiccoloConfig.Int("piccolo_server_port")
	if err != nil {
		log.Panic(err)
	}

	Piccolo.RequestTimeoutMs, err = config.PiccoloConfig.Int("piccolo_request_timeout_ms")
	if err != nil {
		log.Panic(err)
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		log.Panic(err)
	}
	parsePiccolo()

	setUserConfFilePath()
	if err = usrConf.Parse(userConfLocation); err != nil {
		if err = createConfFile(); err != nil {
			log.Panic(err)
		}
	}
	parseUser()
}
