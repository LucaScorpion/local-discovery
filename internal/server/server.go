package server

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"local-discovery/internal/discovery"
	"log"
	"net/http"
)

func StartServer() {
	reg := discovery.NewRegistry()
	mux := http.NewServeMux()

	mux.HandleFunc("/api/agents", handleErrors(agents(reg)))
	mux.HandleFunc("/", handleErrors(staticFiles(reg)))

	log.Println("Server starting on port 4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type handlerWithErrorFunc func(http.ResponseWriter, *http.Request) error

func handleErrors(handler handlerWithErrorFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err == nil {
			return
		}
		log.Println("[ERROR]", err)

		httpErr := httpError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("An unexpected error occurred: %v", err),
		}
		if cast, ok := err.(httpError); ok {
			httpErr = cast
		}

		writer.WriteHeader(httpErr.Status)
		json.NewEncoder(writer).Encode(httpErr)
	}
}
