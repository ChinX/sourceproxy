package oauth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/url"
	"strconv"
	"strings"

	"github.com/chinx/sourceproxy/utils"
)

type Token struct {
	AccessToken  string      `json:"access_token"`
	TokenType    string      `json:"token_type,omitempty"`
	RefreshToken string      `json:"refresh_token,omitempty"`
	ExpiresIn    int         `json:"expires_in,omitempty"`
	Raw          interface{} `json:"-"`
}

type tokenJSON struct {
	AccessToken  string         `json:"access_token"`
	TokenType    string         `json:"token_type"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresIn    expirationTime `json:"expires_in"` // at least PayPal returns string, while most return number
	Expires      expirationTime `json:"expires"`    // broken Facebook spelling of expires_in
}

func (e *tokenJSON) expiresIn() int {
	if v := e.ExpiresIn; v != 0 {
		return int(e.ExpiresIn)
	}
	return int(e.Expires)
}

type expirationTime int32

func (e *expirationTime) UnmarshalJSON(b []byte) error {
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}
	i, err := n.Int64()
	if err != nil {
		return err
	}
	*e = expirationTime(i)
	return nil
}

func RetrieveToken(clientID, clientSecret, tokenURL string, v url.Values) (*Token, error) {
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	bustedAuth := !providerAuthHeaderWorks(tokenURL)
	if bustedAuth {
		if clientID != "" {
			v.Set("client_id", clientID)
		}
		if clientSecret != "" {
			v.Set("client_secret", clientSecret)
		}
	} else {
		header["Authorization"] = "Basic " + basicAuth(url.QueryEscape(clientID), url.QueryEscape(clientSecret))
	}

	r, err := utils.HttpPost(tokenURL, header, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}

	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v\nResponse: %s", r.Status, body)
	}

	var token *Token
	content, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	switch content {
	case "application/x-www-form-urlencoded", "text/plain":
		vals, err := url.ParseQuery(string(body))
		if err != nil {
			return nil, err
		}
		token = &Token{
			AccessToken:  vals.Get("access_token"),
			TokenType:    vals.Get("token_type"),
			RefreshToken: vals.Get("refresh_token"),
			Raw:          vals,
		}
		e := vals.Get("expires_in")
		if e == "" {
			// TODO(jbd): Facebook's OAuth2 implementation is broken and
			// returns expires_in field in expires. Remove the fallback to expires,
			// when Facebook fixes their implementation.
			e = vals.Get("expires")
		}
		expires, _ := strconv.Atoi(e)
		if expires != 0 {
			token.ExpiresIn = expires
		}
	default:
		var tj tokenJSON
		if err = json.Unmarshal(body, &tj); err != nil {
			return nil, err
		}
		token = &Token{
			AccessToken:  tj.AccessToken,
			TokenType:    tj.TokenType,
			RefreshToken: tj.RefreshToken,
			ExpiresIn:    tj.expiresIn(),
			Raw:          make(map[string]interface{}),
		}
		json.Unmarshal(body, &token.Raw) // no error checks for optional fields
	}
	// Don't overwrite `RefreshToken` with an empty value
	// if this was a token refreshing request.
	if token.RefreshToken == "" {
		token.RefreshToken = v.Get("refresh_token")
	}
	return token, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
