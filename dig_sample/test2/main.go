package main

import (
	"fmt"

	"go.uber.org/dig"
)

type BookService struct {
	Repository BookRepository
}

func (s *BookService) Display() {
	name := s.Repository.GetName()
	fmt.Println(name)
}

type BookRepository interface {
	GetName() string
}

type BookRepositoryImpl struct{}

func (b *BookRepositoryImpl) GetName() string {
	return "testbook"
}

func main() {
	c := dig.New()
	c.Provide(func() BookRepository {
		return &BookRepositoryImpl{}
	})

	c.Provide(func(r BookRepository) *BookService {
		return &BookService{Repository: r}
	})

	c.Invoke(func(service *BookService) {
		service.Display()
	})
}
