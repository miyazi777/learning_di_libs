package main

import (
	"fmt"

	"go.uber.org/dig"
)

type Config struct {
	Prefix string
}

func main() {
	c := dig.New()
	c.Provide(func() *Config {
		var cfg Config
		cfg.Prefix = "test"
		return &cfg
	})

	c.Invoke(func(cfg *Config) {
		fmt.Println(cfg.Prefix)
	})
}
