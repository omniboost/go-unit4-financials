package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSalesTaxItemSearch(t *testing.T) {
	// client.SearchPreferences.PageSize = 11
	// client.SearchPreferences.ReturnSearchColumns = true
	req := client.NewSalesTaxItemSearchRequest()
	req.RequestBody().SearchRecord.Basic.ItemID.Operator = "is"
	req.RequestBody().SearchRecord.Basic.ItemID.SearchValue = "SS-NO"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
