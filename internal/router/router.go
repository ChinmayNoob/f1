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
}
