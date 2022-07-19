package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/Anishadahal/mongoAPI/controller"
	"github.com/Anishadahal/mongoAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

//insert one record
func InsertOneMovie(movie model.Netflix) {
	inserted, err := controller.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted values are: ", inserted.InsertedID)
}

//update one record
func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := controller.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

//delete one record
func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := controller.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count: ", deleteCount)
}

//delete all records from mongoDB
func DeleteAllMovies() int64 {
	filter := bson.D{{}}
	deleteResult, err := controller.Collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted.", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all movies from mongoDB
func GetAllMovies() []primitive.M {
	cursor, err := controller.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

//get one movie by id
func GetOneMovie(movieId string) model.Netflix {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	var result model.Netflix

	filter := bson.M{"_id": id}
	err = controller.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched movie: ", result)
	return result
}
