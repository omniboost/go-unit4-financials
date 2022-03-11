package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-netsuite-soap"
)

func TestCustomerSearch(t *testing.T) {
	req := client.NewCustomerSearchRequest()
	// req.RequestBody().SearchRecord.Basic.Email.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.Email.SearchValue = ".net"
	// req.RequestBody().SearchRecord.Basic.InternalID.Operator = "anyOf"
	// req.RequestBody().SearchRecord.Basic.InternalID.SearchValue = []netsuite.RecordRef{{InternalID: "155937"}}
	req.RequestBody().SearchRecord.Basic.ExternalID.Operator = "anyOf"
	req.RequestBody().SearchRecord.Basic.ExternalID.SearchValue = []netsuite.RecordRef{{ExternalID: "70048129"}}
	// req.RequestBody().SearchRecord.Basic.CompanyName.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.CompanyName.SearchValue = "Quality Hotel Leangkollen AS"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
