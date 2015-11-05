package digits

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Account is a Digits user account.
type Account struct {
	AccessToken      AccessToken `json:"access_token"`
	CreatedAt        string      `json:"created_at"`
	Email            Email       `json:"email_address"`
	ID               int64       `json:"id"`
	IDStr            string      `json:"id_str"`
	PhoneNumber      string      `json:"phone_number"`
	VerificationType string      `json:"verification_type"`
}

// Email is a Digits user email.
type Email struct {
	Address  string `json:"address"`
	Verified bool   `json:"is_verified"`
}

// AccessToken is a Digits OAuth1 access token and secret.
type AccessToken struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

// AccountService provides methods for accessing Digits Accounts.
type AccountService struct {
	sling *sling.Sling
}

// NewAccountService returns a new AccountService.
func NewAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling.Path("sdk/"),
	}
}

// Account returns the authenticated user Account.
func (s *AccountService) Account() (*Account, *http.Response, error) {
	account := new(Account)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("account.json").Receive(account, apiError)
	return account, resp, firstError(err, apiError)
}
