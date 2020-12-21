package main

import "api"

func main() {
	a := api.NewServer("test", "customer")
	a.Run()

}
