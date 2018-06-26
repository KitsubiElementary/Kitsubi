package services

func login() {
	s := RESTConnection{"www.google.es", "ds", ""}
	s.connect()
}
