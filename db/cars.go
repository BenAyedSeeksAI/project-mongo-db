package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Car struct {
	Brand      string `bson:"brand"`
	Model      string `bson:"model"`
	BodyWork   string `bson:"body_work"`
	GearType   string `bson:"gear_type"`
	HorsePower int64  `bson:"horsepower"`
}

const (
	AUTOMOBILE_DB   = "automobile"
	CARS_COLLECTION = "cars"
)

func (c Car) DBGetCollection() (*mongo.Client, context.Context, *mongo.Collection, error) {
	client, ctx, err := DBOpenClient()
	if err != nil {
		log.Fatal("Failed to Open database", err.Error())
		return client, ctx, nil, err
	}
	collection := client.Database(AUTOMOBILE_DB).Collection(CARS_COLLECTION)
	return client, ctx, collection, nil
}
func DBGetCars() {
	CarObj := Car{}
	client, ctx, collection, err := CarObj.DBGetCollection()
	defer client.Disconnect(ctx)
	defer ctx.Done()
	if err != nil {
		log.Fatal("Failed to Open database", err.Error())
		return
	}

	collCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(collCtx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(collCtx)

	// Iterate over the results
	var cars []Car
	if err = cursor.All(ctx, &cars); err != nil {
		log.Fatal(err)
	}

	for _, car := range cars {
		fmt.Printf("Car: %+v\n", car)
	}
}
func DBInsertCar(carObj Car) {
	itc := Car{}
	client, ctx, collection, err := itc.DBGetCollection()
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatal("Failed to Open database", err.Error())
		return
	}
	collCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertResult, err := collection.InsertOne(collCtx, carObj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Car with ID:", insertResult.InsertedID)
}
func DBDeleteCar(carID string) {
	CarObj := Car{}
	client, ctx, collection, err := CarObj.DBGetCollection()
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatal("Failed to Open database", err.Error())
		return
	}
	filter := bson.M{"_id": carID}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %d document(s)\n", deleteResult.DeletedCount)
}
