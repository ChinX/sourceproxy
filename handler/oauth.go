package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/chinx/cobweb"
	"github.com/chinx/sourceproxy/modules/oauth"
)

func OAuthURL(rw http.ResponseWriter, req *http.Request, params cobweb.Params) {
	opt, err := parseOAuthConfig(req, params)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	addr := opt.AuthCodeURL(opt.State)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(addr))
}

func AccessToken(rw http.ResponseWriter, req *http.Request, params cobweb.Params) {
	opt, err := parseOAuthConfig(req, params)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	var token *oauth.Token
	if opt.Code != "" {
		token, err = opt.AccessToken(opt.Code)
	} else if opt.RefreshToken != "" {
		token, err = opt.RefreshAccessToken(opt.RefreshToken)
	} else if opt.Username != "" && opt.Password != "" {
		token, err = opt.CredentialsToken(opt.Username, opt.Password)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Lack of necessary parameters"))
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	body, err := json.Marshal(token)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}

func parseOAuthConfig(req *http.Request, params cobweb.Params) (*oauth.Options, error) {
	provider := params.Get("provider")
	endpoint, err := oauth.ProviderEndpoint(provider)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	opt := &oauth.Options{Config: &oauth.Config{Endpoint: endpoint}}
	return opt, json.Unmarshal(body, opt)
}
