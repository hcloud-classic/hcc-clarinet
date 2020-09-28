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
	"hcc/clarinet/lib/errors"
	"hcc/clarinet/lib/logger"
	"hcc/clarinet/model"
)

var serverUUID, nodeUUID, bmcMacAddr, bmcIP, pxeMacAddr, desc, cpuModel string
var active, cpuCores, cpuProcessors, cpuThreads int

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:     `node --uuid "NodeUUID"`,
	Short:   "Get node information",
	Long:    `Show node information by given nodeUUID at table`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["token"] = config.User.Token

		data, err := queryParser.Node(queryArgs)
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

		nodeData := data.(model.Node)
		if nodeData.Errors.Len() >= 0 {
			nodeData.Errors.Print()
		} else {

			t.SetStyle(ts)
			t.SetOutputMirror(os.Stdout)
			t.SetTitle("Node Info\n%s", nodeUUID)

			t.AppendRow(table.Row{"BMC MAC", nodeData.BmcMacAddr})
			t.AppendRow(table.Row{"BMC IP", nodeData.BmcIP})
			t.AppendRow(table.Row{"PXE MAC", nodeData.PXEMacAddr})
			t.AppendRow(table.Row{"MEMORY", nodeData.Memory})
			t.AppendRow(table.Row{"STATUS", nodeData.Status})
			t.AppendRow(table.Row{"ACTIVE", nodeData.Active})
			t.AppendRow(table.Row{"Description", nodeData.Description})
			t.AppendRow(table.Row{"Created At", nodeData.CreatedAt})

			t.Render()
		}

	},
}

var nodeOn = &cobra.Command{
	Use:     "on",
	Short:   "Power on specified node",
	Long:    `Power on specified node`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["token"] = config.User.Token

		data, err := mutationParser.OnOffNode(queryArgs, model.On)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.PowerStateNode)
		if nodeData.Errors.Len() > 0 {
			nodeData.Errors.Print()
			return
		}

		logger.Logger.Println("Flute :" + nodeData.State)
	},
}

var nodeOff = &cobra.Command{
	Use:     "off",
	Short:   "Power off specified node",
	Long:    `Power off specified node`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["force_off"] = "true"
		queryArgs["token"] = config.User.Token

		data, err := mutationParser.OnOffNode(queryArgs, model.Off)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.PowerStateNode)
		if nodeData.Errors.Len() > 0 {
			nodeData.Errors.Print()
			return
		}

		logger.Logger.Println("Flute :" + nodeData.State)
	},
}

var nodeRestart = &cobra.Command{
	Use:     "restart",
	Short:   "Restart specified node",
	Long:    `Restart specified node`,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.OnOffNode(queryArgs, model.Restart)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.PowerStateNode)
		if nodeData.Errors.Len() > 0 {
			nodeData.Errors.Print()
			return
		}

		logger.Logger.Println("Flute :" + nodeData.State)
	},
}

var nodeCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create node or Add detail information of node",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["bmc_mac_addr"] = bmcMacAddr
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["pxe_mac_addr"] = pxeMacAddr
		queryArgs["status"] = status
		queryArgs["description"] = desc
		queryArgs["active"] = strconv.Itoa(active)
		queryArgs["cpu_cores"] = strconv.Itoa(cpuCores)
		queryArgs["memory"] = strconv.Itoa(memory)
		queryArgs["token"] = config.User.Token

		data, err := mutationParser.CreateNode(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.Node)
		if nodeData.Errors.Len() > 0 {
			nodeData.Errors.Print()
			return
		}
	},
}

var nodeUpdate = &cobra.Command{
	Use:     "update",
	Short:   "Update node information",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["server_uuid"] = serverUUID
		queryArgs["bmc_mac_addr"] = bmcMacAddr
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["pxe_mac_addr"] = pxeMacAddr
		queryArgs["args"] = status
		queryArgs["desctiption"] = desc
		queryArgs["active"] = strconv.Itoa(active)
		queryArgs["cpu_cores"] = strconv.Itoa(cpuCores)
		queryArgs["memory"] = strconv.Itoa(memory)
		queryArgs["token"] = config.User.Token

		data, err := mutationParser.UpdateNode(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.Node)
		if nodeData.Errors.Len() > 0 {
			nodeData.Errors.Print()
			return
		}
	},
}

var nodeDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete node",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.DeleteNode(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.Node)
		if nodeData.Errors.Len() > 0 {
			nodeData.Errors.Print()
			return
		}
	},
}

var nodeCreateDetail = &cobra.Command{
	Use:     "detail",
	Short:   "Create detail information of node",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["node_uuid"] = nodeUUID
		queryArgs["cpu_model"] = cpuModel
		queryArgs["cpu_processors"] = strconv.Itoa(cpuProcessors)
		queryArgs["cpu_threads"] = strconv.Itoa(cpuThreads)
		queryArgs["token"] = config.User.Token

		data, err := mutationParser.CreateNodeDetail(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.NodeDetail)
		if nodeData.Errors.Len() > 0 {
			err = nodeData.Errors.Dump()
			if err.Code() == errors.PiccoloGraphQLTokenExpired {
				reRunIfExpired(cmd)
				return
			}
		}
	},
}

var nodeDeleteDetail = &cobra.Command{
	Use:     "detail",
	Short:   "Delte detail information of node",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		//queryArgs["node_uuid"] = nodeUUID
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["token"] = config.User.Token

		data, err := mutationParser.DeleteNodeDetail(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeData := data.(model.NodeDetail)
		if nodeData.Errors.Len() > 0 {
			err = nodeData.Errors.Dump()
			if err.Code() == errors.PiccoloGraphQLTokenExpired {
				reRunIfExpired(cmd)
				return
			}
		}
	},
}

var nodeList = &cobra.Command{
	Use:     "list",
	Short:   "Show node list",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)

		queryArgs["row"] = strconv.Itoa(row)
		queryArgs["page"] = strconv.Itoa(page)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["bmc_mac_addr"] = bmcMacAddr
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["pxe_mac_addr"] = pxeMacAddr
		queryArgs["status"] = status
		queryArgs["cpu_cores"] = strconv.Itoa(cpuCores)
		queryArgs["memory"] = strconv.Itoa(memory)
		queryArgs["description"] = desc
		queryArgs["active"] = strconv.Itoa(active)
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListNode(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeList := data.(model.Nodes)
		if nodeList.Errors.Len() > 0 {
			nodeList.Errors.Print()
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
		t.AppendHeader(table.Row{"No", "UUID", "Server UUID", "BMC MAC", "BMC IP", "PXE MAC",
			"Cores", "Memory", "Description", "Active", "Status"})

		for index, node := range nodeList.Nodes {
			t.AppendRow([]interface{}{
				index + 1, node.UUID, node.ServerUUID, node.BmcMacAddr, node.BmcIP, node.PXEMacAddr,
				node.CPUCores, node.Memory, node.Description, node.Active, node.Status})
		}

		t.AppendFooter(table.Row{"Total", len(nodeList.Nodes)})
		t.Render()

	},
}

var nodeDetail = &cobra.Command{
	Use:     "detail",
	Short:   "Show detail information of node",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)

		queryArgs["node_uuid"] = nodeUUID
		queryArgs["token"] = config.User.Token

		data, err := queryParser.NodeDetail(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeDetailData := data.(model.NodeDetail)
		if nodeDetailData.Errors.Len() > 0 {
			nodeDetailData.Errors.Print()
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
		t.AppendHeader(table.Row{"Node UUID", nodeDetailData.NodeUUID})

		t.AppendRow([]interface{}{"CPU Model", nodeDetailData.CPUModel})
		t.AppendRow([]interface{}{"CPU Processors", nodeDetailData.CPUProcessors})
		t.AppendRow([]interface{}{"CPU Threads", nodeDetailData.CPUThreads})
		t.Render()

	},
}

func ReadyNodeCmd() {
	nodeCmd.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeCmd.MarkFlagRequired("uuid")

	nodeOn.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeOn.MarkFlagRequired("uuid")

	nodeOff.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeOff.MarkFlagRequired("uuid")

	nodeRestart.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeRestart.MarkFlagRequired("uuid")

	nodeCreate.Flags().StringVar(&bmcMacAddr, "bmc_mac_addr", "", "MAC address of BMC")
	nodeCreate.Flags().StringVar(&bmcIP, "bmc_ip", "", "IP address of BMC")
	nodeCreate.Flags().StringVar(&pxeMacAddr, "pxe_mac_addr", "", "PXE MAC address")
	nodeCreate.Flags().StringVar(&status, "status", "", "Status")
	nodeCreate.Flags().StringVar(&desc, "description", "", "Description")
	nodeCreate.Flags().IntVar(&active, "active", 0, "Active state")
	nodeCreate.Flags().IntVar(&cpuCores, "cpu_cores", 0, "Number of CPU cores")
	nodeCreate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	nodeCreate.MarkFlagRequired("bmc_ip")
	nodeCreate.MarkFlagRequired("description")

	nodeUpdate.Flags().StringVar(&nodeUUID, "uuid", "", "Node UUID")
	nodeUpdate.Flags().StringVar(&serverUUID, "server_uuid", "", "Server UUID")
	nodeUpdate.Flags().StringVar(&bmcMacAddr, "bmc_mac_addr", "", "MAC address of BMC")
	nodeUpdate.Flags().StringVar(&bmcIP, "bmc_ip", "", "IP address of BMC")
	nodeUpdate.Flags().StringVar(&pxeMacAddr, "pxe_mac_addr", "", "PXE MAC address")
	nodeUpdate.Flags().StringVar(&status, "status", "", "status")
	nodeUpdate.Flags().StringVar(&desc, "description", "", "Description")
	nodeUpdate.Flags().IntVar(&active, "active", 0, "Active state")
	nodeUpdate.Flags().IntVar(&cpuCores, "cpu", 0, "Number of CPU cores")
	nodeUpdate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	nodeUpdate.MarkFlagRequired("bmc_ip")

	nodeDelete.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeDelete.MarkFlagRequired("uuid")

	nodeList.Flags().IntVar(&row, "row", 0, "rows of node list")
	nodeList.Flags().IntVar(&page, "page", 0, "page of node list")
	nodeList.Flags().StringVar(&serverUUID, "server_uuid", "", "server UUID")
	nodeList.Flags().StringVar(&bmcMacAddr, "bmc_mac_addr", "", "MAC address of BMC")
	nodeList.Flags().StringVar(&bmcIP, "bmc_ip", "", "IP of BMC")
	nodeList.Flags().StringVar(&pxeMacAddr, "pxe_mac_addr", "", "MAC address for PXE boot")
	nodeList.Flags().StringVar(&status, "status", "", "status")
	nodeList.Flags().IntVar(&cpuCores, "cpu_cores", 0, "Number of CPU cores")
	nodeList.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	nodeList.Flags().StringVar(&desc, "description", "", "Descriptions of Node")
	nodeList.Flags().IntVar(&active, "active", 0, "Active status")

	nodeCreateDetail.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeCreateDetail.Flags().StringVar(&cpuModel, "cpu_model", "", "CPU model")
	nodeCreateDetail.Flags().IntVar(&cpuProcessors, "cpu_processors", 0, "Number of processor")
	nodeCreateDetail.Flags().IntVar(&cpuThreads, "cpu_threads", 0, "Number of logical core")
	nodeCreateDetail.MarkFlagRequired("uuid")
	nodeCreateDetail.MarkFlagRequired("cpu_model")
	nodeCreateDetail.MarkFlagRequired("cpu_processors")
	nodeCreateDetail.MarkFlagRequired("cpu_threads")

	nodeDeleteDetail.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeDeleteDetail.MarkFlagRequired("uuid")

	nodeDetail.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeDetail.MarkFlagRequired("uuid")

	nodeCreate.AddCommand(nodeCreateDetail)
	nodeDelete.AddCommand(nodeDeleteDetail)
	nodeCmd.AddCommand(nodeOn, nodeOff, nodeRestart,
		nodeCreate, nodeUpdate, nodeDelete,
		nodeList, nodeDetail)
}
