package models

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Records struct {
    ID           primitive.ObjectID `bson:"_id"`
    Name    string             `bson:"name"`
    DOB  string             `bson:"dob"`
    Address string             `bson:"address"`
	Description string		`bson:"description"`
	CreatedAt time.Time		`bson:createdAt`
}

// {
// 	"id": "xxx", // user ID
// 	"name": "test", // user name
// 	"dob": "", // date of birth
// 	"address": "", // user address
// 	"description": "", // user description
// 	"createdAt": "" // user created date
// 	}