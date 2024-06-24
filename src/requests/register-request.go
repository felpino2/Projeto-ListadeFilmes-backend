package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"ps-backend-felipe-rodrigues/src/components"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// Função para registrar o usuário diretamente através de uma chamada HTTP
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user components.User

	// Decodifica o corpo da requisição JSON e armazena os dados na variável user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid request payload"})
		return
	}

	// Chama a função RegistrarUsuario para processar o registro do usuário
	err = RegistrarUsuario(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	// Define o status HTTP como 201 (Created) e retorna o usuário registrado como JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Função para registrar o usuário (já fornecida)
func RegistrarUsuario(user components.User) error {
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
	handler := http.HandlerFunc(components.RegisterUser)

	// Chama o handler de registro de usuário com a requisição simulada.
	handler.ServeHTTP(rr, req)

	// Verifica o status code
	if status := rr.Code; status != http.StatusCreated {
		return fmt.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	return nil
}
