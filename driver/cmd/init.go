package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
)

var Cmd *cobra.Command = nil

func Init() {
	ReadyServerCmd()
	ReadyNodeCmd()
	ReadySubnetCmd()
	ReadyAIPCmd()

	Cmd = &cobra.Command{Use: "clarinet"}
	Cmd.AddCommand(serverCmd, nodeCmd, subnetCmd, aipCmd)
}

func checkToken(cmd *cobra.Command, args []string) error {
	if config.User.Token == "" {
		config.GetUserInfo()
		userInfo := make(map[string]string)
		userInfo["user_id"] = config.User.UserId
		userInfo["user_passwd"] = config.User.UserPasswd
		if tokenString, err := queryParser.Login(userInfo); err != nil {
			log.Fatalf("Login Failed")
		} else {
			config.SaveTokenString(tokenString.(string))
			config.User.UserId = ""
			config.User.UserPasswd = ""
		}
	}
	return nil
}

func reRunIfExpired(cmd *cobra.Command) {
	log.Println("Token expired.")
	config.User.Token = ""
	config.GetUserInfo()
	cmd.Execute()
}
