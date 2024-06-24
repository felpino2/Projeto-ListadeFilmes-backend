package teste

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"psbackllfa/src/HTTP"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"psbackllfa/src/DataModel"
)

func TestCreateUserHandler(t *testing.T) {
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

	// Cria um novo usuário para o teste
	user := DataModel.User{
		Nome:  "TestUserHandler",
		Senha: "TestPasswordHandler",
	}

	// Serializa o usuário para JSON
	body, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	// Cria uma requisição POST para o handler
	req, err := http.NewRequest("POST", "/createUser", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HTTP.CreateUserHandler)

	// Chama o handler com a requisição simulada
	handler.ServeHTTP(rr, req)

	// Verifica se o status code é 201
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Decodifica a resposta JSON
	var novoUser DataModel.User
	err = json.NewDecoder(rr.Body).Decode(&novoUser)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Verifica se o ID do usuário foi gerado corretamente
	if novoUser.Id_User == primitive.NilObjectID {
		t.Fatalf("Expected valid ObjectID, got NilObjectID")
	}

	// Verifica se o nome e a senha estão corretos
	if novoUser.Nome != "TestUserHandler" || novoUser.Senha != "TestPasswordHandler" {
		t.Fatalf("Expected user with name 'TestUserHandler' and password 'TestPasswordHandler', got name '%s' and password '%s'", novoUser.Nome, novoUser.Senha)
	}
}

func TestLoginUserHandler(t *testing.T) {
	// Cria um usuário para o teste de login
	_, err := DataModel.CreateUser(client, "TestUserHandlerLogin", "TestPasswordHandlerLogin")
	if err != nil {
		t.Fatalf("Failed to create user for login test: %v", err)
	}

	// Cria um usuário de login para o teste
	user := DataModel.User{
		Nome:  "TestUserHandlerLogin",
		Senha: "TestPasswordHandlerLogin",
	}

	// Serializa o usuário para JSON
	body, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	// Cria uma requisição POST para o handler
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HTTP.LoginUserHandler)

	// Chama o handler com a requisição simulada
	handler.ServeHTTP(rr, req)

	// Verifica se o status code é 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decodifica a resposta JSON
	var loggedInUser DataModel.User
	err = json.NewDecoder(rr.Body).Decode(&loggedInUser)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Verifica se o ID do usuário foi retornado corretamente
	if loggedInUser.Id_User == primitive.NilObjectID {
		t.Fatalf("Expected valid ObjectID, got NilObjectID")
	}

	// Verifica se o nome e a senha estão corretos
	if loggedInUser.Nome != "TestUserHandlerLogin" || loggedInUser.Senha != "TestPasswordHandlerLogin" {
		t.Fatalf("Expected user with name 'TestUserHandlerLogin' and password 'TestPasswordHandlerLogin', got name '%s' and password '%s'", loggedInUser.Nome, loggedInUser.Senha)
	}
}
