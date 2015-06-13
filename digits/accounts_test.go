package digits

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_Account(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 11, "id_str": "11", "phone_number": "0123456789", "access_token": {"token": "t", "secret": "s"}}`)
	})

	client := NewClient(httpClient)
	account, _, err := client.Accounts.Account()
	if err != nil {
		t.Errorf("Accounts.Account error %v", err)
	}
	expected := &Account{ID: 11, IDStr: "11", PhoneNumber: "0123456789", AccessToken: AccessToken{Token: "t", Secret: "s"}}
	if !reflect.DeepEqual(expected, account) {
		t.Errorf("Accounts.Account expected:\n%+v\ngot:\n%+v", expected, account)
	}
}

func TestAccountService_APIErrorBadRequest(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"errors": [{"message": "Bad Authentication data.", "code": 215}]}`)
	})

	client := NewClient(httpClient)
	_, _, err := client.Accounts.Account()
	expected := &APIError{
		Errors: []ErrorDetail{
			ErrorDetail{Message: "Bad Authentication data.", Code: 215},
		},
	}
	if !reflect.DeepEqual(expected, err) {
		t.Errorf("Accounts.Account APIError expected:\n%+v\ngot:\n%+v", expected, err)
	}
}

func TestAccountService_APIErrorAuthorizationRequired(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, `{"errors": [{"message": "Could not authenticate you.", "code": 32}]}`)
	})

	client := NewClient(httpClient)
	_, _, err := client.Accounts.Account()
	expected := &APIError{
		Errors: []ErrorDetail{
			ErrorDetail{Message: "Could not authenticate you.", Code: 32},
		},
	}
	if !reflect.DeepEqual(expected, err) {
		t.Errorf("Accounts.Account APIError expected:\n%+v\ngot:\n%+v", expected, err)
	}
}
