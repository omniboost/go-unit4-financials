package financials_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-unit4-financials"
)

func TestAdd(t *testing.T) {
	req := client.NewAddRequest()
	req.RequestBody().Record.Type = "listRel:Customer"
	req.RequestBody().Record.Record = financials.Customer{}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
