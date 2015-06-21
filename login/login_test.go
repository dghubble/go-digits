package login

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dghubble/go-digits/digits"
)

const (
	testConsumerKey          = "mykey"
	testAccountEndpoint      = "https://api.digits.com/1.1/sdk/account.json"
	testAccountRequestHeader = `OAuth oauth_consumer_key="mykey",`
	testAccountJSON          = `{"access_token": {"token": "t", "secret": "s"}, "phone_number": "0123456789"}`
)

func TestValidateEcho_missingAccountEndpoint(t *testing.T) {
	s := NewService(testConsumerKey)
	err := s.validateEcho("", testAccountRequestHeader)
	if err != ErrMissingAccountEndpoint {
		t.Errorf("expected error %v, got %v", ErrMissingAccountEndpoint, err)
	}
}

func TestValidateEcho_missingAccountRequestHeader(t *testing.T) {
	s := NewService(testConsumerKey)
	err := s.validateEcho(testAccountEndpoint, "")
	if err != ErrMissingAccountRequestHeader {
		t.Errorf("expected error %v, got %v", ErrMissingAccountRequestHeader, err)
	}
}

func TestValidateEcho_digitsEndpoint(t *testing.T) {
	cases := []struct {
		endpoint string
		valid    bool
	}{
		{"https://api.digits.com/1.1/sdk/account.json", true},
		{"http://api.digits.com/1.1/sdk/account.json", false},
		{"https://digits.com/1.1/sdk/account.json", false},
		{"https://evil.com/1.1/sdk/account.json", false},
		// respect the path defined in Digits javascript sdk
		{"https://api.digits.com/2.0/future/so/cool.json", true},
	}
	s := NewService(testConsumerKey)
	for _, c := range cases {
		err := s.validateEcho(c.endpoint, testAccountRequestHeader)
		if c.valid && err != nil {
			t.Errorf("expected endpoint %q to be valid, got error %v", c.endpoint, err)
		}
		if !c.valid && err != ErrInvalidDigitsEndpoint {
			t.Errorf("expected endpoint %q to be invalid, got error %v", c.endpoint, err)
		}
	}
}

func TestValidateEcho_headerConsumerKey(t *testing.T) {
	cases := []struct {
		header string
		valid  bool
	}{
		{`OAuth oauth_consumer_key="mykey"`, true},
		// wrong consumer key
		{`OAuth oauth_consumer_key="wrongkey"`, false},
		// empty consumer key
		{`OAuth oauth_consumer_key=""`, false},
		// missing value quotes
		{`OAuth oauth_consumer_key=mykey`, false},
		// no oauth_consumer_key field
		{`OAuth oauth_token="mykey"`, false},
		{"OAuth", false},
	}
	s := NewService(testConsumerKey)
	for _, c := range cases {
		err := s.validateEcho(testAccountEndpoint, c.header)
		if c.valid && err != nil {
			t.Errorf("expected header %q to be valid, got error %v", c.header, err)
		}
		if !c.valid && err != ErrInvalidConsumerKey {
			t.Errorf("expected header %q to be invalid, got error %v", c.header, err)
		}
	}
}

func TestValidateAccountResponse(t *testing.T) {
	emptyAccount := new(digits.Account)
	validAccount := &digits.Account{
		AccessToken: digits.AccessToken{Token: "token", Secret: "secret"},
	}
	successResp := &http.Response{
		StatusCode: 200,
	}
	badResp := &http.Response{
		StatusCode: 400,
	}
	respErr := errors.New("some error decoding Account")

	// success case
	if err := ValidateAccountResponse(validAccount, successResp, nil); err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	// error cases
	errorCases := []error{
		// account missing credentials
		ValidateAccountResponse(emptyAccount, successResp, nil),
		// Digits account API did not return a 200
		ValidateAccountResponse(validAccount, badResp, nil),
		// Network error or JSON unmarshalling error
		ValidateAccountResponse(validAccount, successResp, respErr),
		ValidateAccountResponse(validAccount, badResp, respErr),
	}
	for _, err := range errorCases {
		if err != ErrUnableToGetDigitsAccount {
			t.Errorf("expected %v, got %v", ErrUnableToGetDigitsAccount, err)
		}
	}
}

func TestErrorHandler(t *testing.T) {
	const expectedMessage = "digits: some error"
	rec := httptest.NewRecorder()
	// should pass through errors and codes
	DefaultErrorHandler.ServeHTTP(rec, fmt.Errorf(expectedMessage), http.StatusBadRequest)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected code %v, got %v", http.StatusBadRequest, rec.Code)
	}
	if rec.Body.String() != expectedMessage+"\n" {
		t.Errorf("expected error message %v, got %v", expectedMessage+"\n", rec.Body.String())
	}
}

func TestLoginHandler_successEndToEnd(t *testing.T) {
	digitsProxyClient, _, server := setupDigitsTestServer(testAccountJSON)
	defer server.Close()

	// setup test server which uses go-digits/login for Digits login
	s := NewService(testConsumerKey)
	// proxies all requests to the digits test server
	s.httpClient = digitsProxyClient
	ts := httptest.NewServer(s.LoginHandler(SuccessHandlerFunc(successChecks(t)), ErrorHandlerFunc(errorOnFailure(t))))

	// POST Digits OAuth Echo headers
	resp, err := http.PostForm(ts.URL, url.Values{"accountEndpoint": {testAccountEndpoint}, "accountRequestHeader": {testAccountRequestHeader}})
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected StatusCode %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestLoginHandlerFunc_invalidPOSTArguments(t *testing.T) {
	digitsProxyClient, _, server := setupDigitsTestServer(testAccountJSON)
	defer server.Close()

	// setup test server which uses go-digits/login for Digits login
	s := NewService(testConsumerKey)
	// proxies all requests to the digits test server
	s.httpClient = digitsProxyClient
	ts := httptest.NewServer(s.LoginHandler(SuccessHandlerFunc(errorOnSuccess(t)), DefaultErrorHandler))

	// POST Digits OAuth Echo headers
	resp, _ := http.PostForm(ts.URL, url.Values{"wrongKeyName": {testAccountEndpoint}, "accountRequestHeader": {testAccountRequestHeader}})
	assertBodyString(t, resp.Body, ErrMissingAccountEndpoint.Error()+"\n")
	resp, _ = http.PostForm(ts.URL, url.Values{"accountEndpoint": {"https://evil.com"}, "accountRequestHeader": {testAccountRequestHeader}})
	assertBodyString(t, resp.Body, ErrInvalidDigitsEndpoint.Error()+"\n")
	resp, _ = http.PostForm(ts.URL, url.Values{"accountEndpoint": {testAccountEndpoint}, "accountRequestHeader": {`OAuth oauth_consumer_key="notmyconsumerkey",`}})
	assertBodyString(t, resp.Body, ErrInvalidConsumerKey.Error()+"\n")
	// valid, but incorrect Digits account endpoint
	resp, _ = http.PostForm(ts.URL, url.Values{"accountEndpoint": {"https://api.digits.com/1.1/wrong.json"}, "accountRequestHeader": {testAccountRequestHeader}})
	assertBodyString(t, resp.Body, ErrUnableToGetDigitsAccount.Error()+"\n")
}

func TestLoginHandlerFunc_unauthorized(t *testing.T) {
	digitsProxyClient, _, server := setupUnauthorizedDigitsTestServer()
	defer server.Close()

	// setup test server which uses go-digits/login for Digits login
	s := NewService(testConsumerKey)
	// proxies all requests to the digits test server
	s.httpClient = digitsProxyClient
	ts := httptest.NewServer(s.LoginHandler(SuccessHandlerFunc(errorOnSuccess(t)), DefaultErrorHandler))

	// POST Digits OAuth Echo headers
	resp, _ := http.PostForm(ts.URL, url.Values{"accountEndpoint": {testAccountEndpoint}, "accountRequestHeader": {testAccountRequestHeader}})
	assertBodyString(t, resp.Body, ErrUnableToGetDigitsAccount.Error()+"\n")
}

func TestLoginHandlerFunc_digitsAPIDown(t *testing.T) {
	// setup test server which uses go-digits/login for Digits login
	s := NewService(testConsumerKey)
	ts := httptest.NewServer(s.LoginHandler(SuccessHandlerFunc(errorOnSuccess(t)), DefaultErrorHandler))
	// POST Digits OAuth Echo headers
	resp, _ := http.PostForm(ts.URL, url.Values{"accountEndpoint": {testAccountEndpoint}, "accountRequestHeader": {testAccountRequestHeader}})
	assertBodyString(t, resp.Body, ErrUnableToGetDigitsAccount.Error()+"\n")
}

// success and failure handlers for testing

func successChecks(t *testing.T) func(w http.ResponseWriter, req *http.Request, account *digits.Account) {
	success := func(w http.ResponseWriter, req *http.Request, account *digits.Account) {
		if account.AccessToken.Token != "t" {
			t.Errorf("expected Token value t, got %q", account.AccessToken.Token)
		}
		if account.AccessToken.Secret != "s" {
			t.Errorf("expected Secret value s, got %q", account.AccessToken.Secret)
		}
		if account.PhoneNumber != "0123456789" {
			t.Errorf("expected PhoneNumber 0123456789, got %q", account.PhoneNumber)
		}
	}
	return success
}

func errorOnSuccess(t *testing.T) func(w http.ResponseWriter, req *http.Request, account *digits.Account) {
	success := func(w http.ResponseWriter, req *http.Request, account *digits.Account) {
		t.Errorf("unexpected call to success, %v", account)
	}
	return success
}

func errorOnFailure(t *testing.T) func(w http.ResponseWriter, err error, code int) {
	failure := func(w http.ResponseWriter, err error, code int) {
		t.Errorf("unexpected call to failure, %v %d", err, code)
	}
	return failure
}

// Testing Utils

func setupDigitsTestServer(jsonData string) (*http.Client, *http.ServeMux, *httptest.Server) {
	digitsProxyClient, digitsMux, digitsServer := testServer()
	digitsMux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, jsonData)
	})
	return digitsProxyClient, digitsMux, digitsServer
}

// fake Digits server will always return 401 Unauthorized
func setupUnauthorizedDigitsTestServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	digitsProxyClient, digitsMux, digitsServer := testServer()
	digitsMux.HandleFunc("/1.1/sdk/account.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
	})
	return digitsProxyClient, digitsMux, digitsServer
}

// testServer returns an http Client, ServeMux, and Server. The client proxies
// requests to the server and handlers can be registered on the mux to handle
// requests. The caller must close the test server.
func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

// RewriteTransport rewrites https requests to http to avoid TLS cert issues
// during testing.
type RewriteTransport struct {
	Transport http.RoundTripper
}

// RoundTrip rewrites the request scheme to http and calls through to the
// composed RoundTripper or if it is nil, to the http.DefaultTransport.
func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertBodyString(t *testing.T, rc io.ReadCloser, expected string) {
	defer rc.Close()
	if b, err := ioutil.ReadAll(rc); err == nil {
		if string(b) != expected {
			t.Errorf("expected %q, got %q", expected, string(b))
		}
	} else {
		t.Errorf("error reading Body")
	}
}
