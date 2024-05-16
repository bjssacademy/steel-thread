package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func UpdateSingleUser(writer http.ResponseWriter, request *http.Request) {

    id, err := parseId(request.PathValue("id"))

    if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    user, err := decodeUser(request.Body)

    if err != nil {
        http.Error(writer, "Bad Request Body", http.StatusBadRequest)
        return
    }

    updated, err := service.UpdateUser(id, user)

    if err != nil {
        http.Error(writer, "User not found to update", http.StatusNotFound)
        return
    }

    json.NewEncoder(writer).Encode(updated)

}

func DeleteSingleUser(writer http.ResponseWriter, request *http.Request) {

    id, err := parseId(request.PathValue("id"))

    if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    err = service.DeleteUser(id)
    
    if err != nil {
        http.Error(writer, "Could not delete user", http.StatusBadRequest)
        return
    }

    writer.WriteHeader(http.StatusOK)

}

func GetSingleUser(writer http.ResponseWriter, request *http.Request) {
    
    id, err := parseId(request.PathValue("id"))
    
    if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    user, err := service.GetUser(id)

    if err != nil {
        http.Error(writer, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(writer).Encode(user)

}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
    
    user, err := decodeUser(request.Body)

    if err != nil {
        http.Error(writer, "Bad Request Body", http.StatusBadRequest)
        return
    }

    id, err := service.CreateUser(user)

    if err != nil {
        http.Error(writer, "User not created", http.StatusBadRequest)
        return
    }

    writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "User created successfully: %d", id)

}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
    
    users, err := service.GetUsers()

    if err != nil {
        http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(writer).Encode(users)
    
}

func parseId(idStr string) (id int, err error){
    
    id, err = strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        return 0, err
    }

    return id, nil

}

func decodeUser(body io.ReadCloser) (user model.User, err error) {
    
    err = json.NewDecoder(body).Decode(&user)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        return model.User{}, err
    }

    return user, nil
}