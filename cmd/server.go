package cmd

import (
	"fmt"
	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/model"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server [server options...]",
	Short: "Running server commands",
	Long:  `server: Running server related commands.`,
	Args:  cobra.MinimumNArgs(1),
}

var subnetUUID string
var _os string
var serverName string
var serverDesc string
var cpu int
var memory int
var diskSize int
var userUUID string
var nrNode int

var serverCreate = &cobra.Command{
	Use:   "create",
	Short: "Create server.",
	Long:  `Create server.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]interface{})
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["os"] = _os
		queryArgs["server_name"] = serverName
		queryArgs["server_desc"] = serverDesc
		queryArgs["cpu"] = cpu
		queryArgs["memory"] = memory
		queryArgs["disk_size"] = diskSize
		queryArgs["user_uuid"] = userUUID
		queryArgs["nr_node"] = nrNode
		server, err := mutationParser.CreateServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Creating server for UUID: " + server.(model.Server).UUID)
	},
}

var row int
var page int
