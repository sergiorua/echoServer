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
		fmt.Fprintf(w, `{"result":"","error":%q}`, err)
		return
	}
	dynamic := make(map[string]interface{})
	json.Unmarshal([]byte(reqBody), &dynamic)

	retData, err := json.Marshal(dynamic)
	if err != nil {
		fmt.Fprintf(w, `{"result":"","error":%q}`, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(retData))

	log.Println(string(retData))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getAllEvents)
	log.Fatal(http.ListenAndServe(":8080", router))
}
