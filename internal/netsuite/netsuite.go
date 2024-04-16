package netsuite

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/raloz/suitetalk/internal/netsuite/oauth"
)

type netsuite struct {
	config oauth.Config
	client *http.Client
}

func NewClient(config oauth.Config) netsuite {
	return netsuite{
		config: config,
		client: &http.Client{},
	}
}

func (n netsuite) NewRequest(method string, endpoint string, q map[string]string, body []byte) (*http.Response, error) {

	query := []string{}
	for key, value := range q {
		query = append(query, key+"="+value)
	}

	auth := oauth.CreateAuthorizationHeader(
		method,
		endpoint,
		query,
		n.config,
	)

	target := n.config.CompanyUrl + "/" + endpoint

	if len(query) > 0 {
		queryUrlParams, err := url.QueryUnescape(strings.Join(query, "&"))
		if err != nil {
			return nil, err
		}

		target += "?" + queryUrlParams
	}

	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")

	return n.client.Do(req)
}
