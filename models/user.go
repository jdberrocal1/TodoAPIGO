package models

type User struct {
    Name string `json:"name"`
    ID int `json:"id"`
    Todos Todos `json:"todos"`
}

type Users []User 