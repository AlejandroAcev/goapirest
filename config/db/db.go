package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"   

)

func GetDBCollection() (*mongo.Collection, error) {

	var (
       client     *mongo.Client
       mongoURL = "mongodb://localhost:27017"
   )

	//client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	fmt.Println("Creacion de cliente");
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("GoLogin").Collection("users")
	return collection, nil
}