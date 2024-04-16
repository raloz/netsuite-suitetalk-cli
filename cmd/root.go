/*
Copyright © 2024 Octavio Quiroz <octavioquiroz30@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/raloz/suitetalk/cmd/post"
)

var rootCmd = &cobra.Command{
	Use:   "suitetalk",
	Short: "suitetalk - a simple CLI for REST communication with NetSuite",
	Long: `suitetalk is a simple CLI for REST communication with NetSuite
	
Using suitetalk, you can:

> Use CRUD (create, read, update and delete) operations to perform business processing on NetSuite records
and navigate dynamically between records.
> Get and process the API definition record metadata.
> Execute NetSuite queries and records.`,

	Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.AddCommand(post.PostCmd)
}
