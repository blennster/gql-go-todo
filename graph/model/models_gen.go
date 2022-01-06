// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type TodoFilter struct {
	Done   *bool   `json:"done"`
	Userid *string `json:"userid"`
	Todoid *string `json:"todoid"`
}