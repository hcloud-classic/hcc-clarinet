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
	"github.com/spf13/cobra"
	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/data"
)

// nodeCmd represents the node command
var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Running node related commands",
	Long:  `Running node related commands`,
	Args:  cobra.MinimumNArgs(1),
}

var nodeUUID string

var nodeOn = &cobra.Command{
	Use:   "on",
	Short: "Power on specified node",
	Long:  `Power on specified node`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]interface{})
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.OnOffNode(queryArgs, data.On)
		if err != nil {
			fmt.Println("Wrong command arguments. Abort.")
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
		queryArgs := make(map[string]interface{})
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.OnOffNode(queryArgs, data.Off)
		if err != nil {
			fmt.Println("Wrong command arguments. Abort.")
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
		queryArgs := make(map[string]interface{})
		queryArgs["uuid"] = nodeUUID
		node, err := mutationParser.OnOffNode(queryArgs, data.Restart)
		if err != nil {
			fmt.Println("Wrong command arguments. Abort.")
			return
		}

		fmt.Println("Flute :" + node.(string))
	},
}

func ReadyNodeCmd() {
	nodeOn.Flags().StringVar(&nodeUUID, "node_UUID", "", "UUID of node")
	nodeOff.Flags().StringVar(&nodeUUID, "node_UUID", "", "UUID of node")
	nodeRestart.Flags().StringVar(&nodeUUID, "node_UUID", "", "UUID of node")

	NodeCmd.AddCommand(nodeOn, nodeOff, nodeRestart)
}
