package cmd

import (
	"fmt"

	"github.com/aj-seven/apiron/internal"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [name]",
	Short: "Save an API request for later use",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		method, _ := cmd.Flags().GetString("method")
		headersStr, _ := cmd.Flags().GetString("headers")
		body, _ := cmd.Flags().GetString("body")

		if url == "" || method == "" {
			fmt.Println("⚠ URL and method are required to save a request.")
			return
		}

		headers := internal.ParseHeaders(headersStr)
		internal.SaveRequest(args[0], url, method, headers, body)
		fmt.Printf("✅ Request saved as '%s'\n", args[0])
	},
}

func init() {
	saveCmd.Flags().StringP("url", "u", "", "API URL to save")
	saveCmd.Flags().StringP("method", "m", "GET", "HTTP method")
	saveCmd.Flags().StringP("headers", "H", "", "Request headers in JSON format")
	saveCmd.Flags().StringP("body", "b", "", "Request body in JSON format")
}
