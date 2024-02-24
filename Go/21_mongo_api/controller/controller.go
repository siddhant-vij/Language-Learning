package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go-learning-21/model"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName  = "netflix"
	colName = "watchlist"
)

var collection *mongo.Collection

func loadURI() string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("You must set your 'MONGODB_URI' environment variable.")
	}

	return uri
}

func init() {
	uri := loadURI()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(colName)
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		fmt.Println("Error inserting movie:", err)
	}

	fmt.Println("Inserted movie with ID:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("Error converting ID:", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Error updating movie:", err)
	}

	fmt.Println("Updated movie count:", updated.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("Error converting ID:", err)
	}

	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error deleting movie:", err)
	}

	fmt.Println("Deleted movie count:", deleted.DeletedCount)
}

func deleteAllMovies() int64 {
	deleted, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error deleting all movies:", err)
	}

	fmt.Println("Deleted movie count:", deleted.DeletedCount)
	return deleted.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error getting all movies:", err)
	}

	var movies []primitive.M

	for cursor.Next(context.TODO()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			fmt.Println("Error decoding movie:", err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.TODO())
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode("" + strconv.Itoa(int(count)) + " movies deleted")
}
