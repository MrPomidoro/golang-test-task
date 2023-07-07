package main

import (
	"context"
	_ "github.com/golang-test-task/docs"
	"github.com/golang-test-task/internal/app"
	"github.com/golang-test-task/pkg/common/logging"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Attempt to load environment variables from .env file using godotenv package.
	if err := godotenv.Load(); err != nil {
		// If there is an error, log a fatal error message and exit the program.
		log.Fatal("Error loading .env file")
	}
	// Set log level
	if err := logging.SetLevel(os.Getenv("LOG_LEVEL")); err != nil {
		// If there is an error, log a fatal error message and exit the program.
		log.Fatal(err)
	}
}

// General info
// @Title           Golang test task
// @Version         1.0
// @Description     Golang test task
// @BasePath  /api/v1
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.L(ctx).Info("Golang test task")

	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	newApp := app.NewApp(ctx)
	server, err := newApp.Start()
	if err != nil {
		logging.L(ctx).Fatal(err.Error())
	}

	log.Fatal(server.Run())

}
