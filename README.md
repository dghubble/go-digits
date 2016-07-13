
# go-digits [![Build Status](https://travis-ci.org/dghubble/go-digits.png)](https://travis-ci.org/dghubble/go-digits) [![GoDoc](https://godoc.org/github.com/dghubble/go-digits?status.png)](https://godoc.org/github.com/dghubble/go-digits)
<img align="right" src="https://storage.googleapis.com/dghubble/digits-gopher.png">

go-digits is a Go client library for the [Digits](https://get.digits.com/) API. Check the [usage](#usage) section or the [examples](examples) to learn how to access the Digits API.

### Features

* AccountService for getting Digits accounts
* Get verified user phone numbers and email addresses
* ContactService for finding matching contacts ("friends")
* Digits API Client accepts any OAuth1 `http.Client`

## Install

    go get github.com/dghubble/go-digits/digits

## Docs

Read [GoDoc](https://godoc.org/github.com/dghubble/go-digits/digits)

## Usage

The `digits` package provides a `Client` for accessing the Digits API. Here is an example request for a Digit user's Account.

```go
import (
    "github.com/dghubble/go-digits/digits"
    "github.com/dghubble/oauth1"
)

config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
// OAuth1 http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)

// Digits client
client := digits.NewClient(httpClient)
// get current user's Digits Account
account, resp, err := client.Accounts.Account()
```

### Authentication

The API client accepts any `http.Client` capable of signing OAuth1 requests to handle authorization. See the OAuth1 package [dghubble/oauth1](https://github.com/dghubble/oauth1) for details and examples.

To implement Login with Digits for web or mobile, see the gologin [package](https://github.com/dghubble/gologin) and [examples](https://github.com/dghubble/gologin/tree/master/examples/digits).

## Contributing

See the [Contributing Guide](https://gist.github.com/dghubble/be682c123727f70bcfe7).

## License

[MIT License](LICENSE)


