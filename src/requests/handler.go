package requests

import (
	"encoding/json" // Importa o pacote para trabalhar com JSON
	"go.mongodb.org/mongo-driver/mongo"
	"net/http" // Importa o pacote para trabalhar com HTTP
	"psbackllfa/src/DataModel"
	"strconv" // Importa o pacote para conversões de strings para números
	"strings" // Importa o pacote para manipulações de strings
)

var client *mongo.Client

// Função para criar uma nova lista
func CreateListaHandler(res http.ResponseWriter, req *http.Request) {
	// Verifica se o método HTTP é POST
	if req.Method == "POST" {
		var list DataModel.Lista

		// Decodifica o corpo da requisição JSON para a struct Lista
		err := json.NewDecoder(req.Body).Decode(&list)
		if err != nil {
			// Retorna um erro 400 (Bad Request) se a decodificação falhar
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		// Chama a função para criar uma nova lista
		novaLista, err := DataModel.CreateLista(list.Id_user, list.Nome_da_Lista)
		if err != nil {
			// Retorna um erro 500 (Internal Server Error) se a criação da lista falhar
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		// Codifica a nova lista como JSON e escreve na resposta
		json.NewEncoder(res).Encode(novaLista)
	} else {
		// Retorna um erro 405 (Method Not Allowed) se o método HTTP não for POST
		http.Error(res, "Metodo nao funciona", http.StatusMethodNotAllowed)
	}
}

// Função para inserir um filme em uma lista existente
func InsertFilmeHandler(res http.ResponseWriter, req *http.Request) {
	// Verifica se o método HTTP é PUT
	if req.Method == "PUT" {
		// Extrai a ID da lista da URL
		path := strings.TrimPrefix(req.URL.Path, "/lista_updater/id:")
		listaID, err := strconv.ParseInt(path, 10, 64)
		if err != nil {
			// Retorna um erro 400 (Bad Request) se a ID da lista for inválida
			http.Error(res, "Invalid lista ID", http.StatusBadRequest)
			return
		}

		var requestData struct {
			Filme DataModel.Filme
		}
		// Decodifica o corpo da requisição JSON para a struct Filme
		err = json.NewDecoder(req.Body).Decode(&requestData)
		if err != nil {
			// Retorna um erro 400 (Bad Request) se a decodificação falhar
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		// Busca a lista existente pelo ID
		lista, exists := DataModel.GetListaByID(listaID)
		if !exists {
			// Retorna um erro 404 (Not Found) se a lista não for encontrada
			http.Error(res, "Lista not found", http.StatusNotFound)
			return
		}

		// Atualiza a lista com o novo filme
		updatedLista := DataModel.InsertFilmesLista(lista, requestData.Filme)

		// Salva a lista atualizada
		DataModel.SaveLista(updatedLista)

		// Retorna a lista atualizada como JSON
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(updatedLista)
	} else {
		// Retorna um erro 405 (Method Not Allowed) se o método HTTP não for PUT
		http.Error(res, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Função para atualizar a avaliação (rating) de um filme
func UpdateRatingHandler(res http.ResponseWriter, req *http.Request) {
	// Verifica se o método HTTP é POST
	if req.Method == "POST" {
		var rating DataModel.Rating
		// Decodifica o corpo da requisição JSON para a struct Rating
		err := json.NewDecoder(req.Body).Decode(&rating)
		if err != nil {
			// Retorna um erro 400 (Bad Request) se a decodificação falhar
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		// Chama a função para atualizar a avaliação do filme
		updatedRating, err := DataModel.UpdateRating(rating.Id_User, rating.Id_Filme, rating.Stars)
		if err != nil {
			// Retorna um erro 500 (Internal Server Error) se a atualização da avaliação falhar
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		// Codifica a avaliação atualizada como JSON e escreve na resposta
		json.NewEncoder(res).Encode(updatedRating)
	} else {
		// Retorna um erro 405 (Method Not Allowed) se o método HTTP não for POST
		http.Error(res, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
