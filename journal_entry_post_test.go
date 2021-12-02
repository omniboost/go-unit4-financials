package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"

	netsuite "github.com/omniboost/go-netsuite"
)

func TestJournalEntryPost(t *testing.T) {
	req := client.NewJournalEntryPostRequest()
	req.RequestBody().Subsidiary.ID = "3"
	req.RequestBody().Lines = netsuite.JournalEntryLine{
		Items: netsuite.JournalEntryLineElements{
			{
				Account: netsuite.Account{
					ID: 213,
					// RefName: "722280 Other Exp. : IntExp : IntExp Loan Borrowings (IC)",
					// AcctNumber: "121110",
				},
				Debit:  100,
				Credit: 0,
			},
			{
				Account: netsuite.Account{
					ID: 213,
					// RefName: "722280 Other Exp. : IntExp : IntExp Loan Borrowings (IC)",
					// AcctNumber: "121110",
				},
				Debit:  0,
				Credit: 100,
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
