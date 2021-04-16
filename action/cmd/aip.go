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
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/lib/logger"
	"hcc/clarinet/model"
)

var startIP, endIP, aipUUID, publicIP, privateIP, netmask, extIfaceAddr string

// aipCmd represents the aip command
var aipCmd = &cobra.Command{
	Use:     "aip",
	Short:   "Show current Adaptive IP setting",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]string)
		query["token"] = config.User.Token

		data, err := queryParser.AdaptiveIP(query)
		if err != nil {
			err.Println()
			return
		}

		aipData := data.(model.AdaptiveIP)
		if len(aipData.Errors) > 0 {
			for _, hrr := range aipData.Errors {
				hrr.Println()
			}
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

		t.SetStyle(ts)
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("Adaptive IP Settings")

		t.AppendRow(table.Row{"External iface IP", aipData.ExtIfaceAddress})
		t.AppendRow(table.Row{"Netmask", aipData.Netmask})
		t.AppendRow(table.Row{"Gateway", aipData.Gateway})
		t.AppendRow(table.Row{"Start IP Address", aipData.StartIPAddress})
		t.AppendRow(table.Row{"End IP Address", aipData.EndIPAddress})

		t.Render()

	},
}

var aipCreate = &cobra.Command{
	Use:   "create",
	Short: "Creat Adaptive IP Setting OR Server",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
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
		data, err := mutationParser.CreateAdaptiveIPServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipServerData := data.(model.AdaptiveIPServer)
		if len(aipServerData.Errors) > 0 {
			for _, hrr := range aipServerData.Errors {
				hrr.Println()
			}
			return
		}
	},
}

var aipCreateSetting = &cobra.Command{
	Use:     "setting",
	Short:   "Create Adaptive IP Setting",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["ext_ifaceip_address"] = extIfaceAddr
		queryArgs["netmask"] = netmask
		queryArgs["gateway_address"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.CreateAdaptiveIPSetting(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipData := data.(model.AdaptiveIP)
		if len(aipData.Errors) > 0 {
			for _, hrr := range aipData.Errors {
				hrr.Println()
			}
			return
		}
	},
}

var aipDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Adaptive IP",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
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
		data, err := mutationParser.DeleteAdaptiveIPServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipServerData := data.(model.AdaptiveIPServer)
		if len(aipServerData.Errors) > 0 {
			for _, hrr := range aipServerData.Errors {
				hrr.Println()
			}
			return
		}

		logger.Logger.Printf("Successfully Deleted aipServer - %s\n", serverUUID)
	},
}

var aipList = &cobra.Command{
	Use:   "list",
	Short: "Show available Adaptive IP List or Server",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
}

var aipListAvailable = &cobra.Command{
	Use:     "available",
	Short:   "Show available Adaptive IP List",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["token"] = config.User.Token
		data, err := queryParser.ListAdaptiveIP(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		availableIPList := data.(model.AvailableIPList)
		if len(availableIPList.Errors) > 0 {
			for _, hrr := range availableIPList.Errors {
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
				SeparateColumns: false,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Available IP"})

		for _, aip := range availableIPList.AvailableIPs {
			t.AppendRow([]interface{}{aip})
		}

		t.AppendFooter(table.Row{"Total\t" + strconv.Itoa(len(availableIPList.AvailableIPs))})
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
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		queryArgs["private_ip"] = privateIP
		queryArgs["private_gateway"] = gateway
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListAdaptiveIPServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipServerList := data.(model.AdaptiveIPServers)
		if len(aipServerList.Errors) > 0 {
			for _, hrr := range aipServerList.Errors {
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
		t.AppendHeader(table.Row{"No", "Server UUID", "Public IP", "Private IP", "Private Gateway", "Created At"})

		for index, aipServer := range aipServerList.AdaptiveIPServers {
			t.AppendRow([]interface{}{index + 1, aipServer.ServerUUID,
				aipServer.PublicIP, aipServer.PrivateIP, aipServer.PrivateGateway, aipServer.CreatedAt})
		}

		t.AppendFooter(table.Row{"Total", len(aipServerList.AdaptiveIPServers)})
		t.Render()

	},
}

func ReadyAIPCmd() {

	aipCreateServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipCreateServer.Flags().StringVar(&publicIP, "public_ip", "", "Public IP")
	aipCreateServer.MarkFlagRequired("server_uuid")
	aipCreateServer.MarkFlagRequired("public_ip")

	aipCreateSetting.Flags().StringVar(&extIfaceAddr, "ext_iface_addr", "", "")
	aipCreateSetting.Flags().StringVar(&netmask, "netmask", "", "")
	aipCreateSetting.Flags().StringVar(&gateway, "gateway", "", "")
	aipCreateSetting.Flags().StringVar(&startIP, "start", "", "")
	aipCreateSetting.Flags().StringVar(&endIP, "end", "", "")
	aipCreateSetting.MarkFlagRequired("ext_iface_addr")
	aipCreateSetting.MarkFlagRequired("netmask")
	aipCreateSetting.MarkFlagRequired("gateway")
	aipCreateSetting.MarkFlagRequired("start")
	aipCreateSetting.MarkFlagRequired("end")

	aipCreate.AddCommand(aipCreateServer, aipCreateSetting)

	aipDeleteServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipDeleteServer.MarkFlagRequired("server_uuid")

	aipDelete.AddCommand(aipDeleteServer)

	aipListServer.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipListServer.Flags().StringVar(&publicIP, "public_ip", "", "Public IP of Server")
	aipListServer.Flags().StringVar(&privateIP, "private_ip", "", "Private IP of Server")
	aipListServer.Flags().StringVar(&gateway, "private_gateway", "", "Private Gateway IP")

	aipList.AddCommand(aipListServer, aipListAvailable)

	aipCmd.AddCommand(aipCreate, aipDelete, aipList)
}
