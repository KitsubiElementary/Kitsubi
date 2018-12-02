package services

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TwitterAPI is
type TwitterAPI struct {
	TwitterConsumerAPIkeys       string
	TwitterConsumerAPIkeysSecret string
	TwitterAccessToken           string
	TwitterAccessTokenSecret     string
	OAuthToken                   string
}

func (r *TwitterAPI) getToken() bool {
	response := string(r.connect("/oauth/request_token", "POST", nil))
	response += ""
	return false
}

func (r *TwitterAPI) requestToken() {
}

// Tweet is a function for send Tweets
func (r *TwitterAPI) Tweet(text string) {
	/*	if r.OAuthToken == "" {
		if !r.getToken() {
			return
		}
	}*/
	body, err := json.Marshal(twitterPost{Status: text})
	if err == nil {
		r.connect("/1.1/statuses/update.json", "POST", body)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

/// RandStringRunes is blablba
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func computeHmacSHA1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (r *TwitterAPI) FormatHeaderSignature(oauthNounce string, timestamp string) string {
	var Signature = "oauth_consumer_key=" + r.TwitterConsumerAPIkeys + "&" +
		"oauth_nonce=" + /*base64.StdEncoding.EncodeToString([]byte(*/ oauthNounce /*))*/ + "&" +
		"oauth_signature_method=HMAC-SHA1&oauth_timestamp=" + timestamp + "&" +
		"oauth_token=" + r.TwitterAccessToken + "&oauth_version=1.0"
	Signature = url.PathEscape(Signature)
	Signature = strings.Replace(Signature, "&", "%26", -1)
	Signature = strings.Replace(Signature, "=", "%3D", -1)
	return Signature
}

func (r *TwitterAPI) encryptSignature(signatureClean string) string {
	return strings.Replace(
		/*url.PathEscape(*/
		computeHmacSHA1(signatureClean, url.PathEscape(r.TwitterConsumerAPIkeysSecret+"&"+r.TwitterAccessTokenSecret) /*)*/), "=", "%3D", -1)
}

func (r *TwitterAPI) createAuth(operation string, service string, body string) string {
	service = url.PathEscape(service)
	service = strings.Replace(service, ":", "%3A", -1)
	body = url.PathEscape(string(body))
	body = strings.Replace(body, ":", "%3A", -1)

	//var oauthNounce = base64.StdEncoding.EncodeToString([]byte(RandStringRunes(32)))
	var oauthNounce = RandStringRunes(32)

	var timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	var headerSignature = r.FormatHeaderSignature(oauthNounce, timestamp)
	var signatureClean = strings.ToUpper(operation) + "&" + service + "&" + headerSignature + "&" + string(body)

	fmt.Print("Content Signature to Encrypt: \n" + signatureClean)

	return "OAuth oauth_consumer_key=\"" + r.TwitterConsumerAPIkeys + "\", " +
		"oauth_nonce=\"" + oauthNounce +
		"\", oauth_signature=\"" + r.encryptSignature(signatureClean) +
		"\", " + "oauth_signature_method=\"HMAC-SHA1\", oauth_timestamp=\"" + timestamp + "\", " +
		"oauth_token=\"" + r.TwitterAccessToken + "\", oauth_version=\"1.0\""
}

func (r *TwitterAPI) connect(service string, operation string, body []byte) []byte {

	var header = [][]string{
		{"Authorization", r.createAuth(operation, "https://api.twitter.com"+service, string(body))},
		{"Content-Type", "application/json"}}
	fmt.Print("\nOauth Header: \n" + header[0][1])
	s := restConnection{API: "https://api.twitter.com", Path: service, Operation: "POST", Body: body, Header: header}
	response := string(s.connect())
	response += ""
	return []byte(response)
	//return s.connect()
}

type twitterPost struct {
	Status   string `json:"status"`
	MediaIds string `json:"media_ids"`
}
