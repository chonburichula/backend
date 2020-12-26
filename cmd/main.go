package main

import (
	"backend/api"
)

func main() {
	a := api.NewServer("chulachinburi")

	a.Run(":8080")

}
