package handler

import (
	"net/http"

	"github.com/chinx/sourceproxy/modules/source"
	"github.com/chinx/cobweb"
)

func SourceUser(rw http.ResponseWriter, req *http.Request, params cobweb.Params) {
	provider := req.Header.Get("Provider")
	token := req.Header.Get("Access-Token")
	svc, err := source.NewServer(provider)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	status, result := svc.User(token)
	rw.WriteHeader(status)
	rw.Write(result)
}

func SourceRepos(rw http.ResponseWriter, req *http.Request, params cobweb.Params) {
	provider := req.Header.Get("Provider")
	token := req.Header.Get("Access-Token")
	svc, err := source.NewServer(provider)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	status, result := svc.Repos(token)
	rw.WriteHeader(status)
	rw.Write(result)
}

func SourceRepoBranches(rw http.ResponseWriter, req *http.Request, params cobweb.Params) {
	provider := req.Header.Get("Provider")
	token := req.Header.Get("Access-Token")
	svc, err := source.NewServer(provider)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	status, result := svc.Branches(params.Get("owner"), params.Get("repo"), token)
	rw.WriteHeader(status)
	rw.Write(result)
}

func SourceRepoCommits(rw http.ResponseWriter, req *http.Request, params cobweb.Params) {
	provider := req.Header.Get("Provider")
	token := req.Header.Get("Access-Token")
	svc, err := source.NewServer(provider)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	status, result := svc.Commits(params.Get("owner"), params.Get("repo"), "", token)
	rw.WriteHeader(status)
	rw.Write(result)
}
