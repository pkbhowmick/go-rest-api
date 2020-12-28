# REST API Server in Golang

- This API server provides endpoints to create,read,update & delete users and their repositories (like Github).
  
[![Go Report Card](https://goreportcard.com/badge/github.com/pkbhowmick/go-rest-api)](https://goreportcard.com/report/github.com/pkbhowmick/go-rest-api)

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

|  Method | API Endpoint  | Authentication Type | Description |
|---|---|---|---|
|POST| /api/login | Basic | Return jwt token in response for successful authentication
|GET| /api/users | Basic or Bearer token | Return a list of all users in response| 
|GET| /api/users/{id} | Basic or Bearer token| Return the data of given user id in response| 
|POST| /api/users | Basic or Bearer token |Add an user in the database and return the added user data in response | 
|PUT| /api/users/{id} | Basic or Bearer token |Update the user and return the updated user info in response| 
|DELETE| /api/users/{id} | Basic or Bearer token |Delete the user and return the deleted user data in response| 

## Available Flags

| Flag | Shorthand | Default value | Example | Description
|---|---|---|---|---|
|port|p|8080| go-rest-api start --port=8090 | Start API server in the given port otherwise in default port
|auth|a|true| go-rest-api start --auth=false | If true impose authentication on API server otherwise bypass it

## Sample Curl commands without authentication

```$ go-rest-api start --port=8080 --auth=false``` [Run API server without authentication]

```$ curl -X GET http://localhost:8080/api/users``` [Get all users information]

```$ curl -X GET http://localhost:8080/api/users/1``` [Get user information with id 1]

```$ curl -X POST  -H "Content-Type:application/json" -d '{"id":"6","firstName":"testfirst","lastName":"testlast"}' http://localhost:8080/api/users``` [Create user with given id]

```$ curl -X PUT  -H "Content-Type:application/json" -d '{"firstName":"test","lastName":"test"}' http://localhost:8080/api/users/1 ``` [Modify user data with gigen id]

```$ curl -X DELETE http://localhost:8080/api/users/1``` [Delete user with given id]

## Sample Curl commands with Basic authentication

```$ export ADMIN_USER=admin```

```$ export ADMIN_PASS=demo```

```$ go-rest-api start --port=8080 --auth=true``` [Run API server with authentication]

```$ curl -X GET --user admin:demo http://localhost:8080/api/users``` [Get all users information]

```$ curl -X GET --user admin:demo http://localhost:8080/api/users/1``` [Get user information with id 1]

```$ curl -X POST  --user admin:demo -H "Content-Type:application/json" -d '{"id":"6","firstName":"testfirst","lastName":"testlast"}' http://localhost:8080/api/users``` [Create user with given id]

```$ curl -X PUT  --user admin:demo -H "Content-Type:application/json" -d '{"firstName":"test","lastName":"test"}' http://localhost:8080/api/users/1 ``` [Modify user data with gigen id]

```$ curl -X DELETE --user admin:demo http://localhost:8080/api/users/1``` [Delete user with given id]

## Sample Curl commands with Bearer token(JWT token) authentication

```$ export ADMIN_USER=admin```

```$ export ADMIN_PASS=demo```

```$ export SIGNING_KEY=veryverysecretkey```

```$ go-rest-api start --port=8080 --auth=true``` [Run API server with authentication]

```$ curl -X POST --user admin:demo  http://localhost:8080/api/login``` [Get jwt token via login with basic authentication]

```$ curl -X GET -H "Authorization: Bearer <jwt_token>"  http://localhost:8080/api/users``` [Get all users information]

```$ curl -X GET -H "Authorization: Bearer <jwt_token>" http://localhost:8080/api/users/1``` [Get user information with id 1]

```$ curl -X POST -H "Authorization: Bearer <jwt_token>" -H "Content-Type:application/json" -d '{"id":"6","firstName":"testfirst","lastName":"testlast"}' http://localhost:8080/api/users``` [Create user with given id]

```$ curl -X PUT -H "Authorization: Bearer <jwt_token>" -H "Content-Type:application/json" -d '{"firstName":"test","lastName":"test"}' http://localhost:8080/api/users/1 ``` [Modify user data with gigen id]

```$ curl -X DELETE -H "Authorization: Bearer <jwt_token>" http://localhost:8080/api/users/1``` [Delete user with given id]