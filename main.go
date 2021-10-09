package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"V41SH/instagramAPI/helper"
	"V41SH/instagramAPI/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

// POST method to create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Users

	_ = json.NewDecoder(r.Body).Decode(&user)

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		helper.GetError(err, w)
		return
	}


	json.NewEncoder(w).Encode(result)
}

// GET method to get a user of certain Id
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.Users
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// POST method to create a new post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Posts

	_ = json.NewDecoder(r.Body).Decode(&post)

	result, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GET method to get a post of certain Id
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Posts
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(post)
}

// GET method to get all posts of a user
func getUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var user models.Users
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user.Posts)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Declare all routes
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts/users/{id}", getUserPosts).Methods("GET")

	//config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(":8000", r))
}