package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSuiteqlPost(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.RequestBody().Q = "Select id, email from customer where email = 'leon@omniboost.io'"
	req.RequestBody().Q = "SELECT * FROM customer where customer.firstName = 'Pieter' and lastName = 'Baecke' and customer.subsidiary.id = 46"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))

	customers, err := resp.ToCustomers()
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers)
}
