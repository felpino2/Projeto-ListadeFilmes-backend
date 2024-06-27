package DataModel

/*
import (
	"testing"
)
// TestCreateUser testa a função CreateUser para garantir que um usuário é criado corretamente.
func TestCreateUser(t *testing.T) {
	// Chama a função CreateUser com os parâmetros de teste.
	user, err := CreateUser("caio", "caio123")
	// Verifica se houve um erro ao criar o usuário.
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}
	// Verifica se o nome do usuário criado é o esperado.
	if user.Nome != "caio" {
		t.Fatalf("Esperado nome 'caio', mas obteve '%s'", user.Nome)
	}
	// Registra o usuário criado para visualização.
	t.Logf("User criado: %+v", user)
}

// TestCreateLista testa a função CreateLista para garantir que uma lista é criada corretamente.
func TestCreateLista(t *testing.T) {
	// Cria um usuário primeiro, pois a lista está associada a um usuário.
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}
	// Chama a função CreateLista com o ID do usuário e o nome da lista.
	lista, err := CreateLista(user.Id_User, "TesteListaFilmes")
	if err != nil {
		t.Fatalf("Erro na func de criar lista: %v", err)
	}
	// Verifica se o nome da lista criada é o esperado.
	if lista.Nome_da_Lista != "TesteListaFilmes" {
		t.Fatalf("Esperado nome da lista 'TesteListaFilmes', mas obteve '%s'", lista.Nome_da_Lista)
	}
	// Registra a lista criada para visualização.
	t.Logf("Lista criada: %+v", lista)
}

// TestInsertFilmesLista testa a função InsertFilmesLista para garantir que filmes são inseridos corretamente na lista.
func TestInsertFilmesLista(t *testing.T) {
	// Cria um usuário.
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}
	// Cria uma lista associada ao usuário.
	lista, err := CreateLista(user.Id_User, "TesteListaFilmes")
	if err != nil {
		t.Fatalf("Erro na func de criar lista: %v", err)
	}

	// Define o primeiro filme e insere na lista.
	filme1 := Filme{Id: 1, Nome: "caio", Runtime: 148, NumOrdem: 2, Description: "caio1 Descricao"}
	lista = InsertFilmesLista(lista, filme1)
	// Verifica se o filme foi inserido corretamente.
	if len(lista.Filmes) != 1 || lista.Filmes[0] != filme1 {
		t.Fatalf("Erro ao inserir filme 1 na lista")
	}
	// Registra a lista com o primeiro filme inserido.
	t.Logf("Filme 1 inserido: %+v", lista)

	// Define o segundo filme e insere na lista.
	filme2 := Filme{Id: 2, Nome: "caio2", Runtime: 136, NumOrdem: 1, Description: "caio2 Descricao2"}
	lista = InsertFilmesLista(lista, filme2)
	// Verifica se o segundo filme foi inserido corretamente.
	if len(lista.Filmes) != 2 || lista.Filmes[1] != filme2 {
		t.Fatalf("Erro ao inserir filme 2 na lista")
	}
	// Registra a lista com os dois filmes inseridos.
	t.Logf("Filme 2 inserido: %+v", lista)
}

// TestUpdateRating testa a função UpdateRating para garantir que a avaliação (rating) é atualizada corretamente.
func TestUpdateRating(t *testing.T) {
	// Cria um usuário.
	user, err := CreateUser("caio", "caio123")
	if err != nil {
		t.Fatalf("Erro na func de criar user: %v", err)
	}

	// Define um filme.
	filme1 := Filme{Id: 1, Nome: "caio", Runtime: 148, NumOrdem: 2, Description: "caio1 Descricao"}
	// Atualiza o rating do filme.
	rating, err := UpdateRating(user.Id_User, filme1.Id, 5)
	if err != nil {
		t.Fatalf("Erro ao atualizar rating: %v", err)
	}
	// Verifica se o rating foi atualizado corretamente.
	if rating.Stars != 5 {
		t.Fatalf("Erro no valor do rating atualizado, esperado 5 mas obteve %d", rating.Stars)
	}
	// Registra o rating atualizado.
	t.Logf("Rating atualizado: %+v", rating)
}

// TestSaveLista testa a função SaveLista para garantir que a lista é salva corretamente.
func TestSaveLista(t *testing.T) {
	// Define uma lista.
	lista := Lista{Id_lista: 1, Nome_da_Lista: "Lista Teste"}
	// Salva a lista.
	SaveLista(lista)
	// Recupera a lista pelo ID.
	savedLista, exists := GetListaByID(1)
	// Verifica se a lista foi salva corretamente.
	if !exists {
		t.Fatalf("Erro ao salvar a lista")
	}
	if savedLista.Nome_da_Lista != "Lista Teste" {
		t.Fatalf("Esperado nome da lista 'Lista Teste', mas obteve '%s'", savedLista.Nome_da_Lista)
	}
	// Registra a lista salva.
	t.Logf("Lista salva: %+v", savedLista)
}

// TestGetListaByID testa a função GetListaByID para garantir que a lista é recuperada corretamente pelo ID.
func TestGetListaByID(t *testing.T) {
	// Define e salva uma lista.
	lista := Lista{Id_lista: 2, Nome_da_Lista: "Outra Lista Teste"}
	SaveLista(lista)
	// Recupera a lista pelo ID.
	retrievedLista, exists := GetListaByID(2)
	// Verifica se a lista foi recuperada corretamente.
	if !exists {
		t.Fatalf("Erro ao obter a lista")
	}
	if retrievedLista.Nome_da_Lista != "Outra Lista Teste" {
		t.Fatalf("Esperado nome da lista 'Outra Lista Teste', mas obteve '%s'", retrievedLista.Nome_da_Lista)
	}
	// Registra a lista recuperada.
	t.Logf("Lista obtida: %+v", retrievedLista)
}
*/
