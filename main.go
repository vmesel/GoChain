
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/davecgh/go-spew/spew"
)

var bc = NewBlockchain()

func NewBlockAPI(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	params = 
	bc.NewBlock("New Block Added", )
	json.NewEncoder(w).Encode(&bc)
}

func BlockchainAPI(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(&bc)
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/newblock/{previous}", NewBlockAPI).Methods("GET")
	router.HandleFunc("/", BlockchainAPI).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
