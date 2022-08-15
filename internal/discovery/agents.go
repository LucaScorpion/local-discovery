package discovery

import "net/http"

func agents(reg *registry) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getAgents(reg, writer, request)
		case http.MethodPost:
			postAgent(reg, writer, request)
		default:
			writer.WriteHeader(405)
		}
	}
}

func getAgents(reg *registry, writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(getRemoteIp(request.RemoteAddr)))
}

func postAgent(reg *registry, writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Created agent goes here"))
}
