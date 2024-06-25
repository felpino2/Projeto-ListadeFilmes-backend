package DataModel

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"psbackllfa/src/database"
	"time"
)

type Filme struct {
	Id          int    `json:"Id_Filme" bson:"Id_Filme"`
	Nome        string `json:"Nome_Filme" bson:"Nome_Filme"`
	Runtime     int    `json:"Runtime_Filme" bson:"Runtime_Filme"`
	NumOrdem    int    `json:"NumOrdem" bson:"NumOrdem"`
	Description string `json:"Description" bson:"Description"`
}

type Rating struct {
	Id_Filme int
	Id_User  int64
	Stars    int
}

type User struct {
	Nome  string `bson:"nome" json:"nome"`
	Senha string `bson:"senha" json:"senha"`
}

type Lista struct {
	Id_lista      int64
	Filmes        []Filme
	Nome_da_Lista string
	Id_user       int64
}

// CreateUser cria um novo usuário com um nome e senha fornecidos e insere no MongoDB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decodifica o corpo da requisição JSON e armazena os dados na variável user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Obtém uma referência para a coleção "users" no banco de dados "userdb"
	collection := database.GetUserCollection()

	// Insere o documento user na coleção "users"
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Define o status HTTP como 201(created)
	w.WriteHeader(http.StatusCreated)

	// Codifica a variável user em JSON e escreve na resposta HTTP
	json.NewEncoder(w).Encode(user)
}

// LoginUser verifica as credenciais do usuário e retorna o usuário se for bem-sucedido
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var foundUser User

	// Decodifica o corpo da requisição JSON e armazena os dados na variável user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Obtém uma referência para a coleção "users" no banco de dados "userdb"
	collection := database.GetUserCollection()

	// Busca o usuário no banco de dados pelo DisplayName e Password
	filter := bson.M{"displayname": user.Nome, "password": user.Senha}
	err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Failed to login user", http.StatusInternalServerError)
		return
	}

	// Define o status HTTP como 200 (OK) e retorna o usuário encontrado como JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundUser)
}

// CreateLista cria uma nova lista de filmes para um usuário específico
func CreateLista(iduser int64, nomeList string) (Lista, error) {
	// Gera um ID único para a lista com base no tempo atual em nanossegundos
	idLista := time.Now().UnixNano()
	// Cria uma nova lista com o ID gerado, um slice vazio de filmes, o nome da lista e o ID do usuário
	lista := Lista{Id_lista: idLista, Filmes: []Filme{}, Nome_da_Lista: nomeList, Id_user: iduser}
	// Salva a lista no mapa global de listas
	listas[idLista] = lista
	// Retorna a lista criada
	return lista, nil
}

// InsertFilmesLista adiciona um filme a uma lista existente
func InsertFilmesLista(list Lista, film Filme) Lista {
	// Adiciona o filme ao slice de filmes da lista
	list.Filmes = append(list.Filmes, film)
	// Retorna a lista atualizada
	return list
}

// UpdateRating atualiza a avaliação (rating) de um filme por um usuário específico
func UpdateRating(iduser int64, idfilme int, stars int) (Rating, error) {
	// Verifica se a quantidade de estrelas é válida (entre 0 e 5)
	if stars < 0 || stars > 5 {
		// Retorna um erro se o valor for inválido
		return Rating{}, fmt.Errorf("Estrelas inválidas")
	}
	// Retorna o rating atualizado
	return Rating{idfilme, iduser, stars}, nil
}
