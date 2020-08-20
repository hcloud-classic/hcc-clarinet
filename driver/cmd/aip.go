/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/model"
	"os"
	"strconv"
)

// aipCmd represents the aip command
var aipCmd = &cobra.Command{
	Use:   "aip",
	Short: "Commands for Adaptive IP",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
}

var startIP, endIP, aipUUID, publicIP, privateIP string

var aipCreate = &cobra.Command{
	Use:     "create",
	Short:   "Creat Adaptive IP",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		queryArgs["token"] = config.User.Token
		node, err := mutationParser.CreateAdaptiveIP(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
			return
		}

		fmt.Println(node)
	},
}

var aipUpdate = &cobra.Command{
	Use:     "update",
	Short:   "Update Adaptive IP",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		queryArgs["token"] = config.User.Token
		node, err := mutationParser.UpdateAdaptiveIP(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
			return
		}

		fmt.Println(node)
	},
}

var aipDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete Adaptive IP",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["token"] = config.User.Token
		node, err := mutationParser.DeleteAdaptiveIP(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
			return
		}

		fmt.Println(node)
	},
}

var aipCreateServer = &cobra.Command{
	Use:     "server",
	Short:   "Create Adaptive IP Server",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		queryArgs["token"] = config.User.Token
		node, err := mutationParser.CreateAdaptiveIPServer(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
			return
		}

		fmt.Println(node)
	},
}

var aipDeleteServer = &cobra.Command{
	Use:     "server",
	Short:   "Delete Adaptive IP Server",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["token"] = config.User.Token
		node, err := mutationParser.DeleteAdaptiveIPServer(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
			return
		}

		fmt.Println(node)
	},
}

var aipList = &cobra.Command{
	Use:     "list",
	Short:   "Show available Adaptive IP List",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["token"] = config.User.Token
		aipList, err := queryParser.ListAdaptiveIP(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
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
				SeparateColumns: false,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Available IP"})

		for _, aip := range aipList.([]string) {
			t.AppendRow([]interface{}{aip})
		}

		t.AppendFooter(table.Row{"Total\t" + strconv.Itoa(len(aipList.([]string)))})
		t.Render()
	},
}

var aipListServer = &cobra.Command{
	Use:     "server",
	Short:   "Show Adaptive IP Server List",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		queryArgs["private_ip"] = privateIP
		queryArgs["private_gateway"] = gateway
		queryArgs["token"] = config.User.Token
		aipServerList, err := queryParser.ListAdaptiveIPServer(queryArgs)
		if err != nil {
			reRunIfExpired(cmd)
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
		t.AppendHeader(table.Row{"No", "Server UUID", "Public IP", "Private IP", "Private Gateway"})

		for index, aipServer := range aipServerList.([]model.AdaptiveIPServer) {
			t.AppendRow([]interface{}{index + 1, aipServer.ServerUUID,
				aipServer.PublicIP, aipServer.PrivateIP, aipServer.PrivateGateway})
		}

		t.AppendFooter(table.Row{"Total", len(aipServerList.([]model.AdaptiveIPServer))})
		t.Render()

	},
}

func ReadyAIPCmd() {
	aipCreate.Flags().StringVar(&netIP, "network_address", "", "Network Address")
	aipCreate.Flags().StringVar(&netMask, "netmask", "", "Netmask")
	aipCreate.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipCreate.Flags().StringVar(&startIP, "start_ip_address", "", "Start IP Address")
	aipCreate.Flags().StringVar(&netIP, "end_ip_address", "", "End IP Address")
	aipCreate.MarkFlagRequired("network_address")
	aipCreate.MarkFlagRequired("netmask")
	aipCreate.MarkFlagRequired("gateway")
	aipCreate.MarkFlagRequired("start_ip_address")
	aipCreate.MarkFlagRequired("end_ip_address")

	aipUpdate.Flags().StringVar(&uuid, "uuid", "", "UUID")
	aipUpdate.Flags().StringVar(&netIP, "network_ip", "", "Network Address")
	aipUpdate.Flags().StringVar(&netMask, "netmask", "", "Netmask")
	aipUpdate.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipUpdate.Flags().StringVar(&startIP, "start_ip_address", "", "Start IP Address")
	aipUpdate.Flags().StringVar(&netIP, "end_ip_address", "", "End IP Address")
	aipUpdate.MarkFlagRequired("uuid")

	aipDelete.Flags().StringVar(&uuid, "uuid", "", "UUID")
	aipDelete.MarkFlagRequired("uuid")

	aipCreateServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipCreateServer.Flags().StringVar(&publicIP, "public_ip", "", "Public IP")
	aipCreateServer.MarkFlagRequired("server_uuid")
	aipCreateServer.MarkFlagRequired("public_ip")

	aipCreate.AddCommand(aipCreateServer)

	aipDeleteServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipDeleteServer.MarkFlagRequired("server_uuid")

	aipDelete.AddCommand(aipDeleteServer)

	aipListServer.Flags().IntVar(&row, "row", 0, "Number Of rows to show")
	aipListServer.Flags().IntVar(&page, "page", 0, "Number Of page to show")
	aipListServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipListServer.Flags().StringVar(&publicIP, "public_ip", "", "Public IP of Server")
	aipListServer.Flags().StringVar(&privateIP, "private_ip", "", "Private IP of Server")
	aipListServer.Flags().StringVar(&gateway, "private_gateway", "", "Private Gateway IP")

	aipList.AddCommand(aipListServer)

	aipCmd.AddCommand(aipCreate, aipUpdate, aipDelete, aipList)
}
