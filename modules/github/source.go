package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/chinx/sourceproxy/modules/source"
	"github.com/chinx/sourceproxy/utils"
)

type sourceClient struct{}

func init() {
	source.Registry(provider, &sourceClient{})
}

func (cli *sourceClient) User(token string) (int, []byte) {
	header := defaultHeader()
	header[accessKey] = tokenPrefix + token
	resp, err := utils.HttpGet(userURL, header, nil)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	byteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return http.StatusInternalServerError, []byte(err.Error())
	}

	resp.Body.Close()
	return resp.StatusCode, byteArr
}

func (cli *sourceClient) Repos(token string) (int, []byte) {
	header := defaultHeader()
	header[accessKey] = tokenPrefix + token
	resp, err := utils.HttpGet(reposURL, header, nil)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	byteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return http.StatusInternalServerError, []byte(err.Error())
	}

	resp.Body.Close()
	return resp.StatusCode, byteArr
}

func (cli *sourceClient) Branches(owner, repo, token string) (int, []byte) {
	header := defaultHeader()
	header[accessKey] = tokenPrefix + token
	resp, err := utils.HttpGet(fmt.Sprintf(branchesURL, owner, repo), header, nil)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	byteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return http.StatusInternalServerError, []byte(err.Error())
	}

	resp.Body.Close()
	return resp.StatusCode, byteArr
}

func (cli *sourceClient) Commits(owner, repo, branch, token string) (int, []byte) {
	header := defaultHeader()
	header[accessKey] = tokenPrefix + token
	if branch == "" {
		branch = "master"
	}
	params := url.Values{}
	params.Set("sha", branch)

	resp, err := utils.HttpGet(fmt.Sprintf(commitsURL, owner, repo), header, params)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	byteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return http.StatusInternalServerError, []byte(err.Error())
	}

	resp.Body.Close()
	return resp.StatusCode, byteArr
}

func defaultHeader() map[string]string {
	return map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
}
