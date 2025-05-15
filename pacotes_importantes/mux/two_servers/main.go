package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog Page"))
	})
	// http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World 2!"))
	})
	mux2.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog Page 2"))
	})
	// http.ListenAndServe(":8081", mux2)

	go func() {
		http.ListenAndServe(":8080", mux)
	}()

	go func() {
		http.ListenAndServe(":8081", mux2)
	}()
	fmt.Scanln()
}
