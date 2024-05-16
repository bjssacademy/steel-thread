package postgres

import (
	"acme/model"
	//"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(connectionString string) error {

    db, err := sqlx.Open("postgres", connectionString)
    if err != nil {
        return fmt.Errorf("error connecting to the database: %w", err)
    }
    DB = db

     // Run migrations
     err = goose.Up(db.DB, "./migrations")
     //err = goose.Run("/migrations", db.DB, "postgres")
     if err != nil {
        panic(err)
     }

    // Ping DB to check connection is successful
    if err := DB.Ping(); err != nil {
        return fmt.Errorf("error pinging the database: %w", err)
    }

    fmt.Println("Successfully connected to the database!")
    return nil
}

func GetUsers() ([]model.User, error) {

    users := []model.User{}

    err := sqlx.Select(DB, &users, "SELECT * FROM users")
    if err != nil {
        fmt.Println("Error querying the database:", err)
        return []model.User{}, errors.New("Database could not be queried")
    }

	return users, nil
}

func AddUser(user model.User) (id int, err error) {

    err = DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&id)
    if err != nil {
        fmt.Println("Error inserting user into the database:", err)
        return 0, errors.New("Could not insert user")
    }

    return id, nil   
}