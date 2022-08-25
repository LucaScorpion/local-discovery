package server

import (
	"embed"
	"html/template"
	"local-discovery/internal/discovery"
	"mime"
	"net/http"
	"strings"
)

//go:embed static/*
var static embed.FS

func staticFiles(reg *discovery.Registry) handlerWithErrorFunc {
	return func(writer http.ResponseWriter, request *http.Request) error {
		// Construct the file path.
		filePath := request.URL.Path
		if request.URL.Path == "/" || request.URL.Path == "/index.html" {
			filePath = "/index.gohtml"
		}
		filePath = "static" + filePath

		// Read the file contents.
		content, contentErr := static.ReadFile(filePath)
		if contentErr != nil {
			http.NotFound(writer, request)
			return nil
		}

		// Send the content type.
		fileExt := filePath[strings.LastIndex(filePath, "."):]
		if fileExt == ".gohtml" {
			fileExt = ".html"
		}
		writer.Header().Set("Content-Type", mime.TypeByExtension(fileExt))

		// Parse gohtml templates.
		if strings.HasSuffix(filePath, ".gohtml") {
			tmpl, err := template.New(filePath).Parse(string(content))
			if err != nil {
				return err
			}

			return tmpl.Execute(writer, struct {
				Agents []*discovery.Agent
			}{
				Agents: reg.GetAgents(getRemoteIp(request)),
			})
		}

		writer.Write(content)
		return nil
	}
}
