package main

import (
	"fmt"

	"go.uber.org/dig"
)

type UseCase interface {
	Display()
}

func NewUseCase(r Repository) UseCase {
	return &UseCaseImpl{repository: r}
}

type UseCaseImpl struct {
	repository Repository
}

func (s UseCaseImpl) Display() {
	name := s.repository.GetName()
	fmt.Println(name)
}

type Repository interface {
	GetName() string
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

type RepositoryImpl struct{}

func (r *RepositoryImpl) GetName() string {
	return "test string"
}

func run(s UseCase) {
	s.Display()
}

func main() {
	c := dig.New()
	c.Provide(NewRepository)
	c.Provide(NewUseCase)
	c.Invoke(run)
}
