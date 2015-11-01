package digits

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Contacts represents a cursored subset of a set of contacts.
type Contacts struct {
	NextCursor string `json:"next_cursor"`
	// matched Digits Users are partially hydrated Accounts (i.e. no token, phone)
	Users []Account `json:"users"`
}

// ContactService provides methods for accessing Digits contacts.
type ContactService struct {
	sling *sling.Sling
}

// NewContactService returns a new ContactService.
func NewContactService(sling *sling.Sling) *ContactService {
	return &ContactService{
		sling: sling.Path("contacts/"),
	}
}

// MatchesParams are the parameters for ContactService.Matches
type MatchesParams struct {
	NextCursor string `url:"next_cursor,omitempty"`
	Count      int    `url:"count,omitempty"`
}

// Matches returns Contacts with the cursored Digits Accounts which have logged
// into the Digits application and are known to the authenticated user Account
// via address book uploads.
func (s *ContactService) Matches(params *MatchesParams) (*Contacts, *http.Response, error) {
	contacts := new(Contacts)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("users_and_uploaded_by.json").QueryStruct(params).Receive(contacts, apiError)
	return contacts, resp, firstError(err, apiError)
}
