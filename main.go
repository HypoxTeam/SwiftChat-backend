package main

import (
	"log"
	"os"
)

const (
	defaultPort string = "7777"
)

func main() {
	var (
		port string = stringEnv("PORT", defaultPort)
	)
	
	log.Println(port)
}

func stringEnv(key, defaultValue string) string {
	var value string = os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
