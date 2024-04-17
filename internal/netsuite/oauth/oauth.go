package oauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"slices"
	"strings"
	"time"
)

type Config struct {
	Account    string
	Consumer   Consumer
	Token      Token
	CompanyUrl string
}

type Consumer struct {
	Key    string
	Secret string
}

type Token struct {
	Id     string
	Secret string
}

func CreateAuthorizationHeader(method, baseUrl string, q []string, config Config) string {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	nonce := generateNonce()

	signatureBase := method + "&" + url.QueryEscape(string(config.CompanyUrl)+"/"+baseUrl) + "&"
	params := []string{
		url.QueryEscape("oauth_consumer_key=" + config.Consumer.Key),
		url.QueryEscape("oauth_nonce=" + nonce),
		url.QueryEscape("oauth_signature_method=HMAC-SHA256"),
		url.QueryEscape("oauth_timestamp=" + timestamp),
		url.QueryEscape("oauth_token=" + config.Token.Id),
		url.QueryEscape("oauth_version=1.0"),
	}

	if len(q) > 0 {
		for _, value := range q {
			params = append(params, url.QueryEscape(value))
		}
	}

	slices.Sort(params)

	signatureBase += strings.Join(params, "%26")

	signatureKey := []byte(url.QueryEscape(config.Consumer.Secret) + "&" + url.QueryEscape(config.Token.Secret))

	h := hmac.New(sha256.New, signatureKey)
	h.Write([]byte(signatureBase))

	signature := h.Sum(nil)

	auth := fmt.Sprintf("OAuth realm=\"%s\",oauth_consumer_key=\"%s\",oauth_token=\"%s\",oauth_signature_method=\"%s\",oauth_timestamp=\"%s\",oauth_nonce=\"%s\",oauth_version=\"%s\",oauth_signature=\"%s\"",
		config.Account,
		config.Consumer.Key,
		config.Token.Id,
		"HMAC-SHA256",
		timestamp,
		nonce,
		"1.0",
		url.QueryEscape(base64.StdEncoding.EncodeToString(signature)),
	)

	return auth
}

func generateNonce() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nonce := make([]byte, 11)
	for i := range nonce {
		nonce[i] = charset[rand.Intn(len(charset))]
	}

	return string(nonce)
}
