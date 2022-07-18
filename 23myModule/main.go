package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MOD")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//run in server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world</h1>"))

}
