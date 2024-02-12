package financials_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestElementMasterGet(t *testing.T) {
	req := client.NewElementMasterGetRequest()
	req.RequestBody().Key.CmpCode = os.Getenv("FINANCIALS_COMPANY_CODE")
	req.RequestBody().Key.Level = 3
	req.RequestBody().Key.Code = "W120011A"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
