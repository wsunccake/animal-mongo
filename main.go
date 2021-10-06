package main

import (
	"animal-mongo/utils"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_URI     string = "mongodb://127.0.0.1:27017"
	DATABASE_NAME   string = "zoo"
	COLLECTION_NAME string = "animal"
)

type SpeciesDoc struct {
	ID        string    `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	Food      []string  `bson:"food"`
}

func main() {
	utils.ConnectDatabase(MONGODB_URI)
	defer utils.DisconnectDatabase()

	utils.ListDatabase()
	utils.ListCollection(DATABASE_NAME)

	utils.ConnectCollection(DATABASE_NAME, COLLECTION_NAME)

	animal0 := utils.AnimalDoc{"mickey", 1, "Mouse", map[string]string{}}
	utils.InsertOneDocument(animal0)

	animal1 := utils.AnimalDoc{"doogy", 1, "Dog", map[string]string{}}
	animal2 := utils.AnimalDoc{"kitty", 2, "Cat", map[string]string{}}
	animal3 := utils.AnimalDoc{"piggy", 3, "Pig", map[string]string{"h": "1", "w": "a"}}
	animals := []interface{}{animal1, animal2, animal3}
	utils.InsertManyDocument(animals)

	filter0 := bson.M{"age": bson.M{"$gt": 2}}
	utils.FindDocumnet(filter0)

	rand.Seed(time.Now().UnixNano())
	age := rand.Intn(5)
	id, _ := primitive.ObjectIDFromHex("615d5471587ec242bcdca7e5")
	filter1 := bson.M{"_id": id}
	update1 := bson.D{
		{"$set", bson.D{{"age", age}}},
	}
	utils.UpdateOneDucmnet(filter1, update1)

	filter2 := bson.M{"species": "Mouse"}
	utils.DeleteManyDocument(filter2)
}
