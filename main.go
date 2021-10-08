package main

import (
	"animal-mongo/util"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_URI     string = "mongodb://127.0.0.1:27017"
	DATABASE_NAME   string = "zoo"
	COLLECTION_NAME string = "animal"
)

func main() {
	util.ConnectDatabase(MONGODB_URI)
	defer util.DisconnectDatabase()

	util.ListDatabase()
	util.ListCollection(DATABASE_NAME)

	util.ConnectCollection(DATABASE_NAME, COLLECTION_NAME)

	animal0 := util.AnimalDoc{"mickey", 1, "Mouse", map[string]string{}}
	util.InsertOneDocument(animal0)

	animal1 := util.AnimalDoc{"doogy", 1, "Dog", map[string]string{}}
	animal2 := util.AnimalDoc{"kitty", 2, "Cat", map[string]string{}}
	animal3 := util.AnimalDoc{"piggy", 3, "Pig", map[string]string{"h": "1", "w": "a"}}
	animals := []interface{}{animal1, animal2, animal3}
	util.InsertManyDocument(animals)

	filter00 := bson.M{"age": bson.M{"$gt": 2}}
	results00 := util.FindManyDocumnet(filter00)
	fmt.Printf("result: %v", results00)
	for _, e := range results00 {
		// fmt.Println(i)
		fmt.Printf("%v\n", e)
		fmt.Printf("type: %s\n", reflect.TypeOf(e))
		eMap := e.(bson.D).Map()
		fmt.Printf("%v\n", eMap)
		fmt.Printf("%v\n", reflect.TypeOf(eMap))
		fmt.Printf("%v\n", eMap["age"])
	}

	filter01 := bson.M{"age": bson.M{"$gt": 2}}
	result01 := util.FindOneDocumnet(filter01)
	result01Map := result01.(bson.D).Map()
	fmt.Printf("resultMap: %v\n", result01Map)
	fmt.Printf("resultMap type: %s\n", reflect.TypeOf(result01Map))
	fmt.Printf("%v\n", result01Map["age"])

	output1, err := json.MarshalIndent(result01, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output1)

	rand.Seed(time.Now().UnixNano())
	age := rand.Intn(5)
	id, _ := primitive.ObjectIDFromHex("615d5471587ec242bcdca7e5")
	filter11 := bson.M{"_id": id}
	update11 := bson.D{
		{"$set", bson.D{{"age", age}}},
	}
	util.UpdateOneDucmnet(filter11, update11)

	filter30 := bson.M{"species": "Mouse"}
	util.DeleteManyDocument(filter30)
}
