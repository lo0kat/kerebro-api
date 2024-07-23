package main

import "github.com/lo0kat/kerebro-api/internal/api"

func main() {

	server := api.NewAPIServer(":8080", nil)
	server.Run()
}
