package main

import (
	"awesomeProject/psbackllfa/src/database"
	"awesomeProject/psbackllfa/src/requests"
	//pacote CORS -> permite requisições de diferentes origens (ajuda na implementação do frontend)
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	database.IniciarMongo()

	// Middleware CORS
	corsHandler := cors.Default().Handler

	// Configura a rota para registro de usuários usando a função RegisterUserHandler
	http.HandleFunc("/register", requests.RegisterUserHandler)

	// Inicia o servidor HTTP na porta 5173
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux)))
}
