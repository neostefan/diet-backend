package models

type User struct {
	Id        int8
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"required,min=1,max=100"`
	Password  string `json:"password" binding:"required,min=8,max=20"`
	Allergies string `json:"allergies"`
	Condition string `json:"condition"`
}
