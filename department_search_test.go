package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-netsuite-soap"
)

func TestDepartmentSearch(t *testing.T) {
	req := client.NewDepartmentSearchRequest()
	req.RequestBody().SearchRecord.Type = "listAcct:DepartmentSearch"
	// req.RequestBody().SearchRecord.Basic.Name.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.Name.SearchValue = "1"
	req.RequestBody().SearchRecord.Basic.CustomFieldList = netsuite.SearchCustomFieldList{
		netsuite.SearchStringCustomField{
			ScriptID: "custrecord_nch_departmentname",
			// InternalID:  "2844",
			Operator:    "is",
			SearchValue: "Restaurant 1",
		},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
