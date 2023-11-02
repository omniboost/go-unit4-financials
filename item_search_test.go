package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestItemSearch(t *testing.T) {
	req := client.NewItemSearchRequest()
	// req.RequestBody().SearchRecord.Basic.ItemID.Operator = "is"
	// req.RequestBody().SearchRecord.Basic.ItemID.SearchValue = "831"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

