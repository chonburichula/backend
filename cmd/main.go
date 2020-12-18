package main

import (
	"api"
	"fmt"
)

func main() {
	a := api.NewServer("a")
	fmt.Println(a)
}
