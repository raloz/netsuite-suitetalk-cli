package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "add a new configuration to connect with NetSuite",
	Example: `suitetalk config add -n sandbox -c 12345_SB1 -i key,secret -t access,secret`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Print(name)
		fmt.Print(consumer)
	},
}
