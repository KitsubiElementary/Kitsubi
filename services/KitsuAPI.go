package services

import "encoding/json"
import "time"

// KitsuService is
type KitsuService struct {
	Username string
	Password string
	Id       string
}

func (r *KitsuService) GetUserId() {
	var i map[string]interface{}
	//var i
	err := json.Unmarshal(r.connect("/edge/users?filter[slug]="+r.Username), &i)
	if err == nil {
		r.Id = i["data"].([]interface{})[0].(map[string]interface{})["id"].(string)
	}
}

func (r *KitsuService) GetUserLibrary() []AnimeList {
	var anime []AnimeList

	if r.Id == "" {
		r.GetUserId()
	}
	var i map[string]interface{}
	err := json.Unmarshal(r.connect("/edge/users/"+r.Id+"/library-entries"), &i)
	if err == nil {
		AnimeLists := i["data"].([]interface{})
		for index, element := range AnimeLists {
			Selector := element.(map[string]interface{})
			status := Selector["attributes"].(map[string]interface{})
			time, err := time.Parse(time.RFC3339, status["updatedAt"].(string))
			if err == nil {
				anime = append(anime,
					AnimeList{
						Selector["id"].(string), "", 0, time})
			}
			index++
		}
		return anime
	}
	return anime
}

func (r *KitsuService) connect(service string) []byte {
	s := RESTConnection{"https://kitsu.io/api", service, "GET", ""}
	return s.Connect()
}

type AnimeList struct {
	ID                string
	Name              string
	ChapterInProgress int
	updatedAt         time.Time
}
