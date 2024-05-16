package main

import (
	"acme/api"
	"acme/postgres"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Access-Control-Allow-Origin", "*")
        // Continue with the next handler
        next.ServeHTTP(writer, request)
    })
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello, World!")
}

func getConnectionString() (string, error) {
    env := os.Getenv("APP_ENV")
    var connString string

    switch env {
    case "development":
        connString = os.Getenv("DEV_DB_CONN_STRING")
    case "production":
        connString = os.Getenv("PROD_DB_CONN_STRING")
    default:
        return connString, errors.New("set APP_ENV, DEV_DB_CONN_STRING, and PROD_DB_CONN_STRING")
    }
    return connString, nil
}

func main() {

    // Initialize the database connection
    connectionString, connStrErr := getConnectionString()
    if connStrErr != nil {
        fmt.Println(connStrErr.Error())
        return
    }

    if err := postgres.InitDB(connectionString); err != nil {
        fmt.Println("Error initializing the database:", err)
        return
    }
    defer postgres.DB.Close()

    router := http.NewServeMux()

    router.HandleFunc("GET /", rootHandler)
    router.HandleFunc("GET /api/users", api.GetUsers)
    router.HandleFunc("POST /api/users", api.CreateUser)
    router.HandleFunc("GET /api/users/{id}", api.GetSingleUser)
    router.HandleFunc("DELETE /api/users/{id}", api.DeleteSingleUser)
    router.HandleFunc("PUT /api/users/{id}", api.UpdateSingleUser)
    
    // Starting the HTTP server on port 8080
    fmt.Println("Server listening on port 8080...")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

