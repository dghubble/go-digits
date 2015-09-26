
# go-digits [![Build Status](https://travis-ci.org/dghubble/go-digits.png)](https://travis-ci.org/dghubble/go-digits) [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-digits/digits)](http://gocover.io/github.com/dghubble/go-digits/digits) [![GoDoc](http://godoc.org/github.com/dghubble/go-digits?status.png)](http://godoc.org/github.com/dghubble/go-digits)
<img align="right" src="http://storage.googleapis.com/dghubble/digits-gopher.png">

go-digits is a Go client library for accessing Twitter's [Digits](https://get.digits.com/) API.

If you're trying to implement Login with Digits for web or mobile, see [gologin](https://github.com/dghubble/gologin).

### Features

* Provides a client to the Digits API
* AccountService allows Digits accounts to be retrieved

## Install

    go get github.com/dghubble/go-digits/digits

## Docs

Read [GoDoc](https://godoc.org/github.com/dghubble/go-digits)

## Usage

The `digits` package provides a `Client` and models for the Digits API. By design, it is decoupled from OAuth1 concerns. An `http.Client` which transparently handles OAuth1 request signing should be used to create a Digits Client.

```go
import (
    "github.com/dghubble/go-digits/digits"
    "github.com/dghubble/oauth1"
)

config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
// OAuth1 http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)

// digits client
client := digits.NewClient(httpClient)
// get the current user's Digits Account
account, resp, err := client.Accounts.Account()
```

See [dghubble/oauth1](https://github.com/dghubble/oauth1) for details about OAuth1.

## Contributing

See the [Contributing Guide](https://gist.github.com/dghubble/be682c123727f70bcfe7).

## License

[MIT License](LICENSE)


