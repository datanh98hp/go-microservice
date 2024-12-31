package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// route
	route := chi.NewRouter()
	route.Get("/hello", basicHandler)
	//
	server := &http.Server{
		Addr:    ":3000",
		Handler: route,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

// handle for route
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
