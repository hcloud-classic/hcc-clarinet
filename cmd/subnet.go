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
	//"github.com/jedib0t/go-pretty/table"
	//"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"hcc/clarinet/action/graphql/mutationParser"
	//"hcc/clarinet/action/graphql/queryParser"
	//"hcc/clarinet/model"
	//"os"
	//"strconv"
)

// aipCmd represents the aip command
var SubnetCmd = &cobra.Command{
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
	Use:   "create",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["network_ip"] = netIP
		queryArgs["netmask"] = netMask
		queryArgs["gateway"] = gateway
		queryArgs["next_server"] = nextServer
		queryArgs["name_server"] = nameServer
		queryArgs["domain_name"] = domainName
		queryArgs["server_uuid"] = serverUUID
		queryArgs["leader_node_uid"] = leaderUUID
		queryArgs["os"] = OS
		queryArgs["subnet_name"] = subnetName
		node, err := mutationParser.CreateSubnet(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipCreateDHCPDconf = &cobra.Command{
	Use:   "dhcpconf",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["subnet_uuid"] = subnetUUID
		queryArgs["node_uuids"] = nodeUUID
		node, err := mutationParser.CreateDHCPDConf(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}
var aipUpdate = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
}

var aipUpdateSubnet = &cobra.Command{
	Use:   "subnet",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
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
		queryArgs["leader_node_uid"] = leaderUUID
		queryArgs["os"] = OS
		queryArgs["subnet_name"] = subnetName
		node, err := mutationParser.UpdateSubnet(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var aipDelete = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
}

var aipDeleteSubnet = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = uuid
		node, err := mutationParser.DeleteSubnet(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

func ReadyAIPCmd() {
	aipCreateSubnet.Flags().StringVar(&netIP, "network_ip", "", "Network IP")
	aipCreateSubnet.Flags().StringVar(&netMask, "netmask", "", "Network Mask")
	aipCreateSubnet.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipCreateSubnet.Flags().StringVar(&nextServer, "next_server", "", "Next Server")
	aipCreateSubnet.Flags().StringVar(&nameServer, "name_server", "", "Name Server")
	aipCreateSubnet.Flags().StringVar(&domainName, "domain_name", "", "Domain Name")
	aipCreateSubnet.Flags().StringVar(&serverUUID, "server_uuid", "", "Server UUID")
	aipCreateSubnet.Flags().StringVar(&leaderUUID, "leader_node_uuid", "", "Leader Node UUID")
	aipCreateSubnet.Flags().StringVar(&OS, "os", "", "OS type")
	aipCreateSubnet.Flags().StringVar(&subnetName, "subnet_name", "", "Subnet Name")
	aipCreateSubnet.MarkFlagRequired("network_ip")
	aipCreateSubnet.MarkFlagRequired("netmask")
	aipCreateSubnet.MarkFlagRequired("gateway")
	aipCreateSubnet.MarkFlagRequired("next_server")
	aipCreateSubnet.MarkFlagRequired("name_server")
	aipCreateSubnet.MarkFlagRequired("domain_name")
	aipCreateSubnet.MarkFlagRequired("server_uuid")
	aipCreateSubnet.MarkFlagRequired("leader_node_uuid")
	aipCreateSubnet.MarkFlagRequired("os")
	aipCreateSubnet.MarkFlagRequired("subnet_name")

	aipCreate.AddCommand(aipCreateSubnet)

	aipUpdateSubnet.Flags().StringVar(&uuid, "uuid", "", "UUID")
	aipUpdateSubnet.Flags().StringVar(&netIP, "network_ip", "", "Network IP")
	aipUpdateSubnet.Flags().StringVar(&netMask, "netmask", "", "Network Mask")
	aipUpdateSubnet.Flags().StringVar(&gateway, "gateway", "", "Gateway")
	aipUpdateSubnet.Flags().StringVar(&nextServer, "next_server", "", "Next Server")
	aipUpdateSubnet.Flags().StringVar(&nameServer, "name_server", "", "Name Server")
	aipUpdateSubnet.Flags().StringVar(&domainName, "domain_name", "", "Domain Name")
	aipUpdateSubnet.Flags().StringVar(&serverUUID, "server_uuid", "", "Server UUID")
	aipUpdateSubnet.Flags().StringVar(&leaderUUID, "leader_node_uuid", "", "Leader Node UUID")
	aipUpdateSubnet.Flags().StringVar(&OS, "os", "", "OS type")
	aipUpdateSubnet.Flags().StringVar(&subnetName, "subnet_name", "", "Subnet Name")
	aipUpdateSubnet.MarkFlagRequired("uuid")

	aipUpdate.AddCommand(aipUpdateSubnet)

	aipDeleteSubnet.Flags().StringVar(&uuid, "uuid", "", "UUID")
	aipDeleteSubnet.MarkFlagRequired("uuid")

	aipDelete.AddCommand(aipDeleteSubnet)

	AIPCmd.AddCommand(aipCreate, aipUpdate, aipDelete)
}
