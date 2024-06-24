package components

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ps-backend-felipe-rodrigues/src/database"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// Conectar ao MongoDB Atlas
	uri := "mongodb+srv://felipe:1234@furz-parfum.jwzwsza.mongodb.net/users?retryWrites=true&w=majority&ssl=true"
	client, err := database.ConectarMongo(uri)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Define o cliente MongoDB globalmente
	database.Client = client

	// Cria um request de teste com os dados do usuário
	user := User{
		Name:        "John Doe",
		DisplayName: "johndoe",
		Password:    "password123",
	}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterUser)

	// Chama o handler de registro de usuário
	handler.ServeHTTP(rr, req)

	// Verifica o status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verifica a resposta
	expected := user
	var returnedUser User
	err = json.NewDecoder(rr.Body).Decode(&returnedUser)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if returnedUser != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", returnedUser, expected)
	}
}
