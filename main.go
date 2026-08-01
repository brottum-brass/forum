package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		port, exists := os.LookupEnv("PORT")
		if !exists {
			log.Fatal("Environment variable 'PORT' is not set")
		}

		portInt, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalf("Invalid port number: %v", err)
		}
	*/

	portInt := 8080

	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("./static/html")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portInt), mux))
}
