package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	Title string `json:"title"`
}

type RequestBody struct {
	UserID string `json:"userid"`
}

func HelloName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	requestBody := RequestBody{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(requestBody.UserID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := Data{}
	data.Title = "HELLO, " + requestBody.UserID
	payload, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}


func Hello(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	data.Title = "HELLO WORLD"
	payload, err := json.Marshal(data)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func main() {
	flag.Parse()
	args := flag.Args()
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/helloname", HelloName)
	port := "8080"
	if len(args) > 0 {
		port = args[0]
	}

	fmt.Printf("Listening in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}