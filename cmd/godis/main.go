package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome home %s", vars["name"])
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{name}", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
