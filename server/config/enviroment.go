package config

import "github.com/gofor-little/env"

func LoadEnv() {
	// Load environment variables
	if err := env.Load(".env"); err != nil {
		panic(err)
	}
}
