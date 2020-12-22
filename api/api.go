package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Homepage(res http.ResponseWriter, req *http.Request)  {
	res.Write([]byte(`{"status" : "OK"}`))
}

func StartServer()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Homepage).Methods("GET")
	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}