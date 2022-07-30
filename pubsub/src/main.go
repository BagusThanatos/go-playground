package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"encoding/json"
	"io/ioutil"
  //"os"
)

type Data struct {
	att1 string `json:"att1"`
  att2 string `json:"att2"`
}

type RequestBody struct {
	UserID string `json:"userid"`
}

func Ping(w http.ResponseWriter, r *http.Request) {}



func Pubsub(w http.ResponseWriter, r *http.Request) {
  if  r.Method != "POST" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  body, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }

	data := Data{}
	err = json.Unmarshal(body, &data)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  fmt.Printf("Data: %s %s \n", data.att1, data.att2)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	flag.Parse()
	args := flag.Args()
	http.HandleFunc("/ping", Ping)
  http.HandleFunc("/postgres", Pubsub)
	port := "8080"
	if len(args) > 0 {
		port = args[0]
	}

	fmt.Printf("Listening in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
