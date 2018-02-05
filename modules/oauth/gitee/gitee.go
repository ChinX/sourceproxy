package gitee

import "github.com/chinx/sourceproxy/modules/oauth"

const GITEE = "gitee"

var endpoint = &oauth.Endpoint{
	AuthURL:  "https://gitee.com/oauth/authorize",
	TokenURL: "https://gitee.com/oauth/token",
}

func init()  {
	oauth.RegisterProviderEndpoint(GITEE, endpoint)
	oauth.RegisterBrokenAuthHeaderProvider(endpoint.TokenURL)
}

