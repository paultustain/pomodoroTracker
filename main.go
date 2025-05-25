package main

import (
	"fmt"
	"net/http"
)

const ROOTDIR = "./app"
const PORT = "8080"

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(ROOTDIR))

	mux.Handle("/", fs)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + PORT,
	}
	fmt.Printf("File serving on port: localhost:8080")

	server.ListenAndServe()
}
