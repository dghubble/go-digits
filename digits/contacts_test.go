package digits

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/contacts/users_and_uploaded_by.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"count": "20", "next_cursor": "9876543"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"users": [{"id": 11, "id_str": "11"}], "next_cursor": "9876543"}`)
	})
	expected := &Contacts{
		Users: []Account{
			{ID: 11, IDStr: "11"},
		},
		NextCursor: "9876543",
	}

	client := NewClient(httpClient)
	params := &MatchesParams{Count: 20, NextCursor: "9876543"}
	contacts, _, err := client.Contacts.Matches(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, contacts)
}

func TestContactService_MatchesAPIError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/contacts/users_and_uploaded_by.json", func(w http.ResponseWriter, r *http.Request) {
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
	_, _, err := client.Contacts.Matches(&MatchesParams{})
	if assert.Error(t, err) {
		assert.Equal(t, expected, err)
	}
}
