package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// GOPath Environment Variable should be declarated as Workspace Directory Path
	"./services"
)

func main() {
	var config appConfig
	resp, err := ioutil.ReadFile("./config.json")
	if err == nil {
		json.Unmarshal(resp, &config)
		a := services.KitsuService{Username: config.KitsuUsername, ID: config.KitsuUserID}
		v := a.GetUserEntries()
		fmt.Print(v[0])
	}
}

type appConfig struct {
	KitsuUsername string
	KitsuUserID   string
	TwitterToken  string
}
