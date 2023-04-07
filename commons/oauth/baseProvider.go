package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

type BaseProvider struct {
	Ctx                 context.Context
	Scopes              []string
	ClientId            string
	ClientSecret        string
	RedirectUrl         string
	RedirectLoginUrl    string
	RedirectRegisterUrl string
	AuthUrl             string
	TokenUrl            string
	UserApiUrl          string
}

// Client implements Provider.Client() interface method.
func (p *BaseProvider) Client(token *oauth2.Token) *http.Client {
	return p.oauth2Config().Client(p.Ctx, token)
}

func (p *BaseProvider) sendRawUserDataRequest(req *http.Request, token *oauth2.Token) ([]byte, error) {
	client := p.Client(token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// http.Client.Get doesn't treat non 2xx responses as error
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf(
			"failed to fetch OAuth2 user profile via %s (%d):\n%s",
			p.UserApiUrl,
			res.StatusCode,
			string(result),
		)
	}

	return result, nil
}

func (p *BaseProvider) FetchRawUserData(token *oauth2.Token) ([]byte, error) {
	req, err := http.NewRequestWithContext(p.Ctx, "GET", p.UserApiUrl, nil)
	if err != nil {
		return nil, err
	}

	return p.sendRawUserDataRequest(req, token)
}

func (p *BaseProvider) BuildLoginUrl(state string, opts ...oauth2.AuthCodeOption) string {
	p.SetRedirectUrl(p.RedirectLoginUrl)
	return p.oauth2Config().AuthCodeURL(state, opts...)
}
func (p *BaseProvider) BuildRegisterUrl(state string, opts ...oauth2.AuthCodeOption) string {
	p.SetRedirectUrl(p.RedirectRegisterUrl)
	return p.oauth2Config().AuthCodeURL(state, opts...)
}

func (p *BaseProvider) SetRedirectUrl(url string) {
	p.RedirectUrl = url
}

func (p *BaseProvider) ExchangeCode(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.oauth2Config().Exchange(p.Ctx, code, opts...)
}

func (p *BaseProvider) oauth2Config() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  p.RedirectUrl,
		ClientID:     p.ClientId,
		ClientSecret: p.ClientSecret,
		Scopes:       p.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  p.AuthUrl,
			TokenURL: p.TokenUrl,
		},
	}
}
