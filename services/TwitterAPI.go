package services

import (
	"os"
)

// TwitterAPI is
type TwitterAPI struct {
	Username string
	Token    string
}

func (r *TwitterAPI) getToken() bool {
	os.OpenFile("./config.json", os.O_RDWR|os.O_CREATE, 0755)
	// if Empty Call requestToken
	return false
}

func (r *TwitterAPI) requestToken() {
}

// Tweet is a function for send Tweets
func (r *TwitterAPI) Tweet(text string) {
	if r.Token == "" {
		if !r.getToken() {
			return
		}
	}

	r.connect("statuses/update", "")
}

func (r *TwitterAPI) connect(service string, body string) []byte {
	s := restConnection{API: "https://api.twitter.com/api", Path: service, Operation: "POST", Body: body}
	return s.connect()
}
