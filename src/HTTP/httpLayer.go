package HTTP

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// Função principal para iniciar o serviço HTTP
func RunService() {
	// Configura um novo servidor HTTP
	s := &http.Server{
		Addr:         "localhost:8080", // Define o endereço e a porta do servidor
		Handler:      Router{},         // Define o roteador que vai lidar com as requisições
		ReadTimeout:  10 * time.Second, // Define o tempo máximo para ler uma requisição
		WriteTimeout: 10 * time.Second, // Define o tempo máximo para escrever uma resposta
	}

	// Inicia o servidor e registra uma mensagem de erro caso algo dê errado
	log.Fatal(s.ListenAndServe())
}

// Define a estrutura Router
type Router struct{}

// Mapeia os caminhos das URLs para suas respectivas funções de manipulação
var routeHandlers = map[string]func(http.ResponseWriter, *http.Request){
	"/create_user":       CreateUserHandler,   // Rota para criar um usuário
	"/create_lista":      CreateListaHandler,  // Rota para criar uma lista
	"/lista_updater/id:": insertFilmeHandler,  // Rota para inserir um filme em uma lista
	"/update_rating":     updateRatingHandler, // Rota para atualizar a avaliação de um filme
}

// Implementa a função ServeHTTP para a estrutura Router
func (Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Percorre o mapa de rotas e verifica se a URL da requisição começa com algum dos prefixos
	for prefix, handler := range routeHandlers {
		if strings.HasPrefix(req.URL.Path, prefix) {
			// Chama a função de manipulação correspondente à rota
			handler(w, req)
			return
		}
	}
	// Se nenhuma rota corresponder, retorna um erro 404 (Not Found)
	http.NotFound(w, req)
}
