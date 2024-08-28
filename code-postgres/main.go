package main

import (
	"acme/api"
	"acme/config"
	"acme/db"
	"acme/db/inmemory"
	"acme/db/postgres"
	"acme/service"
	"fmt"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Access-Control-Allow-Origin", "*")
        // Continue with the next handler
        next.ServeHTTP(writer, request)
    })
}

func main() {
    // Load configuration
    config := config.LoadDatabaseConfig()

    // Initialize database
    dbRepo, err := initializeDatabase(config)
    if err != nil {
        fmt.Println("Error initializing the database:", err)
        return
    }
    defer dbRepo.Close()

    // Initialize services
    userService := service.NewUserService(dbRepo)
    userAPI := api.NewUserAPI(userService)

    // Initialize router
    router := http.NewServeMux()

    // Add routes
    router.HandleFunc("GET /", rootHandler)
    router.HandleFunc("GET /api/users", userAPI.GetUsers)
    router.HandleFunc("POST /api/users", userAPI.CreateUser)
    router.HandleFunc("GET /api/users/{id}", userAPI.GetSingleUser)
    router.HandleFunc("DELETE /api/users/{id}", userAPI.DeleteSingleUser)
    router.HandleFunc("PUT /api/users/{id}", userAPI.UpdateSingleUser)

    // Starting the HTTP server on port 8080
    fmt.Println("Server listening on port 8080...")
    err = http.ListenAndServe(":8080", CorsMiddleware(router))
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func initializeDatabase(config config.DatabaseConfig) (db.Repository, error) {
    switch config.Type {
    case "postgres":
        connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.User, config.DBName, config.Password, config.Host, config.SSLMode)
        return postgres.NewPostgresRepository(connectionString)
    case "inmemory":
        return inmemory.NewInMemoryRepository(), nil
    default:
        return nil, fmt.Errorf("unsupported database type: %s", config.Type)
    }
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello, World!")
}
