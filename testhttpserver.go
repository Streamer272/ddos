package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting server...\n")

	http.HandleFunc("/", Tmp)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	log.Printf("Got request %v %v\n", data, err)
	io.WriteString(w, "version 1")
}
