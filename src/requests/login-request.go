package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"ps-backend-felipe-rodrigues/src/components"
)

// Função para lidar com requisições de login de usuários
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user components.User

	// Decodifica o corpo da requisição JSON e armazena os dados na variável user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid request payload"})
		return
	}

	// Chama a função LoginUsuario para processar o login do usuário
	err = LoginUsuario(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	// Define o status HTTP como 200 (OK) e retorna o usuário logado como JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Função para simular uma requisição de login de usuário
func LoginUsuario(user components.User) error {
	// Converte o objeto User em JSON.
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Cria uma requisição HTTP POST simulada com os dados do usuário.
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Define o cabeçalho Content-Type como application/json.
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta da requisição.
	rr := httptest.NewRecorder()

	// Define o handler de login de usuário.
	handler := http.HandlerFunc(components.LoginUser)

	// Chama o handler de login de usuário com a requisição simulada.
	handler.ServeHTTP(rr, req)

	// Verifica o status code
	if status := rr.Code; status != http.StatusOK {
		return fmt.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	return nil
}
