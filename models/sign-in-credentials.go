package models

type SignInCredentials struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
