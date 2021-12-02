package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalEntryLinesGet(t *testing.T) {
	req := client.NewJournalEntryLinesGetRequest()
	req.PathParams().ID = 2248
	// req.QueryParams().Fields = netsuite.Fields{"line"}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
