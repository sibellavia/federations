package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type NewFedAdmin struct {
	Name        string  `json:"name"`
	Email       *string `json:"email,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled     bool    `json:"enabled"`
}

var name, email, description string
var enabled bool

var createFedAdminCmd = &cobra.Command{
	Use:   "createFedAdmin",
	Short: "Create a new Federation Admin",
	Run: func(cmd *cobra.Command, args []string) {
		newFedAdmin := NewFedAdmin{
			Name:        name,
			Email:       &email,
			Description: &description,
			Enabled:     enabled,
		}

		jsonData, err := json.Marshal(newFedAdmin)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		resp, err := http.Post("http://localhost:8083/FHSOperator/NewFedAdmin", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error making POST request:", err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(createFedAdminCmd)

	createFedAdminCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the Federation Admin")
	createFedAdminCmd.Flags().StringVarP(&email, "email", "e", "", "Email of the Federation Admin")
	createFedAdminCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the Federation Admin")
	createFedAdminCmd.Flags().BoolVarP(&enabled, "enabled", "b", false, "Is the Federation Admin enabled")
}
