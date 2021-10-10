package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"nosql1h21-stock-backend/backend/internal/config"
	"nosql1h21-stock-backend/backend/internal/handler"
	"nosql1h21-stock-backend/backend/internal/scratcher"
	"nosql1h21-stock-backend/backend/internal/service"
)

func connectionsClosedForServer(server *http.Server) chan struct{} {
	connectionsClosed := make(chan struct{})
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)
		defer signal.Stop(shutdown)
		<-shutdown

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		log.Println("Closing connections")
		if err := server.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		close(connectionsClosed)
	}()
	return connectionsClosed
}

func connectMongo(dbConn string) (_ *mongo.Client, disconnect func(), _ error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConn))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return client, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	}, nil
}

type Handler interface {
	Method() string
	Path() string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func registerHandler(router chi.Router, handler Handler) {
	router.Method(handler.Method(), handler.Path(), handler)
}

func cacheMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=3600") // Caching for 1 hour
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func scratchIfNeeded(collection *mongo.Collection, drop bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if !drop {
		n, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		if n > 0 {
			return
		}
	}

	scratcher.Scratch(ctx, collection)
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	drop := flag.Bool("d", false, "drop the database at the start (the stocks data will be gotten again)")
	flag.Parse()

	client, disconnect, err := connectMongo(cfg.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer disconnect()

	collection := client.Database("stock_market").Collection("stocks")

	scratchIfNeeded(collection, *drop)

	service := service.NewService(collection)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)

	router.Group(func(router chi.Router) {
		// router.Use(cacheMiddleware)
		registerHandler(router, &handler.StockHandler{Service: service})
		registerHandler(router, &handler.CountriesHandler{Service: service})
		registerHandler(router, &handler.SectorsHandler{Service: service})
		registerHandler(router, &handler.IndustriesHandler{Service: service})
		registerHandler(router, &handler.SearchHandler{Service: service})
		registerHandler(router, &handler.ExportHandler{Service: service})
		registerHandler(router, &handler.ImportHandler{Service: service})
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	connectionsClosed := connectionsClosedForServer(&server)
	log.Println("Server is listening on " + addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Println(err)
	}
	<-connectionsClosed
}
