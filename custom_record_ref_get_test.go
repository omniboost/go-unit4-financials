package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomRecordRefGet(t *testing.T) {
	req := client.NewCustomRecordRefGetRequest()
	req.RequestBody().Xmlns = "urn:messages_2017_1.platform.webservices.financials.com"
	// req.RequestBody().BaseRef.ScriptID = "cseg_nch_property"
	// req.RequestBody().BaseRef.ExternalID = "Apparel"
	req.RequestBody().BaseRef.InternalID = "9"
	req.RequestBody().BaseRef.Type = "customSegment"
	req.RequestBody().BaseRef.XSIType = "platformCore:RecordRef"
	req.RequestBody().BaseRef.Xmlns = "urn:core_2017_1.platform.webservices.financials.com"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
