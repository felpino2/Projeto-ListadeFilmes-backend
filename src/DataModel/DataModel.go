package DataModel

import (
	"fmt"
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
	Id_User int64
	Nome    string
	Senha   string
}

type Lista struct {
	Id_lista      int64
	Filmes        []Filme
	Nome_da_Lista string
	Id_user       int64
}

/*
// CreateUser cria um novo usuário com um nome e senha fornecidos e insere no MongoDB
func CreateUser(client *mongo.Client, nome string, senha string) (User, error) {
	// Gera um ObjectID para o novo usuário
	id := primitive.NewObjectID()

	// Cria o novo usuário
	user := User{
		ID:    id,
		Nome:  nome,
		Senha: senha,
	}

	// Obtém a coleção de usuários
	collection := client.Database("userdb").Collection("users")

	// Insere o novo usuário na coleção
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return User{}, err
	}

	// Retorna o usuário criado
	return user, nil
}
*/

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
