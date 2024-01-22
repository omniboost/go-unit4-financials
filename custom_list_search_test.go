package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomListSearch(t *testing.T) {
	req := client.NewCustomListSearchRequest()
	// req.RequestBody().SearchRecord.Basic.ExternalID.Operator = "anyOf"
	// req.RequestBody().SearchRecord.Basic.ExternalID.SearchValue = []financials.RecordRef{{ExternalID: "70048129"}}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
