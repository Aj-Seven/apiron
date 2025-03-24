package cmd

import (
	"apiron/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved API requests",
	Run: func(cmd *cobra.Command, args []string) {
		internal.ListRequests()
	},
}
