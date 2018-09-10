package services

import "encoding/json"

// KitsuService is
type KitsuService struct {
	Username string
	Password string
	Id       string
}

func (r *KitsuService) GetUserId() {
	var a bodyKitsu
	json.Unmarshal(r.connect("/edge/users?filter[name]="+r.Username), &a)
	r.Id = a.data[0].id
}

func (r *KitsuService) GetUserLibrary() {
	r.connect("/edge/users/" + r.Id + "/library-entries")
}

func (r *KitsuService) connect(service string) []byte {
	s := RESTConnection{"https://kitsu.io/api", service, "GET", ""}
	return s.Connect()
}

type bodyKitsu struct {
	data []userListfield
}

type userListfield struct {
	id string
}
