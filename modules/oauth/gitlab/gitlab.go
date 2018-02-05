package gitlab

import "github.com/chinx/sourceproxy/modules/oauth"

const GITLAB = "gitlab"

var endpoint = &oauth.Endpoint{
	AuthURL:  "https://gitlab.com/oauth/authorize",
	TokenURL: "https://gitlab.com/oauth/token",
}

func init()  {
	oauth.RegisterProviderEndpoint(GITLAB, endpoint)
}
