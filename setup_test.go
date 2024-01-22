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
	user := os.Getenv("FINANCIALS_USER")
	password := os.Getenv("FINANCIALS_PASSWORD")
	companyCode := os.Getenv("FINANCIALS_COMPANY_CODE")
	debug := os.Getenv("DEBUG")

	client = financials.NewClient(nil)
	client.SetUser(user)
	client.SetPassword(password)
	client.SetCompanyCode(companyCode)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	m.Run()
}
