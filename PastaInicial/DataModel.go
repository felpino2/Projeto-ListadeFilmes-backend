package main

import (
	"fmt"
	"time"
)

type Filme struct {
	Id          int
	Nome        string
	Runtime     int
	NumOrdem    int
	Description string
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
	return Lista{Filmes: []Filme{, film.Description, film.Nome, film.NumOrdem, , film.}}
}

func UpdateRating(iduser int64, idfilme int, stars int) (Rating, error) {

	if stars < 0 || stars > 5 {
		return Rating{}, fmt.Errorf("Estrelas inválidas")
	}
	return Rating{idfilme, iduser, stars}, nil
}

func main() {
	nome := "juleano"
	senha := "12232"

	fmt.Print(CreateUser(nome, senha))
}
