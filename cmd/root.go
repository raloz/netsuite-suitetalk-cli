/*
Copyright © 2024 Octavio Quiroz <octavioquiroz30@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/raloz/suitetalk/cmd/create"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "suitetalk",
	Short: "suitetalk - a simple CLI for RESTful communication with NetSuite",
	Long: `Suitetalk CLI is a simple http client in your terminal to send request to NetSuite.
It allows you manage your connections to your NetSuite instances (almost similar to postman envs),
you can send CRUD requests to make any action as NetSuite Suitetalk Rest allows.

With the suitetalk cli commands you can manage your connections, create, update, delete, and copy
record and transactions.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetUsageTemplate(`Usage:
suitetalk [command]

Available Commands:
create      Send a POST request to NetSuite
help        Help about any command

Flags:
-h, --help   help for suitetalk

Use "suitetalk [command] --help" for more information about a command.
`)
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while suitetalk was executed '%s'", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(create.CreateCmd)
}
