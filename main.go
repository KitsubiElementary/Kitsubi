package main

import (
	"fmt"

	"./services"
)

func main() {
	a := (services.KitsuService{"sifi", "123", ""})
	a.GetUserId()
	fmt.Print("test")
}
