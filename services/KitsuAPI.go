package services

import "encoding/json"
import "time"

// KitsuService is
type KitsuService struct {
	Username string
	ID       string
}

func (r *KitsuService) getUserID() {
	var i map[string]interface{}
	//var i
	err := json.Unmarshal(r.connect("edge/users?filter[slug]="+r.Username), &i)
	if err == nil {
		r.ID = i["data"].([]interface{})[0].(map[string]interface{})["id"].(string)
	}
}

// GetUserEntries : Get Last Entries of a User
func (r *KitsuService) GetUserEntries() []AnimeList {
	var anime []AnimeList

	if r.ID == "" {
		r.getUserID()
	}
	var i map[string]interface{}
	err := json.Unmarshal(r.connect("edge/users/"+r.ID+"/library-entries"), &i)
	if err == nil {
		AnimeLists := i["data"].([]interface{})
		for index, element := range AnimeLists {
			Selector := element.(map[string]interface{})
			status := Selector["attributes"].(map[string]interface{})
			time, err := time.Parse(time.RFC3339, status["updatedAt"].(string))
			if err == nil {
				anime = append(anime,
					AnimeList{
						ID:                Selector["id"].(string),
						Name:              r.GetAnimeInfo(Selector["id"].(string)).Name,
						ChapterInProgress: 0,
						updatedAt:         time})
			}
			index++
		}
		return anime
	}
	return anime
}

// GetAnimeInfo get Anime Information by Id
func (r *KitsuService) GetAnimeInfo(id string) AnimeInfo {
	r.connect("edge/library-entries/")
	return AnimeInfo{Name: ""}
}

func (r *KitsuService) connect(service string) []byte {
	s := restConnection{API: "https://kitsu.io/api/", Path: service, Operation: "GET"}
	return s.connect()
}

// AnimeList is a struct formatted usually with a Array of Current Anime Status from a User
type AnimeList struct {
	ID                string
	Name              string
	ChapterInProgress int
	updatedAt         time.Time
}

// AnimeInfo is a Struct that represent Anime's information from Kitsu's Database
type AnimeInfo struct {
	Name string
}
