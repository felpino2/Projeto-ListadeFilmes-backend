package teste

import (
	"context"
	"psbackllfa/src/DataModel"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI = "mongodb+srv://felpino2:senhaDB@clusterfelpino.p7mxxkv.mongodb.net/?retryWrites=true&w=majority&appName=ClusterFelpino"
)

var client *mongo.Client

func TestCreateUser(t *testing.T) {
	// Conecta ao MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Verifica a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Testa a criação de um novo usuário
	user, err := DataModel.CreateUser(client, "TestUser", "TestPassword")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Verifica se o ID do usuário foi gerado corretamente
	if user.Id_User == primitive.NilObjectID {
		t.Fatalf("Expected valid ObjectID, got NilObjectID")
	}

	// Verifica se o nome e a senha estão corretos
	if user.Nome != "TestUser" || user.Senha != "TestPassword" {
		t.Fatalf("Expected user with name 'TestUser' and password 'TestPassword', got name '%s' and password '%s'", user.Nome, user.Senha)
	}
}

func TestLoginUser(t *testing.T) {
	// Testa o login do usuário criado
	user, err := DataModel.LoginUser(client, "TestUser", "TestPassword")
	if err != nil {
		t.Fatalf("Failed to login user: %v", err)
	}

	// Verifica se o ID do usuário foi retornado corretamente
	if user.Id_User == primitive.NilObjectID {
		t.Fatalf("Expected valid ObjectID, got NilObjectID")
	}

	// Verifica se o nome e a senha estão corretos
	if user.Nome != "TestUser" || user.Senha != "TestPassword" {
		t.Fatalf("Expected user with name 'TestUser' and password 'TestPassword', got name '%s' and password '%s'", user.Nome, user.Senha)
	}
}
