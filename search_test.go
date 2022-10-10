package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	req := client.NewSearchRequest()
	// req.RequestBody().SearchRecord.Type = "listRel:CustomerSearch"
	// req.RequestBody().SearchRecord.Basic.CompanyName.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.CompanyName.SearchValue = "test"
	// req.RequestBody().SearchRecord.Basic.Email.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.Email.SearchValue = "@ektos.net"

	req.RequestBody().SearchRecord.Type = "platformCommon:CustomRecordSearchBasic"
	req.RequestBody().SearchRecord.RecType.InternalID = "767"
	req.RequestBody().SearchRecord.RecType.ScriptID = "cseg_nch_property"
	req.RequestBody().SearchRecord.RecType.Type = "customRecordType"
	req.RequestBody().SearchRecord.RecType.XSIType = "platformCore:CustomizationRef"
	req.RequestBody().SearchRecord.RecType.Name.Text = "Property"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
