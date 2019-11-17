package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	corsOptions := handlers.CORSOption(
		handlers.AllowedOrigins([]string{"*"}),
	)

	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Pong!")
	})
	http.ListenAndServe(":5000", handlers.CORS(corsOptions)(r))
}
