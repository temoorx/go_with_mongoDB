package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://testGo:taimoor@test.dhhl2xd.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchList"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection is ready")
}
