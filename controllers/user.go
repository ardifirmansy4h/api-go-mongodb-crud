package controllers

import (
	"context"
	"fmt"
	"go-mongodb/db"
	"go-mongodb/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(c echo.Context) error {
	client := db.ConnectMongo()
	defer client.Disconnect(context.TODO())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var tb = db.ClientMongo("users", client)
	_, err := tb.InsertOne(ctx, models.User{
		Name:      "Ardi Wibowo",
		Age:       "25",
		CreatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println("error insert")
	}
	return c.String(http.StatusOK, "Success input data")
}

func Update(c echo.Context) error {
	
	client := db.ConnectMongo()
	defer client.Disconnect(context.TODO())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	col := db.ClientMongo("users", client)
	id, _ := primitive.ObjectIDFromHex("63ebd4709724b3d10097f76a")
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set" : bson.M{"age" : "30"},
	}
	result, _ := col.UpdateOne(ctx, filter, update)
	return c.JSON(http.StatusOK, result)
	
}

func Delete(c echo.Context) error {
	age := c.QueryParam("age")
	client := db.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	col := db.ClientMongo("users", client)
	result, _ := col.DeleteOne(
		ctx, bson.M{
			"age": age,
		},
	)

	return c.JSON(http.StatusOK, result.DeletedCount)
}

func Get(c echo.Context) error {
	age := c.QueryParam("age")
	client := db.ConnectMongo()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var users = make([]models.User, 0)
	col := db.ClientMongo("users", client)
	result, err := col.Find(
		ctx, bson.M{
			"age": age,
		},
	)
	if err != nil {
		return err
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var user models.User
		err = result.Decode(&user)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return c.JSON(http.StatusOK, users)
}