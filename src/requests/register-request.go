package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"psbackllfa/src/DataModel"
	"psbackllfa/src/database"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// Função para registrar o usuário diretamente através de uma chamada HTTP
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user DataModel.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received user: %v", user)

	collection := database.Client.Database("TRABALHOCAIO").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully"))
}

// Função para registrar o usuário (já fornecida)
func RegistrarUsuario(user DataModel.User) error {
	// Converte o objeto User em JSON.
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Cria uma requisição HTTP POST simulada com os dados do usuário.
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Define o cabeçalho Content-Type como application/json.
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta da requisição.
	rr := httptest.NewRecorder()

	// Define o handler de registro de usuário.
	handler := http.HandlerFunc(DataModel.CreateUser)

	// Chama o handler de registro de usuário com a requisição simulada.
	handler.ServeHTTP(rr, req)

	// Verifica o status code
	if status := rr.Code; status != http.StatusCreated {
		return fmt.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	return nil
}
