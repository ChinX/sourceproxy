package oauth

import (
	"fmt"
	"net/url"
	"strings"
)

var brokenAuthHeaderProviders = []string{
	"https://accounts.google.com/",
	"https://api.codeswholesale.com/oauth/token",
	"https://api.dropbox.com/",
	"https://api.dropboxapi.com/",
	"https://api.instagram.com/",
	"https://api.netatmo.net/",
	"https://api.odnoklassniki.ru/",
	"https://api.pushbullet.com/",
	"https://api.soundcloud.com/",
	"https://api.twitch.tv/",
	"https://app.box.com/",
	"https://connect.stripe.com/",
	"https://graph.facebook.com", // see https://github.com/golang/oauth2/issues/214
	"https://login.microsoftonline.com/",
	"https://login.salesforce.com/",
	"https://login.windows.net",
	"https://oauth.sandbox.trainingpeaks.com/",
	"https://oauth.trainingpeaks.com/",
	"https://oauth.vk.com/",
	"https://openapi.baidu.com/",
	"https://slack.com/",
	"https://test-sandbox.auth.corp.google.com",
	"https://test.salesforce.com/",
	"https://user.gini.net/",
	"https://www.douban.com/",
	"https://www.googleapis.com/",
	"https://www.linkedin.com/",
	"https://www.strava.com/oauth/",
	"https://www.wunderlist.com/oauth/",
	"https://api.patreon.com/",
	"https://sandbox.codeswholesale.com/oauth/token",
	"https://api.sipgate.com/v1/authorization/oauth",
}

// brokenAuthHeaderDomains lists broken providers that issue dynamic endpoints.
var brokenAuthHeaderDomains = []string{
	".force.com",
	".myshopify.com",
	".okta.com",
	".oktapreview.com",
}

var providerMap = map[string]*Endpoint{}

func RegisterBrokenAuthHeaderProvider(tokenURL string) {
	brokenAuthHeaderProviders = append(brokenAuthHeaderProviders, tokenURL)
}

func providerAuthHeaderWorks(tokenURL string) bool {
	for _, s := range brokenAuthHeaderProviders {
		if strings.HasPrefix(tokenURL, s) {
			return false
		}
	}

	if u, err := url.Parse(tokenURL); err == nil {
		for _, s := range brokenAuthHeaderDomains {
			if strings.HasSuffix(u.Host, s) {
				return false
			}
		}
	}
	return true
}

func RegisterProviderEndpoint(provider string, endpoint *Endpoint) {
	if _, ok := providerMap[provider]; ok {
		fmt.Printf("provider '%s' is already exists\n", provider)
	}
	providerMap[provider] = endpoint
}

func ProviderEndpoint(provider string) (*Endpoint, error) {
	endpoint, ok := providerMap[provider]
	if !ok {
		return nil, fmt.Errorf("provider '%s' is not set", provider)
	}
	return endpoint, nil
}
