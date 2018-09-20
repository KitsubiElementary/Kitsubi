package services

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// restConnection is blablba
type restConnection struct {
	API       string
	Path      string
	Operation string
	Body      string
}

func (r *restConnection) connect() []byte {
	if r.Operation != "GET" {
		return httpPost(r)
	} else {
		return httpGet(r)
	}
}

func httpPost(request *restConnection) []byte {

	resp, err := http.Post(request.API+request.Path, "json", bytes.NewBufferString(request.Body))
	if err == nil {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return contents
		}
	}
	return nil
}

func httpGet(request *restConnection) []byte {
	resp, err := http.Get(request.API + request.Path)
	if err == nil {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return contents
		}
	}
	return nil
}
