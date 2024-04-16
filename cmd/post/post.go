/*
Copyright © 2024 Octavio Quiroz <octavioquiroz30@gmail.com>
*/
package post

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/raloz/suitetalk/internal/netsuite"
	"github.com/raloz/suitetalk/internal/netsuite/oauth"
	"github.com/spf13/cobra"
)

var recordType string
var data string

// postCmd represents the post command
var PostCmd = &cobra.Command{
	Use:   "post",
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

		n := netsuite.NewClient(oauth.Config{
			Account: "Account_ID",
			Consumer: oauth.Consumer{
				Key:    "MyConsumerKey",
				Secret: "MyConsumerSecret",
			},
			Token: oauth.Token{
				Id:     "MyTokenId",
				Secret: "MyTokenSecret",
			},
			CompanyUrl: "https://account-id.suitetalk.api.netsuite.com",
		})

		r, err := n.NewRequest(http.MethodPost, endpoint, map[string]string{}, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		if r.StatusCode >= 300 {
			fmt.Fprintf(os.Stderr, "posting to netsuite: %s\n", body)
			os.Exit(1)
		}

		fmt.Println(string(body))
	},
}

func init() {
	PostCmd.Flags().StringVarP(&recordType, "type", "t", "", "record type to create [customer, vendor, employee, item, etc]")
	PostCmd.Flags().StringVarP(&data, "data", "d", "", "payload for send as request body to netsuite")

	PostCmd.MarkFlagRequired("type")
	PostCmd.MarkFlagRequired("data")
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
