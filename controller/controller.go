package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/mongodb/mongo-go-driver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//insert a record

func inserOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a movie in db with id", inserted.InsertedID)
}

//update a record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"watched": true
	}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("modified count", result.ModifiedCount)
}

//delete a record
func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deleteCount , err :=collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted count: ", deleteCount)
}

//delete all record

func deleteAllMovie(movieId string) int64{

	deleteResult , err :=collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
   fmt.Println("Number of movies delete", deleteResult.DeletedCount)
    return deleteResult.DeletedCount

}