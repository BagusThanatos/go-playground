package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"encoding/json"
)

type Data struct {
	Title string `json:"title"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	data.Title = "HELLO WORLD"
	payload, err := json.Marshal(data)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(data.Title)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func main() {
	flag.Parse()
	args := flag.Args()
	http.HandleFunc("/hello", Hello)

	port := "8080"
	if len(args) > 0 {
		port = args[0]
	}

	fmt.Printf("Listening in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}