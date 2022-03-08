package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAll(t *testing.T) {
	req := client.NewGetAllRequest()
	req.RequestBody().Record.RecordType = "currency"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
