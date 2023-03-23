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
	RunpodAPIKey     string
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
		RunpodAPIKey:     os.Getenv("runpodAPIKey"),
	}
}

func LoadEnvAndGetSecrets() *SecretData {

	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {

		emptySecretData := SecretData{
			DatabaseHost:     "",
			DatabasePort:     "",
			DatabaseName:     "",
			DatabaseUsername: "",
			DatabasePassword: "",
			HashSalt:         "",
			PasswordSecret:   "",
			DemoAPIKey:       "",
			RunpodAPIKey:     "",
		}

		return &emptySecretData
	}

	// Get the secret data
	secretData := getSecretData()

	return &secretData
}
