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

	port := os.Getenv("PORT")

	cfg := &config{
		address: port,
	}

	app := &application{
		config: *cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
