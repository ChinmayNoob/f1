package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ChinmayNoob/f1/internal/handler"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/ChinmayNoob/f1/internal/service"
	"github.com/ChinmayNoob/f1/pkg/db"
)

func main() {
	ctx := context.Background()
	db.InitDB(ctx)
	defer db.Conn.Close(ctx)

	constructorRepo := repository.NewConstructorRepository()
	constructorService := service.NewConstructorService(constructorRepo)
	constructorHandler := handler.NewConstructorHandler(ctx, constructorService)

	driverRepo := repository.NewDriverRepository()
	driverService := service.NewDriverService(driverRepo)
	driverHandler := handler.NewDriverHandler(ctx, driverService)

	mux := http.NewServeMux()

	mux.HandleFunc("/constructors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			constructorHandler.GetConstructor(w, r)
		case http.MethodPost:
			constructorHandler.CreateConstructor(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/constructors/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			constructorHandler.GetConstructorByID(w, r)
		case http.MethodPut:
			constructorHandler.UpdateConstructor(w, r)
		case http.MethodDelete:
			constructorHandler.DeleteConstructor(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/drivers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			driverHandler.GetDriver(w, r)
		case http.MethodPost:
			driverHandler.CreateDriver(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/drivers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			driverHandler.GetDriverByID(w, r)
		case http.MethodPut:
			driverHandler.UpdateDriver(w, r)
		case http.MethodDelete:
			driverHandler.DeleteDriver(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
