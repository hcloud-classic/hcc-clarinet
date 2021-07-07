package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/model"
)

var Cmd *cobra.Command = nil

func Init() {
	ReadyServerCmd()
	ReadyNodeCmd()
	ReadySubnetCmd()
	ReadyAIPCmd()

	Cmd = &cobra.Command{Use: "clarinet"}
	Cmd.AddCommand(serverCmd, nodeCmd, subnetCmd, aipCmd, logoutCmd, versionCmd)
}

func checkToken(cmd *cobra.Command, args []string) error {
	if config.User.Token == "" {
		config.GetUserInfo()
		userInfo := make(map[string]string)
		userInfo["id"] = config.User.UserId
		userInfo["password"] = config.User.UserPasswd
		if loginData, err := queryParser.Login(userInfo); err != nil {
			log.Fatalf("Login Failed")
		} else {
			config.SaveTokenString(loginData.(model.Login).Token)
			config.User.UserId = ""
			config.User.UserPasswd = ""
		}
	} else {
		token := make(map[string]string)
		token["token"] = config.User.Token
		if isValid, err := queryParser.CheckToken(token); err != nil {
			log.Fatalf("Token Check Failed")
		} else {
			if !isValid.(model.Valid).IsValid {
				log.Println("Invalid token, Enter user info to login")
				config.User.Token = ""
				checkToken(cmd, args)
			}
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
