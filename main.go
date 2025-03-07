package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	})
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Okey\n"))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", loggingMiddleware(fs))
	http.Handle("/healthz", loggingMiddleware(http.HandlerFunc(healthCheckHandler)))

	addr := ":" + port
	fmt.Printf("Server is running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
