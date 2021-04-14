package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var collection = ConnectDB()

func main() {
	// Init Mux Router
	router := mux.NewRouter()

	// Create a handler for undefined routes
	router.NotFoundHandler = notFound()

	// Handle routes
	router.HandleFunc("/notes", getNotes).Methods("GET")
	router.HandleFunc("/notes", createNote).Methods("POST")

	// Log requests
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	// Handle Requests
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}

func notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Attempted to get path %s - not implemented", r.URL.Path)
		sendErrorResponse(w, http.StatusNotImplemented, "Not Implemented")
		return
	})
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Create an empty array of Notes
	var notes []Note

	// Query Database
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		log.Printf("Failed to query collection: %s", err)
		return
	}

	// Defer closing the cursor
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		// Define a Note model
		var note Note

		// Attempt to decode into the model
		err := cur.Decode(&note)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			log.Printf("Unable to decode Model: %s", err)
			return
		}

		// Append that Note back to the Array
		notes = append(notes, note)
	}

	if err := cur.Err(); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		log.Printf("Problem with Cursor: %s", err)
		return
	}

	// Write the Response
	json.NewEncoder(w).Encode(notes)
	return
}

func createNote(w http.ResponseWriter, r *http.Request) {
	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Create a Note
	var note Note

	// Decode Request Body into Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		log.Printf("Failed to decode request: %s", err)
		return
	}

	// Insert Note into the Collection
	result, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		log.Printf("Failed to insert record: %s", err)
		return
	}

	// Write the Response
	json.NewEncoder(w).Encode(result)
	return
}

func sendErrorResponse(w http.ResponseWriter, sc int, e string) {
	w.WriteHeader(sc)
	json.NewEncoder(w).Encode(&ErrorResponse{sc, e})
}
