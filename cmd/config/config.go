/*
Copyright © 2024 Octavio Quiroz <octavioquiroz30@gmail.com>
*/
package config

import (
	"github.com/spf13/cobra"
)

var name string
var company string
var consumer []string
var token []string

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "manage your NetSuite instances configuration",
	Long: `The config command allow you manage your configurations.
You can add a new NetSuite instance config, list all your existing configs or
remove them.

All the configurations are saved in you ~/.suitetalk/configs/ directory, every 
configuration is a json file and has a unique name.`,
	Args: cobra.NoArgs,
}

func init() {
	ConfigCmd.AddCommand(addCmd)

	//adding the flags to create a new configuration file
	addCmd.Flags().StringVarP(&name, "name", "n", "", "the name represent how identify the configuration in a friendly way")
	addCmd.Flags().StringVarP(&company, "company", "i", "", "the netsuite account id")
	addCmd.Flags().StringArrayVarP(&consumer, "consumer", "c", []string{}, "the consumer key,secret pair values")
	addCmd.Flags().StringArrayVarP(&token, "token", "t", []string{}, "the token access,secret pair values")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("company")
	addCmd.MarkFlagRequired("consumer")
	addCmd.MarkFlagRequired("token")
}
