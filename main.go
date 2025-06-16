package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/search", SearchHandler)
	mux.HandleFunc("/bar", BarHandler)

	handler := cors.Default().Handler(mux)

	fmt.Println("running backend on port 3001")
	log.Fatal(http.ListenAndServe(":3001", handler))
}