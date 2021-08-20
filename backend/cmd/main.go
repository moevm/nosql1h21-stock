package main

import (
	"context"
	"fmt"
	"net/http"
	"nosql1h21-stock-backend/backend/internal/repository"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"nosql1h21-stock-backend/backend/internal/config"
	"nosql1h21-stock-backend/backend/internal/handler"
	"nosql1h21-stock-backend/backend/internal/service"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal().Err(err).Msg("Configuration error")
	}

	r := chi.NewRouter()

	mongoConnectCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(mongoConnectCtx, options.Client().ApplyURI(cfg.DBConnString))
	if err != nil {
		logger.Fatal().Err(err).Msg("MongoDB connection error")
	}
	defer mongoClient.Disconnect(mongoConnectCtx)
	if err := mongoClient.Ping(mongoConnectCtx, readpref.Primary()); err != nil {
		logger.Fatal().Err(err).Msg("MongoDB pinging error")
	}
	logger.Info().Msg("MongoDB is successfully connected and pinged.")

	collection := mongoClient.Database("stock_market").Collection("stocks")

	stockService := service.NewStockService(&logger, collection)
	validTickersRepo := repository.NewCache()
	validTickersService := service.NewValidTickersService(&logger, validTickersRepo, collection)

	stockHandler := handler.NewStockHandler(&logger, stockService)
	validTickersHandler := handler.NewValidTickersHandler(&logger, validTickersService)

	r.Route("/", func(r chi.Router) {
		r.Use(middleware.RequestLogger(&handler.LogFormatter{Logger: &logger}))
		r.Use(middleware.Recoverer)
		r.Method(http.MethodGet, handler.StockPath, stockHandler)
		r.Method(http.MethodGet, handler.ValidTickersPath, validTickersHandler)
	})

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdown)

	go func() {
		logger.Info().Msgf("Server is listening on :%d", cfg.Port)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Server error")
		}
	}()

	<-shutdown

	logger.Info().Msg("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Server shutdown error")
	}

	logger.Info().Msg("Server stopped gracefully")
}
