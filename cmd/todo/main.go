package main

import (
	"fmt"

	//TODO clarify this
	"github.com/alinux78/todo/internal/db"
)

func main() {
	fmt.Println("starting server")
	err := database.Init()
	defer func() {
		fmt.Println("closing database")
		database.Close()
		fmt.Println("server stopped")
	}()

	if err != nil {
		panic(err)
	}
}
