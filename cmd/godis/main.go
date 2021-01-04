package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/IktaS/godis/internal/short"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func main() {
	dbInt, err := strconv.Atoi(os.Getenv("DB_INT"))
	if err != nil {
		log.Fatal(err)
	}
	repo := short.NewRedisRepo(context.Background(), &redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbInt,
	})
	err = (*repo).Init()
	if err != nil {
		log.Fatal("cannot initialize repostory")
	}
	log.Fatal(http.ListenAndServe(":8080", routes(repo)))
}

func routes(r *short.Repo) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/save", useRepo(saveHandler, r))
	router.HandleFunc("/{key}", useRepo(getHandler, r))
	return router
}

func useRepo(fn func(http.ResponseWriter, *http.Request, *short.Repo), repo *short.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, repo)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got Home!")
}

func getHandler(w http.ResponseWriter, r *http.Request, repo *short.Repo) {
	vars := mux.Vars(r)
	val, err := (*repo).Get(vars["key"])
	if err != nil {
		log.Printf("Error getting from repo: %v", err)
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(val)
}

func saveHandler(w http.ResponseWriter, r *http.Request, repo *short.Repo) {
	var shortlink short.Link
	err := json.NewDecoder(r.Body).Decode(&shortlink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = (*repo).Save(&shortlink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data saved!")
}
