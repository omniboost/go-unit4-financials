package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDocumentMasterList(t *testing.T) {
	req := client.NewDocumentMasterListRequest()
	req.RequestBody().Filter.MaxKeys = 10
	// req.RequestBody().Filter.FullWildKey.CmpCode = os.Getenv("FINANCIALS_COMPANY_CODE")
	// req.RequestBody().Filter.FullWildKey.Level = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}


