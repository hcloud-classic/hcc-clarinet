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
	//"fmt"
	//"github.com/jedib0t/go-pretty/table"
	//"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	//"hcc/clarinet/action/graphql/mutationParser"
	//"hcc/clarinet/action/graphql/queryParser"
	//"hcc/clarinet/model"
	//"os"
	//"strconv"
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

var startIP, endIP string

func ReadyAIPCmd() {

}
