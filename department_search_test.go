package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDepartmentSearch(t *testing.T) {
	req := client.NewDepartmentSearchRequest()
	// req.RequestBody().SearchRecord.Basic.Name.Operator = "is"
	// req.RequestBody().SearchRecord.Basic.Name.SearchValue = "HTL : 1 : 110"
	// req.RequestBody().SearchRecord.Basic.CustomFieldList = netsuite.SearchCustomFieldList{
	// 	netsuite.SearchStringCustomField{
	// 		ScriptID: "custrecord_nch_departmentname",
	// 		// InternalID:  "2844",
	// 		Operator:    "is",
	// 		SearchValue: "Restaurant 1",
	// 	},
	// }
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
