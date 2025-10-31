package router

import (
	"net/http"

	"github.com/ChinmayNoob/f1/internal/handler"
)

func SetupRoutes(
	mux *http.ServeMux,
	constructorHandler *handler.ConstructorHandler,
	driverHandler *handler.DriverHandler,
	circuitHandler *handler.CircuitHandler,
	seasonHandler *handler.SeasonHandler,
	raceHandler *handler.RaceHandler,
	resultHandler *handler.ResultHandler,
	standingHandler *handler.StandingHandler,
) {
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

	mux.HandleFunc("/circuits", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			circuitHandler.GetCircuit(w, r)
		case http.MethodPost:
			circuitHandler.CreateCircuit(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/circuits/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			circuitHandler.GetCircuitByID(w, r)
		case http.MethodPut:
			circuitHandler.UpdateCircuit(w, r)
		case http.MethodDelete:
			circuitHandler.DeleteCircuit(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/seasons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			seasonHandler.GetSeason(w, r)
		case http.MethodPost:
			seasonHandler.CreateSeason(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/seasons/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			seasonHandler.GetSeasonByID(w, r)
		case http.MethodPut:
			seasonHandler.UpdateSeason(w, r)
		case http.MethodDelete:
			seasonHandler.DeleteSeason(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/races", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			raceHandler.GetRace(w, r)
		case http.MethodPost:
			raceHandler.CreateRace(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/races/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			raceHandler.GetRaceByID(w, r)
		case http.MethodPut:
			raceHandler.UpdateRace(w, r)
		case http.MethodDelete:
			raceHandler.DeleteRace(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			resultHandler.GetResult(w, r)
		case http.MethodPost:
			resultHandler.CreateResult(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/results/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			resultHandler.GetResultByID(w, r)
		case http.MethodPut:
			resultHandler.UpdateResult(w, r)
		case http.MethodDelete:
			resultHandler.DeleteResult(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/driver-standings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			standingHandler.GetDriverStanding(w, r)
		case http.MethodPost:
			standingHandler.CreateDriverStanding(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/driver-standings/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			standingHandler.GetDriverStandingByID(w, r)
		case http.MethodPut:
			standingHandler.UpdateDriverStanding(w, r)
		case http.MethodDelete:
			standingHandler.DeleteDriverStanding(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/constructor-standings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			standingHandler.GetConstructorStanding(w, r)
		case http.MethodPost:
			standingHandler.CreateConstructorStanding(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/constructor-standings/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			standingHandler.GetConstructorStandingByID(w, r)
		case http.MethodPut:
			standingHandler.UpdateConstructorStanding(w, r)
		case http.MethodDelete:
			standingHandler.DeleteConstructorStanding(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
