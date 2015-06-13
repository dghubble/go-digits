
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

## Docs

Read [GoDoc](https://godoc.org/github.com/dghubble/go-digits)

### Login with Digits

1. Add a "Use My Phone Number" button to your login page and follow the [Digits for Web instructions](https://dev.twitter.com/twitter-kit/web/digits) to add the Digits Javascript snippet (example)[examples/login/home.html].

2. Add the go-digits imports

```go
import (
    "github.com/dghubble/go-digits/digits"
    "github.com/dghubble/go-digits/login"
)
```

3. Create a `LoginService` struct with your Digits Consumer Key.

```go
var digitsService = login.NewLoginService("digitsConsumerKey")
```

4. Register the LoginHandler to receive POST's from your login page.

```go
http.Handle("/digits_login", digitsService.LoginHandlerFunc(successHandler, dgtsLogin.ErrorHandler))
```

5. Receive the validated `Digits.Account` in the `successHandler`, which you may implement to issue any sort of session you prefer.

```
func successHandler(w http.ResponseWriter, req *http.Request, account *digits.Account) {
    session := sessionStore.New(sessionName)
    session.Values["digitsID"] = account.ID
    session.Save(w)
    http.Redirect(w, req, "/profile", http.StatusFound)
}
```

#### Examples

A [minimal web app](examples/login) is included which lets you paste in your Digits Consumer Key and run to see a working SMS phone number login site.

    cd examples/login
    # fill in YOUR_DIGITS_CONSUMER_KEY in app.go and home.html
    go run app.go

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


