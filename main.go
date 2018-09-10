package main

import (
	"fmt"

	"./services"
)

func main() {
	a := (services.KitsuService{"sifiro", "123", ""})
	a.GetUserLibrary()
	fmt.Print("test")
}
