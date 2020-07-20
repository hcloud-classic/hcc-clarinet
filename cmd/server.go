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

var serverArgs struct {
	ServerInfo model.Server
	NrNode     int
	Row        int
	Page       int
}

var serverCreate = &cobra.Command{
	Use:   "create",
	Short: "Create server.",
	Long:  `Create server with given information. memroy & disk size assign to GB.`,
	Example: `	clarinet server create --subnet_uuid "string" --os "string" --server_name "string" -- server_desc "description string" --cpu 4 --memory 2 --disk_size 10 --user_uuid "string" --nr_node 3`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]interface{})
		queryArgs["subnet_uuid"] = serverArgs.ServerInfo.SubnetUUID
		queryArgs["os"] = serverArgs.ServerInfo.OS
		queryArgs["server_name"] = serverArgs.ServerInfo.ServerName
		queryArgs["server_desc"] = serverArgs.ServerInfo.ServerDesc
		queryArgs["cpu"] = serverArgs.ServerInfo.CPU
		queryArgs["memory"] = serverArgs.ServerInfo.Memory
		queryArgs["disk_size"] = serverArgs.ServerInfo.DiskSize
		queryArgs["user_uuid"] = serverArgs.ServerInfo.UserUUID
		queryArgs["nr_node"] = serverArgs.NrNode
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
		queryArgs := make(map[string]interface{})
		var servers interface{}
		var err error

		if (serverArgs.Row != 0) != (serverArgs.Page != 0) {
			fmt.Println("List Server need [BOTH | NEITHER] of page and row.")
			return
		} else if serverArgs.Row != 0 && serverArgs.Page != 0 {
			queryArgs["row"] = serverArgs.Row
			queryArgs["page"] = serverArgs.Page
		}

		queryArgs["subnet_uuid"] = serverArgs.ServerInfo.SubnetUUID
		queryArgs["os"] = serverArgs.ServerInfo.OS
		queryArgs["server_name"] = serverArgs.ServerInfo.ServerName
		queryArgs["server_desc"] = serverArgs.ServerInfo.ServerDesc
		queryArgs["cpu"] = serverArgs.ServerInfo.CPU
		queryArgs["memory"] = serverArgs.ServerInfo.Memory
		queryArgs["disk_size"] = serverArgs.ServerInfo.DiskSize
		queryArgs["status"] = serverArgs.ServerInfo.Status
		queryArgs["user_uuid"] = serverArgs.ServerInfo.UserUUID

		servers, err = queryParser.ListServer(queryArgs)

		fmt.Println(servers)
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
			serverUUIDArg := make(map[string]interface{})
			serverUUIDArg["server_uuid"] = server.UUID
			t.AppendRow([]interface{}{
				index + 1, server.UUID, server.ServerName, server.CPU, server.Memory, server.DiskSize,
				server.Status})
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

		if serverArgs.ServerInfo.Status == "" {
			fmt.Println("Empty server uuid")
			return
		}

		queryArgs := make(map[string]interface{})
		queryArgs["uuid"] = serverArgs.ServerInfo.UUID
		queryArgs["subnet_uuid"] = serverArgs.ServerInfo.SubnetUUID
		queryArgs["os"] = serverArgs.ServerInfo.OS
		queryArgs["server_name"] = serverArgs.ServerInfo.ServerName
		queryArgs["server_desc"] = serverArgs.ServerInfo.ServerDesc
		queryArgs["cpu"] = serverArgs.ServerInfo.CPU
		queryArgs["memory"] = serverArgs.ServerInfo.Memory
		queryArgs["status"] = serverArgs.ServerInfo.Status
		queryArgs["disk_size"] = serverArgs.ServerInfo.DiskSize
		queryArgs["user_uuid"] = serverArgs.ServerInfo.UserUUID

		_, err := mutationParser.UpdateServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully update server (" + serverArgs.ServerInfo.UUID + ") information.")
	},
}

var serverDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete one of server by UUID.",
	Long:  `Delete one of server by UUID.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]interface{})
		queryArgs["uuid"] = serverArgs.ServerInfo.UUID
		_, err := mutationParser.DeleteServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully delete server (" + serverArgs.ServerInfo.UUID + ").")
	},
}

func ReadyServerCmd() {
	// server create flags setup
	serverCreate.Flags().StringVar(&serverArgs.ServerInfo.SubnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverCreate.Flags().StringVar(&serverArgs.ServerInfo.OS, "os", "", "Type of OS")
	serverCreate.Flags().StringVar(&serverArgs.ServerInfo.ServerName, "server_name", "", "Name of server")
	serverCreate.Flags().StringVar(&serverArgs.ServerInfo.ServerDesc, "server_desc", "", "Description of server")
	serverCreate.Flags().IntVar(&serverArgs.ServerInfo.CPU, "cpu", 0, "Number of CPU cores")
	serverCreate.Flags().IntVar(&serverArgs.ServerInfo.Memory, "memory", 0, "Size of memory")
	serverCreate.Flags().IntVar(&serverArgs.ServerInfo.DiskSize, "disk_size", 0, "Size of disk")
	serverCreate.Flags().StringVar(&serverArgs.ServerInfo.UserUUID, "user_uuid", "", "UUID of user")
	serverCreate.Flags().IntVar(&serverArgs.NrNode, "nr_node", 0, "Number of nodes")

	serverCreate.MarkFlagRequired("subnet_uuid")
	serverCreate.MarkFlagRequired("os")
	serverCreate.MarkFlagRequired("server_name")
	serverCreate.MarkFlagRequired("server_desc")
	serverCreate.MarkFlagRequired("cpu")
	serverCreate.MarkFlagRequired("memory")
	serverCreate.MarkFlagRequired("disk_size")
	serverCreate.MarkFlagRequired("user_uuid")
	serverCreate.MarkFlagRequired("nr_node")

	// server list flags setup
	serverList.Flags().IntVar(&serverArgs.Row, "row", 0, "rows of server list")
	serverList.Flags().IntVar(&serverArgs.Page, "page", 0, "page of server list")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.UUID, "uuid", "", "UUID of server")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.SubnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.OS, "os", "", "Type of OS")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.ServerName, "server_name", "", "Name of server")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.ServerDesc, "server_desc", "", "Description of server")
	serverList.Flags().IntVar(&serverArgs.ServerInfo.CPU, "cpu", 0, "Number of CPU cores")
	serverList.Flags().IntVar(&serverArgs.ServerInfo.Memory, "memory", 0, "Size of memory")
	serverList.Flags().IntVar(&serverArgs.ServerInfo.DiskSize, "disk_size", 0, "Size of disk")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.Status, "status", "", "Server Status [Running | Stop]")
	serverList.Flags().StringVar(&serverArgs.ServerInfo.UserUUID, "user_uuid", "", "UUID of user")

	//server update flags setup
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.UUID, "uuid", "", "UUID of server")
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.SubnetUUID, "subnet_uuid", "", "UUID of subnet")
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.OS, "os", "", "Type of OS")
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.ServerName, "server_name", "", "Name of server")
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.ServerDesc, "server_desc", "", "Description of server")
	serverUpdate.Flags().IntVar(&serverArgs.ServerInfo.CPU, "cpu", 0, "Number of CPU cores")
	serverUpdate.Flags().IntVar(&serverArgs.ServerInfo.Memory, "memory", 0, "Size of memory")
	serverUpdate.Flags().IntVar(&serverArgs.ServerInfo.DiskSize, "disk_size", 0, "Size of disk")
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.Status, "status", "", "Server Status [Running | Stop]")
	serverUpdate.Flags().StringVar(&serverArgs.ServerInfo.UserUUID, "user_uuid", "", "UUID of user")

	serverUpdate.MarkFlagRequired("uuid")

	//server delte flags setup
	serverDelete.Flags().StringVar(&serverArgs.ServerInfo.UUID, "uuid", "", "UUID of server")

	serverDelete.MarkFlagRequired("uuid")

	ServerCmd.AddCommand(serverCreate, serverList, serverUpdate, serverDelete)
}
