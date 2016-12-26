package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "strconv"
	"io/ioutil"
	"io"

    "github.com/gorilla/mux"

    "TodoAPIGO/controllers"
    "TodoAPIGO/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
} 

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users := controllers.GetUsers()
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId,_ := strconv.Atoi(vars["userId"])
    user := controllers.GetUser(userId)   
    if(user.ID != 0){
        json.NewEncoder(w).Encode(user)
    } else {
        fmt.Fprintln(w, "No user found")
    }   
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if (err != nil) {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &user); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    result := controllers.CreateUser(user)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        panic(err)
    }
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId,_ := strconv.Atoi(vars["userId"])
    user := controllers.DeleteUser(userId)   
    if(user != nil){
        fmt.Fprintln(w, "No user found")
    } 
}