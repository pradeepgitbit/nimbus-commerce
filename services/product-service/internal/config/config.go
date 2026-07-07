package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Using Docker/Kubernetes environment variables")
	}
}