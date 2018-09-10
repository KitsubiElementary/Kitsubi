package services

import "encoding/json"

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

func (r *KitsuService) GetUserLibrary() {
	if r.Id == "" {
		r.GetUserId()
	}
	var i map[string]interface{}
	err := json.Unmarshal(r.connect("/edge/users/"+r.Id+"/library-entries"), &i)
	if err == nil {

	}

}

func (r *KitsuService) connect(service string) []byte {
	s := RESTConnection{"https://kitsu.io/api", service, "GET", ""}
	return s.Connect()
}

type data struct {
	data *[]userListfield
}

type userListfield struct {
	id *string
}
