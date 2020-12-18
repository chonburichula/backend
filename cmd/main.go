package main

import (
	"api"
)

func main() {
	a := api.NewServer("ordersystem", "customer")
	a.Run()

}
