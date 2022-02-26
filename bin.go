package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bin struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Name              string             `json:"name" bson:"name"`
	Location          Location           `json:"location" bson:"location"`
	Longitude         float64            `json:"longitude" bson:"longitude"`
	Latitude          float64            `json:"latitude" bson:"latitude"`
	AcceptedMaterials []string           `json:"acceptedMaterials" bson:"acceptedMaterials"`
}
