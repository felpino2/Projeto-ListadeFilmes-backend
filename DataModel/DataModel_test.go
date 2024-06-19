package DataModel

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}
	t.Logf("User criado: %+v", user)
}

func TestCreateLista(t *testing.T) {
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}
	lista, err := CreateLista(user.Id_User, "TesteListaFilmes")
	if err != nil {
		t.Fatalf("Erro na func de criar lista: %v", err)
	}
	t.Logf("Lista criada: %+v", lista)
}

func TestInsertFilmesLista(t *testing.T) {
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}
	lista, err := CreateLista(user.Id_User, "TesteListaFilmes")
	if err != nil {
		t.Fatalf("Erro na func de criar lista: %v", err)
	}

	filme1 := Filme{Id: 1, Nome: "caio", Runtime: 148, NumOrdem: 2, Description: "caio1 Descricao"}
	lista = InsertFilmesLista(lista, filme1)
	if len(lista.Filmes) != 1 || lista.Filmes[0] != filme1 {
		t.Fatalf("Erro ao inserir filme 1 na lista")
	}
	t.Logf("Filme 1 inserido: %+v", lista)

	filme2 := Filme{Id: 2, Nome: "caio2", Runtime: 136, NumOrdem: 1, Description: "caio2 Descricao2"}
	lista = InsertFilmesLista(lista, filme2)
	if len(lista.Filmes) != 2 || lista.Filmes[1] != filme2 {
		t.Fatalf("Erro ao inserir filme 2 na lista")
	}
	t.Logf("Filme 2 inserido: %+v", lista)
}

func TestUpdateRating(t *testing.T) {
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}

	filme1 := Filme{Id: 1, Nome: "caio", Runtime: 148, NumOrdem: 2, Description: "caio1 Descricao"}
	rating, err := UpdateRating(user.Id_User, filme1.Id, 5)
	if err != nil {
		t.Fatalf("Erro ao atualizar rating: %v", err)
	}
	if rating.Stars != 5 {
		t.Fatalf("Erro no valor do rating atualizado")
	}
	t.Logf("Rating atualizado: %+v", rating)
}
