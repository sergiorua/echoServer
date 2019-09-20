package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error decoding")
		return
	}
	dynamic := make(map[string]interface{})
	json.Unmarshal([]byte(reqBody), &dynamic)

	retData, err := json.Marshal(dynamic)
	if err != nil {
		fmt.Println("Error decoding")
		return
	}
	fmt.Fprintf(w, string(retData))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getAllEvents)
	log.Fatal(http.ListenAndServe(":8080", router))
}
