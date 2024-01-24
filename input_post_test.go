package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInputPost(t *testing.T) {
	req := client.NewInputPostRequest()

	b := []byte(`{
            "Header": {
                "CurCode": "EUR",
				"Date": "2024-01-22T00:00:00.000",
                "Description": "MEWS 2024-01-22",
                "Key": {
                    "CmpCode": "1110",
                    "Code": "SK-INT"
                },
                "OriginalCompany": "1110",
                "Period": "2024/01"
            },
            "Lines": [
                {
                    "AccountCode": "1103",
                    "Description": "Guest Ledger",
                    "DocValue": 35591.94,
                    "LineSense": "debit",
                    "LineType": "analysis",
                    "Number": 1,
                    "Taxes": []
                },
                {
                    "AccountCode": "1103",
                    "Description": "Revenue",
                    "DocValue": 35591.94,
                    "LineSense": "credit",
                    "LineType": "analysis",
                    "Number": 2,
                    "Taxes": []
                }
			]
		}`)

	err := json.Unmarshal(b, &req.RequestBody().Transaction)
	if err != nil {
		t.Error(err)
	}
	req.RequestBody().PostOptions.Postto = "books"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
