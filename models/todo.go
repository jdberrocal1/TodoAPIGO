package models

type Todo struct {
    Name string `json:"name"`
    Description string `json:"description"`
    UserID int `json:"userId"`
    ID int `json:"id"`
}

type Todos []Todo