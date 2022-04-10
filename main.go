package main

import (
	"github.com/Alxus228/Key-Value/server"
	"github.com/Alxus228/Key-Value/storage"
)

func main() {
	//We use interface here, because we might switch between storage implementations in the future
	var a server.Storage = storage.New()
	err := server.Run(&a)

	if err != nil {
		panic(err)
	}
}
