package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
