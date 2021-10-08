package test

import (
	"animal-mongo/util"
	"testing"
)

const (
	MONGODB_URI     string = "mongodb://127.0.0.1:27017"
	DATABASE_NAME   string = "testZoo"
	COLLECTION_NAME string = "animal"
)

func TestConnect(t *testing.T) {
	util.ConnectDatabase(MONGODB_URI)
	defer util.DisconnectDatabase()
}
