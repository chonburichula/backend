package main

import "github.com/chonburichula/backend/api"

func main() {
	a := api.NewServer("chulachinburi")

	a.Run(":8080")
}
