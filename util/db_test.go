package util

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	MONGODB_URI     string = "mongodb://127.0.0.1:27017"
	DATABASE_NAME   string = "testZoo"
	COLLECTION_NAME string = "animal"
)

func TestInsertDocument(t *testing.T) {
	ConnectDatabase(MONGODB_URI)
	defer DisconnectDatabase()

	ConnectCollection(DATABASE_NAME, COLLECTION_NAME)
	animal := AnimalDoc{"mickey", 1, "Mouse", map[string]string{}}
	InsertOneDocument(animal)

	results := FindManyDocumnet(bson.M{})
	t.Logf("result: %v", results)
	if len(results) < 1 {
		t.Errorf("fail to insert document")
	}
}
