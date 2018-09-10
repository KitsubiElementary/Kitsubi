package services

import (
	"io/ioutil"
	"net/http"
)

// RESTConnection is blablba
type RESTConnection struct {
	Api       string
	Path      string
	Operation string
	Body      string
}

func (r *RESTConnection) Connect() []byte {
	if r.Operation != "GET" {

	} else {
		return httpGet(r)
	}

	return nil
}

func httpGet(request *RESTConnection) []byte {
	resp, err := http.Get(request.Api + request.Path)
	if err == nil {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return contents
		}
	}
	return nil
}
