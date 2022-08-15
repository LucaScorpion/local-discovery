package server

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"local-discovery/internal/discovery"
	"log"
	"net/http"
)

//go:embed index.gohtml
var indexHtml string

func StartServer() {
	reg := discovery.NewRegistry()
	mux := http.NewServeMux()
	indexTemplate, _ := template.New("index").Parse(indexHtml)

	mux.HandleFunc("/api/agents", handleErrors(agents(reg)))
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}

		err := indexTemplate.Execute(writer, struct {
			Agents []*discovery.Agent
		}{
			Agents: reg.GetAgents(getRemoteIp(request)),
		})

		if err != nil {
			log.Println(err)
		}
	})

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
