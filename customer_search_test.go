package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerSearch(t *testing.T) {
	req := client.NewContactSearchRequest()
	req.RequestBody().SearchRecord.Type = "listRel:CustomerSearch"
	req.RequestBody().SearchRecord.Basic.Email.Operator = "contains"
	req.RequestBody().SearchRecord.Basic.Email.SearchValue = ".net"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
