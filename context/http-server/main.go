package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	case <-time.After(5 * time.Second):
		log.Println("Request finalizada com sucesso")
		w.Write([]byte("Request finalizada com sucesso"))
	}
}
