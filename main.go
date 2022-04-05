package main

import (
	"github.com/Alxus228/Key-Value/server"
	"github.com/Alxus228/Key-Value/storage"
)

func main() {
	//using interface, because we might switch between storage implementations in future
	var a server.Storage = storage.New()
	err := server.Run(&a)

	if err != nil {
		panic(err)
	}
}
