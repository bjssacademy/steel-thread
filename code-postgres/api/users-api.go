package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserAPI struct {
    userService *service.UserService
}

func NewUserAPI(userService *service.UserService) *UserAPI {
    return &UserAPI{
        userService: userService,
    }
}

func (api *UserAPI) UpdateSingleUser(writer http.ResponseWriter, request *http.Request) {
    id, err := api.parseID(request.PathValue("id"))
    if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    var user model.User
    err = json.NewDecoder(request.Body).Decode(&user)
    if err != nil {
        http.Error(writer, "Bad Request Body", http.StatusBadRequest)
        return
    }

    updatedUser, err := api.userService.UpdateUser(id, user)
    if err != nil {
        http.Error(writer, "User not found to update", http.StatusNotFound)
        return
    }

    json.NewEncoder(writer).Encode(updatedUser)
}

func (api *UserAPI) DeleteSingleUser(writer http.ResponseWriter, request *http.Request) {
    id, err := api.parseID(request.PathValue("id"))
    if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    err = api.userService.DeleteUser(id)
    if err != nil {
        http.Error(writer, "Could not delete user", http.StatusBadRequest)
        return
    }

    writer.WriteHeader(http.StatusOK)
}

func (api *UserAPI) GetSingleUser(writer http.ResponseWriter, request *http.Request) {
    id, err := api.parseID(request.PathValue("id"))
    if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    user, err := api.userService.GetUser(id)
    if err != nil {
        http.Error(writer, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(writer).Encode(user)
}

func (api *UserAPI) CreateUser(writer http.ResponseWriter, request *http.Request) {
    var user model.User
    err := json.NewDecoder(request.Body).Decode(&user)
    if err != nil {
        http.Error(writer, "Bad Request Body", http.StatusBadRequest)
        return
    }

    id, err := api.userService.CreateUser(user)
    if err != nil {
        http.Error(writer, "User not created", http.StatusBadRequest)
        return
    }

    writer.WriteHeader(http.StatusCreated)
    fmt.Fprintf(writer, "User created successfully: %d", id)
}

func (api *UserAPI) GetUsers(writer http.ResponseWriter, request *http.Request) {
    users, err := api.userService.GetUsers()
    if err != nil {
        http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(writer).Encode(users)
}

func (api *UserAPI) parseID(idStr string) (int, error) {
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        return 0, err
    }
    return id, nil
}
