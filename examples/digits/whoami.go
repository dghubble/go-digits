package main

import (
	"fmt"
	"github.com/dghubble/go-digits/digits"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type config struct {
	ConsumerKey       string `envconfig:"CONSUMER_KEY"`
	ConsumerSecret    string `envconfig:"CONSUMER_SECRET"`
	AccessToken       string `envconfig:"ACCESS_TOKEN"`
	AccessTokenSecret string `envconfig:"ACCESS_TOKEN_SECRET"`
}

// Main requests an Account as a consumer on behalf of a Digits user.
func main() {
	var c config
	err := envconfig.Process("DIGITS", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: remove empty string check once required tag is available
	// https://github.com/kelseyhightower/envconfig/pull/19
	if c.ConsumerKey == "" || c.ConsumerSecret == "" || c.AccessToken == "" || c.AccessTokenSecret == "" {
		log.Fatal("Missing required environment variable")
	}

	authConfig := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessTokenSecret)
	// http.Client which will automatically authorize Request
	httpClient := authConfig.Client(token)

	// digits client
	client := digits.NewClient(httpClient)
	account, _, _ := client.Accounts.Account()
	fmt.Printf("Digits ACCOUNT:\n%+v\n", account)
}
