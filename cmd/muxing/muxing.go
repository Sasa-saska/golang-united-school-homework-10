package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello, %s!", vars["PARAM"])
	})
	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "I got message:\n%s", string(body))

	}).Methods("POST")

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		var a, b int
		a, err := strconv.Atoi(r.Header.Get("a"))
		if err != nil {
			log.Fatal(err)
		}
		b, err = strconv.Atoi(r.Header.Get("b"))
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set(strings.ToLower("a+b"), strconv.Itoa(a+b))
	}).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
