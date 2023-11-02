package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountSearch(t *testing.T) {
	req := client.NewAccountSearchRequest()
	// req.RequestBody().SearchRecord.Basic.Number.Operator = "is"
	// req.RequestBody().SearchRecord.Basic.Number.SearchValue = "1524"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
