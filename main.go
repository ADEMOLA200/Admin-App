package main

import (
	"log"

	"github.com/ADEMOLA200/Admin-App.git/app"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	fx.New(app.Module).Run()
}
