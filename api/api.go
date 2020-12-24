package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkbhowmick/go-rest-api/auth"
	"github.com/pkbhowmick/go-rest-api/model"
)

var users map[string]model.User

func InitializeDB() {
	users = make(map[string]model.User)
	users["1"] = model.User{
		"1",
		"Pulak",
		"Kanti",
		[]model.Repository{
			{
				"1001",
				"go-rest-api",
				"public",
				1,
				time.Now(),
			},
		},
		time.Now(),
	}
	users["2"] = model.User{
		"2",
		"Mehedi",
		"Hasan",
		[]model.Repository{
			{
				"1002",
				"go-api-server",
				"public",
				2,
				time.Now(),
			},
		},
		time.Now(),
	}
	users["3"] = model.User{
		"3",
		"Prangon",
		"Majumdar",
		[]model.Repository{
			{
				"1003",
				"go-http-api-server",
				"private",
				3,
				time.Now(),
			},
		},
		time.Now(),
	}
	users["4"] = model.User{
		"4",
		"Sakib",
		"Alamin",
		[]model.Repository{
			{
				"1004",
				"go-httpapi-server",
				"private",
				5,
				time.Now(),
			},
		},
		time.Now(),
	}
	users["5"] = model.User{
		"5",
		"Sahadat",
		"Sahin",
		[]model.Repository{
			{
				"1005",
				"go-http-server",
				"public",
				5,
				time.Now(),
			},
		},
		time.Now(),
	}
}

func userToArray() []model.User {
	items := make([]model.User, 0)
	for _, item := range users {
		items = append(items, item)
	}
	return items
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	allUsers := userToArray()
	err := json.NewEncoder(res).Encode(allUsers)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	if user, ok := users[id]; ok {
		err := json.NewEncoder(res).Encode(user)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	errMsg := "User with id " + id + " doesn't exist"
	http.Error(res, errMsg, http.StatusNotFound)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(res, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if user.ID == "" || user.FirstName == "" || user.LastName == "" {
		http.Error(res, "Missing required fields", http.StatusBadRequest)
		return
	}
	if _, ok := users[user.ID]; ok {
		http.Error(res, "User with given ID already exist", http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()
	user.Repositories = make([]model.Repository, 0)
	users[user.ID] = user
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(res, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	var newUser, oldUser model.User
	oldUser, ok := users[id]
	if !ok {
		http.Error(res, "User doesn't exist", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	oldUser.FirstName = newUser.FirstName
	oldUser.LastName = newUser.LastName
	users[id] = oldUser
	err = json.NewEncoder(res).Encode(&oldUser)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	if user, ok := users[id]; ok {
		delete(users, id)
		err := json.NewEncoder(res).Encode(&user)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(res, "User doesn't exist", http.StatusNotFound)
}

func Homepage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"status" : "OK"}`))
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	token, err := auth.GenerateToken("testuser")
	if err != nil {
		http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
		return
	}
	res.Write([]byte(`{"token" : "` + token + `"}`))
}

func StartServer() {
	InitializeDB()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Homepage).Methods("GET")
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/login", auth.BasicAuth(Login)).Methods("POST")

	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
