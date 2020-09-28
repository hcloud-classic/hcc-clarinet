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
	"hcc/clarinet/model"
	"os"
	"strconv"
)

// aipCmd represents the aip command
var AIPCmd = &cobra.Command{
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
	Use:   "create",
	Short: "Creat Adaptive IP",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		node, err := mutationParser.CreateAdaptiveIP(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update Adaptive IP",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		node, err := mutationParser.UpdateAdaptiveIP(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Adaptive IP",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		node, err := mutationParser.DeleteAdaptiveIP(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipCreateServer = &cobra.Command{
	Use:   "server",
	Short: "Create Adaptive IP Server",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["adaptiveip_uuid"] = aipUUID
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		node, err := mutationParser.CreateAdaptiveIPServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipDeleteServer = &cobra.Command{
	Use:   "server",
	Short: "Delete Adaptive IP Server",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		node, err := mutationParser.DeleteAdaptiveIPServer(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipList = &cobra.Command{
	Use:   "list",
	Short: "Creat Adaptive IP",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		aipList, err := queryParser.ListAdaptiveIP(queryArgs)
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
		t.AppendHeader(table.Row{"No", "UUID", "IP", "Netmask", "Gateway", "Start IP", "End IP", "Created At"})

		for index, aip := range aipList.([]model.AdaptiveIP) {
			t.AppendRow([]interface{}{index + 1, aip.UUID, aip.NetworkAddress, aip.Netmask,
				aip.Gateway, aip.StartIPAddress, aip.EndIPAddress, aip.CreatedAt})
		}

		t.AppendFooter(table.Row{"Total", len(aipList.([]model.AdaptiveIP))})
		t.Render()
	},
}

var aipListServer = &cobra.Command{
	Use:   "server",
	Short: "Show Adaptive IP Server List",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		queryArgs["private_ip"] = privateIP
		queryArgs["private_gateway"] = gateway
		aipServerList, err := queryParser.ListAdaptiveIPServer(queryArgs)
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
		t.AppendHeader(table.Row{"No", "AIP UUID", "Server UUID", "Public IP", "Private IP", "Private Gateway"})

		for index, aipServer := range aipServerList.([]model.AdaptiveIPServer) {
			t.AppendRow([]interface{}{index + 1, aipServer.AdaptiveIPUUID, aipServer.ServerUUID,
				aipServer.PublicIP, aipServer.PrivateIP, aipServer.PrivateGateway})
		}

		t.AppendFooter(table.Row{"Total", len(aipServerList.([]model.AdaptiveIPServer))})
		t.Render()

	},
}

func ReadyAIPCmd() {
	aipCreate.Flags().StringVar(&netIP, "network_ip", "", "Network Address")
	aipCreate.Flags().StringVar(&netMask, "netmask", "", "Netmask")
	aipCreate.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipCreate.Flags().StringVar(&startIP, "start_ip_address", "", "Start IP Address")
	aipCreate.Flags().StringVar(&netIP, "end_ip_address", "", "End IP Address")
	nodeDelete.MarkFlagRequired("network_address")
	nodeDelete.MarkFlagRequired("netmask")
	nodeDelete.MarkFlagRequired("gateway")
	nodeDelete.MarkFlagRequired("start_ip_address")
	nodeDelete.MarkFlagRequired("end_ip_address")

	aipUpdate.Flags().StringVar(&uuid, "uuid", "", "UUID")
	aipUpdate.Flags().StringVar(&netIP, "network_ip", "", "Network Address")
	aipUpdate.Flags().StringVar(&netMask, "netmask", "", "Netmask")
	aipUpdate.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipUpdate.Flags().StringVar(&startIP, "start_ip_address", "", "Start IP Address")
	aipUpdate.Flags().StringVar(&netIP, "end_ip_address", "", "End IP Address")

	aipDelete.Flags().StringVar(&uuid, "uuid", "", "UUID")
	aipDelete.MarkFlagRequired("uuid")

	aipList.Flags().IntVar(&row, "row", 0, "Number of rows to show")
	aipList.Flags().IntVar(&page, "page", 0, "Nuber of page to show")
	aipList.Flags().StringVar(&netIP, "network_ip", "", "Network Address")
	aipList.Flags().StringVar(&netMask, "netmask", "", "Netmask")
	aipList.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipList.Flags().StringVar(&startIP, "start_ip_address", "", "Start IP Address")
	aipList.Flags().StringVar(&netIP, "end_ip_address", "", "End IP Address")

	aipCreateServer.Flags().StringVar(&aipUUID, "adaptiveip_uuid", "", "UUID of Adatative IP")
	aipCreateServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipCreateServer.Flags().StringVar(&publicIP, "public_ip", "", "Public IP")
	aipCreateServer.MarkFlagRequired("adaptiveip_uuid")
	aipCreateServer.MarkFlagRequired("server_uuid")
	aipCreateServer.MarkFlagRequired("public_ip")

	aipCreate.AddCommand(aipCreateServer)

	aipDeleteServer.Flags().StringVar(&uuid, "server_uuid", "", "UUID of Server")
	aipDeleteServer.MarkFlagRequired("server_uuid")

	aipDelete.AddCommand(aipDeleteServer)

	aipListServer.Flags().IntVar(&row, "row", 0, "Number Of rows to show")
	aipListServer.Flags().IntVar(&page, "page", 0, "Number Of page to show")
	aipListServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipListServer.Flags().StringVar(&publicIP, "public_ip", "", "Public IP of Server")
	aipListServer.Flags().StringVar(&privateIP, "private_ip", "", "Private IP of Server")
	aipListServer.Flags().StringVar(&gateway, "private_gateway", "", "Private Gateway IP")
	aipListServer.MarkFlagRequired("suber_uuid")

	aipList.AddCommand(aipListServer)

	AIPCmd.AddCommand(aipCreate, aipUpdate, aipDelete, aipList)
}
