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
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/model"
)

var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Commands for Subnet",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
}

var netIP, netMask, gateway, nextServer, nameServer, domainName, leaderUUID, subnetName string

var subnetCreate = &cobra.Command{
	Use:     "create",
	Short:   "Creat Subnet",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["next_server"] = nextServer
		queryArgs["name_server"] = nameServer
		queryArgs["domain_name"] = domainName
		queryArgs["server_uuid"] = serverUUID
		queryArgs["leader_node_uuid"] = leaderUUID
		queryArgs["os"] = OS
		queryArgs["subnet_name"] = subnetName
		queryArgs["token"] = config.User.Token

		fmt.Print("Create Subnet .... ")

		data, err := mutationParser.CreateSubnet(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		subnetData := data.(model.Subnet)
		if len(subnetData.Errors) > 0 {
			for _, hrr := range subnetData.Errors {
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
		t.AppendHeader(table.Row{"UUID", subnetData.UUID})
		t.AppendRow([]interface{}{"Subnet Name", subnetData.SubnetName})
		t.AppendRow([]interface{}{"Domain Name", subnetData.DomainName})
		t.AppendRow([]interface{}{"Network IP", subnetData.NetworkIP})
		t.AppendRow([]interface{}{"Netmask", subnetData.Netmask})
		t.AppendRow([]interface{}{"Gateway", subnetData.Gateway})
		t.AppendRow([]interface{}{"DNS", subnetData.NameServer})
		t.AppendRow([]interface{}{"Next Server", subnetData.NextServer})
		t.AppendRow([]interface{}{"Server UUID", subnetData.ServerUUID})
		t.AppendRow([]interface{}{"Leader UUID", subnetData.LeaderNodeUUID})
		t.AppendRow([]interface{}{"OS", subnetData.OS})
		t.AppendRow([]interface{}{"Created At", subnetData.CreatedAt})
		t.Render()
	},
}

var subnetUpdate = &cobra.Command{
	Use:     "update",
	Short:   "Update Subnet",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["next_server"] = nextServer
		queryArgs["name_server"] = nameServer
		queryArgs["domain_name"] = domainName
		queryArgs["server_uuid"] = serverUUID
		queryArgs["leader_node_uuid"] = leaderUUID
		queryArgs["os"] = OS
		queryArgs["subnet_name"] = subnetName
		queryArgs["token"] = config.User.Token

		fmt.Print("Update Subnet .... ")

		data, err := mutationParser.UpdateSubnet(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		subnetData := data.(model.Subnet)
		if len(subnetData.Errors) > 0 {
			for _, hrr := range subnetData.Errors {
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
		t.AppendHeader(table.Row{"UUID", subnetData.UUID})
		t.AppendRow([]interface{}{"Subnet Name", subnetData.SubnetName})
		t.AppendRow([]interface{}{"Domain Name", subnetData.DomainName})
		t.AppendRow([]interface{}{"Network IP", subnetData.NetworkIP})
		t.AppendRow([]interface{}{"Netmask", subnetData.Netmask})
		t.AppendRow([]interface{}{"Gateway", subnetData.Gateway})
		t.AppendRow([]interface{}{"DNS", subnetData.NameServer})
		t.AppendRow([]interface{}{"Next Server", subnetData.NextServer})
		t.AppendRow([]interface{}{"Server UUID", subnetData.ServerUUID})
		t.AppendRow([]interface{}{"Leader UUID", subnetData.LeaderNodeUUID})
		t.AppendRow([]interface{}{"OS", subnetData.OS})
		t.AppendRow([]interface{}{"Created At", subnetData.CreatedAt})
		t.Render()
	},
}

var subnetDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete Subnet",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		queryArgs["token"] = config.User.Token

		fmt.Print("Delete Subnet .... ")

		data, err := mutationParser.DeleteSubnet(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		subnetData := data.(model.Subnet)
		if len(subnetData.Errors) > 0 {
			for _, hrr := range subnetData.Errors {
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
		t.AppendHeader(table.Row{"UUID", subnetData.UUID})
		t.AppendRow([]interface{}{"Subnet Name", subnetData.SubnetName})
		t.AppendRow([]interface{}{"Domain Name", subnetData.DomainName})
		t.AppendRow([]interface{}{"Network IP", subnetData.NetworkIP})
		t.AppendRow([]interface{}{"Netmask", subnetData.Netmask})
		t.AppendRow([]interface{}{"Gateway", subnetData.Gateway})
		t.AppendRow([]interface{}{"DNS", subnetData.NameServer})
		t.AppendRow([]interface{}{"Next Server", subnetData.NextServer})
		t.AppendRow([]interface{}{"Server UUID", subnetData.ServerUUID})
		t.AppendRow([]interface{}{"Leader UUID", subnetData.LeaderNodeUUID})
		t.AppendRow([]interface{}{"OS", subnetData.OS})
		t.AppendRow([]interface{}{"Created At", subnetData.CreatedAt})
		t.Render()
	},
}

var subnetCreateDHCPDConf = &cobra.Command{
	Use:     "dhcpconf",
	Short:   "Create DHCPD Configuration file",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["node_uuids"] = "[" + nodeUUID + "]"
		queryArgs["token"] = config.User.Token

		fmt.Print("Create DHCP Configuration .... ")

		data, err := mutationParser.CreateDHCPDConf(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		subnetData := data.(model.DHCPDConfResult)
		if len(subnetData.Errors) > 0 {
			for _, hrr := range subnetData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")
	},
}

var subnetList = &cobra.Command{
	Use:     "list",
	Short:   "Show subnet list",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["next_server"] = nextServer
		queryArgs["name_server"] = nameServer
		queryArgs["domain_name"] = domainName
		queryArgs["server_uuid"] = serverUUID
		queryArgs["leader_node_uuid"] = leaderUUID
		queryArgs["os"] = OS
		queryArgs["subnet_name"] = subnetName
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListSubnet(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		subnetList := data.(model.Subnets)
		if len(subnetList.Errors) > 0 {
			for _, hrr := range subnetList.Errors {
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
		t.AppendHeader(table.Row{"No", "UUID", "IP", "Netmask", "Gateway", "Next Server", "Name Server",
			"Domain Name", "Server UUID", "Leader UUID", "OS", "Subnet Name", "Create At"})

		for index, subnet := range subnetList.Subnets {
			t.AppendRow([]interface{}{index + 1, subnet.UUID, subnet.NetworkIP, subnet.Netmask, subnet.Gateway,
				subnet.NextServer, subnet.NameServer, subnet.DomainName, subnet.ServerUUID,
				subnet.LeaderNodeUUID, subnet.OS, subnet.SubnetName, subnet.CreatedAt})
		}

		t.AppendFooter(table.Row{"Total", len(subnetList.Subnets)})
		t.Render()
	},
}

func ReadySubnetCmd() {
	subnetCreate.Flags().StringVar(&netIP, "network_ip", "", "Network IP [x.x.x.0]")
	subnetCreate.Flags().StringVar(&netMask, "netmask", "", "Network Mask")
	subnetCreate.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	subnetCreate.Flags().StringVar(&nextServer, "next_server", "", "Next Server")
	subnetCreate.Flags().StringVar(&nameServer, "name_server", "", "Name Server")
	subnetCreate.Flags().StringVar(&domainName, "domain_name", "", "Domain Name")
	subnetCreate.Flags().StringVar(&serverUUID, "server_uuid", "", "Server UUID")
	subnetCreate.Flags().StringVar(&leaderUUID, "leader_node_uuid", "", "Leader Node UUID")
	subnetCreate.Flags().StringVar(&OS, "os", "", "OS type")
	subnetCreate.Flags().StringVar(&subnetName, "subnet_name", "", "Subnet Name")
	subnetCreate.MarkFlagRequired("network_ip")
	subnetCreate.MarkFlagRequired("netmask")
	subnetCreate.MarkFlagRequired("gateway")
	subnetCreate.MarkFlagRequired("next_server")
	subnetCreate.MarkFlagRequired("name_server")
	subnetCreate.MarkFlagRequired("domain_name")
	subnetCreate.MarkFlagRequired("os")
	subnetCreate.MarkFlagRequired("subnet_name")

	subnetCreateDHCPDConf.Flags().StringVar(&subnetUUID, "subnet_uuid", "", "Subnet UUID")
	subnetCreateDHCPDConf.Flags().StringVar(&nodeUUID, "node_uuids", "", "Node UUIDs")
	subnetCreateDHCPDConf.MarkFlagRequired("subnet_uuid")
	subnetCreateDHCPDConf.MarkFlagRequired("node_uuids")

	subnetCreate.AddCommand(subnetCreateDHCPDConf)

	subnetUpdate.Flags().StringVar(&uuid, "uuid", "", "UUID")
	subnetUpdate.Flags().StringVar(&netIP, "network_ip", "", "Network IP [x.x.x.0]")
	subnetUpdate.Flags().StringVar(&netMask, "netmask", "", "Network Mask")
	subnetUpdate.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	subnetUpdate.Flags().StringVar(&nextServer, "next_server", "", "Next Server")
	subnetUpdate.Flags().StringVar(&nameServer, "name_server", "", "Name Server")
	subnetUpdate.Flags().StringVar(&domainName, "domain_name", "", "Domain Name")
	subnetUpdate.Flags().StringVar(&serverUUID, "server_uuid", "", "Server UUID")
	subnetUpdate.Flags().StringVar(&leaderUUID, "leader_node_uuid", "", "Leader Node UUID")
	subnetUpdate.Flags().StringVar(&OS, "os", "", "OS type")
	subnetUpdate.Flags().StringVar(&subnetName, "subnet_name", "", "Subnet Name")
	subnetUpdate.MarkFlagRequired("uuid")

	subnetDelete.Flags().StringVar(&uuid, "uuid", "", "UUID")
	subnetDelete.MarkFlagRequired("uuid")

	subnetList.Flags().IntVar(&row, "row", 0, "")
	subnetList.Flags().IntVar(&page, "page", 0, "")
	subnetList.Flags().StringVar(&netIP, "network_ip", "", "Network IP [x.x.x.0]")
	subnetList.Flags().StringVar(&netMask, "netmask", "", "Network Mask")
	subnetList.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	subnetList.Flags().StringVar(&nextServer, "next_server", "", "Next Server")
	subnetList.Flags().StringVar(&nameServer, "name_server", "", "Name Server")
	subnetList.Flags().StringVar(&domainName, "domain_name", "", "Domain Name")
	subnetList.Flags().StringVar(&serverUUID, "server_uuid", "", "Server UUID")
	subnetList.Flags().StringVar(&leaderUUID, "leader_node_uuid", "", "Leader Node UUID")
	subnetList.Flags().StringVar(&OS, "os", "", "OS type")
	subnetList.Flags().StringVar(&subnetName, "subnet_name", "", "Subnet Name")

	subnetCmd.AddCommand(subnetCreate, subnetUpdate, subnetDelete, subnetList)
}
