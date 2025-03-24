package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "apiron",
	Short: "A lightweight interactive API testing CLI",
}

func Execute() {
	// Load existing saved requests
	viper.SetConfigFile("requests.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		fmt.Println("⚠ Error loading saved requests:", err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(requestCmd)
	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
}

