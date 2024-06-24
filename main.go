package main

import (
	"log"
	"net/http"
)

func healthz(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("OK"))
}

func main() {
	const addr = "localhost:8080"
	const root = "."
	serveMux := http.NewServeMux()
	serveMux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir(root))))
	serveMux.HandleFunc("/healthz", healthz)
	server := &http.Server{Handler: serveMux, Addr: addr}
	log.Fatal(server.ListenAndServe())
}
