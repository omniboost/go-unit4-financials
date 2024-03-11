package financials_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestTaxGet(t *testing.T) {
	req := client.NewTaxGetRequest()
	req.RequestBody().Key.CmpCode = os.Getenv("FINANCIALS_COMPANY_CODE")
	req.RequestBody().Key.Code = "U000-NSTB-DE"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

