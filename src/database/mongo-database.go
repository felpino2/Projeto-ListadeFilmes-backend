package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client é uma variável global que armazenará a conexão com o banco de dados MongoDB.
var Client *mongo.Client

// ConectarMongo é uma função que estabelece uma conexão com o MongoDB.
// Ela recebe uma URI de conexão como parâmetro e retorna um cliente MongoDB ou um erro, caso ocorra.
func ConectarMongo(uri string) (*mongo.Client, error) {
	// Define as opções do cliente MongoDB usando a URI fornecida.
	clientOptions := options.Client().ApplyURI(uri)

	// Tenta conectar ao MongoDB usando as opções definidas.
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Tenta fazer um "ping" no banco de dados para verificar se a conexão foi bem-sucedida.
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetUserCollection é uma função que retorna uma referência para a coleção "user_profiles"
// no banco de dados "users" usando o cliente MongoDB global.
func GetUserCollection() *mongo.Collection {
	// Acessa o banco de dados "users" e retorna a coleção "user_profiles".
	return Client.Database("userDB").Collection("user_profiles")
}
