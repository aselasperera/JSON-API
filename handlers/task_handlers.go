// handlers/task_handlers.go
package handlers

import (
	"context"
	"json-api/database"
	"json-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tasks []models.Task
	cursor, err := database.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	var task models.Task
	if err := database.TaskCollection.FindOne(ctx, bson.M{"_id": taskID}).Decode(&task); err != nil {
		return c.Status(404).SendString("Task not found")
	}

	return c.JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	task.ID = primitive.NewObjectID()
	result, err := database.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(result)
}

func UpdateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"completed":   task.Completed,
		},
	}

	_, err = database.TaskCollection.UpdateOne(ctx, bson.M{"_id": taskID}, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func DeleteTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	_, err = database.TaskCollection.DeleteOne(ctx, bson.M{"_id": taskID})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
