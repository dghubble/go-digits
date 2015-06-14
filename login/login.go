package login

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/dghubble/go-digits/digits"
	"github.com/dghubble/sling"
)

const (
	accountEndpointKey      = "accountEndpoint"
	accountRequestHeaderKey = "accountRequestHeader"
)

// Errors for missing data, invalid data, and errors gettting Digits accounts
var (
	ErrMissingAccountEndpoint      = fmt.Errorf("digits: missing Digits OAuth Echo %s in POST form", accountEndpointKey)
	ErrMissingAccountRequestHeader = fmt.Errorf("digits: missing Digits OAuth Echo %s in POST form", accountRequestHeaderKey)
	ErrInvalidDigitsEndpoint       = errors.New("digits: invalid Digits endpoint")
	ErrInvalidConsumerKey          = errors.New("digits: incorrect Digits OAuth Echo Auth Header Consumer Key")
	ErrUnableToGetDigitsAccount    = errors.New("digits: unable to get Digits account")
	consumerKeyRegexp              = regexp.MustCompile("oauth_consumer_key=\"(.*?)\"")
)

// Service provides a Digits login handler which validates logins and
// retrieves Digits accounts.
type Service struct {
	consumerKey string
	httpClient  *http.Client
}

// NewService returns a new login Service.
func NewService(consumerKey string) *Service {
	return &Service{
		consumerKey: consumerKey,
		httpClient:  http.DefaultClient,
	}
}

// LoginHandlerFunc receives POST'ed Digits OAuth Echo headers, validates them,
// retrieves the Digits user account, and calls the given success or failure
// handler function.
func (s *Service) LoginHandlerFunc(success func(http.ResponseWriter, *http.Request, *digits.Account), failure func(http.ResponseWriter, error, int)) http.Handler {
	return s.loginHandler(successHandlerFunc(success), errorHandlerFunc(failure))
}

// loginHandler is the implementation of LoginHandlerFunc which accepts success
// and failure http.Handler's.
func (s *Service) loginHandler(success successHandler, failure errorHandler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		// validate POST'ed Digits OAuth Echo data
		req.ParseForm()
		accountEndpoint := req.PostForm.Get(accountEndpointKey)
		accountRequestHeader := req.PostForm.Get(accountRequestHeaderKey)
		err := s.validateEcho(accountEndpoint, accountRequestHeader)
		if err != nil {
			failure.ServeHTTP(w, err, http.StatusBadRequest)
			return
		}

		// fetch Digits account from the API
		account, resp, err := requestAccount(s.httpClient, accountEndpoint, accountRequestHeader)

		// validate the Digits Account
		err = validateAccountResponse(account, resp, err)
		if err != nil {
			failure.ServeHTTP(w, err, http.StatusBadRequest)
			return
		}
		success.ServeHTTP(w, req, account)
	}

	return http.HandlerFunc(fn)
}

// requestAccount makes a request to the Digits account endpoint using the
// provided Authorization header.
func requestAccount(client *http.Client, accountEndpoint, authorizationHeader string) (*digits.Account, *http.Response, error) {
	request, err := http.NewRequest("GET", accountEndpoint, nil)
	if err != nil {
		return nil, nil, ErrInvalidDigitsEndpoint
	}
	request.Header.Set("Authorization", authorizationHeader)
	account := new(digits.Account)
	resp, err := sling.New().Client(client).Do(request, account, nil)
	return account, resp, err
}

// validateEcho checks that the Digits OAuth Echo arguments are valid. If the
// endpoint does not match the Digits API or the header does not include the
// correct consumer key, a non-nil error is returned.
func (s *Service) validateEcho(accountEndpoint, accountRequestHeader string) error {
	if accountEndpoint == "" {
		return ErrMissingAccountEndpoint
	}
	if accountRequestHeader == "" {
		return ErrMissingAccountRequestHeader
	}
	// check accountEndpoint matches expected protocol/domain
	if !strings.HasPrefix(accountEndpoint, digits.DigitsAPI) {
		return ErrInvalidDigitsEndpoint
	}
	// validate the OAuth Echo data's auth header consumer key
	matches := consumerKeyRegexp.FindStringSubmatch(accountRequestHeader)
	if len(matches) != 2 || matches[1] != s.consumerKey {
		return ErrInvalidConsumerKey
	}
	return nil
}

// validateAccountResponse checks that the response to the Digits account
// endpoint was successful and that a valid Account was received. Otherwise
// return a non-nil error.
func validateAccountResponse(account *digits.Account, resp *http.Response, err error) error {
	if err != nil || resp.StatusCode != http.StatusOK || account == nil {
		return ErrUnableToGetDigitsAccount
	}
	if token := account.AccessToken; token.Token == "" || token.Secret == "" {
		// JSON deserialized Digits account is missing fields
		return ErrUnableToGetDigitsAccount
	}
	return nil
}

// ErrorHandler replies to requests with the given error message and code.
// This handler is appropriate for most login handler uses, unless custom
// error messages/codes should be returned.
func ErrorHandler(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

// Internal Handlers

// successHandler is called when account login succeeds.
type successHandler interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request, account *digits.Account)
}

type successHandlerFunc func(w http.ResponseWriter, req *http.Request, account *digits.Account)

func (f successHandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request, account *digits.Account) {
	f(w, req, account)
}

// errorHandler is called when account login fails.
type errorHandler interface {
	ServeHTTP(w http.ResponseWriter, err error, code int)
}

type errorHandlerFunc func(w http.ResponseWriter, err error, code int)

func (f errorHandlerFunc) ServeHTTP(w http.ResponseWriter, err error, code int) {
	f(w, err, code)
}
