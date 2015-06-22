// Package login handles Digits token based logins, typically for mobile clients.
package login

import (
	"net/http"

	"github.com/dghubble/go-digits/digits"
)

// AuthClientSource is an interface for sources of oauth1 token authorized
// http.Client's. This interface avoids a hard dependency on a particular
// oauth1 implementation.
type AuthClientSource interface {
	GetClient(token, tokenSecret string) *http.Client
}

// TokenHandlerConfig configures a TokenHandler.
type TokenHandlerConfig struct {
	AuthConfig AuthClientSource
	Success    SuccessHandler
	Failure    ErrorHandler
}

// TokenHandler receives a POSTed Digits token/secret and fetches the Digits
// Account. If successful, handling is delegated to the SuccessHandler.
// Otherwise, the ErrorHandler is called.
type TokenHandler struct {
	authConfig AuthClientSource
	success    SuccessHandler
	failure    ErrorHandler
}

// NewTokenHandler returns a new TokenHandler.
func NewTokenHandler(config TokenHandlerConfig) *TokenHandler {
	return &TokenHandler{
		authConfig: config.AuthConfig,
		success:    config.Success,
		failure:    config.Failure,
	}
}

// ServeHTTP receives a POSTed Digits token/secret and fetches the Digits
// Account. If successful, handling is delegated to the SuccessHandler.
// Otherwise, the ErrorHandler is called.
func (h *TokenHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	accessToken := req.PostForm.Get("digitsAccessToken")
	accessTokenSecret := req.PostForm.Get("digitsAccessTokenSecret")
	httpClient := h.authConfig.GetClient(accessToken, accessTokenSecret)
	digitsClient := digits.NewClient(httpClient)

	// fetch Digits Account
	account, resp, err := digitsClient.Accounts.Account()
	err = ValidateAccountResponse(account, resp, err)
	if err != nil {
		h.failure.ServeHTTP(w, err, http.StatusBadRequest)
		return
	}
	h.success.ServeHTTP(w, req, account)
}
