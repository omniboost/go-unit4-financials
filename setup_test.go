package netsuite_test

import (
	"os"
	"testing"

	netsuite "github.com/omniboost/go-netsuite-soap"
)

var (
	client *netsuite.Client
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

	client = netsuite.NewClient(nil)
	// client.SetApplicationID(applicationID)
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
