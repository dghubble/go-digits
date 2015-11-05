
# Examples

## Digits API

A user grants a Digits consumer application access to his/her Digits resources (phone number, contacts). As a result, the application receives an OAuth1 access token and secret.

To make requests as an application, on behalf of the user, a client requires the application consumer key and secret and the user access token and secret.

    export DIGITS_CONSUMER_KEY=xxx
    export DIGITS_CONSUMER_SECRET=xxx
    export DIGITS_ACCESS_TOKEN=xxx
    export DIGITS_ACCESS_SECRET=xxx

Get the dependencies for the examples

    cd examples
    go get .

## Accounts API

Get the current user's Digits `Account`. Compile and run

    go run accounts.go
