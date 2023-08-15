package data

import "encoding/json"

type User struct {
	Name        string      `json:"name" binding:"required"`
	Email       string      `json:"email" binding:"required"`
	Age         json.Number `json:"age" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
}
