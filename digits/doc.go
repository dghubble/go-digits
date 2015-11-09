/*
Package digits provides a Client for the Digits API.

The digits package provides a Client for accessing Digits API services. Here
is an example request for a Digit user's Account.

	import (
		"github.com/dghubble/go-digits/digits"
		"github.com/dghubble/oauth1"
	)

	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessTokenSecret")
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(token)

	// Digits client
	client := digits.NewClient(httpClient)
	// get current user's Digits Account
	account, resp, err := client.Accounts.Account()

The API client accepts any http.Client capable of signing OAuth1 requests to
handle authorization.

See the OAuth1 package https://github.com/dghubble/oauth1 for authorization
details and examples.

To implement Login with Digits, see https://github.com/dghubble/gologin.
*/
package digits
