package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	dbName      = "db_test_cron"
	collection  = "users"
)

type Data struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func main() {

	mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client

	e := echo.New()

	//CRON Scheduler

	c := cron.New()
	c.AddFunc("*/1 * * * *", insertData)
	c.Start()

	e.Logger.Fatal(e.Start(":8080"))

}

func insertData() {
	collection := mongoClient.Database(dbName).Collection(collection)
	ctx := context.TODO()

	data := bson.M{
		"name": "John Doe",
		"age":  30,
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println("Error insert data", err)
	} else {
		log.Println("Success insert data")
	}
}
