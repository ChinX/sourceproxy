package oauth

import (
	"bytes"
	"net/url"
	"strings"
)

type Endpoint struct {
	AuthURL  string
	TokenURL string
}

type Options struct {
	*Config      		`json:",inline"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Code         string `json:"code"`
	State        string `json:"state"`
}

type Config struct {
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Endpoint     *Endpoint `json:"-"`
	RedirectURL  string    `json:"redirect_uri"`
	Scope        string    `json:"scope"`
}

func (c *Config) AuthCodeURL(state string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.AuthURL)
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientID},
		"redirect_uri":  CondVal(c.RedirectURL),
		"scope":         CondVal(c.Scope),
		"state":         CondVal(state),
	}
	if strings.Contains(c.Endpoint.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

func (c *Config) CredentialsToken(username, password string) (*Token, error) {
	return RetrieveToken(c.ClientID, c.ClientSecret, c.Endpoint.TokenURL, url.Values{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
		"scope":      CondVal(c.Scope),
	})
}

func (c *Config) AccessToken(code string) (*Token, error) {
	return RetrieveToken(c.ClientID, c.ClientSecret, c.Endpoint.TokenURL, url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": CondVal(c.RedirectURL),
	})
}

func (c *Config) RefreshAccessToken(refreshToken string) (*Token, error) {
	return RetrieveToken(c.ClientID, c.ClientSecret, c.Endpoint.TokenURL, url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	})
}

func CondVal(v string) []string {
	if v == "" {
		return nil
	}
	return []string{v}
}
