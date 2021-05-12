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

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
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
		if len(nodeData.Errors) > 0 {
			for _, hrr := range nodeData.Errors {
				hrr.Println()
			}
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

		fmt.Print("Power On .... ")

		data, err := mutationParser.OnOffNode(queryArgs, model.On)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.PowerStateNode)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")
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

		fmt.Print("Power Off .... ")

		data, err := mutationParser.OnOffNode(queryArgs, model.Off)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.PowerStateNode)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")
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

		fmt.Print("Restart .... ")

		data, err := mutationParser.OnOffNode(queryArgs, model.Restart)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.PowerStateNode)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println("[SUCCESS]")
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
		queryArgs["bmc_ip"] = bmcIP
		queryArgs["description"] = desc
		queryArgs["token"] = config.User.Token

		fmt.Print("Create Node .... ")

		data, err := mutationParser.CreateNode(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.Node)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
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
		t.AppendHeader(table.Row{"UUID", nodeData.UUID})
		t.AppendRow([]interface{}{"Server UUID", nodeData.ServerUUID})
		t.AppendRow([]interface{}{"BMC IP", nodeData.BmcIP})
		t.AppendRow([]interface{}{"BMC MAC", nodeData.BmcMacAddr})
		t.AppendRow([]interface{}{"PXE MAC", nodeData.PXEMacAddr})
		t.AppendRow([]interface{}{"CORES", nodeData.CPUCores})
		t.AppendRow([]interface{}{"MEMORY", nodeData.Memory})
		t.AppendRow([]interface{}{"POWER", nodeData.Status})
		t.AppendRow([]interface{}{"ACTIVE", nodeData.Active})
		t.AppendRow([]interface{}{"DESCRIPTION", nodeData.Description})
		t.Render()
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

		fmt.Print("Update Node .... ")

		data, err := mutationParser.UpdateNode(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.Node)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
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
		t.AppendHeader(table.Row{"UUID", nodeData.UUID})
		t.AppendRow([]interface{}{"Server UUID", nodeData.ServerUUID})
		t.AppendRow([]interface{}{"BMC IP", nodeData.BmcIP})
		t.AppendRow([]interface{}{"BMC MAC", nodeData.BmcMacAddr})
		t.AppendRow([]interface{}{"PXE MAC", nodeData.PXEMacAddr})
		t.AppendRow([]interface{}{"CORES", nodeData.CPUCores})
		t.AppendRow([]interface{}{"MEMORY", nodeData.Memory})
		t.AppendRow([]interface{}{"POWER", nodeData.Status})
		t.AppendRow([]interface{}{"ACTIVE", nodeData.Active})
		t.AppendRow([]interface{}{"DESCRIPTION", nodeData.Description})
		t.Render()
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

		fmt.Print("Delete Node .... ")

		data, err := mutationParser.DeleteNode(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.Node)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
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
		t.AppendHeader(table.Row{"UUID", nodeData.UUID})
		t.AppendRow([]interface{}{"Server UUID", nodeData.ServerUUID})
		t.AppendRow([]interface{}{"BMC IP", nodeData.BmcIP})
		t.AppendRow([]interface{}{"BMC MAC", nodeData.BmcMacAddr})
		t.AppendRow([]interface{}{"PXE MAC", nodeData.PXEMacAddr})
		t.AppendRow([]interface{}{"CORES", nodeData.CPUCores})
		t.AppendRow([]interface{}{"MEMORY", nodeData.Memory})
		t.AppendRow([]interface{}{"POWER", nodeData.Status})
		t.AppendRow([]interface{}{"ACTIVE", nodeData.Active})
		t.AppendRow([]interface{}{"DESCRIPTION", nodeData.Description})
		t.Render()
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

		fmt.Print("Create Node Detail Info .... ")

		data, err := mutationParser.CreateNodeDetail(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.NodeDetail)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
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
		t.AppendHeader(table.Row{"UUID", nodeData.NodeUUID})
		t.AppendRow([]interface{}{"NODE DETAIL", nodeData.NodeDetail})
		t.AppendRow([]interface{}{"NIC DETAIL", nodeData.NicDetail})
		t.Render()
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
		queryArgs["node_uuid"] = nodeUUID
		queryArgs["token"] = config.User.Token

		fmt.Print("Create Node Detail Info .... ")

		data, err := mutationParser.DeleteNodeDetail(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		nodeData := data.(model.NodeDetail)
		if len(nodeData.Errors) > 0 {
			fmt.Println("[FAIL]")
			for _, hrr := range nodeData.Errors {
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
		t.AppendHeader(table.Row{"UUID", nodeData.NodeUUID})
		t.AppendRow([]interface{}{"NODE DETAIL", nodeData.NodeDetail})
		t.AppendRow([]interface{}{"NIC DETAIL", nodeData.NicDetail})
		t.Render()
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
		queryArgs["active"] = strconv.Itoa(active)
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListNode(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		nodeList := data.(model.Nodes)
		if len(nodeList.Errors) > 0 {
			fmt.Println(nodeList.Errors)
			for _, hrr := range nodeList.Errors {
				hrr.Println()
			}
			return
		}

		fmt.Println(nodeList)

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
			"Cores", "Memory", "Active", "Status"})

		for index, node := range nodeList.Nodes {
			t.AppendRow([]interface{}{
				index + 1, node.UUID, node.ServerUUID, node.BmcMacAddr, node.BmcIP, node.PXEMacAddr,
				node.CPUCores, node.Memory, node.Active, node.Status})
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

		nodeDetailData := data.(model.NodeDetailData)
		if len(nodeDetailData.Errors) > 0 {
			for _, hrr := range nodeDetailData.Errors {
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
		t.AppendHeader(table.Row{"No", "CPU ID", "Cores", "Manufacture", "Speed(MHZ)", "Model",
			"Socket", "Threads", "State", "Health"})

		for index, cpu := range nodeDetailData.CPUs {
			t.AppendRow([]interface{}{
				index + 1, cpu.ID, cpu.Cores, cpu.Manufacture, cpu.MaxSpeed,
				cpu.Model, cpu.Socket, cpu.Threads, cpu.Status.State, cpu.Status.Health})
		}
		t.AppendSeparator()
		t.AppendRow([]interface{}{"NO", "MEMORY ID", "CAPACITY", "MANUFACTURE", "SPEED(MHZ)", "DEVICE LOCATOR",
			"PART NUMBER", "SERIAL NUMBER", "STATE", "HEALTH"})
		t.AppendSeparator()

		for index, memory := range nodeDetailData.Memories {
			t.AppendRow([]interface{}{
				index + 1, memory.ID, memory.CapacityMB, memory.Manufacture, memory.Speed, memory.DeviceLocator,
				memory.PartNumber, memory.SerialNumber, memory.Status.State, memory.Status.Health})
		}

		t.AppendSeparator()
		t.AppendRow([]interface{}{"NO", "NIC ID", "MAC", "", "SPEED(GB)", "MODEL", "TYPE", "", "", ""})
		t.AppendSeparator()

		for index, nic := range nodeDetailData.NICs {
			t.AppendRow([]interface{}{
				index + 1, nic.ID, nic.Mac, "", nic.Speed, nic.Model, nic.Type, "", "", ""})
		}
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

	nodeCreate.Flags().StringVar(&bmcIP, "bmc_ip", "", "IP address of BMC [x.x.x.x/x]")
	nodeCreate.Flags().StringVar(&desc, "description", "", "Description")
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
	nodeUpdate.MarkFlagRequired("uuid")

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
