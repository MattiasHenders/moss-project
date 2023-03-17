package main

import (
	"os"

	server "github.com/MattiasHenders/moss-communication-server/internal"
	"github.com/MattiasHenders/moss-communication-server/pkg/secrets"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func runAllCronJobs() {
	c := cron.New()

	if os.Getenv("ENV") == "production" {
		// Cron Jobs here...
		zap.L().Info("Running cron jobs...")
	}

	c.Start()
	defer c.Stop()
}

func main() {

	_ = secrets.LoadEnvAndGetSecrets()

	// Database
	// _, dbErr := db.NewDatabase(
	// 	secretData.DatabaseHost,
	// 	secretData.DatabasePort,
	// 	secretData.DatabaseUsername,
	// 	secretData.DatabasePassword,
	// 	secretData.DatabaseName,
	// )
	// if dbErr != nil {
	// 	panic(dbErr)
	// }

	// Logger
	var logger *zap.Logger
	var err error
	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "staging" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	// Run Cron Jobs
	runAllCronJobs()

	// Server
	port := os.Getenv("PORT")

	server.Start(port)
}
