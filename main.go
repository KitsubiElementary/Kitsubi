package main

import (
	"fmt"

	"./services"
)

func main() {
	a := (services.KitsuService{"sifiro", "123", ""})
	v := a.GetUserLibrary()
	fmt.Print(v[0])
}
