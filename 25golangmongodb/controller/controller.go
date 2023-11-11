package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://suman:password@localhost:27017/?authMechanism=DEFAULT"

const dbName = "netflix"
const colName = "watchlist"

// imp
var collection *mongo.Collection

// Special Method
func init() {

	// Client options
	clientOptions := options.Client().ApplyURI(connectionString)

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
