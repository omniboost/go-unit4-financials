package financials_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestInputPost(t *testing.T) {
	req := client.NewInputPostRequest()

	b := []byte(`
        {
            "Header": {
                "CurCode": "EUR",
                "Date": "2024-01-30T00:00:00.000",
                "Description": "MEWS 2024-01-30",
                "Key": {
                    "CmpCode": "1110",
                    "Code": "SK-INT"
                },
                "OriginalCompany": "1110",
                "Period": "2024/01"
            },
            "Lines": [
                {
                    "AccountCode": "108.123011",
                    "Description": "Deposits (Negative)",
                    "DocValue": 10441.97,
                    "LineSense": "debit",
                    "LineType": "analysis",
                    "Number": 1,
                    "Taxes": [
                        {
                            "Code": "U070",
                            "Value": -9.59
                        }
                    ]
                },
                {
                    "AccountCode": "108.123011",
                    "Description": "Deposits (Positive)",
                    "DocValue": 24093.79,
                    "LineSense": "credit",
                    "LineType": "analysis",
                    "Number": 2,
                    "Taxes": [
                        {
                            "Code": "U070",
                            "Value": 7.99
                        }
                    ]
                },
                {
                    "AccountCode": "108.123014",
                    "Description": "Omzet fallback MEWS 0%",
                    "DocValue": 40,
                    "LineSense": "debit",
                    "LineType": "analysis",
                    "Number": 5,
                    "Taxes": [
                        {
                            "Code": "U190",
                            "Value": -6.39
                        }
                    ]
                }
            ]
        }
		`)

	err := json.Unmarshal(b, &req.RequestBody().Transaction)
	if err != nil {
		t.Error(err)
	}
	req.RequestBody().PostOptions.Postto = "books"

	b, _ = json.MarshalIndent(req.RequestBody(), "", "  ")
	log.Println(string(b))
	os.Exit(12)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
