package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Add(value1 int, value2 int) int {
	return value1 + value2
}

func Subtract(value1 int, value2 int) int {
	return value1 - value2
}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

//calculate  return x+2
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	Add(1, 2)
	Subtract(3, 5)
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe("1234", router))
}
