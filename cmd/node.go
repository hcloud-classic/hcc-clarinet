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

// nodeCmd represents the node command
var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Running node related commands",
	Long:  `Running node related commands`,
	Args:  cobra.MinimumNArgs(1),
}

var serverUUID, nodeUUID, bmcMacAddr, bmcIP, pxeMacAddr, desc, active, cpuModel string
var cpuCores, cpuProcessors, cpuThreads int

var nodeOn = &cobra.Command{
	Use:   "on",
	Short: "Power on specified node",
	Long:  `Power on specified node`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.OnOffNode(queryArgs, model.On)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Flute :" + node.(string))
	},
}

var nodeOff = &cobra.Command{
	Use:   "off",
	Short: "Power off specified node",
	Long:  `Power off specified node`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.OnOffNode(queryArgs, model.Off)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Flute :" + node.(string))
	},
}

var nodeRestart = &cobra.Command{
	Use:   "restart",
	Short: "Restart specified node",
	Long:  `Restart specified node`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.OnOffNode(queryArgs, model.Restart)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Flute :" + node.(string))
	},
}

var nodeCreate = &cobra.Command{
	Use:   "create",
	Short: "Create node or Add detail information of node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["bmc_mac_addr"] = bmcMacAddr
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["pxe_mac_addr"] = pxeMacAddr
		queryArgs["status"] = status
		queryArgs["description"] = desc
		queryArgs["active"] = active
		queryArgs["cpu_cores"] = strconv.Itoa(cpuCores)
		queryArgs["memory"] = strconv.Itoa(memory)
		node, err := mutationParser.CreateNode(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var nodeUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update node information",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		queryArgs["bmc_mac_addr"] = bmcMacAddr
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["pxe_mac_addr"] = pxeMacAddr
		queryArgs["args"] = status
		queryArgs["desctiption"] = desc
		queryArgs["active"] = active
		queryArgs["cpu_cores"] = strconv.Itoa(cpuCores)
		queryArgs["memory"] = strconv.Itoa(memory)
		node, err := mutationParser.UpdateNode(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var nodeDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.DeleteNode(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var nodeCreateDetail = &cobra.Command{
	Use:   "detail",
	Short: "Create detail information of node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["node_uuid"] = nodeUUID
		queryArgs["cpu_model"] = cpuModel
		queryArgs["cpu_processors"] = strconv.Itoa(cpuProcessors)
		queryArgs["cpu_threads"] = strconv.Itoa(cpuThreads)
		node, err := mutationParser.CreateNodeDetail(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(node)
	},
}

var nodeDeleteDetail = &cobra.Command{
	Use:   "detail",
	Short: "Delte detail information of node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["node_uuid"] = nodeUUID
		node, err := mutationParser.DeleteNodeDetail(queryArgs)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(node)
	},
}

var nodeList = &cobra.Command{
	Use:   "list",
	Short: "Show node list",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
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
		queryArgs["active"] = active

		nodes, err := queryParser.ListNode(queryArgs)

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
		t.AppendHeader(table.Row{"No", "UUID", "BMC MAC", "BMC IP", "PXE MAC",
			"Cores", "Memory", "Description", "Active", "Status"})

		for index, node := range nodes.([]model.Node) {
			serverUUIDArg := make(map[string]string)
			serverUUIDArg["node_uuid"] = node.UUID
			t.AppendRow([]interface{}{
				index + 1, node.UUID, node.BmcMacAddr, node.BmcIP, node.PXEMacAddr,
				node.CPUCores, node.Memory, node.Description, node.Active, node.Status})
		}

		t.AppendFooter(table.Row{"Total", len(nodes.([]model.Node))})
		t.Render()

	},
}

var nodeDetail = &cobra.Command{
	Use:   "detail",
	Short: "Show detail information of node",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)

		queryArgs["node_uuid"] = nodeUUID

		nodeDetail, err := queryParser.NodeDetail(queryArgs)

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
		nd := nodeDetail.(model.NodeDetail)
		t.AppendHeader(table.Row{"Node UUID", nd.NodeUUID})

		t.AppendRow([]interface{}{"CPU Model", nd.CPUModel})
		t.AppendRow([]interface{}{"CPU Processors", nd.CPUProcessors})
		t.AppendRow([]interface{}{"CPU Threads", nd.CPUThreads})
		t.Render()

	},
}

func ReadyNodeCmd() {
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
	nodeCreate.Flags().StringVar(&active, "active", "", "Active state")
	nodeCreate.Flags().IntVar(&cpuCores, "cpu_cores", 0, "Number of CPU cores")
	nodeCreate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	nodeCreate.MarkFlagRequired("bmc_mac_addr")
	nodeCreate.MarkFlagRequired("bmc_ip")
	nodeCreate.MarkFlagRequired("pxe_mac_addr")
	nodeCreate.MarkFlagRequired("status")
	nodeCreate.MarkFlagRequired("description")
	nodeCreate.MarkFlagRequired("active")
	nodeCreate.MarkFlagRequired("cpu")
	nodeCreate.MarkFlagRequired("memory")

	nodeUpdate.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeUpdate.Flags().StringVar(&bmcMacAddr, "bmc_mac_addr", "", "MAC address of BMC")
	nodeUpdate.Flags().StringVar(&bmcIP, "bmc_ip", "", "IP address of BMC")
	nodeUpdate.Flags().StringVar(&pxeMacAddr, "pxe_mac_addr", "", "PXE MAC address")
	nodeUpdate.Flags().StringVar(&status, "status", "", "status")
	nodeUpdate.Flags().StringVar(&desc, "description", "", "Description")
	nodeUpdate.Flags().StringVar(&active, "active", "", "Active state")
	nodeUpdate.Flags().IntVar(&cpuCores, "cpu", 0, "Number of CPU cores")
	nodeUpdate.Flags().IntVar(&memory, "memory", 0, "Size of memory")
	nodeUpdate.MarkFlagRequired("uuid")

	nodeDelete.Flags().StringVar(&nodeUUID, "uuid", "", "UUID of node")
	nodeDelete.MarkFlagRequired("uuid")

	nodeCreateDetail.Flags().StringVar(&nodeUUID, "node_uuid", "", "UUID of node")
	nodeCreateDetail.Flags().StringVar(&cpuModel, "cpu_model", "", "CPU model")
	nodeCreateDetail.Flags().IntVar(&cpuProcessors, "cpu_processors", 0, "Number of processor")
	nodeCreateDetail.Flags().IntVar(&cpuThreads, "cpu_threads", 0, "Number of logical core")
	nodeCreateDetail.MarkFlagRequired("node_uuid")
	nodeCreateDetail.MarkFlagRequired("cpu_model")
	nodeCreateDetail.MarkFlagRequired("cpu_processors")
	nodeCreateDetail.MarkFlagRequired("cpu_threads")

	nodeCreate.AddCommand(nodeCreateDetail)

	nodeDeleteDetail.Flags().StringVar(&nodeUUID, "node_uuid", "", "UUID of node")
	nodeDeleteDetail.MarkFlagRequired("node_uuid")

	nodeDelete.AddCommand(nodeDeleteDetail)

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
	nodeList.Flags().StringVar(&active, "active", "", "Active status")

	nodeDetail.Flags().StringVar(&nodeUUID, "node_uuid", "", "UUID of node")
	nodeDetail.MarkFlagRequired("node_uuid")

	NodeCmd.AddCommand(nodeOn, nodeOff, nodeRestart,
		nodeCreate, nodeUpdate, nodeDelete,
		nodeList, nodeDetail)
}
