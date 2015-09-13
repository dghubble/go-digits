
# go-digits [![Build Status](https://travis-ci.org/dghubble/go-digits.png)](https://travis-ci.org/dghubble/go-digits) [![GoDoc](http://godoc.org/github.com/dghubble/go-digits?status.png)](http://godoc.org/github.com/dghubble/go-digits)
<img align="right" src="http://storage.googleapis.com/dghubble/digits-gopher.png">

go-digits provides unofficial Go packages for (Twitter) Digits. Add Digits (phone number login to your Go server or make Accounts API requests.

### Packages

#### login [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-digits/login)](http://gocover.io/github.com/dghubble/go-digits/login) [![GoDoc](http://godoc.org/github.com/dghubble/go-digits/login?status.png)](http://godoc.org/github.com/dghubble/go-digits/login)

* Provides a login handler for adding Digits phone number login to web apps
* Register a `WebHandler` on your `ServeMux` to handle Digits web logins
* Register a `TokenHandler` on your `ServeMux` to handle Digits token logins
* Works with any session library you prefer. No context dependencies.

#### digits [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-digits/digits)](http://gocover.io/github.com/dghubble/go-digits/digits) [![GoDoc](http://godoc.org/github.com/dghubble/go-digits/digits?status.png)](http://godoc.org/github.com/dghubble/go-digits/digits)

* Provides a client to the Digits API.
* AccountService allows Digits accounts to be retrieved

## Install

    go get github.com/dghubble/go-digits/digits
    go get github.com/dghubble/go-digits/login

## Docs

Read [GoDoc](https://godoc.org/github.com/dghubble/go-digits)

### Phone Numer Login via Digits Web

Get started with the [example app](examples/login). Paste in your Digits consumer key and run it locally to see phone number login in action.

Alternately, add Login with Digits to your existing web app:

1. Follow the [Digits for Web instructions](https://dev.twitter.com/twitter-kit/web/digits) to add a "Use My Phone Number" button and Digits JS snippet to your login page.
2. Add the go-digits imports
    
    ```go
    import (
        "github.com/dghubble/go-digits/digits"
        "github.com/dghubble/go-digits/login"
    )
    ```

3. Register a `WebHandler` to receive POST's from your login page.

    ```go
    handlerConfig := login.Config{
        ConsumerKey: "YOUR_DIGITS_CONSUMER_KEY",
        Success: login.SuccessHandlerFunc(issueWebSession),
        Failure: login.DefaultErrorHandler,
    }
    http.Handle("/digits_login", login.NewWebHandler(handlerConfig))
    ```

4. Receive the validated `Digits.Account` in a `SuccessHandler`. Issue a session your backend supports.

    ```
    func issueWebSession(w http.ResponseWriter, r *http.Request, account *digits.Account) {
        session := sessionStore.New(sessionName)
        session.Values["digitsID"] = account.ID
        session.Values["phoneNumer"] = account.PhoneNumber
        session.Save(w)
        http.Redirect(w, r, "/profile", http.StatusFound)
    }
    ```

### Digits API

The `digits` package provides a client and models for the Digits API. By design, it is decoupled from OAuth concerns. An `http.Client` which transparently handles OAuth1 request signing should be passed to create a Digits Client. See [dghubble/oauth1](https://github.com/dghubble/oauth1).

```go
import "github.com/dghubble/go-digits/digits"
import "github.com/dghubble/oauth1"

config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessTokenSecret")
// OAuth1 http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)

// digits client
client := digits.NewClient(httpClient)
// get the current user's Digits Account
account, resp, err := client.Accounts.Account()
```

## Roadmap

* Configurable POST field names

## License

[MIT License](LICENSE)


