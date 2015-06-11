
# Examples

## Account API

A Digits user grants a consumer application access to its resources (phone number, contacts, etc.) with an OAuth1 access token.

To demonstrate making a request as the application, on behalf of the user, set the consumer and access tokens and secrets as environment variables.

    export DIGITS_CONSUMER_KEY=xxx
    export DIGITS_CONSUMER_SECRET=xxx
    export DIGITS_ACCESS_TOKEN=xxx
    export DIGITS_ACCESS_TOKEN_SECRET=xxx

Run

    cd examples
    go run whoami.go

to show account details of the user.

## Web App with Digits Login

TODO
