package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeeter() {

	fmt.Println("Hey there mod users")

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to go lang</h1>"))
}
