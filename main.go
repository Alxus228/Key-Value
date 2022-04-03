package main

import (
	"fmt"
	"github.com/Alxus228/Key-Value/storage"
)

func main() {
	a := storage.New()
	a.Put("something", "123")
	a.Put("bla bla bla", "one-two-three")
	fmt.Println(a.GetAll())
	fmt.Println(a.Get("nothing"))
	fmt.Println(a.Get("something"))
	a.Delete("something")
	fmt.Println(a.Get("something"))
	fmt.Println(a.Get("bla bla bla"))
}
