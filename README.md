# REST API Server in Golang

- This API server provides endpoints to create,read,update & delete users and their repositories (like Github).

## To Start API Server
```$ git clone https://github.com/pkbhowmick/go-rest-api.git```

```$ cd go-rest-api```

```$ go install```

```$ go-rest-api start```

## Command to run unit test for API endpoints
```$ cd api```

```$ go test```

## Data Model

- User Model
``````
type User struct {
	ID           string       `json:"id"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Repositories []Repository `json:"repositories"`
	CreatedAt    time.Time    `json:"createdAt"`
}
``````
- Repository Model
``````
type Repository struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Visibility string    `json:"visibility"`
	Star       int       `json:"star"`
	CreatedAt  time.Time `json:"createdAt"`
}
``````

## Available API Endpoints

|  Method | API Endpoint  | Description |
|---|---|---|
|GET| /api/users | Return a list of all users in response| 
|GET| /api/users/{id} | Return the data of given user id in response| 
|POST| /api/users | Add an user in the database and return the added user data in response | 
|PUT| /api/users/{id} | Update the user and return the updated user info in response| 
|DELETE| /api/users/{id} | Delete the user and return the deleted user data in response| 