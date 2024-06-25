package main

import (
	"psbackllfa/src/database"
	"psbackllfa/src/requests"
	//pacote CORS -> permite requisições de diferentes origens (ajuda na implementação do frontend)
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	uri := "mongodb+srv://livialemos:LMPRY@trabalhocaio.fay3v2j.mongodb.net/?retryWrites=true&w=majority&appName=TRABALHOCAIO"
	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado.
	client, err := database.ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	database.Client = client
	// Middleware CORS
	corsHandler := cors.Default().Handler

	// Configura a rota para registro de usuários usando a função RegisterUserHandler
	http.HandleFunc("/register", requests.RegisterUserHandler)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux)))
}
