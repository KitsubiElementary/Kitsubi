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
	Body      []byte
	Header    [][]string
}

func (r *restConnection) connect() []byte {
	switch r.Operation {
	case "GET":
		return httpGet(r)
	case "POST":
		return httpPost(r)
	default:
		return nil
	}
}

func httpPost(request *restConnection) []byte {
	// "application/json"
	client := &http.Client{}
	req, err := http.NewRequest("POST", request.API+request.Path, bytes.NewBuffer(request.Body))
	for index, element := range request.Header {
		req.Header.Add(element[0], element[1])
		index++
	}
	resp, err := client.Do(req)
	if err == nil {
		defer req.Body.Close()
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
