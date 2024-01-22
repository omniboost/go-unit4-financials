package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetCustomizationID(t *testing.T) {
	req := client.NewGetCustomizationIDRequest()
	req.RequestBody().CustomizationType.GetCustomizationType = "transactionBodyCustomField"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
