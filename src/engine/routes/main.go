package routes

import (
	"bigdawgs/handlers"
	"fmt"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", handlers.HealthzHandler)

}

func ListenAndServe(port string) error {
	mux := http.NewServeMux()
	Register(mux)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	return http.ListenAndServe(addr, mux)
}
