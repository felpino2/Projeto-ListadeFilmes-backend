package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConectarMongo(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

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

func GetUserCollection() *mongo.Collection {
	return Client.Database("DBuser").Collection("usuario")
}

func AuthenticateUser(username, password string) bool {
	// Exemplo simples: autenticação com credenciais fixas
	if username == "testuser" && password == "testpassword" {
		return true
	}
	// Aqui você deve verificar o usuário e a senha no banco de dados
	return false
}
