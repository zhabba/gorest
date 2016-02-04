package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting API version `%s`", apiVersion)
	router := NewRouter()
	log.Fatal(http.ListenAndServe("0.0.0.0:6060", router))
}
