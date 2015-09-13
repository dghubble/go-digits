
# Digits API

## Account API

A Digits user grants a consumer application access to his/her resources (phone number, contacts, etc.) and an application receives an `Account.AccessToken` (an OAuth1 access token).

Get the `digits` package, the examples, and their dependencies.

    go get github.com/dghubble/go-digits/digits
    cd $GOPATH/src/github.com/dghubble/go-digits/examples/digits
    go get .

To demonstrate making a request as the application, on behalf of the user, set the consumer and access tokens and secrets as environment variables.

    export DIGITS_CONSUMER_KEY=xxx
    export DIGITS_CONSUMER_SECRET=xxx
    export DIGITS_ACCESS_TOKEN=xxx
    export DIGITS_ACCESS_TOKEN_SECRET=xxx

Compile and run

    go run requests.go

to show account details of the user.
