package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetServerTime(t *testing.T) {
	req := client.NewGetServerTimeRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
