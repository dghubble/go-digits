
# Digits Login

Twitter's [Digits](http://get.digits.com/) service provides phone number login via SMS confirmation codes for any web app or mobile app.

**Warning**: Digits login has moved to github.com/dghubble/gologin/digits.

## Digits for Web

The Digits Javascipt snippet launches a web popup for user phone number and confirmation code entry. OAuth Echo headers [source](https://dev.twitter.com/twitter-kit/web/digits) are returned to the page.

To use Digits to authenticate to your own backend, these headers should be posted to your server, validated, and used to fetch the user's Digits account and access token. That's where `go-digits` `login` comes in.

## Web App with Digits Login

Package `login` provides a `WebHandler` which receives POSTed OAuth Echo headers, validates them, and fetches the `Digits.Account`. Handling is then delegated to your own `SuccessHandler` or `ErrorHandler`. Typical behavior is to issue a web session when successful or use the `DefaultErrorHandler` for failures.

[app.go](app.go) shows an example web app which issues a client-side cookie sessions and persists no data.

### Getting Started

Get the `login` package, the examples, and their dependencies.

    go get github.com/dghubble/go-digits/login
    cd $GOPATH/src/github.com/dghubble/go-digits/examples/login
    go get .

Create a Digits application to obtain your consumer key/secret. Paste in your **Consumer Key** for `YOUR_DIGITS_CONSUMER_KEY` in `app.go` and `home.html`.

Note: Currently a Digits application must be created by making a dummy iOS or Android app via the Fabric [iOS Mac App](https://fabric.io/downloads/xcode) or [Android Studio Plugin](https://fabric.io/downloads).

Compile and run the app from the `examples/login` directory:

    go run app.go

### Web Login in Action

1. Clicking the "Login with Digits" button launches the phone number login popup.

![Phone Number Login](http://storage.googleapis.com/dghubble/digits-phone-number.png)

2. User enters a phone number and received SMS confirmation number. The Javascript snippet recevies OAuth Echo headers and POSTs them to the Go server.

3. If valid, the Digits `Account` is provided to the `SuccessHandler`. A cookie session is issued so the user is considered logged in.

