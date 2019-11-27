package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
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

var row int
var page int

var serverList = &cobra.Command{
	Use:   "list [row and page options]",
	Short: "Get list of servers with row and page options.",
	Long: `Get list of servers with row and page options.`,
	Args: cobra.MaximumNArgs(2),
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

func ReadyServerCmd() {
	serverList.Flags().IntVar(&row, "row", 0, "rows of server list")
	serverList.Flags().IntVar(&page, "page", 0, "page of server list")

	ServerCmd.AddCommand(serverList)
}
