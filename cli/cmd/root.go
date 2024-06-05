package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fedadmin-cli",
	Short: "FedAdmin CLI is a tool to manage federations",
	Long:  `A Command Line Interface to manage federations and federation admins`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
