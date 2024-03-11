package financials_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestTaxList(t *testing.T) {
	req := client.NewTaxListRequest()
	req.RequestBody().Filter.MaxKeys = 10
	req.RequestBody().Filter.Key.CmpCode = os.Getenv("FINANCIALS_COMPANY_CODE")
	req.RequestBody().Filter.Key.Code = "*"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}


