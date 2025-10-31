package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ChinmayNoob/f1/internal/handler"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/ChinmayNoob/f1/internal/router"
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

	circuitRepo := repository.NewCircuitRepository()
	circuitService := service.NewCircuitService(circuitRepo)
	circuitHandler := handler.NewCircuitHandler(ctx, circuitService)

	seasonRepo := repository.NewSeasonRepository()
	seasonService := service.NewSeasonService(seasonRepo)
	seasonHandler := handler.NewSeasonHandler(ctx, seasonService)

	raceRepo := repository.NewRaceRepository()
	raceService := service.NewRaceService(raceRepo)
	raceHandler := handler.NewRaceHandler(ctx, raceService)

	resultRepo := repository.NewResultRepository()
	resultService := service.NewResultService(resultRepo)
	resultHandler := handler.NewResultHandler(ctx, resultService)

	standingRepo := repository.NewStandingRepository()
	standingService := service.NewStandingService(standingRepo)
	standingHandler := handler.NewStandingHandler(ctx, standingService)

	mux := http.NewServeMux()

	router.SetupRoutes(mux, constructorHandler, driverHandler, circuitHandler, seasonHandler, raceHandler, resultHandler, standingHandler)

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
