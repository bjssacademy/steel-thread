package service

import (
	"acme/db"
	"acme/model"
	"fmt"
)

type UserService struct {
    repository db.Repository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo db.Repository) *UserService {
    return &UserService{
        repository: repo,
    }
}

func (s *UserService) GetUsers() ([]model.User, error) {
    users, err := s.repository.GetUsers()
    if err != nil {
        return nil, fmt.Errorf("error getting users from DB: %w", err)
    }
    return users, nil
}

func (s *UserService) DeleteUser(id int) error {
    err := s.repository.DeleteUser(id)
    if err != nil {
        return fmt.Errorf("error deleting user from DB: %w", err)
    }
    return nil
}

func (s *UserService) GetUser(id int) (model.User, error) {
    user, err := s.repository.GetUser(id)
    if err != nil {
        return model.User{}, fmt.Errorf("error getting user from DB: %w", err)
    }
    return user, nil
}

func (s *UserService) UpdateUser(id int, user model.User) (model.User, error) {
    updatedUser, err := s.repository.UpdateUser(id, &user)
    if err != nil {
        return model.User{}, fmt.Errorf("error updating user in DB: %w", err)
    }
    return updatedUser, nil
}

func (s *UserService) CreateUser(user model.User) (int, error) {
    id, err := s.repository.AddUser(user)
    if err != nil {
        return 0, fmt.Errorf("error creating user in DB: %w", err)
    }
    return id, nil
}
