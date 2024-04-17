package requests

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/raloz/suitetalk/internal/netsuite"
	"github.com/raloz/suitetalk/internal/netsuite/oauth"
)

func SendPost(endpoint string, data []byte) {
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
}
