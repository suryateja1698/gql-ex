// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
