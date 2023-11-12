package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suman/25golangmongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const connectionString = "mongodb://suman:password@localhost:27017/?authMechanism=DEFAULT"

const dbName = "netflix"
const colName = "watchlist"

// imp
var collection *mongo.Collection

// Special Method
func init() {

	// credential := options.Credential{
	// 	AuthSource: "netflix",
	// 	Username:   "suman",
	// 	Password:   "password",
	// }
	// Client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	clientOptions := options.Client().ApplyURI("mongodb://suman:password@localhost:27017/?authMechanism=SCRAM-SHA-1")

	// Create connection
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb connection success")
	collection = client.Database(dbName).Collection(colName)

	// If collection is ready
	fmt.Println("Collection instance is ready")
}

// MongoDB Helpers
func insertOneRecord(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted : ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated one movie on db with id : ", id)
	fmt.Println("Updatecount: ", res.ModifiedCount)

}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Single recorded deleted with id : ", id)

}

func deleteAllMovies() int64 {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	print("Deleted all movies , count : ", res.DeletedCount)
	return res.DeletedCount
}

func getAllMovies() []primitive.M {
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var movies []primitive.M

	for cursor.Next(ctx) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	fmt.Println("All movies, count: ", movies)
	return movies
}

// Actual Controller File
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	insertOneRecord(movie)

	json.NewEncoder(w).Encode(movie)

}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["movieId"])

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["movieId"])
	json.NewEncoder(w).Encode(params["movieId"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
