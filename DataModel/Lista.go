package DataModel

// Declaramos uma variável global chamada 'listas' que é um mapa.
// Este mapa associa um ID de lista (int64) a uma estrutura de dados 'Lista'.
// O mapa é usado para armazenar todas as listas criadas em memória.
var listas = make(map[int64]Lista)

// Função GetListaByID
// Esta função recebe um ID de lista como parâmetro e retorna a lista correspondente e um valor booleano.
// O valor booleano indica se a lista foi encontrada (true) ou não (false).
func GetListaByID(id int64) (Lista, bool) {
	// Procuramos a lista no mapa 'listas' usando o ID fornecido.
	lista, exists := listas[id]
	// Retornamos a lista encontrada (ou uma lista vazia se não encontrada) e o valor booleano.
	return lista, exists
}

// Função SaveLista
// Esta função recebe uma lista como parâmetro e a salva no mapa 'listas'.
// A lista é armazenada usando o ID da lista como chave.
func SaveLista(lista Lista) {
	// Salvamos ou atualizamos a lista no mapa 'listas' usando o ID da lista como chave.
	listas[lista.Id_lista] = lista
}
