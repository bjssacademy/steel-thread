package postgres

import (
	"acme/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
    DB *sqlx.DB
}

func NewPostgresRepository(connectionString string) (*PostgresRepository, error) {

    db, err := sqlx.Open("postgres", connectionString)
    if err != nil {
        return nil, fmt.Errorf("error connecting to the database: %w", err)
    }

	// Run migrations
    err = goose.Up(db.DB, "./migrations")
    if err != nil {
        panic(err)
    }

    // Ping DB to check connection is successful
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error pinging the database: %w", err)
    }

    fmt.Println("Successfully connected to the database!")
    return &PostgresRepository{DB: db}, nil
}

// GetUsers retrieves all users from the database.
func (repo *PostgresRepository) GetUsers() ([]model.User, error) {
    users := []model.User{}

    err := sqlx.Select(repo.DB, &users, "SELECT * FROM users")
    if err != nil {
        fmt.Println("Error querying the database:", err)
        return []model.User{}, errors.New("Database could not be queried")
    }

	return users, nil
}

// GetUser retrieves a user by ID from the database.
func (repo *PostgresRepository) GetUser(id int) (model.User, error) {
    // Implement logic to fetch a user by ID from the database
    return model.User{}, nil
}

// AddUser adds a new user to the database.
func (repo *PostgresRepository) AddUser(user model.User) (id int, err error) {
    err = repo.DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&id)
    if err != nil {
        fmt.Println("Error inserting user into the database:", err)
        return 0, errors.New("Could not insert user")
    }

    return id, nil
}

// UpdateUser updates an existing user in the database.
func (repo *PostgresRepository) UpdateUser(id int, user *model.User) (model.User, error) {
    // Implement logic to update a user in the database
    return model.User{}, nil
}

// DeleteUser deletes a user from the database.
func (repo *PostgresRepository) DeleteUser(id int) error {
    // Implement logic to delete a user from the database
    return nil
}

func (repo *PostgresRepository) Close() {

}