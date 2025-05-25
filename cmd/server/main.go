// @title Service Catalog API
// @version 1.0
// @description This is the API for service catalog assignment.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq" // 🔥 REQUIRED: Register postgres driver
	_ "github.com/sonishivam10/service-catalog/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sonishivam10/service-catalog/internal/config"
	"github.com/sonishivam10/service-catalog/internal/handler"
	"github.com/sonishivam10/service-catalog/internal/middleware"
	"github.com/sonishivam10/service-catalog/internal/repository"
	"github.com/sonishivam10/service-catalog/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to DB
	db, err := sqlx.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Printf("Connected to DB at %s", cfg.DatabaseURL)

	// Setup repository, service, and handlers
	repo := repository.NewPostgresRepository(db)
	svc := service.NewServiceUsecase(repo)

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/healthz", handler.HealthCheck).Methods("GET")
	r.Use(middleware.AuthMiddleware) // apply to all routes

	handler.NewServiceHandler(r, svc)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited gracefully")

}
