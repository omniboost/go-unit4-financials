package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestElementMasterAdd(t *testing.T) {
	req := client.NewElementMasterAddRequest()
	b := []byte(`
        {
            "AccountType": "look_left",
            "Addresses": [
                {
                    "Address1": "19 Lonatt Close",
                    "Address2": "",
                    "Category": "RECHNUNG",
                    "Country": "GB",
                    "DefaultAddress": true,
                    "EMail": null,
                    "Name": "Michael Munnelly",
                    "PostCode": "RG31 5HG",
                    "Tel": null
                }
            ],
            "CmpCode": "1110",
            "Code": "9900016228",
            "CustomerSupplier": true,
            "IsCustomer": true,
            "IsSupplier": false,
            "Level": 3,
            "Matchable": true,
            "Name": "Michael Munnelly",
            "PayStatus": "available",
            "Terms": "01LL"
        }
		`)

	err := json.Unmarshal(b, &req.RequestBody().Element)
	if err != nil {
		t.Error(err)
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
