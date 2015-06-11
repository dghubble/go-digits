
# go-digits [![Build Status](https://travis-ci.org/dghubble/go-digits.png)](https://travis-ci.org/dghubble/go-digits) [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-digits/login)](http://gocover.io/github.com/dghubble/go-digits/login) [![GoDoc](http://godoc.org/github.com/dghubble/go-digits?status.png)](http://godoc.org/github.com/dghubble/go-digits)

go-digits provides Twitter Digits Go packages for implementing Login with Digits and making Digits API requests from your servers.

### Features

* Package `login` provides a login service for adding Login with Digits to an auth system
    * Register a `LoginHandlerFunc` on your `ServeMux` to handle Digits login validation and Digits account retrieval.
    * Use any session library you prefer to issue a user session upon success
* Package `digits` provides a client to the Digits API.
    * Get Digits accounts

## Install

    go get github.com/dghubble/go-digits

## Documentation

Read [GoDoc](https://godoc.org/github.com/dghubble/go-digits)

### Digits API

The `digits` package provides a client and models for the Digits API. By design, it is decoupled from OAuth concerns. An `http.Client` which transparently handles OAuth1 request signing should be passed to create a Digits Client. See [dghubble/oauth1](https://github.com/dghubble/oauth1).

```go
import "github.com/dghubble/go-digits/digits"
import "github.com/dghubble/oauth1"

authConfig := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessTokenSecret")
// http.Client will automatically authorize Requests
httpClient := authConfig.Client(token)

// digits client
client := digits.NewClient(authClient)
// get the current user's Digits Account
account, resp, err := client.Accounts.Account()
```

## License

[MIT License](LICENSE)


