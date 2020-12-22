package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkbhowmick/go-rest-api/model"
	"log"
	"net/http"
	"time"
)

var users map[string]model.User

func initializeDB()  {
	users = make(map[string]model.User)
}

func userToArray() []model.User{
	items := make([]model.User,0)
	for _,item := range users{
		items = append(items, item)
	}
	return items
}

func GetUsers(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	allUsers := userToArray()
	err := json.NewEncoder(res).Encode(allUsers)
	if err!=nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUser(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	id := params["id"]
	if user, ok := users[id]; ok {
		err := json.NewEncoder(res).Encode(user)
		if err!=nil{
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(res, "User doesn't exist", http.StatusNotFound)
}

func CreateUser(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if  err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if user.ID == "" || user.FirstName == "" || user.LastName == "" {
		http.Error(res, "Missing required fields", http.StatusBadRequest)
		return
	}
	if _, ok := users[user.ID]; ok {
		http.Error(res, "User with given ID already exist",http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()
	users[user.ID] = user
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(&user)
	if err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUser(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
}

func DeleteUser(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
}

func Homepage(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	res.Write([]byte(`{"status" : "OK"}`))
}

func StartServer()  {
	initializeDB()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",Homepage).Methods("GET")
	router.HandleFunc("/api/users",GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}",GetUser).Methods("GET")
	router.HandleFunc("/api/users",CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}",UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}",DeleteUser).Methods("DELETE")

	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}