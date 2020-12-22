package config

import (
	"github.com/Terry-Mao/goconf"

	errors "github.com/hcloudclassic/hcc_errors"
)

var conf = goconf.New()
var usrConf = goconf.New()
var config = ClarinetConfig{}
var err error

func parseUser() {
	User = user{}
	config.UserConfig = usrConf.Get("user")
	if config.UserConfig == nil {
		errors.NewHccError(errors.ClarinetInternalParsingError, "user config").Fatal()
	} else {
		User.Token, err = config.UserConfig.String("token")
		if err != nil {
			errors.NewHccError(errors.ClarinetInternalParsingError, "user token").Fatal()
		}
	}
}

func parsePiccolo() {
	config.PiccoloConfig = conf.Get("piccolo")
	if config.PiccoloConfig == nil {
		errors.NewHccError(errors.ClarinetInternalParsingError, "piccolo config").Fatal()
	}

	Piccolo = piccolo{}
	Piccolo.ServerAddress, err = config.PiccoloConfig.String("piccolo_server_address")
	if err != nil {
		errors.NewHccError(errors.ClarinetInternalParsingError, "piccolo server address").Fatal()
	}

	Piccolo.ServerPort, err = config.PiccoloConfig.Int("piccolo_server_port")
	if err != nil {
		errors.NewHccError(errors.ClarinetInternalParsingError, "piccolo server port").Fatal()
	}

	Piccolo.RequestTimeoutMs, err = config.PiccoloConfig.Int("piccolo_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.ClarinetInternalParsingError, "piccolo timeout").Fatal()
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		errors.NewHccError(errors.ClarinetInternalParsingError, err.Error()).Fatal()
	}
	parsePiccolo()

	setUserConfFilePath()
	if err = usrConf.Parse(userConfLocation); err != nil {
		if err = createConfFile(); err != nil {
			errors.NewHccError(errors.ClarinetInternalParsingError, err.Error()).Fatal()
		}
	}
	parseUser()
}
