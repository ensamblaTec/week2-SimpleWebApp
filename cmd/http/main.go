package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":90",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
