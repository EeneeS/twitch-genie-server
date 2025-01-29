package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

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
