package main

import (
	"fmt"
	"github.com/Alxus228/Key-Value/storage"
)

func main() {
	fmt.Println(storage.Put("idea"))
	fmt.Println(storage.Get(0))
	storage.Delete(0)
	fmt.Println(storage.Get(0))
}
