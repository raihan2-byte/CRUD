package data

import "time"

type Nama struct {
	ID          int
	Name        string
	Email       string
	Age         int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
}
