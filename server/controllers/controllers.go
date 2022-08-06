package controllers

import (
	"context"
	"fmt"
	"log"
	"server/database"
	"server/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const SECRET = "secret"

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
		ID: primitive.NewObjectID(),
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

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: user.ID.Hex(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claim.SignedString([]byte(SECRET))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login.",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"token": token})
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie {
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour), // Set the expiry time in the past to remove cookie
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetUser(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if err!=nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	userId, _ := primitive.ObjectIDFromHex(claims.Issuer)
	userError := database.Collection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&user)

	if userError!=nil{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found.",
		})
	}

	return c.JSON(user)
}