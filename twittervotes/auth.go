package main

import (
	"log"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/joeshaw/envdecode"
)

// TwitterAccess modelates the fields to the twitter access.
type TwitterAccess struct {
	ConsumerKey    string `env:"SP_TWITTER_KEY,required"`
	ConsumerSecret string `env:"SP_TWITTER_SECRET,required"`
	AccessToken    string `env:"SP_TWITTER_ACCESSTOKEN,required"`
	AccessSecret   string `env:"SP_TWITTER_ACESSSECRET,required"`
}

// Credentials
var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func setupTwitterAuth() {
	var ta TwitterAccess

	if err := envdecode.Decode(&ta); err != nil {
		log.Fatalln(err)
	}

	creds = &oauth.Credentials{
		Token:  ta.AccessToken,
		Secret: ta.AccessSecret,
	}

	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  ta.ConsumerKey,
			Secret: ta.ConsumerSecret,
		},
	}
}
