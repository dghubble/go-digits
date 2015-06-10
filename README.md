
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

## License

[MIT License](LICENSE)


