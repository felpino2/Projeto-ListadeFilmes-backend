package requests

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"psbackllfa/src/DataModel"
	"psbackllfa/src/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Trying to log in user: %v", loginReq)

	collection := database.Client.Database("TRABALHOCAIO").Collection("users")
	filter := bson.M{"nome": loginReq.Username, "senha": loginReq.Password}

	var foundUser DataModel.User
	err := collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("User not found or wrong password: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			log.Printf("Error finding user: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	log.Printf("User logged in successfully: %v", foundUser)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
