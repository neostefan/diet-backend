package models

//calories, protein, cost accepted constraint values

type UserConstraints struct {
	Min string `json:"minimum" binding:"required"`
	Max string `json:"maximum" binding:"required"`
}
