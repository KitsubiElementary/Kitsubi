package main

import (
	"encoding/json"
	"io/ioutil"
	// GOPath Environment Variable should be declarated as Workspace Directory Path
	"./services"
)

func main() {
	var config appConfig
	resp, err := ioutil.ReadFile("./config.json")

	if err == nil {
		err := json.Unmarshal(resp, &config)
		if err == nil {
			//kitsu := services.KitsuService{Username: config.Client.KitsuUsername, ID: config.Client.KitsuUserID}
			twitter := services.TwitterAPI{
				TwitterConsumerAPIkeys:       config.Client.TwitterConsumerAPIkeys,
				TwitterConsumerAPIkeysSecret: config.Client.TwitterConsumerAPIkeysSecret,
				TwitterAccessToken:           config.TwitterAppToken.TwitterAccessToken,
				TwitterAccessTokenSecret:     config.TwitterAppToken.TwitterAccessTokenSecret}

			//v := kitsu.GetUserEntries()
			twitter.Tweet("blabla")
		}
	}

}

type appConfig struct {
	Client          clientConfig     `json:"client"`
	TwitterAppToken twitterAppConfig `json:"twitterAppToken"`
}
type clientConfig struct {
	KitsuUsername                string
	KitsuUserID                  string
	TwitterConsumerAPIkeys       string
	TwitterConsumerAPIkeysSecret string
}

type twitterAppConfig struct {
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
}
