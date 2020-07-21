package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/model"
	"os"
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
var cpu string
var memory string
var diskSize string
var status string
var userUUID string
var nrNode string
var row string
var page string
var uuid string

var serverCreate = &cobra.Command{
	Use:   "create",
	Short: "Create server.",
	Long:  `Create server with given information. memroy & disk size assign to GB.`,
	Example: `	clarinet server create --subnet_uuid "string" --os "string" --server_name "string" -- server_desc "description string" --cpu 4 --memory 2 --disk_size 10 --user_uuid "string" --nr_node 3`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
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

		fmt.Println("Create server SUCCESS\nUUID\t" + server.(model.Server).UUID)
	},
}

var serverList = &cobra.Command{
	Use:   "list",
	Short: "Get list of servers.",
	Long:  `Get list of servers with filters.`,
	Example: `	clarinet server list				get all server list.
	clarinet server list --row 1 --page 5		get paged all server list
							row & page cannot use alone.
	clarinet server list [--filter] [--row & --page]
							get filtered server list.
`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		var servers interface{}
		var err error

		queryArgs["row"] = row
		queryArgs["page"] = page
		queryArgs["uuid"] = uuid
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["os"] = _os
		queryArgs["server_name"] = serverName
		queryArgs["server_desc"] = serverDesc
		queryArgs["cpu"] = cpu
		queryArgs["memory"] = memory
		queryArgs["disk_size"] = diskSize
		queryArgs["status"] = status
		queryArgs["user_uuid"] = userUUID

		servers, err = queryParser.ListServer(queryArgs)

		if err != nil {
			fmt.Println(err)
			return
		}

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"No", "UUID", "Server Name", "Cores", "Memory", "Disk", "Nodes", "Status"})

		for index, server := range servers.([]model.Server) {
			serverUUIDArg := make(map[string]string)
			serverUUIDArg["server_uuid"] = server.UUID
			numNodesServer, err := queryParser.NumNodesServer(serverUUIDArg)
			if err != nil {
				fmt.Println(err)
				return
			}
			t.AppendRow([]interface{}{
				index + 1, server.UUID, server.ServerName, server.CPU, server.Memory, server.DiskSize,
				numNodesServer.(model.ServerNodeNum).Number, server.Status})
		}

		t.AppendFooter(table.Row{"Total", len(servers.([]model.Server))})
		t.Render()
	},
}

var serverUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update server information.",
	Long:  `Update server info by given information.`,
	Example: `	clarinet server update --uuid "uuidstring" [flags]`,

	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		if uuid == "" {
			fmt.Println("Empty server uuid")
			return
		}

		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["os"] = _os
		queryArgs["server_name"] = serverName
		queryArgs["server_desc"] = serverDesc
		queryArgs["cpu"] = cpu
		queryArgs["memory"] = memory
		queryArgs["status"] = status
		queryArgs["disk_size"] = diskSize
		queryArgs["user_uuid"] = userUUID

		server, err := mutationParser.UpdateServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully update server (" + server.(model.Server).UUID + ") information.")
	},
}

var serverDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete one of server by UUID.",
	Long:  `Delete one of server by UUID.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		_, err := mutationParser.DeleteServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully delete server (" + uuid + ").")
	},
}

func ReadyServerCmd() {
	serverCreate.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverCreate.Flags().StringVar(&_os, "os", "", "Type of OS")
	serverCreate.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverCreate.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverCreate.Flags().StringVar(&cpu, "cpu", "0", "Number of CPU cores")
	serverCreate.Flags().StringVar(&memory, "memory", "0", "Size of memory")
	serverCreate.Flags().StringVar(&diskSize, "disk_size", "0", "Size of disk")
	serverCreate.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")
	serverCreate.Flags().StringVar(&nrNode, "nr_node", "0", "Number of nodes")

	serverList.Flags().StringVar(&row, "row", "0", "rows of server list")
	serverList.Flags().StringVar(&page, "page", "0", "page of server list")
	serverList.Flags().StringVar(&uuid, "uuid", "", "UUID of server")
	serverList.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverList.Flags().StringVar(&_os, "os", "", "Type of OS")
	serverList.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverList.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverList.Flags().StringVar(&cpu, "cpu", "0", "Number of CPU cores")
	serverList.Flags().StringVar(&memory, "memory", "0", "Size of memory")
	serverList.Flags().StringVar(&diskSize, "disk_size", "0", "Size of disk")
	serverList.Flags().StringVar(&status, "status", "", "Server Status [Running | Stop]")
	serverList.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")

	serverUpdate.Flags().StringVar(&uuid, "uuid", "", "UUID of server")
	serverUpdate.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverUpdate.Flags().StringVar(&_os, "os", "", "Type of OS")
	serverUpdate.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverUpdate.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverUpdate.Flags().StringVar(&cpu, "cpu", "0", "Number of CPU cores")
	serverUpdate.Flags().StringVar(&memory, "memory", "0", "Size of memory")
	serverUpdate.Flags().StringVar(&diskSize, "disk_size", "0", "Size of disk")
	serverUpdate.Flags().StringVar(&status, "status", "", "Server Status [Running | Stop]")
	serverUpdate.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")

	serverDelete.Flags().StringVar(&uuid, "uuid", "", "UUID of server")

	ServerCmd.AddCommand(serverCreate, serverList, serverUpdate, serverDelete)
}
