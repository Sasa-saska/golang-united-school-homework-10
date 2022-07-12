package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", GetParam)
	router.HandleFunc("/bad", Bad)
	router.HandleFunc("/data", PostParam)
	router.HandleFunc("/headers", PostHeaders)

	fmt.Printf("Starting API server on %s:%d\n", host, port)
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

func GetParam(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r) // to take params from url path like this -> /name/{PARAM}
	value := m["PARAM"]
	w.Write([]byte(fmt.Sprintf("Hello, %v!", value))) // creates concatination
}

func Bad(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func PostParam(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("I got message:\n%v", string(body))))
}

func PostHeaders(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	a := headers["A"]
	b := headers["B"]

	f, err := strconv.Atoi(a[0])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	s, err := strconv.Atoi(b[0])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Add("a+b", fmt.Sprintf("%v", f+s))
}