package main

import (
	"fmt"

	"go.uber.org/dig"
)

type Book struct {
	name string
}

type User struct {
	book1 Book `name:"book1"`
	book2 Book `name:"book2"`
}

type UserRepo struct {
	user User
}

func (r UserRepo) Display() {
	fmt.Println(r.user.book1.name)
	fmt.Println(r.user.book2.name)
}

func NewBook1() Book {
	return Book{name: "book1"}
}

func NewBook2() Book {
	return Book{name: "book2"}
}

//func BuildUser(book1 Book, book2 Book) User {
//	return User{book1: book1, book2: book2}
//}

func BuildUserRepo(u User) UserRepo {
	return UserRepo{user: u}
}

func main() {
	c := dig.New()
	c.Provide(NewBook1)
	c.Provide(NewBook2)
	//c.Provide(BuildUser)
	c.Provide(BuildUserRepo)
	c.Invoke(func(r UserRepo) {
		r.Display()
	})
}
