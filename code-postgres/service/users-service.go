package service

import (
	"acme/db"
	"acme/model"
	"acme/postgres"
	"errors"
	"fmt"
)

func GetUsers() ([]model.User, error) {

	users, err := postgres.GetUsers()

	if err != nil {
		fmt.Println("Error getting users from DB:", err)
		return nil, errors.New("There was an error getting the users from the database.")
	}

	return users, nil

}

func DeleteUser(id int) error {
	err := db.DeleteUser(id)

	if err != nil{
		fmt.Println("Error deleting user from DB:", err)
		return errors.New("Could not delete user")
	}

	return nil
}

func GetUser(id int) (user model.User, err error){
	user, err = db.GetUser(id)

	if err != nil{
		fmt.Println("Error getting user from DB:", err)
		return model.User{}, errors.New("Could not find user")
	}

	return user, nil

}

func UpdateUser(id int, user model.User) (updateUser model.User, err error){
	updated, err := db.UpdateUser(id, user)

	if err != nil {
		fmt.Println("Error updating user in DB:", err)
		return model.User{}, errors.New("Could not update user")
	}

	return updated, nil

}

func CreateUser(user model.User) (id int, err error) {
	id, err = postgres.AddUser(user)

	if err != nil {
		fmt.Println("Error creating user in DB:", err)
		return 0, errors.New("Could not create user")
	}

	return id, nil
}