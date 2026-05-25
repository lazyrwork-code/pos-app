package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Product struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Price    float64       `json:"price" bson:"price"`
	Stock    int           `json:"stock" bson:"stock"`
	Category string        `json:"category" bson:"category"`
	Unit     string        `json:"unit" bson:"unit"`
}
