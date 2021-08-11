package controllers

import (
	connectionhelper "api_golang/db"
	models "api_golang/models"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Sends a GET request and returns the appropriate response.
// PARAM: id
func GetRecord(c *fiber.Ctx) error {
	recordCollection := connectionhelper.MI.DB.Collection(os.Getenv("COLL"))

	// get param
	paramID := c.Params("id")

	// convert parameter to object id
	id, err := primitive.ObjectIDFromHex(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	record := &models.Records{}
	query := bson.D{{Key: "_id", Value: id}}

	recordCollection.FindOne(c.Context(), query).Decode(record)

	err = recordCollection.FindOne(c.Context(), query).Err()

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Record Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"record": record,
		},
	})

}

// Sends a POST request and returns its ID.
// PARAM: Name, DOB, Address, Description
func AddRecord(c *fiber.Ctx) error {

	recordCollection := connectionhelper.MI.DB.Collection(os.Getenv("COLL"))

	data := new(models.Records)

	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	data.CreatedAt = time.Now()

	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
    

	result, err := recordCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Record failed to insert",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "Record inserted successfully",
	})

}

// Sends a DELETE request and returns Success if query is successful.
// PARAM: id
func DeleteRecord(c *fiber.Ctx) error {
	recordCollection := connectionhelper.MI.DB.Collection(os.Getenv("COLL"))

	// get param
	paramID := c.Params("id")

	// convert parameter to object id
	id, err := primitive.ObjectIDFromHex(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// find and delete record
	query := bson.D{{Key: "_id", Value: id}}

	err = recordCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Record Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete record",
			"error":   err,
		})
	}

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Record deleted successfully",
	})
}

// Sends a PUT request and returns Success if query is successful.
// PARAM: id
func UpdateRecord(c *fiber.Ctx) error {
	recordCollection := connectionhelper.MI.DB.Collection(os.Getenv("COLL"))

	// find parameter
	paramID := c.Params("id")

	// convert parameterID to objectId
	id, err := primitive.ObjectIDFromHex(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// var data Request
	data := new(models.Records)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	// updateData
	var dataToUpdate bson.D

	if data.Name != "" {
		// todo.Title = *data.Title
		dataToUpdate = append(dataToUpdate, bson.E{Key: "name", Value: data.Name})
	}

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	// update
	err = recordCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Record Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update todo",
			"error":   err,
		})
	}

	// get updated data
	record := &models.Records{}

	recordCollection.FindOne(c.Context(), query).Decode(record)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"record": record,
		},
	})
}
