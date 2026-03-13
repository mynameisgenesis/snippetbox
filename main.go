package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet..."))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
