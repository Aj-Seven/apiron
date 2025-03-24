package cmd

import (
	"fmt"

	"github.com/aj-seven/apiron/internal"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a saved API request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if internal.DeleteRequest(args[0]) {
			fmt.Printf("🗑️ Deleted request '%s'\n", args[0])
		} else {
			fmt.Println("⚠ Request not found!")
		}
	},
}
