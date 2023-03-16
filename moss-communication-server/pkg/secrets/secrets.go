package secrets

import (
	"os"

	"github.com/joho/godotenv"
)

type SecretData struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	HashSalt         string
	PasswordSecret   string
	DemoAPIKey       string
}

func getSecretData() SecretData {
	return SecretData{
		DatabaseHost:     os.Getenv("databaseHost"),
		DatabasePort:     os.Getenv("databasePort"),
		DatabaseName:     os.Getenv("databaseName"),
		DatabaseUsername: os.Getenv("databaseUsername"),
		DatabasePassword: os.Getenv("databasePassword"),
		HashSalt:         os.Getenv("hashSalt"),
		PasswordSecret:   os.Getenv("passwordSecret"),
		DemoAPIKey:       os.Getenv("demoAPIKey"),
	}
}

func LoadEnvAndGetSecrets() *SecretData {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Get the secret data
	secretData := getSecretData()

	return &secretData
}
