package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// @title Twitch Genie API
// @version 1.0
// @description This is the Twitch Genie API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:6969

// @BasePath /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &config{
		address: os.Getenv("PORT"),
	}

	app := &application{
		config: *cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
