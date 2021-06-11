package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/model"
)

var subnetUUID string
var OS string
var serverName string
var serverDesc string
var cpu int
var memory int
var diskSize int
var status string
var userUUID string
var nrNode int
var row int
var page int
var uuid string

var serverCmd = &cobra.Command{
	Use:     `server --uuid "serverUUID"`,
	Short:   "Get server Information",
	Long:    `Show server information by given serverUUID at table`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["token"] = config.User.Token

		data, err := queryParser.Server(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		t := table.NewWriter()

		ts := table.Style{
			Box: table.StyleBoxLight,
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
		}

		serverData := data.(model.Server)
		if len(serverData.Errors) > 0 {
			for _, hrr := range serverData.Errors {
				hrr.Println()
			}
		} else {

			t.SetStyle(ts)
			t.SetOutputMirror(os.Stdout)
			t.SetTitle("Server Info\n%s", uuid)

			t.AppendRow(table.Row{"Name", serverData.ServerName})
			t.AppendRow(table.Row{"OS", serverData.OS})
			t.AppendRow(table.Row{"CPU", serverData.CPU})
			t.AppendRow(table.Row{"Memory", serverData.Memory})
			t.AppendRow(table.Row{"Disk", serverData.DiskSize})
			t.AppendRow(table.Row{"Status", serverData.Status})
			t.AppendRow(table.Row{"Description", serverData.ServerDesc})
			t.AppendRow(table.Row{"Created At", serverData.CreatedAt})

			t.Render()
		}
		queryArgs["server_uuid"] = uuid
		delete(queryArgs, "uuid")

		data, err = queryParser.ListServerNode(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		serverNodeList := data.(model.ServerNodes)
		if len(serverNodeList.Errors) > 0 {
			for _, hrr := range serverNodeList.Errors {
				hrr.Println()
			}
			return
		}

		t = table.NewWriter()
		t.SetStyle(ts)
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{"Node UUID", "CPU Model", "sockets", "Cores", "Threads", "Memory", "Created At"})
		for _, node := range serverNodeList.NodeList {
			t.AppendRow(table.Row{node.NodeUUID, node.CPUModel, node.CPUSocket, node.CPUCores, node.CPUThreads, node.Memory, node.CreatedAt})
		}
		t.Render()
	},
}

var serverCreate = &cobra.Command{
	Use:   "create",
	Short: "Create server.",
	Long:  `Create server with given information. memroy & disk size assign to GB.`,
	Example: `	clarinet server create --subnet_uuid "string" --os "string" --server_name "string" -- server_desc "description string" --cpu 4 --memory 2 --disk_size 10 --user_uuid "string" --nr_node 3`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["os"] = OS
		queryArgs["server_name"] = serverName
		queryArgs["server_desc"] = serverDesc
		queryArgs["cpu"] = strconv.Itoa(cpu)
		queryArgs["memory"] = strconv.Itoa(memory)
		queryArgs["disk_size"] = strconv.Itoa(diskSize)
		queryArgs["user_uuid"] = userUUID
		queryArgs["nr_node"] = strconv.Itoa(nrNode)
		queryArgs["token"] = config.User.Token

		fmt.Print("Create Server .... ")

		data, err := mutationParser.CreateServer(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		serverData := data.(model.Server)
		if len(serverData.Errors) > 0 {
			for _, hrr := range serverData.Errors {
				hrr.Println()
			}
			return
		}

		serverUUIDArg := make(map[string]string)
		serverUUIDArg["server_uuid"] = serverData.UUID
		serverUUIDArg["token"] = config.User.Token

		data, err = queryParser.NumNodesServer(serverUUIDArg)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		numNodeData := data.(model.ServerNodeNum)
		if len(numNodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range numNodeData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")

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
		t.AppendHeader(table.Row{"UUID", serverData.UUID})
		t.AppendRow([]interface{}{"NAME", serverData.ServerName})
		t.AppendRow([]interface{}{"CORES", serverData.CPU})
		t.AppendRow([]interface{}{"MEMORY", serverData.Memory})
		t.AppendRow([]interface{}{"DISK", serverData.DiskSize})
		t.AppendRow([]interface{}{"NODES", numNodeData.Number})
		t.AppendRow([]interface{}{"STATUS", serverData.Status})
		t.Render()
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
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["uuid"] = uuid
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["os"] = OS
		queryArgs["server_name"] = serverName
		queryArgs["server_desc"] = serverDesc
		queryArgs["cpu"] = strconv.Itoa(cpu)
		queryArgs["memory"] = strconv.Itoa(memory)
		queryArgs["disk_size"] = strconv.Itoa(diskSize)
		queryArgs["status"] = status
		queryArgs["user_uuid"] = userUUID
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		serverList := data.(model.Servers)
		if len(serverList.Errors) > 0 {
			for _, hrr := range serverList.Errors {
				hrr.Println()
			}
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

		for index, server := range serverList.Server {
			serverUUIDArg := make(map[string]string)
			serverUUIDArg["server_uuid"] = server.UUID
			serverUUIDArg["token"] = config.User.Token
			num, _ := queryParser.NumNodesServer(serverUUIDArg)
			t.AppendRow([]interface{}{
				index + 1, server.UUID, server.ServerName, server.CPU, server.Memory, server.DiskSize,
				num.(model.ServerNodeNum).Number, server.Status})
		}

		t.AppendFooter(table.Row{"Total", len(serverList.Server)})
		t.Render()
	},
}

var serverUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update server information.",
	Long:  `Update server info by given information.`,
	Example: `	clarinet server update --uuid "uuidstring" [flags]`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {

		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["os"] = OS
		queryArgs["server_name"] = serverName
		queryArgs["server_desc"] = serverDesc
		queryArgs["cpu"] = strconv.Itoa(cpu)
		queryArgs["memory"] = strconv.Itoa(memory)
		queryArgs["status"] = status
		queryArgs["disk_size"] = strconv.Itoa(diskSize)
		queryArgs["user_uuid"] = userUUID
		queryArgs["token"] = config.User.Token

		fmt.Print("Update Server .... ")

		data, err := mutationParser.UpdateServer(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		serverData := data.(model.Server)
		if len(serverData.Errors) > 0 {
			for _, hrr := range serverData.Errors {
				hrr.Println()
			}
			return
		}

		serverUUIDArg := make(map[string]string)
		serverUUIDArg["server_uuid"] = serverData.UUID
		serverUUIDArg["token"] = config.User.Token

		data, err = queryParser.NumNodesServer(serverUUIDArg)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		numNodeData := data.(model.ServerNodeNum)
		if len(numNodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range numNodeData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")

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
		t.AppendHeader(table.Row{"UUID", serverData.UUID})
		t.AppendRow([]interface{}{"NAME", serverData.ServerName})
		t.AppendRow([]interface{}{"CORES", serverData.CPU})
		t.AppendRow([]interface{}{"MEMORY", serverData.Memory})
		t.AppendRow([]interface{}{"DISK", serverData.DiskSize})
		t.AppendRow([]interface{}{"NODES", numNodeData.Number})
		t.AppendRow([]interface{}{"STATUS", serverData.Status})
		t.Render()
	},
}

var serverDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete one of server by UUID.",
	Long:    `Delete one of server by UUID.`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["token"] = config.User.Token

		fmt.Print("Delete Server .... ")

		data, err := mutationParser.DeleteServer(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		serverData := data.(model.Server)
		if len(serverData.Errors) > 0 {
			for _, hrr := range serverData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")

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
		t.AppendHeader(table.Row{"UUID", serverData.UUID})
		t.AppendRow([]interface{}{"NAME", serverData.ServerName})
		t.AppendRow([]interface{}{"CORES", serverData.CPU})
		t.AppendRow([]interface{}{"MEMORY", serverData.Memory})
		t.AppendRow([]interface{}{"DISK", serverData.DiskSize})
		t.AppendRow([]interface{}{"STATUS", serverData.Status})
		t.Render()
	},
}

func ReadyServerCmd() {

	serverCmd.AddCommand(serverCreate, serverList, serverUpdate, serverDelete)

	serverCmd.Flags().StringVar(&uuid, "uuid", "", "UUID of server")
	serverCmd.MarkFlagRequired("uuid")

	serverCreate.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverCreate.Flags().StringVar(&OS, "os", "", "Type of OS")
	serverCreate.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverCreate.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverCreate.Flags().IntVar(&cpu, "cpu", 0, "Number of CPU cores")
	serverCreate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	serverCreate.Flags().IntVar(&diskSize, "disk_size", 0, "Size of disk")
	serverCreate.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")
	serverCreate.Flags().IntVar(&nrNode, "nr_node", 0, "Number of nodes")
	serverCreate.MarkFlagRequired("subnet_uuid")

	serverList.Flags().IntVar(&row, "row", 0, "rows of server list")
	serverList.Flags().IntVar(&page, "page", 0, "page of server list")
	serverList.Flags().StringVar(&uuid, "uuid", "", "UUID of server")
	serverList.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverList.Flags().StringVar(&OS, "os", "", "Type of OS")
	serverList.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverList.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverList.Flags().IntVar(&cpu, "cpu", 0, "Number of CPU cores")
	serverList.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	serverList.Flags().IntVar(&diskSize, "disk_size", 0, "Size of disk")
	serverList.Flags().StringVar(&status, "status", "", "Server Status [Running | Stop]")
	serverList.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")

	serverUpdate.Flags().StringVar(&uuid, "uuid", "", "UUID of server")
	serverUpdate.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverUpdate.Flags().StringVar(&OS, "os", "", "Type of OS")
	serverUpdate.Flags().StringVar(&serverName, "server_name", "", "Name of server")
	serverUpdate.Flags().StringVar(&serverDesc, "server_desc", "", "Description of server")
	serverUpdate.Flags().IntVar(&cpu, "cpu", 0, "Number of CPU cores")
	serverUpdate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	serverUpdate.Flags().IntVar(&diskSize, "disk_size", 0, "Size of disk")
	serverUpdate.Flags().StringVar(&status, "status", "", "Server Status [Running | Stop]")
	serverUpdate.Flags().StringVar(&userUUID, "user_uuid", "", "UUID of user")
	serverUpdate.MarkFlagRequired("uuid")

	serverDelete.Flags().StringVar(&uuid, "uuid", "", "UUID of server")
	serverDelete.MarkFlagRequired("uuid")
}
