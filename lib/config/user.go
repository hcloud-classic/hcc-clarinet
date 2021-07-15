package config

type user struct {
	Token string `goconf:"user:token"`
}

// User : user config structure
var User user
