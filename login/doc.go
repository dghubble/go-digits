/*
Package login provides a login service for implementing Login with Digits.

Package login provides a configurable Digits login service and login handler.
The login handler receives POST'ed OAuth Echo headers, uses them to get a
validated Digits.Account, and calls a custom success handler (or error
handler) to issue a session.

Get started with the 100 line web app example https://github.com/dghubble/go-digits/tree/master/examples/login
Paste in your Digits consumer key and run it to see login by phone number in
action.

To add Login with Digits to your existing web app:

1. Follow the Digits for Web instructions to add a "Use My Phone Number"
button and Digits JS snippet to your login page. https://dev.twitter.com/twitter-kit/web/digits


2. Add the go-digits imports

	import (
	    "github.com/dghubble/go-digits/digits"
	    "github.com/dghubble/go-digits/login"
	)

3. Create a LoginService struct with your Digits Consumer Key.

	var dgts = login.NewLoginService("digitsConsumerKey")

4. Register a LoginHandler to receive POST's from your login page.

	http.Handle("/digits_login", dgts.LoginHandlerFunc(successHandler,
	    login.ErrorHandler))

5. Receive the validated Digits.Account in a successHandler and issue any kind of session you prefer.

	func successHandler(w http.ResponseWriter, r *http.Request, account *digits.Account) {
	    session := sessionStore.New(sessionName)
	    session.Values["digitsID"] = account.ID
	    session.Save(w)
	    http.Redirect(w, r, "/profile", http.StatusFound)
	}
*/
package login
