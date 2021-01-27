package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/IktaS/godis/internal/short"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
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
	http.ListenAndServe(":8080", routes(repo))
}

func routes(r *short.Repo) *mux.Router {
	router := mux.NewRouter()
	var dir string

	flag.StringVar(&dir, "dir", "./web/assets/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router.HandleFunc("/favicon.ico", doNothing)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
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
	renderTemplate(w, "home", nil)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}

func getHandler(w http.ResponseWriter, r *http.Request, repo *short.Repo) {
	vars := mux.Vars(r)
	val, err := (*repo).Get(vars["key"])
	if err != nil {
		log.Printf("Error getting from repo: %v", err)
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(val)
	http.Redirect(w, r, val, http.StatusFound)
}

func saveHandler(w http.ResponseWriter, r *http.Request, repo *short.Repo) {
	var shortlink short.Link
	err := json.NewDecoder(r.Body).Decode(&shortlink)
	if err != nil {
		log.Printf("Error decoding : %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = (*repo).Save(&shortlink)
	if err != nil {
		log.Printf("Error saving : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data saved!")
}

var templates = template.Must(template.ParseFiles("./web/html/home.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
