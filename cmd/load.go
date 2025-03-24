package cmd

import (
	"apiron/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [name]",
	Short: "Load and send a saved API request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request, found := internal.LoadRequest(args[0])
		if !found {
			fmt.Println("⚠ Request not found!")
			return
		}

		fmt.Printf("📤 Sending saved request: %s %s\n", request.Method, request.URL)
		internal.SendRequest(request.URL, request.Method, request.Headers, []byte(request.Body))
	},
}

