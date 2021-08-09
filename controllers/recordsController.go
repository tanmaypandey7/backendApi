package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateRecord - Insert a new document in the collection.
func CreateRecord(task Records) error {
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

// MarkCompleted - MarkCompleted
func MarkCompleted(code string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)

	//Perform UpdateOne operation & validate against the error.
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func DeleteOne(code string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ISSUES)
	//Perform DeleteOne operation & validate against the error.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
