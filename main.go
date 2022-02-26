package main

import (
	"fmt"

	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
)

// getBinsByDistance - returns an array of bins within a given distance (metres)
// and that accept a given material
func getBinsByDistance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var binRequest BinRequest
	_ = json.NewDecoder(r.Body).Decode(&binRequest)
	fmt.Println("Find Bin Request: ", binRequest)

	coll := conn.Database(DBName).Collection(PointCollection)
	var results []Bin
	filter := bson.D{
		{"location",
			bson.D{
				{"$near", bson.D{
					{"$geometry", NewPoint(binRequest.Longitude, binRequest.Latitude)},
					{"$maxDistance", binRequest.Distance},
				}},
			}},
		{"acceptedMaterials", binRequest.Material},
	}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	for cur.Next(ctx) {
		var bin Bin
		err := cur.Decode(&bin)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		results = append(results, bin)
	}
	json.NewEncoder(w).Encode(results)
}

// addBin - adds a new bin to the collection in the DB.
func addBin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Receive and decode HTTP POST Request
	var bin Bin
	_ = json.NewDecoder(r.Body).Decode(&bin)
	bin.ID = primitive.NewObjectID()
	bin.Location = NewPoint(bin.Longitude, bin.Latitude)

	// Connect to MongoDB collection
	coll := conn.Database(DBName).Collection(PointCollection)
	insertResult, err := coll.InsertOne(ctx, bin)

	if err != nil {
		fmt.Printf("Could not insert new bin. ID: %s\n", bin.ID)
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Printf("Inserted new bin. ID: %s\n", insertResult.InsertedID)

	// Send newly added bin in response
	json.NewEncoder(w).Encode(bin)
}

// updateBin - updates a bin with a given ID from the DB
func updateBin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	binId, _ := primitive.ObjectIDFromHex(params["id"])

	// Receive and decode HTTP PUT Request
	var bin Bin
	_ = json.NewDecoder(r.Body).Decode(&bin)

	updateFilter := bson.M{}

	// Check which fields are being updated
	if bin.Name != "" {
		updateFilter["name"] = bin.Name
	}
	if bin.Longitude != 0.0 {
		updateFilter["longitude"] = bin.Longitude
	}
	if bin.Latitude != 0.0 {
		updateFilter["latitude"] = bin.Latitude
	}
	if bin.AcceptedMaterials != nil {
		updateFilter["acceptedMaterials"] = bin.AcceptedMaterials
	}

	updateFilter = bson.M{"$set": updateFilter}

	// Connect to MongoDB collection
	coll := conn.Database(DBName).Collection(PointCollection)
	updateResult, err := coll.UpdateByID(ctx, binId, updateFilter)

	if err != nil {
		fmt.Printf("Could not update bin. ID: %s\n", binId)
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Printf("Updated bin. ID: %s\n", updateResult.UpsertedID)

	json.NewEncoder(w).Encode(updateResult.UpsertedCount)
}

// deleteBin - removes a bin with a given ID from the DB
func deleteBin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	binId, _ := primitive.ObjectIDFromHex(params["id"])

	// Connect to MongoDB collection
	coll := conn.Database(DBName).Collection(PointCollection)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": binId})

	if err != nil {
		fmt.Printf("Could not delete bin. ID: %s\n", binId)
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Printf("Deleted bin. ID: %s\n", binId)

	json.NewEncoder(w).Encode(deleteResult.DeletedCount)
}

// main - connects to MongoDB and awaits REST requests
func main() {
	if err := createDBSession(); err != nil {
		fmt.Println(err)
		return
	}
	if err := createIndex(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to MongoDB")

	r := mux.NewRouter()

	// Set up Endpoints/Router Handling
	r.HandleFunc("/api/bins", getBinsByDistance).Methods("GET")
	r.HandleFunc("/api/bins", addBin).Methods("POST")
	r.HandleFunc("/api/bins/{id}", updateBin).Methods("PUT")
	r.HandleFunc("/api/bins/{id}", deleteBin).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
