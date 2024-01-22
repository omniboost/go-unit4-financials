package financials_test

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
	// req.RequestBody().SearchRecord.RecType.InternalID = "1568"
	// req.RequestBody().SearchRecord.RecType.ScriptID = "custbody_exp_type_journal"
	// req.RequestBody().SearchRecord.RecType.Type = "customRecordType"
	// req.RequestBody().SearchRecord.RecType.XSIType = "platformCore:CustomizationRef"
	// req.RequestBody().SearchRecord.RecType.Name.Text = "Property"

	// client.SearchPreferences.BodyFieldsOnly = false
	// req.RequestBody().SearchRecord.Type = "platformCommon:CustomListSearchBasic"
	// req.RequestBody().SearchRecord.InternalID.Operator = "anyOf"
	// req.RequestBody().SearchRecord.InternalID.SearchValue = []financials.RecordRef{{InternalID: "171"}}
	// req.RequestBody().SearchRecord.RecType.InternalID = "1568"
	// req.RequestBody().SearchRecord.RecType.ScriptID = "custbody_exp_type_journal"
	// req.RequestBody().SearchRecord.RecType.Type = "customRecordType"
	// req.RequestBody().SearchRecord.RecType.XSIType = "platformCore:CustomizationRef"
	// req.RequestBody().SearchRecord.RecType.Name.Text = "Property"

	req.RequestBody().SearchRecord.Type = "platformCommon:CustomRecordSearchBasic"
	req.RequestBody().SearchRecord.RecType.ScriptID = "custbody_journal_memo"
	req.RequestBody().SearchRecord.RecType.Type = "customTransactionType"
	req.RequestBody().SearchRecord.RecType.XSIType = "platformCore:CustomizationRef"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
