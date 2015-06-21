/*
Package digits provides a client and models for the Digits API.

The digits package provides a client and models for the Digits API. By design,
the package is decoupled from OAuth concerns. An http.Client which
transparently handles OAuth1 request signing should be passed to create a
Digits Client. See https://github.com/dghubble/oauth1.

	import "github.com/dghubble/go-digits/digits"
	import "github.com/dghubble/oauth1"

	authConfig := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessTokenSecret")
	// http.Client will automatically authorize Requests
	httpClient := authConfig.Client(token)

	// digits client
	client := digits.NewClient(httpClient)
	// get the current user's Digits Account
	account, resp, err := client.Accounts.Account()

*/
package digits
