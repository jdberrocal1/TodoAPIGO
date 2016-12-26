package controllers

import (
    "fmt"

    "TodoAPIGO/models"
)

var currentId int

var users = models.Users{}

// Give us some seed data
func init() {
    fmt.Println("Initializing User Controller")
}

func GetUsers() models.Users {
    return users
}

func GetUser(id int) models.User {
    for _, user := range users {
        if user.ID == id {
            return user
        }
    }
    // return empty user if not found
    return models.User{}
}

func CreateUser(user models.User) models.User {
    currentId += 1
    user.ID = currentId
    users = append(users, user)
    return user
}

func DeleteUser(id int) error {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find User with id of %d to delete", id)
}