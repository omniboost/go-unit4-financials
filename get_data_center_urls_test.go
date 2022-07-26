package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetDataCenterURLs(t *testing.T) {
	req := client.NewGetDataCenterURLsRequest()
	req.RequestBody().Account = "6239966_SB1"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
