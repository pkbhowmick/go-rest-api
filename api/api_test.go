package api

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
)

type Test struct {
	Method             string
	Url                string
	Body               io.Reader
	ExpectedStatusCode int
}

type TestWithID struct {
	Method             string
	Url                string
	Body               io.Reader
	ExpectedStatusCode int
	UserID             string
}

func TestCreateUser(t *testing.T) {
	InitializeDB()
	tests := []Test{
		{
			"POST",
			"/api/users",
			bytes.NewReader([]byte(`{"id": "6","firstName": "test","lastName": "test"}`)),
			201,
		},
		{
			"POST",
			"/api/users",
			bytes.NewReader([]byte(`{"id": "10"}`)),
			400,
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateUser)
		handler.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	InitializeDB()
	tests := []TestWithID{
		{"DELETE", "/api/users", nil, 200, "1"},
		{"DELETE", "/api/users", nil, 404, "6"},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		params := make(map[string]string)
		params["id"] = test.UserID
		req = mux.SetURLVars(req, params)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(DeleteUser)
		handler.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestGetUser(t *testing.T) {
	InitializeDB()
	tests := []TestWithID{
		{"GET", "/api/users", nil, 200, "1"},
		{"GET", "/api/users", nil, 404, "8"},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		params := make(map[string]string)
		params["id"] = test.UserID
		req = mux.SetURLVars(req, params)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetUser)
		handler.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestGetUsers(t *testing.T) {
	InitializeDB()
	tests := []Test{
		{"GET", "/api/users", nil, 200},
		{"GET", "/api/users/", nil, 200},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetUsers)
		handler.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	InitializeDB()
	tests := []TestWithID{
		{
			"PUT",
			"/api/users",
			bytes.NewReader([]byte(`{"firstName": "test","lastName": "test"}`)),
			200,
			"1",
		},
		{
			"PUT",
			"/api/users",
			bytes.NewReader([]byte(`{"firstName": "test","lastName": "test"}`)),
			404,
			"10",
		},
		{
			"PUT",
			"/api/users",
			bytes.NewReader([]byte(`"firstName": "test","lastName": "test"`)),
			400,
			"2",
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		params := make(map[string]string)
		params["id"] = test.UserID
		req = mux.SetURLVars(req, params)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(UpdateUser)
		handler.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}
