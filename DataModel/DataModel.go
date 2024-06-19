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
	Filmes        []Filme
	Nome_da_Lista string
	Id_user       int64
}

// BANCO DE DADOS
func CreateUser(nome string, senha string) (User, error) {
	id := time.Now().UnixNano()
	return User{id, nome, senha}, nil
}

// BANCO DE DADOS
func SignInUser(nome string, senha string) (User, error) {
	return User{}, nil
}

// SEARCH need Banco de DADOS

func CreateLista(iduser int64, nomeList string) (Lista, error) {
	return Lista{Filmes: []Filme{}, Nome_da_Lista: nomeList, Id_user: iduser}, nil
}

func InsertFilmesLista(list Lista, film Filme) Lista {
	list.Filmes = append(list.Filmes, film)
	return list
}

func UpdateRating(iduser int64, idfilme int, stars int) (Rating, error) {

	if stars < 0 || stars > 5 {
		return Rating{}, fmt.Errorf("Estrelas inválidas")
	}
	return Rating{idfilme, iduser, stars}, nil
}
