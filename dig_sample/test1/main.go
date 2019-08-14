package main

import (
	"fmt"

	"go.uber.org/dig"
)

type Book struct {
	Name string
}

func main() {
	c := dig.New()
	c.Provide(func() *Book {
		book := Book{Name: "test"}
		return &book
	})

	c.Invoke(func(book *Book) {
		fmt.Println(book.Name)
	})
}
