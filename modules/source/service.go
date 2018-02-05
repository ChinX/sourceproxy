package source

import (
	"fmt"
	"log"
)

var proxyServers = map[string]ProxyServer{}

type ProxyServer interface {
	User(token string) (int, []byte)
	Repos(token string) (int, []byte)
	Branches(owner, repo, token string) (int, []byte)
	Commits(owner, repo, branch, token string) (int, []byte)
}

func Registry(provider string, svr ProxyServer) {
	opt := "SourceProxy.Registry"
	if svr == nil {
		log.Println(opt, "ProxyServer must not nil")
	}

	if _, exist := proxyServers[provider]; exist {
		log.Println(opt, "ProxyServer \"%s\" is already exists")
	}

	proxyServers[provider] = svr
}

func NewServer(provider string) (ProxyServer, error) {
	opt := "SourceProxy.Registry"

	svr, exist := proxyServers[provider]
	if !exist {
		return nil, fmt.Errorf(opt+" ProxyServer \"%s\" is not set", provider)
	}

	return svr, nil
}
