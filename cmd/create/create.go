/*
Copyright © 2024 Octavio Quiroz <octavioquiroz30@gmail.com>
*/
package create

import (
	"fmt"
	"os"

	"github.com/raloz/suitetalk/pkg/requests"
	"github.com/spf13/cobra"
)

var recordType string
var data string

// CreateCmd represents the post command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Send a POST request to NetSuite",
	Long: `Create a brand new record sending a POST request just indicating 
the record type and the payload`,
	Args:    cobra.NoArgs,
	Example: `suitetalk post --type vendor --data '{"companyName": "John Doe"}'`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, err := getEndpoint(recordType)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		//Sending the post to create the new record into NetSuite
		requests.SendPost(endpoint, []byte(data))
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&recordType, "type", "t", "", "record type to create [customer, vendor, employee, item, etc]")
	CreateCmd.Flags().StringVarP(&data, "data", "d", "", "payload for send as request body to netsuite")

	CreateCmd.MarkFlagRequired("type")
	CreateCmd.MarkFlagRequired("data")
}

func getEndpoint(recordType string) (string, error) {
	records := map[string]string{
		"vendor": "services/rest/record/v1/vendor",
	}

	endpoint, ok := records[recordType]
	if !ok {
		return "", fmt.Errorf("error getting endpoint: invalid record type")
	}

	return endpoint, nil
}
