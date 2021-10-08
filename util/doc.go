package util

import "time"

type SpeciesDoc struct {
	ID        string    `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	Food      []string  `bson:"food"`
}

type AnimalDoc struct {
	Name    string            `bson:"name,omitempty"`
	Age     int               `bson:"age,omitempty"`
	Species string            `bson:"species"`
	Size    map[string]string `bson:"size"`
}
