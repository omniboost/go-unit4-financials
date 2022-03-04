package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"

	netsuite "github.com/omniboost/go-netsuite-rest"
)

func TestInvoicePost(t *testing.T) {
	req := client.NewInvoicePostRequest()
	req.RequestBody().Subsidiary.ID = "49"
	// req.RequestBody().Entity.ID = 145
	// req.RequestBody().Location.ID = 5
	req.RequestBody().Item = netsuite.InvoiceItem{
		Items: netsuite.InvoiceItemItems{
			{
				Account: netsuite.Account{
					ID: "635",
				},
				Amount: 80000,
				Item: netsuite.InvoiceItemItemItem{
					ID: 131,
				},
				ItemSubType: "Resale",
				ItemType:    "NonInvtPart",
			},
		},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
