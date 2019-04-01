package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/PeterBooker/colligo/internal/config"
	"github.com/PeterBooker/colligo/internal/store"
	"github.com/mrjones/oauth"
)

var (
	twitterRequestTokenURL   = "https://api.twitter.com/oauth/request_token"
	twitterAuthorizeTokenURL = "https://api.twitter.com/oauth/authorize"
	twitterAccessTokenURL    = "https://api.twitter.com/oauth/access_token"
	twitterCallbackURL       = "https://colligo.dev/api/v1/callback/twitter"
)

// TwitterGetToken ...
func TwitterGetToken(c config.Config, s *store.Store, w http.ResponseWriter, r *http.Request) error {

	consumer := oauth.NewConsumer(
		c.Twitter.Key,
		c.Twitter.Secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   twitterRequestTokenURL,
			AuthorizeTokenUrl: twitterAuthorizeTokenURL,
			AccessTokenUrl:    twitterAccessTokenURL,
		},
	)

	values := r.URL.Query()
	verifier := values.Get("oauth_verifier")
	token := values.Get("oauth_token")

	visitor, found := s.Get(token)
	if !found {
		return errors.New("Token not found")
	}

	accessToken, err := consumer.AuthorizeToken(visitor.Token, verifier)
	if err != nil {
		return err
	}

	log.Printf("Access Token: %s\n", accessToken)

	return nil

}

// TwitterSendUser ...
func TwitterSendUser(c config.Config, s *store.Store, w http.ResponseWriter, r *http.Request) (string, error) {

	values := r.URL.Query()
	cb := values.Get("callback")

	if cb == "" {
		return "", errors.New("You must provide a callback URL")
	}

	consumer := oauth.NewConsumer(
		c.Twitter.Key,
		c.Twitter.Secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   twitterRequestTokenURL,
			AuthorizeTokenUrl: twitterAuthorizeTokenURL,
			AccessTokenUrl:    twitterAccessTokenURL,
		},
	)

	requestToken, loginURL, err := consumer.GetRequestTokenAndUrl(twitterCallbackURL)
	if err != nil {
		return "", err
	}

	log.Printf("Token: %s, loginURL: %s\n", requestToken, loginURL)

	v := &store.Visitor{
		Token:    requestToken,
		Callback: cb,
	}

	s.Set(requestToken.Token, v)

	return loginURL, nil

}
