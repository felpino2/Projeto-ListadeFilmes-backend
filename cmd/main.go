package main

import (
	//pacote CORS -> permite requisições de diferentes origens (ajuda na implementação do frontend)
	"github.com/rs/cors"
	"log"
	"net/http"
	"ps-backend-felipe-rodrigues/src/database"
	"ps-backend-felipe-rodrigues/src/requests"
)

func main() {
	log.Println("Server is starting...")

	// Define a URI de conexão com o MongoDB Atlas
	uri := "mongodb+srv://felipe:1234@furz-parfum.jwzwsza.mongodb.net/mydatabase?retryWrites=true&w=majority&ssl=true"

	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado
	client, err := database.ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Define o cliente MongoDB globalmente
	database.Client = client

	// Middleware CORS
	corsHandler := cors.Default().Handler

	// Configura a rota para registro de usuários usando a função RegisterUserHandler
	http.HandleFunc("/register", requests.RegisterUserHandler)

	// Inicia o servidor HTTP na porta 5173
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux)))
}

/*user := components.User{
	Name:        "Felipe Queiroz",
	DisplayName: "FelipePirocade2cm",
	Password:    "peidofrouxoquenemfazbarulhomasfede123",
}*/
