package models

import (
	"time"
)

type Records struct {
	Name        string    `bson:"name"`
	DOB         string    `bson:"dob"`
	Address     string    `bson:"address"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"createdAt"`
}
