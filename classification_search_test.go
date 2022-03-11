package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClassificationSearch(t *testing.T) {
	req := client.NewClassificationSearchRequest()
	req.RequestBody().SearchRecord.Basic.Name.Operator = "contains"
	req.RequestBody().SearchRecord.Basic.Name.SearchValue = "e"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
