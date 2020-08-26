package config

type user struct {
	UserId     string `goconf:"user:user_id"`
	UserPasswd string `goconf:"user:user_passwd"`
	Token      string `goconf:"user:token"`
}

// Flute : user config structure
var User user
