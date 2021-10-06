package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnimalDoc struct {
	Name    string            `bson:"name,omitempty"`
	Age     int               `bson:"age,omitempty"`
	Species string            `bson:"species"`
	Size    map[string]string `bson:"size"`
}

var (
	err        error
	trash      interface{}
	client     *mongo.Client
	ctx        context.Context
	collection *mongo.Collection
)

func ConnectDatabase(uri string) {
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, trash = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect mongodb")
}

func DisconnectDatabase() {
	client.Disconnect(ctx)
	fmt.Println("disconnect mongodb")
}

func ListDatabase() {
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list database:", databases)
}

func ListCollection(db string) {
	collections, err := client.Database(db).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list collection:", collections)
}

func ConnectCollection(database_name string, collection_name string) {
	collection = client.Database(database_name).Collection(collection_name)
}

func InsertOneDocument(doc interface{}) {
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert one document: ", result.InsertedID)
}

func InsertManyDocument(docs []interface{}) {
	result, err := collection.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert multiple document: ", result.InsertedIDs)
}

func FindDocumnet(filter interface{}) {
	cur, currErr := collection.Find(ctx, filter)
	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)
	var animals []AnimalDoc
	if err := cur.All(ctx, &animals); err != nil {
		panic(err)
	}
	fmt.Println("find document:", animals)
}

func UpdateOneDucmnet(filter interface{}, update interface{}) {
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update document: %v\n", result.ModifiedCount)
}

func DeleteManyDocument(filter interface{}) {
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete documents %v\n", result.DeletedCount)
}
