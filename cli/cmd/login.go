package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var username, password string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login as a Federation Admin",
	Run: func(cmd *cobra.Command, args []string) {
		// This is where you would handle the login logic
		fmt.Printf("Logging in as %s with password %s\n", username, password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
}
