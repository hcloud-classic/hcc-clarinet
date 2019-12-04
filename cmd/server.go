package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/model"
	"os"
)

var ServerCmd = &cobra.Command{
	Use:   "server [server options...]",
	Short: "Running server commands",
	Long: `server: Running server related commands.`,
	Args: cobra.MinimumNArgs(1),
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
	Long: `Create server.`,
	Args: cobra.MinimumNArgs(0),
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

var serverList = &cobra.Command{
	Use:   "list",
	Short: "Get list of servers with row and page options.",
	Long: `Get list of servers with row and page options.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]interface{})
		if row != 0 || page != 0 {
			if row <= 0 || page <= 0 {
				fmt.Println("Please provide row and page options correctly.")
				return
			}
			queryArgs["row"] = row
			queryArgs["page"] = page
		}
		servers, err := queryParser.AllServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		serverNum, err := queryParser.NumServer()
		if err != nil {
			fmt.Println(err)
			return
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"No", "UUID", "Server Name", "Cores", "Memory", "Disk", "Nodes", "Status"})
		for i, server := range servers.([]model.Server) {
			serverUUIDArg := make(map[string]interface{})
			serverUUIDArg["server_uuid"] = server.UUID
			numNodesServer, err := queryParser.NumNodesServer(serverUUIDArg)
			if err != nil {
				fmt.Println(err)
				return
			}
			t.AppendRow([]interface{}{i + 1, server.UUID, server.ServerName, server.CPU, server.Memory, server.DiskSize,
				numNodesServer.(model.ServerNodeNum).Number, server.Status})
		}
		t.AppendFooter(table.Row{"Total Server Num", serverNum.(model.ServerNum).Number, "", "", "", "", "", ""})
		t.Render()
	},
}

var uuid string

var serverDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete one of server by UUID.",
	Long: `Delete one of server by UUID.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]interface{})
		queryArgs["uuid"] = uuid
		_, err := mutationParser.DeleteServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Delete command ended successfully.")
	},
}

func ReadyServerCmd() {
	serverCreate.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverCreate.Flags().StringVar(&_os, "os", "", "Type of OS")
	serverCreate.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverCreate.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverCreate.Flags().IntVar(&cpu, "cpu", 0, "Number of CPU cores")
	serverCreate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	serverCreate.Flags().IntVar(&diskSize, "disk_size", 0, "Size of disk")
	serverCreate.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")
	serverCreate.Flags().IntVar(&nrNode, "nr_node", 0, "Number of nodes")

	serverList.Flags().IntVar(&row, "row", 0, "rows of server list")
	serverList.Flags().IntVar(&page, "page", 0, "page of server list")

	serverDelete.Flags().StringVar(&uuid, "uuid", "", "UUID of server")

	ServerCmd.AddCommand(serverCreate, serverList, serverDelete)
}
