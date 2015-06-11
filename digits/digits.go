/*
Package digits provides a client and models for Digits API services.
*/
package digits

import (
	"github.com/dghubble/sling"
	"net/http"
)

// DigitsAPI required protocol and domain
const DigitsAPI = "https://api.digits.com"
const apiVersion = "/1.1/"

// Client is a Digits client for making Digits API request
type Client struct {
	sling *sling.Sling
	// Digits API Services
	Accounts *AccountService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(DigitsAPI + apiVersion)
	return &Client{
		sling:    base,
		Accounts: NewAccountService(base.New()),
	}
}
