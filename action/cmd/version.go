package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show build date of clarinet",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()

		t.SetOutputMirror(os.Stdout)
		t.SetTitle("Clarinet")
		t.AppendRow(table.Row{"Build Date", "07 July, 2021"})

		t.Render()
	},
}
