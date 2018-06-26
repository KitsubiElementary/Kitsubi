package services

import (
	"io/ioutil"
	"net/http"
)

type RESTConnection struct {
	URL       string
	Operation string
	Body      string
}

func (r *RESTConnection) connect() []byte {
	switch op := r.Operation; op {
	case "GET":
		return HTTPGet(r.URL)
		break
	}
	return nil
}

func HTTPGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return contents
		}
	}
	return nil
}
