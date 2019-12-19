package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"Readr/server/models"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING = "mongodb+srv://xd:passwordasd@cluster0-3mnec.mongodb.net/test?retryWrites=true&w=majority"
const dbName = "test"
const collName = "todolist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(CONNECTION_STRING)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(collName)
}

