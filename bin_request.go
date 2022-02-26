package main

type BinRequest struct {
	Distance  int     `json:"distance" bson:"distance"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Material  string  `json:"material" bson:"material"`
}
