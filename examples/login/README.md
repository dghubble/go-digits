
# Digits Login

Digits provides phone number login via SMS confirmation codes for any web app or mobile app. The Digits javascipt snippet launches a web popup which accepts a phone number and confirmation code and returns OAuth Echo headers [source](https://dev.twitter.com/twitter-kit/web/digits). These headers **must** be validated on your Go server before issuing a web session.

## Web App with Digits Login

Package `login` provides a handler which receives POST'ed OAuth Echo headers, uses them to get a validated `Digits.Account`, and calls your success handler (or error handler) to issue a session.

[app.go](app.go) shows a tiny web app which issues a client-side cookie session upon successful Digits login. It does not write the Digits `Account` details to a database, though it could do so in the success handler.

## Getting Started

Get the `login` package, the examples, and their dependencies.

    go get github.com/dghubble/go-digits/login
    cd $GOPATH/src/github.com/dghubble/go-digits/examples/login
    go get .

Create a Digits application to obtain your consumer key/secret. Currently this must be done by creating a dummy iOS or Android app via the Fabric [iOS Mac App](https://fabric.io/downloads/xcode) or [Android Studio Plugin](https://fabric.io/downloads). Once you've obtained a Digits consumer key, paste it in for `YOUR_DIGITS_CONSUMER_KEY` in `app.go` and `home.html`.

Compile and run the app:

    go run app.go

## Results

Clicking the "Login with Digits" button should launch the phone number login flow.

![Phone Number Login](http://storage.googleapis.com/dghubble/digits-phone-number.png)

After the user enters a phone number and enters the SMS confirmation number, the Javascript snippet POST's the OAuth Echo data to the tiny Go server.

If the echo data validates correctly, a Digits `Account` is provided to the successHandler, which issues a cookie session and redirects to a page that is only visible to logged in users.


