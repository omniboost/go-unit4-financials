package financials_test

import (
	"os"
	"testing"

	financials "github.com/omniboost/go-unit4-financials"
)

var (
	client *financials.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tokenID := os.Getenv("TOKEN_ID")
	tokenSecret := os.Getenv("TOKEN_SECRET")
	// applicationID := os.Getenv("APPLICATION_ID")
	accountID := os.Getenv("ACCOUNT_ID")
	debug := os.Getenv("DEBUG")

	client = financials.NewClient(nil)
	client.SetClientID(clientID)
	client.SetClientSecret(clientSecret)
	client.SetTokenID(tokenID)
	client.SetTokenSecret(tokenSecret)
	client.SetAccountID(accountID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	m.Run()
}
