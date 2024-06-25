package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func IniciarMongo() (client *mongo.Client) {
	log.Println("Server is starting...")

	// Define a URI de conexão com o MongoDB Atlas.
	uri := "mongodb+srv://livialemos:LMPRY@trabalhocaio.fay3v2j.mongodb.net/?retryWrites=true&w=majority&appName=TRABALHOCAIO"

	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado.
	client, err := ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	Client = client
	return (client)
}
