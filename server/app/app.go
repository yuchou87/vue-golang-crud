package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize With Router
func (a *App) Initialize(user, password, dbname, host string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", user, password, dbname, host)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.Handle("/ping", pingHandler()).Methods("GET")
	a.Router = router
}

// Run with CORS
func (a *App) Run(addr string) {
	corsOptions := handlers.CORSOption(handlers.AllowedOrigins([]string{"*"}))
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(corsOptions)(a.Router)))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
