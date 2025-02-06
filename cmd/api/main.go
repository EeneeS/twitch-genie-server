package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/eenees/twitch-genie-server/internal/repositories"
	"github.com/eenees/twitch-genie-server/internal/utils/auth"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

  mongoUri := os.Getenv("MONGO_URI")
  db, err := connectMongoDB(mongoUri)
  if err != nil {
    log.Fatal("Error connecting to mongoDB:", err)
  }
  defer db.Disconnect(context.Background())

	authSecret := os.Getenv("JWT_SECRET")

	app := &application{
		config: *cfg,
		// repo:   *repositories.NewMockRepository(),
    repo: *repositories.NewRepository(db),
		auth:   *auth.NewJWTAuthenticator(authSecret),
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}

func connectMongoDB(uri string) (*mongo.Client, error) {
  options := options.Client().ApplyURI(uri)
  client, err := mongo.Connect(options)
  if err != nil {
    return nil, err
  }

  err = client.Ping(context.Background(), nil)
  if err != nil {
    return nil, err
  }

  fmt.Println("[database] connected")
  return client, nil
}
