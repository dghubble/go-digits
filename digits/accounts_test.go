package digits

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountService_Account(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 11, "id_str": "11", "phone_number": "0123456789", "email_address":{"address":"user@example.com","is_verified":true}, "access_token": {"token": "t", "secret": "s"}, "verification_type":"sms"}`)
	})
	expected := &Account{
		AccessToken:      AccessToken{Token: "t", Secret: "s"},
		Email:            Email{Address: "user@example.com", Verified: true},
		ID:               11,
		IDStr:            "11",
		PhoneNumber:      "0123456789",
		VerificationType: "sms",
	}

	client := NewClient(httpClient)
	account, _, err := client.Accounts.Account()
	assert.Nil(t, err)
	assert.Equal(t, expected, account)
}

func TestAccountService_APIError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"errors": [{"message": "Bad Authentication data.", "code": 215}]}`)
	})
	expected := &APIError{
		Errors: []ErrorDetail{
			ErrorDetail{Message: "Bad Authentication data.", Code: 215},
		},
	}

	client := NewClient(httpClient)
	_, _, err := client.Accounts.Account()
	if assert.Error(t, err) {
		assert.Equal(t, expected, err)
	}
}
