package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	req := client.NewAuthenticateRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
