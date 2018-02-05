package utils

import (
	"crypto/tls"
	"io"
	"net/http"
)

func HttpGet(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodGet, addr, header, body)
}

func HttpPost(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodPost, addr, header, body)
}

func HttpPut(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodPut, addr, header, body)
}

func HttpPatch(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodPatch, addr, header, body)
}

func HttpDelete(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodDelete, addr, header, body)
}

func request(method string, addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, addr, body)
	if err != nil {
		return nil, err
	}

	if header != nil {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}

	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
	}

	return client.Do(req)
}
