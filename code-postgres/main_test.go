package main

import (
	"acme/api"
	"acme/config"
	"acme/db/inmemory"
	"acme/model"
	"acme/service"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRootHandler(t *testing.T) {
    //ARRANGE
	// Create a new HTTP request
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Create a handler
    handler := http.HandlerFunc(rootHandler)

	//ACT
    // Serve the request
    handler.ServeHTTP(rr, req)

	//ASSERT
    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := "Hello, World!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestGetUsersHandler(t *testing.T) {
    // ARRANGE
	// Create a new HTTP request
    req, err := http.NewRequest("GET", "/api/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Create a handler

    //We no longer create the mock directly, instead we load an env file
    // This is nice, because we can just change the .env file and point it at an actual DB
    //or in memory, or whatever

    //dbRepo := inmemory.NewInMemoryRepository();
    config := config.LoadDatabaseConfig(".env.inmemory")
    dbRepo, err := initializeDatabase(config)
    if err != nil {
        t.Fatalf("Error initializing the database: %v", err)
        return
    }
    defer dbRepo.Close()

    userService := service.NewUserService(dbRepo)
    userAPI := api.NewUserAPI(userService)
    handler := http.HandlerFunc(userAPI.GetUsers)

	//Arrange our expected response
	expected := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	//ACT
    // Serve the request
    handler.ServeHTTP(rr, req)

	//ASSERT
    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
	var actual []model.User
    if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }

}

func TestRootHandlerWithServer(t *testing.T) {
    //ARRANGE
	// Create a new server with the handler
    server := httptest.NewServer(http.HandlerFunc(rootHandler))
    defer server.Close()

	//ACT
    // Send a GET request to the server
    resp, err := http.Get(server.URL + "/")
    if err != nil {
        t.Fatalf("Failed to send GET request: %v", err)
    }
    defer resp.Body.Close()

    // Check the status code
    if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := "Hello, World!"
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }
    if string(bodyBytes) != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", string(bodyBytes), expected)
    }
}

func TestGetUsersHandlerWithServer(t *testing.T) {
    //ARRANGE
	// Create a new server with the handler
    dbRepo := inmemory.NewInMemoryRepository();
    userService := service.NewUserService(dbRepo)
    userAPI := api.NewUserAPI(userService)

    server := httptest.NewServer(http.HandlerFunc(userAPI.GetUsers))
    defer server.Close()

	expected := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}

	//ACT
    // Send a GET request to the server
    resp, err := http.Get(server.URL + "/api/users")
    if err != nil {
        t.Fatalf("Failed to send GET request: %v", err)
    }
    defer resp.Body.Close()

	//ASSERT
    // Check the status code
    if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Unmarshal the response body
    var actual []model.User
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }
	
    err = json.Unmarshal(bodyBytes, &actual)
    if err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    // Compare the actual response with the expected response using deepequal
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }
}