package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-digits/digits"
	"github.com/dghubble/oauth1"
)

func main() {
	flags := flag.NewFlagSet("digits-example", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Digits Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Digits Consumer Secret")
	accessToken := flags.String("access-token", "", "Digits Access Token")
	accessSecret := flags.String("access-secret", "", "Digits Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "DIGITS")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Digits client
	client := digits.NewClient(httpClient)

	// get current user's Digits Account
	account, _, err := client.Accounts.Account()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ACCOUNT:\n%+v\n", account)
	}

	// get Digits users who have signed up for the Digits App and are known to
	// the current user
	matchParams := &digits.MatchesParams{Count: 20}
	contacts, _, _ := client.Contacts.Matches(matchParams)
	fmt.Printf("CONTACT MATCHES:\n%+v\n", contacts)
}
