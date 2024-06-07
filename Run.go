package main

import (
	"encoding/json"
	"io"
	"net/http"
	//	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"go.mongodb.org/mongo-driver/mongo"
)

func Run() {

	s := http.Server{
		Addr:    "localhost:8080",
		Handler: Router{},
	}

	s.ListenAndServe()
}

type Router struct{}

func (Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var u User
	body, _ := io.ReadAll(req.Body)
	json.Unmarshal(body, &u)
	//Insert(u, "users")
}

type User struct {
	Name  string `json:"name" bson:"name"`
	RA    string `json:"ra" bson:"ra"`
	Email string `json:"email" bson:"email"`
	Pass  string `json:"pass" bson:"email"`
}

/*
func Insert(u any, collection string) (primitive.ObjectId, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("nome-do-banco").Collection(collection)

	res, err := collection.InsertOne(ctx, u)

	return res.InsertedID, nil

}
*/
