package main

import (
	"log"
	"net/http"
	"github.com/jrakhman/router"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}