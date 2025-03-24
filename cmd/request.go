package cmd

import (
	"apiron/internal"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Start interactive API request builder",
	Run: func(cmd *cobra.Command, args []string) {
		runInteractiveMode()
	},
}

func runInteractiveMode() {
	fmt.Println("🚀 Welcome to APIron Interactive Mode!")

	urlPrompt := promptui.Prompt{
		Label: "Enter API URL",
		Validate: func(input string) error {
			if !strings.HasPrefix(input, "http") {
				return fmt.Errorf("invalid URL")
			}
			return nil
		},
	}
	url, _ := urlPrompt.Run()

	methodPrompt := promptui.Select{
		Label: "Select HTTP Method",
		Items: []string{"GET", "POST", "PUT", "DELETE"},
	}
	_, method, _ := methodPrompt.Run()

	headers := internal.GetHeadersFromUser()
	body := internal.GetBodyFromUser(method)

	internal.SendRequest(url, method, headers, []byte(body)) // Convert string to []byte

	savePrompt := promptui.Prompt{
		Label:     "Save this request? (y/n)",
		IsConfirm: true,
	}
	save, _ := savePrompt.Run()
	if strings.ToLower(save) == "y" {

		namePrompt := promptui.Prompt{
			Label: "Enter a name for the request",
		}
		name, _ := namePrompt.Run()
		if name == "" {
			name = "request"
		}
		internal.SaveRequest(name, url, method, headers, body) // Convert string to []byte
	}
}
