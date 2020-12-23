package main

import (
	"api"
)

func main() {

	a := api.NewServer("chulachinburi")
	a.Run(":8080")

}
