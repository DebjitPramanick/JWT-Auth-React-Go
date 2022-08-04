package controllers

import (
	"context"
	"fmt"
	"log"
	"server/database"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CheckHealth(c *fiber.Ctx) error {
	return c.SendString("APP IS RUNNING.")
}

func Register(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	pass, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User {
		Name: data["name"],
		Password: pass,
		Email: data["email"],
	}

	result, err := database.Collection.InsertOne(context.TODO(), &user)

	if err!=nil{
		log.Fatal(err)
		return err
	}
	fmt.Println("Created user with email & ID: ", user.Email, result.InsertedID)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User
	userError := database.Collection.FindOne(context.TODO(), bson.M{"email": data["email"]}).Decode(&user)

	if userError == mongo.ErrNoDocuments {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password is not correct.",
		})
	}

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}


	return c.JSON(data)
}