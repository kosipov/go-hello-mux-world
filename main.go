package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// RootHandler is mapping "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

// GetSampleHandler is mapping "/sample"
func GetSampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /sample")
}

// PostSampleHandler is mapping "/sample"
func PostSampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST /sample")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/sample", GetSampleHandler).Methods("GET")
	r.HandleFunc("/sample", PostSampleHandler).Methods("POST")

	// Server configuration
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run
	log.Fatal(srv.ListenAndServe())
}
