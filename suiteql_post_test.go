package netsuite_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSuiteqlPost(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.RequestBody().Q = "Select id, email from customer where email = 'leon@omniboost.io'"
	req.RequestBody().Q = "SELECT * FROM customer where customer.firstName = 'Pieter' and lastName = 'Baecke' and customer.subsidiary.id = 46"
	req.RequestBody().Q = "SELECT * FROM department where name like '%196%'"
	req.RequestBody().Q = "SELECT * FROM department"
	req.RequestBody().Q = "SELECT * FROM account"
	req.RequestBody().Q = "SELECT * FROM classification"

	// req.RequestBody().Q = "SELECT * FROM classification where name like '%196%'"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
	os.Exit(12)

	customers, err := resp.ToDepartments(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers)
}

func TestSuiteqlPostCustomers(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	// req.QueryParams().Limit = 12
	req.RequestBody().Q = "SELECT * FROM customer where companyName = 'Agate Utvikling AS_170122'"
	req.RequestBody().Q = "SELECT * FROM customer where lastmodifieddate > '20.03.2020' and lastmodifieddate < '20.03.2022'"
	req.RequestBody().Q = "SELECT * FROM customer where parent = 13299"
	req.RequestBody().Q = "SELECT * FROM customer where id IN(2607)"

	// req.RequestBody().Q = "SELECT * FROM classification where name like '%196%'"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	customers, err := resp.ToCustomers(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers[1])
}

func TestSuiteqlPostAddresses(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.QueryParams().Limit = 12
	req.RequestBody().Q = "SELECT * FROM EntityAddress where nkey IN (458404, 458304)"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	customers, err := resp.ToAddresses(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers[1])
}
