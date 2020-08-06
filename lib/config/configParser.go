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
	cnt := 0
RETRY:
	if cnt > 3 {
		log.Panic("Cannot Create User configure file (~/.hcc/clarinet/user.conf)")
	}
	config.UserConfig = usrConf.Get("user")
	if config.UserConfig == nil {
		cnt++
		GetUserInfo()
		goto RETRY
	}

	User = user{}
	User.UserId, err = config.UserConfig.String("user_id")
	if err != nil {
		cnt++
		GetUserInfo()
		goto RETRY

	}
	User.UserPasswd, err = config.UserConfig.String("user_passwd")
	if err != nil {
		cnt++
		GetUserInfo()
		goto RETRY
	}

	return

}

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

func parseHarp() {
	config.HarpConfig = conf.Get("harp")
	if config.HarpConfig == nil {
		log.Panic("no harp section")
	}

	Harp = harp{}
	Harp.ServerAddress, err = config.HarpConfig.String("harp_server_address")
	if err != nil {
		log.Panic(err)
	}

	Harp.ServerPort, err = config.HarpConfig.Int("harp_server_port")
	if err != nil {
		log.Panic(err)
	}

	Harp.RequestTimeoutMs, err = config.HarpConfig.Int("harp_request_timeout_ms")
	if err != nil {
		log.Panic(err)
	}
}

func parseViolin() {
	config.ViolinConfig = conf.Get("violin")
	if config.ViolinConfig == nil {
		log.Panic("no violin section")
	}

	Violin = violin{}
	Violin.ServerAddress, err = config.ViolinConfig.String("violin_server_address")
	if err != nil {
		log.Panic(err)
	}

	Violin.ServerPort, err = config.ViolinConfig.Int("violin_server_port")
	if err != nil {
		log.Panic(err)
	}

	Violin.RequestTimeoutMs, err = config.ViolinConfig.Int("violin_request_timeout_ms")
	if err != nil {
		log.Panic(err)
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		log.Panic(err)
	}

	setUserConfFilePath()
	if err = usrConf.Parse(userConfLocation); err != nil {
		GetUserInfo()
	}
	parseUser()
	parseFlute()
	parseHarp()
	parseViolin()
}
