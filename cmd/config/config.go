/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "manage your NetSuite instances configuration",
	Long: `The config command allow you manage your configurations.
You can add a new NetSuite instance config, list all your existing configs or
remove them.

All the configurations are saved in you ~/.suitetalk/configs/ directory, every 
configuration is a json file and has a unique name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello word")
	},
}

func init() {

}
