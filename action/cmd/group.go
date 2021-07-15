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

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Commands for group",
	Long:  `Commands for create and update, delete groups`,
	Args:  cobra.MinimumNArgs(1),
}

var groupCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create the Group",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["group_id"] = strconv.Itoa(int(groupID))
		queryArgs["group_name"] = name
		queryArgs["token"] = config.User.Token

		fmt.Print("Creating Group .... ")

		data, err := mutationParser.CreateGroup(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		groupData := data.(model.Group)
		if len(groupData.Errors) > 0 {
			for _, hrr := range groupData.Errors {
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
		t.AppendHeader(table.Row{"ID", groupData.ID})
		t.AppendRow([]interface{}{"Name", groupData.Name})
		t.Render()
	},
}

var groupUpdate = &cobra.Command{
	Use:     "update",
	Short:   "Update the Group",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["group_id"] = id
		queryArgs["group_name"] = name
		queryArgs["token"] = config.User.Token

		fmt.Print("Updating Group .... ")

		data, err := mutationParser.UpdateGroup(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		groupData := data.(model.Group)
		if len(groupData.Errors) > 0 {
			for _, hrr := range groupData.Errors {
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
		t.AppendHeader(table.Row{"Group ID", groupData.ID})
		t.AppendRow([]interface{}{"Group Name", groupData.Name})
		t.Render()
	},
}

var groupDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete the Group",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["group_id"] = id
		queryArgs["token"] = config.User.Token

		fmt.Print("Deleting Group .... ")

		data, err := mutationParser.DeleteGroup(queryArgs)
		if err != nil {
			fmt.Println("[FAIL]")
			err.Println()
			return
		}

		groupData := data.(model.Group)
		if len(groupData.Errors) > 0 {
			for _, hrr := range groupData.Errors {
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
		t.AppendHeader(table.Row{"ID", id})
		t.Render()
	},
}

var groupList = &cobra.Command{
	Use:     "list",
	Short:   "Show group list",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["token"] = config.User.Token

		data, err := queryParser.AllGroup(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		groupList := data.(model.GroupList)
		if len(groupList.Errors) > 0 {
			for _, hrr := range groupList.Errors {
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
		t.AppendHeader(table.Row{"Group ID", "Group Name"})

		for index, group := range groupList.Groups {
			t.AppendRow([]interface{}{index + 1, group.ID, group.Name})
		}

		t.AppendFooter(table.Row{"Total", len(groupList.Groups)})
		t.Render()
	},
}

func ReadyGroupCmd() {
	groupCreate.Flags().Int64Var(&groupID, "group_id", 0, "Group ID")
	groupCreate.Flags().StringVar(&name, "group_name", "", "Group Name")
	groupCreate.MarkFlagRequired("group_id")
	groupCreate.MarkFlagRequired("group_name")

	groupUpdate.Flags().Int64Var(&groupID, "group_id", 0, "Group ID")
	groupUpdate.Flags().StringVar(&name, "group_name", "", "Group Name")
	groupUpdate.MarkFlagRequired("group_id")
	groupUpdate.MarkFlagRequired("group_name")

	groupDelete.Flags().StringVar(&id, "group_id", "", "ID")
	groupDelete.MarkFlagRequired("group_id")

	groupList.Flags().StringVar(&id, "id", "", "ID")
	groupList.Flags().Int64Var(&groupID, "group_id", 0, "Group ID")
	groupList.Flags().StringVar(&authentication, "authentication", "", "Authentication (admin/group)")
	groupList.Flags().StringVar(&name, "name", "", "Name")
	groupList.Flags().StringVar(&email, "email", "", "E-Mail")

	groupCmd.AddCommand(groupCreate, groupUpdate, groupDelete, groupList)
}
