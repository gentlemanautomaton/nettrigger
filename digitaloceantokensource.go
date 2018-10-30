package nettrigger

import "golang.org/x/oauth2"

type digitalOceanTokenSource struct {
	AccessToken string
}

func (t *digitalOceanTokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
