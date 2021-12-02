package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalEntryGet(t *testing.T) {
	req := client.NewJournalEntryGetRequest()
	req.PathParams().ID = 1083
	// req.QueryParams().Fields = netsuite.Fields{"line"}
	req.QueryParams().ExpandSubResources = false
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
