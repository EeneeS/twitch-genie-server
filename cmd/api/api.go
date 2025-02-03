package main

import (
	"fmt"
	_ "github.com/eenees/twitch-genie-server/docs"
	"github.com/eenees/twitch-genie-server/internal/handlers"
	"github.com/eenees/twitch-genie-server/internal/middlewares"
	"github.com/eenees/twitch-genie-server/internal/repositories"
	"github.com/eenees/twitch-genie-server/internal/services"
	"github.com/eenees/twitch-genie-server/internal/utils/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
	repo   repositories.Repository
	auth   auth.JWTAuthenticator
}

type config struct {
	address string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	app.setupMiddleWare(r)

	tokenService := services.NewTokenService(&app.repo, &app.auth)
	tokenHandler := handlers.NewTokenHandler(tokenService)

	channelService := services.NewChannelService(&app.repo)
	channelHandler := handlers.NewChannelHandler(channelService)

	r.Route("/v1", func(r chi.Router) {
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.address)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		r.Get("/health", app.healthCheckHandler)

		r.Post("/exchange-token", tokenHandler.ExchangeToken)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware(&app.auth))
			r.Get("/moderated-channels", channelHandler.GetModeratedChannels)
		})

		r.Get("/ws", WsHandler)

	})

	return r
}

func (app *application) setupMiddleWare(r chi.Router) {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:6969"}, // Change to match your frontend
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsOptions).Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %v", app.config.address)

	return server.ListenAndServe()
}
