package github

import "github.com/chinx/sourceproxy/modules/oauth"

const GITHUB = "github"

var endpoint = &oauth.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

func init()  {
	oauth.RegisterProviderEndpoint(GITHUB, endpoint)
}