package digits

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Digits API and version
const DigitsAPI = "https://api.digits.com"
const apiVersion = "/1.1/"

// Client is a Digits client for making Digits API requests.
type Client struct {
	sling *sling.Sling
	// Digits API Services
	Accounts *AccountService
	Contacts *ContactService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(DigitsAPI + apiVersion)
	return &Client{
		sling:    base,
		Accounts: NewAccountService(base.New()),
		Contacts: NewContactService(base.New()),
	}
}
