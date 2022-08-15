package discovery

import "net/http"

func agents(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getAgents(writer, request)
	case http.MethodPost:
		postAgent(writer, request)
	default:
		writer.WriteHeader(405)
	}
}

func getAgents(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(getRemoteIp(request.RemoteAddr)))
}

func postAgent(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Created agent goes here"))
}
